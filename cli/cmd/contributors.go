// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"docwiz/internal/git"
	"docwiz/internal/io"
	"docwiz/internal/style"

	"github.com/caarlos0/log"
	"github.com/spf13/cobra"
)

type ContributorsCmdParameter struct {
	baseParameter

	// repoPath specifies the path to the Git repository, from which information like tags will be gathered.
	// The default value is the current directory ("./").
	repoPath string
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
  docwiz contributors --disable-copyright`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Infof("creating %s", contributorsParameter.output)
			output, err := io.NewSafeFile(contributorsParameter.output)
			if err != nil {
				log.WithError(err).Fatal("fail to create file")
			}
			defer output.Close()

			defer func() {
				if err := recover(); err != nil {
					output.Rollback()
					log.WithError(err.(error)).Fatal("error happen and rollback!")
				}
			}()

			log.WithField("path", contributorsParameter.repoPath).Info("parsing .git directory")
			r, err := git.New(contributorsParameter.repoPath)
			if err != nil {
				log.WithError(err).Fatal("fail to read git repository")
			}

			log.Infof("generating %s", style.Bold(contributorsParameter.output))
			err = r.GenerateContributors(output)
			if err != nil {
				log.WithError(err).Fatal("fail to generate contributors")
			}

			if !contributorsParameter.disableCopyright {
				output.Write(COPYRIGHT)
			}
			log.Info("thanks for using docwiz!")
		},
	}
)

func init() {
	docwizCmd.AddCommand(contributorsCmd)
	contributorsCmd.PersistentFlags().StringVarP(&contributorsParameter.output, "output", "o", "CONTRIBUTORS.md", "Path to the output contributors file")
	contributorsCmd.PersistentFlags().StringVarP(&contributorsParameter.repoPath, "repo", "r", ".", "Path to the target Git repository")
	contributorsCmd.PersistentFlags().BoolVarP(&contributorsParameter.disableCopyright, "disable-copyright", "d", false, "Disable copyright information in the contributors")
}
