// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"docwiz/internal/io"

	"docwiz/internal/os"
	"docwiz/internal/style"
	"docwiz/internal/template"
	"fmt"

	"path/filepath"

	"github.com/caarlos0/log"
	"github.com/spf13/cobra"
)

// RoadMapCmdParameter stores parameters for the "roadmap" command.
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

			log.Infof("creating %s", roadMapParameter.output)
			output, err := io.NewSafeFile(roadMapParameter.output)
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

			log.WithField("kind", roadMapParameter.kind).
				WithField("theme", roadMapParameter.theme).
				WithField("language", roadMapParameter.language).
				WithField("target", tpl).Info("loading template")
			tmpl, err := template.Default(tpl)
			if err != nil {
				log.WithError(err).Fatal("fail to load template")
			}

			version := roadMapParameter.data["version"]
			log.Info("executing template")
			log.IncreasePadding()
			log.WithField("version", version).Info("parameters")
			if len(version) == 0 {
				log.Warnf("you can use %s to specify the version parameter", style.Keyword("-d version=1.0.0"))
			}
			log.DecreasePadding()

			err = tmpl.Execute(output, map[string]any{
				"Version": version,
			})

			if err != nil {
				log.WithError(err).Fatal("fail to load template")
			}

			if !roadMapParameter.disableCopyright {
				output.Write(COPYRIGHT)
			}
			log.Infof("generating %s", style.Bold(roadMapParameter.output))
			log.Info("thanks for using docwiz!")
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
