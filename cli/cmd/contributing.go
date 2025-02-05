// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"docwiz/internal/git"
	"docwiz/internal/io"
	"docwiz/internal/log"
	"docwiz/internal/os"
	. "docwiz/internal/template"
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/spf13/cobra"
)

type contributingCmdParameter struct {
	// The path and filename of the output file
	output    string
	theme     string
	copyright bool
	repoPath  string
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

			repo, err := git.New(contributingParameter.repoPath)
			if err == nil {
				name = repo.Name()
				owner = repo.Owner()
			} else {
				log.Warnf("fail to read git repository, err: %s", err.Error())
			}

			contributingPath := filepath.Join(os.TemplatePath, "CONTRIBUTING")

			tpl := filepath.Join(contributingPath, fmt.Sprintf("%s.tpl", contributingParameter.theme))
			output, err := io.NewSafeFile(contributingParameter.output)
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

			tmpl, err := template.New(filepath.Base(tpl)).Funcs(DocwizFuncMap(contributingPath)).ParseFiles(tpl)

			if err != nil {
				panic(err)
			}

			err = tmpl.Execute(output, map[string]any{
				"ProjectName":  name,
				"ProjectOwner": owner,
			})
			if err != nil {
				panic(err)
			}

			if contributingParameter.copyright {
				output.Write([]byte(COPYRIGHT))
			}
		},
	}
)

func init() {
	docwizCmd.AddCommand(contributingCmd)
	contributingCmd.PersistentFlags().StringVarP(&contributingParameter.output, "output", "o", "CONTRIBUTING.md", "Path to save the generated contributing file")
	contributingCmd.PersistentFlags().StringVarP(&contributingParameter.theme, "theme", "t", "default", "Theme for the contributing template")
	contributingCmd.PersistentFlags().BoolVarP(&contributingParameter.copyright, "copyright", "c", true, "Include copyright information in the contributing")
	contributingCmd.PersistentFlags().StringVarP(&contributingParameter.repoPath, "repo", "r", ".", "Path to the target Git repository")
}
