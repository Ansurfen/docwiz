// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"docwiz/internal/git"
	"docwiz/internal/io"
	"docwiz/internal/os"
	. "docwiz/internal/template"
	"docwiz/internal/tui"
	"docwiz/internal/walk"
	"fmt"
	"html/template"
	"io/fs"
	"path/filepath"

	"github.com/spf13/cobra"
)

type readmeCmdParameter struct {
	baseParameter
	scan     bool
	template string
}

var (
	readmeParameter readmeCmdParameter
	readmeCmd       = &cobra.Command{
		Use:   "readme",
		Short: "Generate a README.md file for your project",
		Long: `The 'readme' command allows you to generate a README.md file
based on predefined templates. You can specify the programming language, 
theme, and whether to include copyright information.`,
		Example: `docwiz readme -s
  docwiz readme -l go -t default -o README.md
  docwiz readme -l python -o docs/README.md
  docwiz readme`,
		Run: func(cmd *cobra.Command, args []string) {
			if readmeParameter.scan {
				ignore, err := git.CompileIgnoreFile(".docwizignore")
				if err != nil {
					ignore = &git.GitIgnore{}
				}
				if len(readmeParameter.output) == 0 {
					readmeParameter.output = "README.md"
				}

				tpl := filepath.Join(os.TemplatePath, "README", "README.tpl")
				if readmeParameter.language != defaultLanguage {
					tpl = filepath.Join(os.TemplatePath, "README", readmeParameter.language, "README.tpl")
				}

				ctx := &walk.Context{
					Ignore:   ignore,
					Output:   readmeParameter.output,
					Template: tpl,
					Walkers: []walk.Walker{
						&walk.AndroidWalker{},
						&walk.BashWalker{},
						&walk.CWalker{},
						&walk.ClojureWalker{},
						&walk.CMakeWalker{},
						&walk.CPPWalker{},
						&walk.CrystalWalker{},
						&walk.CSharpWalker{},
						&walk.CSSWalker{},
						&walk.CudaWalker{},
						&walk.DartWalker{},
						&walk.DockerWalker{},
						&walk.ElixirWalker{},
						&walk.ElmWalker{},
						&walk.ErlangWalker{},
						&walk.FortranWalker{},
						&walk.GDScriptWalker{},
						&walk.GitWalker{},
						&walk.GoWalker{},
						&walk.GradleWalker{},
						&walk.GraphQLWalker{},
						&walk.GroovyWalker{},
						&walk.HaskellWalker{},
						&walk.HTMLWalker{},
						&walk.JavaWalker{},
						&walk.JavaScriptWalker{},
						&walk.JSPWalker{},
						&walk.JuliaWalker{},
						&walk.JupyterWalker{},
						&walk.KotlinWalker{},
						&walk.LaTeXWalker{},
						&walk.LuaWalker{},
						&walk.MarkdownWalker{},
						&walk.NimWalker{},
						&walk.NixWalker{},
						&walk.ObjectiveCWalker{},
						&walk.OCamlWalker{},
						&walk.PerlWalker{},
						&walk.PHPWalker{},
						&walk.PowerShellWalker{},
						&walk.PythonWalker{},
						&walk.QTWalker{},
						&walk.RWaler{},
						&walk.ReScriptWalker{},
						&walk.RubyWalker{},
						&walk.RustWalker{},
						&walk.ScalaWalker{},
						&walk.SolidityWalker{},
						&walk.SwiftWalker{},
						&walk.TypeScriptWalker{},
						&walk.VSCodeWalker{},
						&walk.YAMLWalker{},
						&walk.YarnWalker{},
						&walk.ZigWalker{},
					},
				}
				walk.Walk(".", ctx)

				gen := &generator{
					output: readmeParameter.output,
					action: func() {

						tmpl, err := template.New(filepath.Base(tpl)).Funcs(DocwizFuncMap(tpl)).ParseFiles(tpl)
						if err != nil {
							panic(err)
						}
						output, err := io.NewSafeFile(readmeParameter.output)
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

						err = tmpl.Execute(output, map[string]any{
							"ProjectName":        ctx.ProjectName,
							"ProjectOwner":       ctx.ProjectOwner,
							"ProjectStack":       ctx.ProjectStack,
							"ProjectDescription": ctx.ProjectDescription,
							"Sections":           ctx.Sections,
						})

						if err != nil {
							panic(err)
						}

						if !readmeParameter.disableCopyright {
							output.Write(COPYRIGHT)
						}
					},
				}
				gen.run()
			} else {
				var language string
				if len(readmeParameter.language) != 0 {
					language = readmeParameter.language
				} else {
					languages := []string{}
					readmePath := filepath.Join(os.TemplatePath, "README")
					filepath.Walk(readmePath, func(path string, info fs.FileInfo, err error) error {
						if err != nil {
							return err
						}
						if path == readmePath {
							return nil
						}
						if info.IsDir() {
							path = filepath.Base(path)
							languages = append(languages, path)
						}
						return nil
					})
					m := tui.NewSelectModel(tui.SelectModelConfigure{
						Title:       "language",
						Description: "What language is your project based on? (or none)",
						Placeholder: "Search or press enter to select the language",
						SelectTitle: "Pick a language",
						Candicates:  languages,
					})

					if err := m.Run(); err != nil {
						panic(err)
					}
					language = m.Value()
				}

				var (
					templateDir = filepath.Join(os.TemplatePath, "README")
					questions   = tui.DefaultQuestion
				)

				if language != "none" {
					templateDir = filepath.Join(templateDir, language)
					var index tui.IndexFile

					err := io.ReadJSON(filepath.Join(templateDir, "index.json"), &index)
					if err != nil {
						panic(err)
					}
					questions = append(questions, index.Questions...)
				}

				m := tui.NewReadmeModel(questions...)

				fmt.Println("🥳 Welcome to use docwiz to create readme.md (use tab to enable default)")
				if err := m.Run(); err != nil {
					panic(err)
				}

				gen := &generator{
					output: readmeParameter.output,
					action: func() {
						tpl := filepath.Join(templateDir, fmt.Sprintf("%s.tpl", readmeParameter.theme))
						if readmeParameter.language != defaultLanguage {
							tpl = filepath.Join(templateDir, readmeParameter.language, fmt.Sprintf("%s.tpl", readmeParameter.theme))
						}

						tmpl, err := template.New(filepath.Base(tpl)).Funcs(DocwizFuncMap(templateDir)).ParseFiles(tpl)
						if err != nil {
							panic(err)
						}

						output, err := io.NewSafeFile(readmeParameter.output)
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

						err = tmpl.Execute(output, m.Value())
						if err != nil {
							panic(err)
						}

						if !readmeParameter.disableCopyright {
							output.Write(COPYRIGHT)
						}
					},
				}
				gen.run()
			}
		},
	}
)

func init() {
	docwizCmd.AddCommand(readmeCmd)
	readmeCmd.PersistentFlags().StringVarP(&readmeParameter.output, "output", "o", "README.md", "Path to save the generated README file")
	readmeCmd.PersistentFlags().StringVarP(&readmeParameter.template, "template", "T", "", "Programming language for the README template")
	readmeCmd.PersistentFlags().StringVarP(&readmeParameter.theme, "theme", "t", "default", "Theme of the README template")
	readmeCmd.PersistentFlags().BoolVarP(&readmeParameter.disableCopyright, "disable-copyright", "d", false, "Disable copyright information in the README")
	readmeCmd.PersistentFlags().BoolVarP(&readmeParameter.scan, "scan", "s", false, "Automatically scan and generate")
	readmeCmd.PersistentFlags().StringVarP(&readmeParameter.language, "language", "l", "en_us", "Set the language for contributing file (e.g. zh_cn)")
}
