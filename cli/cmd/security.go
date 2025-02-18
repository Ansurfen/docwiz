// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"docwiz/internal/git"
	"docwiz/internal/io"
	"docwiz/internal/os"
	"docwiz/internal/template"
	"fmt"
	"path/filepath"

	"github.com/caarlos0/log"
	"github.com/spf13/cobra"
)

// securityCmdParameter stores parameters for the "security" command.
type securityCmdParameter struct {
	baseParameter

	// repoPath specifies the path to the Git repository, from which information like tags will be gathered.
	// The default value is the current directory ("./").
	repoPath string
	email    string
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

			log.WithField("path", securityParameter.repoPath).Info("parsing .git directory")
			repo, err := git.New(securityParameter.repoPath)
			if err == nil {
				name = repo.Name()
				owner = repo.Owner()
			} else {
				log.Warnf("fail to read git repository, err: %s", err.Error())
			}

			securityPath := filepath.Join(os.TemplatePath, "SECURITY")
			if securityParameter.language != defaultLanguage {
				securityPath = filepath.Join(securityPath, securityParameter.language)
			}
			tpl := filepath.Join(securityPath, fmt.Sprintf("%s.tpl", securityParameter.theme))

			log.Infof("creating %s", securityParameter.output)
			output, err := io.NewSafeFile(securityParameter.output)
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

			log.WithField("target", tpl).Info("loading template")
			tmpl, err := template.Default(tpl)

			if err != nil {
				log.WithError(err).Fatal("fail to load template")
			}

			log.Info("executing template")
			log.IncreasePadding()
			log.WithField("ProjectName", name).
				WithField("ProjectOwner", owner).
				WithField("Email", securityParameter.email).Info("parameters")
			log.DecreasePadding()
			err = tmpl.Execute(output, map[string]any{
				"ProjectName":  name,
				"ProjectOwner": owner,
				"Email":        securityParameter.email,
			})
			if err != nil {
				log.WithError(err).Fatal("fail to execute template")
			}

			if !securityParameter.disableCopyright {
				output.Write(COPYRIGHT)
			}

			log.Info("thanks for using docwiz!")
		},
	}
)

func init() {
	docwizCmd.AddCommand(securityCmd)
	securityCmd.PersistentFlags().StringVarP(&securityParameter.output, "output", "o", "SECURITY.md", "Path to save the generated security file")
	securityCmd.PersistentFlags().StringVarP(&securityParameter.theme, "theme", "t", "default", "Theme for the security template")
	securityCmd.PersistentFlags().BoolVarP(&securityParameter.disableCopyright, "disable-copyright", "d", false, "Disable copyright information in the security")
	securityCmd.PersistentFlags().StringVarP(&securityParameter.repoPath, "repo", "r", ".", "Path to the target Git repository")
	securityCmd.PersistentFlags().StringVarP(&securityParameter.email, "email", "e", "", "Email to contact and report issues")
	securityCmd.PersistentFlags().StringVarP(&securityParameter.language, "language", "l", "en_us", "Set the language for contributing file (e.g. zh_cn)")
}
