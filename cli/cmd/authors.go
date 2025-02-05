// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// authorsCmdParameter to store parameters related to the authors command.
type authorsCmdParameter struct {
	theme string

	// The path and filename of the output file
	output string

	// The type of license (e.g., MIT, Apache, etc.)
	license string

	maintainers []string

	contributors []string

	specialContributors []string

	copyright bool
}

var (
	authorsParameter authorsCmdParameter
	authorsCmd       = &cobra.Command{
		Use:   "authors",
		Short: "Generate an AUTHORS file with maintainers and contributors.",
		Long: `The 'authors' command generates an AUTHORS file that includes 
		maintainers, contributors, and special contributors based on the provided details.`,
		Example: `  docwiz authors -o AUTHORS.md -t default -m '{"name":"Alice","duty":"Lead Developer"}' \
	-c '{"name":"Bob","duty":"Contributor"}' -s '{"name":"Charlie","duty":"Special Contributor"}'`,
		Run: func(cmd *cobra.Command, args []string) {
			var maintainers, contributors, specialContributors []User

			for _, m := range authorsParameter.maintainers {
				var user User
				err := json.Unmarshal([]byte(m), &user)
				if err != nil {
					panic(err)
				}
				maintainers = append(maintainers, user)
			}

			for _, c := range authorsParameter.contributors {
				var user User
				err := json.Unmarshal([]byte(c), &user)
				if err != nil {
					panic(err)
				}
				contributors = append(contributors, user)
			}

			for _, s := range authorsParameter.specialContributors {
				var user User
				err := json.Unmarshal([]byte(s), &user)
				if err != nil {
					panic(err)
				}
				specialContributors = append(specialContributors, user)
			}

			execPath, err := os.Executable()
			if err != nil {
				panic(err)
			}

			tpl := filepath.Join(execPath, fmt.Sprintf("../template/AUTHORS/%s.tpl", authorsParameter.theme))

			output, err := os.Create(authorsParameter.output)
			if err != nil {
				panic(err)
			}
			defer output.Close()

			tmpl, err := template.ParseFiles(tpl)
			if err != nil {
				panic(err)
			}

			err = tmpl.Execute(output, map[string]any{
				"Maintainers":         maintainers,
				"Contributors":        contributors,
				"SpecialContributors": specialContributors,
				"License":             authorsParameter.license,
				"Copyright":           authorsParameter.copyright,
			})
			if err != nil {
				panic(err)
			}
		},
	}
)

// Initialization function to set up flags for the authors command
func init() {
	// Add the authors command to the docwizCmd (presumably another main command)
	docwizCmd.AddCommand(authorsCmd)
	// Define the flags for the authors command with their default values and descriptions
	authorsCmd.PersistentFlags().StringVarP(&authorsParameter.theme, "theme", "t", "default", "Template theme to use for the AUTHORS file")
	authorsCmd.PersistentFlags().StringVarP(&authorsParameter.output, "output", "o", "AUTHORS.md", "Path to the output authors file")
	authorsCmd.PersistentFlags().StringVarP(&authorsParameter.license, "license", "l", "", "License type to include in the AUTHORS file (e.g., MIT, Apache)")
	authorsCmd.PersistentFlags().StringArrayVarP(&authorsParameter.maintainers, "maintainers", "m", []string{}, "List of maintainers in JSON format")
	authorsCmd.PersistentFlags().StringArrayVarP(&authorsParameter.contributors, "contributors", "c", []string{}, "List of contributors in JSON format")
	authorsCmd.PersistentFlags().StringArrayVarP(&authorsParameter.specialContributors, "special-contributors", "s", []string{}, "List of special contributors in JSON format")
}

type User struct {
	Name     string `json:"name"`
	Duty     string `json:"duty"`
	HomePage string `json:"homePage"`
	Profile  string `json:"profile"`

	// personal account (true); a business or institution (false)
	IsIndividual bool `json:"isIndividual"`

	//
	Data map[string]string
}
