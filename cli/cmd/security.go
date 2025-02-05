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

type securityCmdParameter struct {
	// The path and filename of the output file
	output    string
	theme     string
	copyright bool
	repoPath  string
	email     string
}

var (
	securityParameter securityCmdParameter
	securityCmd       = &cobra.Command{
		Use:   "security",
		Short: "Generate a security guide for your project",
		Long: `The 'security' command helps you generate a security guide for your project, 
	providing templates for common security best practices such as handling vulnerabilities, 
	data privacy, and secure coding guidelines.`,
		Example: "  docwiz security",
		Run: func(cmd *cobra.Command, args []string) {
			var (
				name  string
				owner string
			)

			repo, err := git.New(securityParameter.repoPath)
			if err == nil {
				name = repo.Name()
				owner = repo.Owner()
			} else {
				log.Warnf("fail to read git repository, err: %s", err.Error())
			}

			securityPath := filepath.Join(os.TemplatePath, "SECURITY")
			tpl := filepath.Join(securityPath, fmt.Sprintf("%s.tpl", securityParameter.theme))

			output, err := io.NewSafeFile(securityParameter.output)
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

			tmpl, err := template.New(filepath.Base(tpl)).Funcs(DocwizFuncMap(securityPath)).ParseFiles(tpl)

			if err != nil {
				panic(err)
			}

			err = tmpl.Execute(output, map[string]any{
				"ProjectName":  name,
				"ProjectOwner": owner,
				"Email":        securityParameter.email,
			})
			if err != nil {
				panic(err)
			}

			if securityParameter.copyright {
				output.Write([]byte(COPYRIGHT))
			}
		},
	}
)

func init() {
	docwizCmd.AddCommand(securityCmd)
	securityCmd.PersistentFlags().StringVarP(&securityParameter.output, "output", "o", "SECURITY.md", "Path to save the generated security file")
	securityCmd.PersistentFlags().StringVarP(&securityParameter.theme, "theme", "t", "default", "Theme for the security template")
	securityCmd.PersistentFlags().BoolVarP(&securityParameter.copyright, "copyright", "c", true, "Include copyright information in the security")
	securityCmd.PersistentFlags().StringVarP(&securityParameter.repoPath, "repo", "r", ".", "Path to the target Git repository")
	securityCmd.PersistentFlags().StringVarP(&securityParameter.email, "email", "e", "", "Email to contact and report issues")
}
