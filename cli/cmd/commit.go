// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"docwiz/internal/log"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// commitCmdParameter stores parameters related to the "commit" command.
type commitCmdParameter struct {
	// message is the commit message provided by the user.
	// It can include a conventional commit prefix, such as "feat:", "fix:", or "docs:".
	message string

	// pure determines whether to output only the processed commit message
	// without executing the git commit command.
	// If set to true, the formatted message is printed instead of being committed.
	pure bool

	// exec specifies whether to execute the "git commit" command directly.
	// If set to true, the command is run using "git commit -m <message>".
	exec bool
}

var (
	emojiMap        map[string]string
	commitParameter commitCmdParameter
	commitCmd       = &cobra.Command{
		Use:   "commit",
		Short: "Commit changes with an optional emoji prefix.",
		Long: `The 'commit' command allows you to make a git commit while optionally 
enhancing the commit message with relevant emojis based on predefined mappings.`,
		Example: `  docwiz commit -m "fix: corrected database query"
  docwiz commit -m "feat: added new API endpoint" -e
  docwiz commit -m "docs: updated README" -p`,
		Run: func(cmd *cobra.Command, args []string) {
			execPath, err := os.Executable()
			if err != nil {
				log.Fata(err)
			}
			index := filepath.Join(execPath, "../template/COMMIT/index.json")
			data, err := os.ReadFile(index)
			if err != nil {
				log.Fata(err)
			}

			err = json.Unmarshal(data, &emojiMap)
			if err != nil {
				log.Fata(err)
			}

			msg := addGitEmoji(commitParameter.message)

			if commitParameter.exec {
				cmd := exec.Command("git", "commit", "-m", msg)

				cmd.Stdout = os.Stdout
				cmd.Stdin = os.Stdin
				cmd.Stderr = os.Stderr

				if err = cmd.Run(); err != nil {
					log.Fata(err)
				}
				return
			}

			if commitParameter.pure {
				fmt.Println(msg)
			} else {
				fmt.Printf(`git commit -m "%s"`, msg)
			}

		},
	}
)

func init() {
	docwizCmd.AddCommand(commitCmd)
	commitCmd.PersistentFlags().StringVarP(&commitParameter.message, "message", "m", "", "Commit message to use")
	commitCmd.PersistentFlags().BoolVarP(&commitParameter.pure, "pure", "p", false, "Output only the processed commit message")
	commitCmd.PersistentFlags().BoolVarP(&commitParameter.exec, "exec", "e", false, "Execute the git commit command directly")
}

func addGitEmoji(message string) string {
	// Record the emojis that have been found
	emojiSet := make(map[string]struct{})
	// Split the message into words and convert them to lowercase
	words := strings.Fields(strings.ToLower(message))

	// Iterate over the words to map keywords to emojis
outer:
	for _, word := range words {
		// Loop through each keyword-emoji pair in the emoji map
		for key, emoji := range emojiMap {
			// If the word contains the keyword, add the emoji to the set
			if strings.Contains(word, key) {
				emojiSet[emoji] = struct{}{}
				// If there are more than 2 emojis, stop further processing
				if len(emojiSet) >= 2 {
					break outer
				}
			}
		}
	}

	// Convert the emoji set to a slice and concatenate the emojis
	var emojiList []string
	for emoji := range emojiSet {
		emojiList = append(emojiList, emoji)
	}

	// Prepend the emojis to the message if any emojis are found
	if len(emojiList) > 0 {
		return strings.Join(emojiList, " ") + " " + message
	}
	return message
}
