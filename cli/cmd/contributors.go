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

type ContributorsCmdParameter struct {
	output string

	repoPath string

	copyright bool
}

var (
	contributorsParameter ContributorsCmdParameter
	contributorsCmd       = &cobra.Command{
		Use:   "contributors",
		Short: "Generate a contributors list from a Git repository.",
		Long: `The 'contributors' command scans the Git history of a repository 
to extract and list all contributors who have committed changes.`,
		Example: `  docwiz contributors -o CONTRIBUTORS.md
  docwiz contributors -r /path/to/repo -o contributors.txt
  docwiz contributors --no-copyright`,
		Run: func(cmd *cobra.Command, args []string) {
			output, err := io.NewSafeFile(contributorsParameter.output)
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

			r := git.New(contributorsParameter.repoPath)
			err = r.GenerateContributors(output)
			if err != nil {
				panic(err)
			}

			if contributorsParameter.copyright {
				output.Write([]byte(COPYRIGHT))
			}
		},
	}
)

func init() {
	docwizCmd.AddCommand(contributorsCmd)
	contributorsCmd.PersistentFlags().StringVarP(&contributorsParameter.output, "output", "o", "CONTRIBUTORS.md", "Path to the output contributors file")
	contributorsCmd.PersistentFlags().StringVarP(&contributorsParameter.repoPath, "repo", "r", ".", "Path to the target Git repository")
	contributorsCmd.PersistentFlags().BoolVarP(&contributorsParameter.copyright, "copyright", "c", true, "Include copyright information in the contributors")
}
