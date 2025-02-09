// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"docwiz/internal/io"
	"docwiz/internal/log"
	"docwiz/internal/os"
	"docwiz/internal/template"
	"fmt"

	"path/filepath"

	"github.com/spf13/cobra"
)

type RoadMapCmdParameter struct {
	baseParameter
	kind string
	data map[string]string
}

var (
	roadMapParameter RoadMapCmdParameter
	roadMapCmd       = &cobra.Command{
		Use:   "roadmap",
		Short: "Generate a roadmap file for your project",
		Long: `The 'roadmap' command allows you to generate a roadmap (e.g., for a product or project) 
based on predefined templates and provide information like versioning, kind, theme, etc.`,
		Example: `  docwiz roadmap -k quarter -t default -o ROADMAP.md -d version=1.0.0`,
		Run: func(cmd *cobra.Command, args []string) {
			roadMapPath := filepath.Join(os.TemplatePath, "ROADMAP")

			if roadMapParameter.language != defaultLanguage {
				roadMapPath = filepath.Join(roadMapPath, roadMapParameter.language)
			}
			tpl := filepath.Join(roadMapPath, fmt.Sprintf("%s.%s.tpl", roadMapParameter.kind, roadMapParameter.theme))

			gen := &generator{
				output: roadMapParameter.output,
				action: func() {
					output, err := io.NewSafeFile(roadMapParameter.output)
					if err != nil {
						log.Fata(err)
					}
					defer output.Close()

					defer func() {
						if err := recover(); err != nil {
							output.Rollback()
							fmt.Println(err)
						}
					}()

					tmpl, err := template.Default(tpl)
					if err != nil {
						log.Fata(err)
					}

					err = tmpl.Execute(output, map[string]any{
						"Version": roadMapParameter.data["version"],
					})

					if err != nil {
						log.Fata(err)
					}

					if !roadMapParameter.disableCopyright {
						output.Write(COPYRIGHT)
					}
				},
			}
			gen.run()
		},
	}
)

func init() {
	docwizCmd.AddCommand(roadMapCmd)
	roadMapCmd.PersistentFlags().StringVarP(&roadMapParameter.output, "output", "o", "ROADMAP.md", "Path to save the generated roadmap file")
	roadMapCmd.PersistentFlags().StringVarP(&roadMapParameter.theme, "theme", "t", "default", "Theme for the roadmap template")
	roadMapCmd.PersistentFlags().StringVarP(&roadMapParameter.kind, "kind", "k", "quarter", "Kind of roadmap (e.g., quarter, version, etc.)")
	roadMapCmd.PersistentFlags().StringToStringVarP(&roadMapParameter.data, "data", "d", nil, "Additional data to inject into the template (e.g., version=1.0.0)")
	roadMapCmd.PersistentFlags().BoolVarP(&roadMapParameter.disableCopyright, "disable-copyright", "", false, "Disable copyright information in the roadmap")
	roadMapCmd.PersistentFlags().StringVarP(&roadMapParameter.language, "language", "l", "en_us", "Set the language for contributing file (e.g. zh_cn)")
}
