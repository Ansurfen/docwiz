// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"docwiz/internal/cfg"
	"docwiz/internal/io"
	"docwiz/internal/os"
	"docwiz/internal/style"
	"docwiz/internal/template"
	"docwiz/internal/tui"
	"docwiz/internal/walk"
	androidwalk "docwiz/internal/walk/android"
	bashwalk "docwiz/internal/walk/bash"
	cwalk "docwiz/internal/walk/c"
	clojurewalk "docwiz/internal/walk/clojure"
	cmakewalk "docwiz/internal/walk/cmake"
	cppwalk "docwiz/internal/walk/cpp"
	crystalwalk "docwiz/internal/walk/crystal"
	csharpwalk "docwiz/internal/walk/csharp"
	csswalk "docwiz/internal/walk/css"
	cudawalk "docwiz/internal/walk/cuda"
	dartwalk "docwiz/internal/walk/dart"
	dockerwalk "docwiz/internal/walk/docker"
	elixirwalk "docwiz/internal/walk/elixir"
	elmwalk "docwiz/internal/walk/elm"
	erlangwalk "docwiz/internal/walk/erlang"
	fortranwalk "docwiz/internal/walk/fortran"
	gdscriptwalk "docwiz/internal/walk/gdscript"
	gitwalk "docwiz/internal/walk/git"
	gowalk "docwiz/internal/walk/go"
	gradlewalk "docwiz/internal/walk/gradle"
	graphqlwalk "docwiz/internal/walk/graphql"
	groovywalk "docwiz/internal/walk/groovy"
	haskellwalk "docwiz/internal/walk/haskell"
	htmlwalk "docwiz/internal/walk/html"
	javawalk "docwiz/internal/walk/java"
	jswalk "docwiz/internal/walk/js"
	jspwalk "docwiz/internal/walk/jsp"
	juliawalk "docwiz/internal/walk/julia"
	jupyterwalk "docwiz/internal/walk/jupyter"
	kotlinwalk "docwiz/internal/walk/kotlin"
	latexwalk "docwiz/internal/walk/latex"
	luawalk "docwiz/internal/walk/lua"
	mdwalk "docwiz/internal/walk/md"
	nimwalk "docwiz/internal/walk/nim"
	nixwalk "docwiz/internal/walk/nix"
	objectivecwalk "docwiz/internal/walk/oc"
	ocamlwalk "docwiz/internal/walk/ocaml"
	perlwalk "docwiz/internal/walk/perl"
	phpwalk "docwiz/internal/walk/php"
	powershellwalk "docwiz/internal/walk/powershell"
	pythonwalk "docwiz/internal/walk/python"
	qtwalk "docwiz/internal/walk/qt"
	rwalk "docwiz/internal/walk/r"
	rescriptwalk "docwiz/internal/walk/rescript"
	rubywalk "docwiz/internal/walk/ruby"
	rustwalk "docwiz/internal/walk/rust"
	scalawalk "docwiz/internal/walk/scala"
	soliditywalk "docwiz/internal/walk/solidity"
	swiftwalk "docwiz/internal/walk/swift"
	tswalk "docwiz/internal/walk/ts"
	vscodewalk "docwiz/internal/walk/vscode"
	yamlwalk "docwiz/internal/walk/yaml"
	yarnwalk "docwiz/internal/walk/yarn"
	zigwalk "docwiz/internal/walk/zig"
	"fmt"

	"io/fs"
	"path/filepath"

	"github.com/caarlos0/log"
	"github.com/spf13/cobra"
)

