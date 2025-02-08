// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"docwiz/internal/io"
	"docwiz/internal/os"
	"docwiz/internal/tui"

	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type gitignoreCmdParameter struct {
	output   string
	template string
}

const NoneGitignore = "none"

var (
	gitignoreParameter gitignoreCmdParameter
	gitignoreCmd       = &cobra.Command{
		Use:   "gitignore",
		Short: "Generate a .gitignore file based on predefined templates.",
		Long: `The 'gitignore' command helps generate a .gitignore file 
for your project by selecting from a list of predefined templates. 
You can specify a template directly or pick one interactively.`,
		Example: `  docwiz gitignore -t Go -o .gitignore
  docwiz gitignore -t Python
  docwiz gitignore`,
		Run: func(cmd *cobra.Command, args []string) {
			gitignorePath := filepath.Join(os.TemplatePath, "GITIGNORE")
			index := map[string]string{}
			key2File := map[string]string{}

			err := filepath.Walk(gitignorePath, func(path string, info fs.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() || filepath.Ext(path) != ".gitignore" {
					return nil
				}
				relPath, err := filepath.Rel(gitignorePath, path)
				if err != nil {
					return err
				}

				parts := strings.Split(relPath, string(filepath.Separator))
				filename := parts[len(parts)-1]
				tags := []string{}
				if len(parts)-2 >= 0 {
					tags = parts[:len(parts)-1]
				}
				tagsStr := ""
				var key string
				if len(tags) != 0 {
					tagsStr = fmt.Sprintf("(%s)", strings.Join(tags, ", "))
					key = fmt.Sprintf("%s %s", getFileNameWithoutExtension(filename), tagsStr)
				} else {
					key = getFileNameWithoutExtension(filename)
				}

				index[key] = path
				key2File[key] = filename
				return nil
			})

			if err != nil {
				panic(err)
			}

			var key string

			if gitignoreParameter.template != NoneGitignore {
				key = gitignoreParameter.template
			} else {
				gitignores := []string{NoneGitignore}
				for k := range index {
					gitignores = append(gitignores, k)
				}

				m := tui.NewSelectModel(tui.SelectModelConfigure{
					Title:       "Gitignore",
					Description: "What gitignore is your project based on? (or none)",
					Placeholder: "Search or press enter to select the gitignore",
					SelectTitle: "Pick a gitignore",
					Candicates:  gitignores,
				})
				if err = m.Run(); err != nil {
					panic(err)
				}

				key = m.Value()
			}

			if v := index[key]; len(v) != 0 {
				if len(gitignoreParameter.output) == 0 {
					gitignoreParameter.output = ".gitignore"
				}
				gen := &generator{
					output: gitignoreParameter.output,
					action: func() {
						io.WriteFileFrom(v, gitignoreParameter.output)
					},
				}
				gen.run()
			}

			if err != nil {
				panic(err)
			}
		},
	}
)

func init() {
	docwizCmd.AddCommand(gitignoreCmd)
	gitignoreCmd.PersistentFlags().StringVarP(&gitignoreParameter.output, "output", "o", "", "Path to save the generated .gitignore file")
	gitignoreCmd.PersistentFlags().StringVarP(&gitignoreParameter.template, "template", "t", "none", "Specify a predefined gitignore template (e.g., Go, Python)")
}

func getFileNameWithoutExtension(filePath string) string {
	fileName := filepath.Base(filePath)
	ext := filepath.Ext(fileName)
	return fileName[:len(fileName)-len(ext)]
}
