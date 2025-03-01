// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"docwiz/internal/io"
	"docwiz/internal/style"

	"docwiz/internal/os"
	"docwiz/internal/tui"
	"fmt"
	"html/template"
	"path/filepath"
	"time"

	"github.com/caarlos0/log"
	"github.com/spf13/cobra"
)

// licenseCmdParameter to store parameters related to the license command.
type licenseCmdParameter struct {
	// The path and filename of the output file
	output string

	// The type of license (e.g., MIT, Apache, etc.)
	license string

	// The year for the project
	year uint

	// The author of the project
	author string
}

const NoneLicense = "none"

// Declare global variables for storing command parameters and the command itself.
var (
	licenseParameter licenseCmdParameter
	licenseCmd       = &cobra.Command{
		Use:   "license",
		Short: "Generate a license file for your project",
		Long: `The 'license' command allows you to generate a license file 
based on predefined templates. You can specify a license type, author, 
and year, or select interactively from available licenses.`,
		Example: `  docwiz license -l MIT -a "John Doe" -y 2025 -o LICENSE
  docwiz license -l Apache -o LICENSE.txt
  docwiz license`,
		Run: func(cmd *cobra.Command, args []string) {
			indexFile := filepath.Join(os.TemplatePath, "LICENSE", "index.json")

			index := map[string]string{}

			err := io.ReadJSON(indexFile, &index)
			if err != nil {
				log.WithError(err).Fatal("reading json")
			}

			var key string

			if licenseParameter.license != NoneLicense {
				key = licenseParameter.license
			} else {
				license := []string{NoneLicense}

				for k := range index {
					license = append(license, k)
				}

				m := tui.NewSelectModel(tui.SelectModelConfigure{
					Title:       "License",
					Description: "What license is your project based on? (or none)",
					Placeholder: "Search or press enter to select the language",
					SelectTitle: "Pick a license",
					Candicates:  license,
				})

				if err = m.Run(); err != nil {
					log.WithError(err).Fatal("running select model")
				}
				key = m.Value()
			}

			if v := index[key]; len(v) != 0 {
				tpl := filepath.Join(os.TemplatePath, fmt.Sprintf("LICENSE/%s.tpl", v))

				log.Infof("creating %s", licenseParameter.output)
				output, err := io.NewSafeFile(licenseParameter.output)
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
				tmpl, err := template.ParseFiles(tpl)
				if err != nil {
					log.WithError(err).Fatal("fail to load template")
				}

				log.Info("executing template")
				log.IncreasePadding()
				log.WithField("Year", licenseParameter.year).
					WithField("Author", licenseParameter.author).
					Info("parameters")
				if len(licenseParameter.author) == 0 {
					log.Warnf("you can use %s to specify the author parameter", style.Keyword("-a docwiz"))
				}
				log.DecreasePadding()
				err = tmpl.Execute(output, map[string]any{
					"Year":   licenseParameter.year,
					"Author": licenseParameter.author,
				})
				if err != nil {
					log.WithError(err).Fatal("fail to execute template")
				}
			}

			log.Infof("generating %s", style.Bold(licenseParameter.output))
			log.Info("thanks for using docwiz!")
		},
	}
)

// Initialization function to set up flags for the license command
func init() {
	// Add the license command to the docwizCmd (presumably another main command)
	docwizCmd.AddCommand(licenseCmd)
	// Define the flags for the license command with their default values and descriptions
	licenseCmd.PersistentFlags().StringVarP(&licenseParameter.output, "output", "o", "LICENSE", "Path to save the generated license file")
	licenseCmd.PersistentFlags().StringVarP(&licenseParameter.license, "license", "l", NoneLicense, "Specify the license type (e.g., MIT, Apache)")
	licenseCmd.PersistentFlags().StringVarP(&licenseParameter.author, "author", "a", "", "Author name for the license")
	licenseCmd.PersistentFlags().UintVarP(&licenseParameter.year, "year", "y", uint(time.Now().Year()), "Year for the license")
}
