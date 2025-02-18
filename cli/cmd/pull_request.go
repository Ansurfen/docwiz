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

// pullRequestCmdParameter stores the parameters for the "pull-request" command.
type pullRequestCmdParameter struct {
	baseParameter
}

var (
	pullRequestParameter pullRequestCmdParameter
	pullRequestCmd       = &cobra.Command{
		Use:     "pull-request",
		Aliases: []string{"pr"},
		Short:   "Generate a pull request template file",
		Long: `This command generates a pull request template file based on the selected theme. 
You can customize the template rendering by specifying the theme. The template will be saved to the output path specified.`,
		Example: `  docwiz pull-request --theme default --output PULL_REQUEST_TEMPLATE.md
  
  # You can also use the 'pr' alias for the same functionality
  docwiz pr`,
		Run: func(cmd *cobra.Command, args []string) {
			prPath := filepath.Join(os.TemplatePath, "PULL_REQUEST")
			tpl := filepath.Join(prPath, fmt.Sprintf("%s.tpl", pullRequestParameter.theme))

			log.Infof("creating %s", pullRequestParameter.output)
			output, err := io.NewSafeFile(pullRequestParameter.output)
			if err != nil {
				log.WithError(err).Fatalf("fail to create %s", pullRequestParameter.output)
			}
			defer output.Close()

			defer func() {
				if err := recover(); err != nil {
					output.Rollback()
					log.WithError(err.(error)).Fatal("error happen and rollback!")
				}
			}()

			log.Info("loading template")
			tmpl, err := template.Default(tpl)
			if err != nil {
				log.WithError(err).Fatal("fail to template")
			}

			log.Info("executing template")
			err = tmpl.Execute(output, nil)
			if err != nil {
				log.WithError(err).Fatal("fail to execute template")
			}

			log.Infof("generating %s", style.Bold(readmeParameter.output))
			log.Info("thanks for using docwiz!")
		},
	}
)

func init() {
	docwizCmd.AddCommand(pullRequestCmd)
	pullRequestCmd.PersistentFlags().StringVarP(&pullRequestParameter.output, "output", "o", "PULL_REQUEST_TEMPLATE.md", "Path to save the generated template file")
	pullRequestCmd.PersistentFlags().StringVarP(&pullRequestParameter.theme, "theme", "t", "default", "Theme for pull request template rendering")
}
