// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"docwiz/internal/git"
	"docwiz/internal/io"
	"docwiz/internal/log"

	"github.com/spf13/cobra"
)

type changelogCmdParameter struct {
	baseParameter

	// repoPath specifies the path to the Git repository, from which information like tags will be gathered.
	// The default value is the current directory ("./").
	repoPath string
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
				log.Fata(err)
			}
			defer output.Close()

			defer func() {
				if err := recover(); err != nil {
					output.Rollback()
					log.Fata(err)
				}
			}()

			r, err := git.New(changelogParameter.repoPath)
			if err != nil {
				log.Fata(err)
			}

			err = r.GenerateChangelog(output)
			if err != nil {
				log.Fata(err)
			}

			if !changelogParameter.disableCopyright {
				output.Write(COPYRIGHT)
			}
		},
	}
)

func init() {
	docwizCmd.AddCommand(changelogCmd)
	changelogCmd.PersistentFlags().StringVarP(&changelogParameter.output, "output", "o", "CHANGELOG.md", "Path to the output changelog file")
	changelogCmd.PersistentFlags().StringVarP(&changelogParameter.repoPath, "repository", "r", ".", "Path to the target Git repository")
	changelogCmd.PersistentFlags().BoolVarP(&changelogParameter.disableCopyright, "disable-copyright", "d", false, "Disable copyright information in the changelog")
}
