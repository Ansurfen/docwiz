// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"docwiz/internal/git"
	"docwiz/internal/io"
	"fmt"

	"github.com/spf13/cobra"
)

type changelogCmdParameter struct {
	output    string
	repoPath  string
	copyright bool
}

var (
	changelogParameter changelogCmdParameter
	changelogCmd       = &cobra.Command{
		Use:   "changelog",
		Short: "Generate a changelog from the Git repository history.",
		Long:  "The 'changelog' command analyzes the commit history of a Git repository and generates a changelog file based on the commits and tags.",
		Example: `  docwiz changelog -o CHANGELOG.md -r /path/to/repo
  docwiz changelog --output my_changelog.md --repository .`,
		Run: func(cmd *cobra.Command, args []string) {
			output, err := io.NewSafeFile(changelogParameter.output)
			if err != nil {
				panic(err)
			}
			defer output.Close()

			defer func() {
				if err := recover(); err != nil {
					output.Rollback()
					fmt.Println(err)
				}
			}()

			r := git.New(changelogParameter.repoPath)

			err = r.GenerateChangelog(output)
			if err != nil {
				panic(err)
			}

			if changelogParameter.copyright {
				output.Write([]byte(COPYRIGHT))
			}
		},
	}
)

func init() {
	docwizCmd.AddCommand(changelogCmd)
	changelogCmd.PersistentFlags().StringVarP(&changelogParameter.output, "output", "o", "CHANGELOG.md", "Path to the output changelog file")
	changelogCmd.PersistentFlags().StringVarP(&changelogParameter.repoPath, "repository", "r", ".", "Path to the target Git repository")
	changelogCmd.PersistentFlags().BoolVarP(&changelogParameter.copyright, "copyright", "c", true, "Include copyright information in the changelog")
}
