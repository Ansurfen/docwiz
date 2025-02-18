// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"docwiz/internal/io"
	"docwiz/internal/style"

	"docwiz/internal/os"
	"docwiz/internal/template"
	"fmt"

	"path/filepath"

	"github.com/caarlos0/log"
	"github.com/spf13/cobra"
)

type CodeOfConductCmdParameter struct {
	baseParameter
	email string
}

var (
	conductParameter CodeOfConductCmdParameter
	conductCmd       = &cobra.Command{
		Use:     "code-of-conduct",
		Aliases: []string{"coc"},
		Short:   "Generate a code of conduct for your project",
		Long: `The 'conduct' command generates a code of conduct file for your project, 
which includes guidelines for respectful behavior, inclusivity, and maintaining a positive community environment.`,
		Example: "  docwiz conduct",
		Run: func(cmd *cobra.Command, args []string) {
			conductPath := filepath.Join(os.TemplatePath, "CODE_OF_CONDUCT")

			tpl := filepath.Join(conductPath, fmt.Sprintf("%s.tpl", conductParameter.theme))
			if conductParameter.language != defaultLanguage {
				tpl = filepath.Join(conductPath, conductParameter.language, fmt.Sprintf("%s.tpl", conductParameter.theme))
			}

			log.Infof("creating %s", authorsParameter.output)
			output, err := io.NewSafeFile(conductParameter.output)
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

			log.WithField("target", tpl).Info("loading template")
			tmpl, err := template.Default(tpl)
			if err != nil {
				log.WithError(err).Fatal("fail to load template")
			}

			log.Info("executing template")
			err = tmpl.Execute(output, map[string]any{
				"Email": conductParameter.email,
			})
			if err != nil {
				log.WithError(err).Fatal("fail to execute template")
			}

			if conductParameter.disableCopyright {
				output.Write([]byte(COPYRIGHT))
			}
			log.Infof("generating %s", style.Bold(conductParameter.output))
			log.Info("thanks for using docwiz!")
		},
	}
)

func init() {
	docwizCmd.AddCommand(conductCmd)
	conductCmd.PersistentFlags().StringVarP(&conductParameter.output, "output", "o", "CODE_OF_CONDUCT.md", "Path to save the generated conduct file")
	conductCmd.PersistentFlags().StringVarP(&conductParameter.theme, "theme", "t", "default", "Theme for the conduct template")
	conductCmd.PersistentFlags().BoolVarP(&conductParameter.disableCopyright, "disable-copyright", "d", false, "Disable copyright information in the conduct")
	conductCmd.PersistentFlags().StringVarP(&conductParameter.email, "email", "e", "", "Email to contact and report issues")
	conductCmd.PersistentFlags().StringVarP(&conductParameter.language, "language", "l", "en_us", "Set the language for contributing file (e.g. zh_cn)")
}
