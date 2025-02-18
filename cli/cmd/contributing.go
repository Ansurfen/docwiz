// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"docwiz/internal/git"
	"docwiz/internal/io"
	"docwiz/internal/style"

	"docwiz/internal/os"
	"docwiz/internal/template"
	"fmt"

	"path/filepath"

	"github.com/caarlos0/log"
	"github.com/spf13/cobra"
)

type contributingCmdParameter struct {
	baseParameter

	// repoPath specifies the path to the Git repository, from which information like tags will be gathered.
	// The default value is the current directory ("./").
	repoPath string
}

var (
	contributingParameter contributingCmdParameter
	contributingCmd       = &cobra.Command{
		Use:   "contributing",
		Short: "Generate a contributing file for your project",
		Long: `The 'contributing' command allows you to generate a contributing guide (e.g., for open-source projects) based on predefined templates. 
You can provide information like contribution guidelines, code of conduct, and setup instructions.`,
		Example: "  docwiz contributing",
		Run: func(cmd *cobra.Command, args []string) {
			var (
				name  string
				owner string
			)

			log.WithField("path", contributingParameter.repoPath).Info("parsing .git directory")
			repo, err := git.New(contributingParameter.repoPath)
			if err == nil {
				name = repo.Name()
				owner = repo.Owner()
			} else {
				log.WithError(err).Warnf("fail to read git repository")
			}

			contributingPath := filepath.Join(os.TemplatePath, "CONTRIBUTING")

			tpl := filepath.Join(contributingPath, fmt.Sprintf("%s.tpl", contributingParameter.theme))
			if contributingParameter.language != defaultLanguage {
				tpl = filepath.Join(contributingPath, contributingParameter.language, fmt.Sprintf("%s.tpl", contributingParameter.theme))
			}

			log.Infof("creating %s", contributingParameter.output)
			output, err := io.NewSafeFile(contributingParameter.output)
			if err != nil {
				log.WithError(err).Fatalf("fail to create file")
			}
			defer output.Close()

			defer func() {
				if err := recover(); err != nil {
					output.Rollback()
					log.WithError(err.(error)).Fatal("error happen and rollback!")
				}
			}()

			log.WithField("theme", contributorsParameter.theme).
				WithField("language", contributingParameter.language).
				WithField("target", tpl).Info("loading template")
			tmpl, err := template.Default(tpl)

			if err != nil {
				log.WithError(err).Fatal("fail to load template")
			}

			log.Info("executing template")
			log.IncreasePadding()
			log.WithField("ProjectName", name).WithField("ProjectOwner", owner).Info("parameters")
			log.DecreasePadding()
			err = tmpl.Execute(output, map[string]any{
				"ProjectName":  name,
				"ProjectOwner": owner,
			})
			if err != nil {
				log.WithError(err).Fatal("fail to execute template")
			}

			if !contributingParameter.disableCopyright {
				output.Write(COPYRIGHT)
			}

			log.Infof("generating %s", style.Bold(changelogParameter.output))
			log.Info("thanks for using docwiz!")
		},
	}
)

func init() {
	docwizCmd.AddCommand(contributingCmd)
	contributingCmd.PersistentFlags().StringVarP(&contributingParameter.output, "output", "o", "CONTRIBUTING.md", "Path to save the generated contributing file")
	contributingCmd.PersistentFlags().StringVarP(&contributingParameter.theme, "theme", "t", "default", "Theme for the contributing template")
	contributingCmd.PersistentFlags().BoolVarP(&contributingParameter.disableCopyright, "disable-copyright", "d", false, "Disable copyright information in the contributing")
	contributingCmd.PersistentFlags().StringVarP(&contributingParameter.repoPath, "repo", "r", ".", "Path to the target Git repository")
	contributingCmd.PersistentFlags().StringVarP(&contributingParameter.language, "language", "l", "en_us", "Set the language for contributing file (e.g. zh_cn)")
}
