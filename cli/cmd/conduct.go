// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"docwiz/internal/io"
	"docwiz/internal/os"
	. "docwiz/internal/template"
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/spf13/cobra"
)

type CodeOfConductCmdParameter struct {
	// The path and filename of the output file
	output    string
	theme     string
	copyright bool
	email     string
}

var (
	conductParameter CodeOfConductCmdParameter
	conductCmd       = &cobra.Command{
		Use:   "conduct",
		Short: "Generate a code of conduct for your project",
		Long: `The 'conduct' command generates a code of conduct file for your project, 
which includes guidelines for respectful behavior, inclusivity, and maintaining a positive community environment.`,
		Example: "  docwiz conduct",
		Run: func(cmd *cobra.Command, args []string) {
			conductPath := filepath.Join(os.TemplatePath, "CODE_OF_CONDUCT")

			tpl := filepath.Join(conductPath, fmt.Sprintf("%s.tpl", conductParameter.theme))
			output, err := io.NewSafeFile(conductParameter.output)
			if err != nil {
				panic(err)
			}
			defer output.Close()

			defer func() {
				if err := recover(); err != nil {
					output.Rollback()
				}
			}()

			tmpl, err := template.New(filepath.Base(tpl)).Funcs(DocwizFuncMap(conductPath)).ParseFiles(tpl)

			if err != nil {
				panic(err)
			}

			err = tmpl.Execute(output, map[string]any{
				"Email": conductParameter.email,
			})
			if err != nil {
				panic(err)
			}

			if conductParameter.copyright {
				output.Write([]byte(COPYRIGHT))
			}
		},
	}
)

func init() {
	docwizCmd.AddCommand(conductCmd)
	conductCmd.PersistentFlags().StringVarP(&conductParameter.output, "output", "o", "CODE_OF_CONDUCT.md", "Path to save the generated conduct file")
	conductCmd.PersistentFlags().StringVarP(&conductParameter.theme, "theme", "t", "default", "Theme for the conduct template")
	conductCmd.PersistentFlags().BoolVarP(&conductParameter.copyright, "copyright", "c", true, "Include copyright information in the conduct")
	conductCmd.PersistentFlags().StringVarP(&conductParameter.email, "email", "e", "", "Email to contact and report issues")
}