// readmeCmdParameter stores the parameters for the "readme" command.
type readmeCmdParameter struct {
	baseParameter

	// scan indicates whether to automatically scan and generate the README file.
	// If true, it will scan the project and generate the README.
	scan bool

	// template defines the programming language template for the README.
	// It is used to determine which template should be applied when generating the README file.
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
				ignore, _ := cfg.LoadDocWizIgnore(".docwizignore")

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
						&androidwalk.Walker{},
						&bashwalk.Walker{},
						&cwalk.Walker{},
						&clojurewalk.Walker{},
						&cmakewalk.Walker{},
						&cppwalk.Walker{},
						&crystalwalk.Walker{},
						&csharpwalk.Walker{},
						&csswalk.Walker{},
						&cudawalk.Walker{},
						&dartwalk.Walker{},
						&dockerwalk.Walker{},
						&elixirwalk.Walker{},
						&elmwalk.Walker{},
						&erlangwalk.Walker{},
						&fortranwalk.Walker{},
						&gdscriptwalk.Walker{},
						&gitwalk.Walker{},
						&gowalk.Walker{},
						&gradlewalk.Walker{},
						&graphqlwalk.Walker{},
						&groovywalk.Walker{},
						&haskellwalk.Walker{},
						&htmlwalk.Walker{},
						&javawalk.Walker{},
						&jswalk.Walker{},
						&jspwalk.Walker{},
						&juliawalk.Walker{},
						&jupyterwalk.Walker{},
						&kotlinwalk.Walker{},
						&latexwalk.Walker{},
						&luawalk.Walker{},
						&mdwalk.Walker{},
						&nimwalk.Walker{},
						&nixwalk.Walker{},
						&objectivecwalk.Walker{},
						&ocamlwalk.Walker{},
						&perlwalk.Walker{},
						&phpwalk.Walker{},
						&powershellwalk.Walker{},
						&pythonwalk.Walker{},
						&qtwalk.Walker{},
						&rwalk.Walker{},
						&rescriptwalk.Walker{},
						&rubywalk.Walker{},
						&rustwalk.Walker{},
						&scalawalk.Walker{},
						&soliditywalk.Walker{},
						&swiftwalk.Walker{},
						&tswalk.Walker{},
						&vscodewalk.Walker{},
						&yamlwalk.Walker{},
						&yarnwalk.Walker{},
						&zigwalk.Walker{},
					},
				}
				walk.Walk(".", ctx)
				tmpl, err := template.New(tpl).LoadStdlib().Parse()
				if err != nil {
					log.WithError(err).Fatal("loading template")
				}

				log.Infof("creating %s", readmeParameter.output)
				output, err := io.NewSafeFile(readmeParameter.output)
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

				err = tmpl.Execute(output, map[string]any{
					"ProjectName":        ctx.ProjectName,
					"ProjectOwner":       ctx.ProjectOwner,
					"ProjectStack":       ctx.ProjectStack,
					"ProjectDescription": ctx.ProjectDescription,
					"Sections":           ctx.Sections,
				})

				if err != nil {
					log.WithError(err).Fatal("executing template")
				} else {
					log.Info("executing template")
				}

				if !readmeParameter.disableCopyright {
					output.Write(COPYRIGHT)
				}
			} else {
				var chosedTemplate string
				if len(readmeParameter.template) != 0 {
					chosedTemplate = readmeParameter.language
				} else {
					templates := []string{}
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
							templates = append(templates, path)
						}
						return nil
					})
					m := tui.NewSelectModel(tui.SelectModelConfigure{
						Title:       "language",
						Description: "What language is your project based on? (or none)",
						Placeholder: "Search or press enter to select the language",
						SelectTitle: "Pick a language",
						Candicates:  templates,
					})

					if err := m.Run(); err != nil {
						log.WithError(err).Fatal("")
					}
					chosedTemplate = m.Value()
				}

				var (
					templateDir = filepath.Join(os.TemplatePath, "README")
					questions   = tui.DefaultQuestion
				)

				if chosedTemplate != "none" {
					templateDir = filepath.Join(templateDir, chosedTemplate)
					var index tui.IndexFile

					err := io.ReadJSON(filepath.Join(templateDir, "index.json"), &index)
					if err != nil {
						log.WithError(err).Fatal("reading json")
					}
					questions = append(questions, index.Questions...)
				}

				m := tui.NewReadmeModel(questions...)

				log.Info("🥳 Welcome to use docwiz to create readme.md (use tab to enable default)")
				if err := m.Run(); err != nil {
					log.WithError(err).Fatal("running readme model")
				}

				tpl := filepath.Join(templateDir, fmt.Sprintf("%s.tpl", readmeParameter.theme))
				if readmeParameter.language != defaultLanguage {
					tpl = filepath.Join(templateDir, readmeParameter.language, fmt.Sprintf("%s.tpl", readmeParameter.theme))
				}

				tmpl, err := template.New(tpl).LoadStdlib().Parse()
				if err != nil {
					log.WithError(err).Fatal("loading template")
				}

				output, err := io.NewSafeFile(readmeParameter.output)
				if err != nil {
					log.WithError(err).Fatalf("creating %s", readmeParameter.output)
				}
				defer output.Close()

				defer func() {
					if err := recover(); err != nil {
						output.Rollback()
						log.WithError(err.(error)).Fatal("error happen and rollback!")
					}
				}()

				err = tmpl.Execute(output, m.Value())
				if err != nil {
					log.WithError(err).Fatal("executing template")
				} else {
					log.Info("executing template")
				}

				if !readmeParameter.disableCopyright {
					output.Write(COPYRIGHT)
				}
			}
			log.Infof("generating %s", style.Bold(readmeParameter.output))
			log.Info("thanks for using docwiz!")
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
