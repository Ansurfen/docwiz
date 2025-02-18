// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"docwiz/internal/io"
	"docwiz/internal/style"

	"docwiz/internal/os"
	"docwiz/internal/template"
	"fmt"

	"path/filepath"
	"strings"

	"github.com/caarlos0/log"
	"github.com/spf13/cobra"
)

// authorsCmdParameter to store parameters related to the authors command.
type authorsCmdParameter struct {
	baseParameter

	// The type of license (e.g., MIT, Apache, etc.)
	license string

	// maintainers is a list of primary maintainers responsible for overseeing the project.
	// They have authority over major decisions, code reviews, and releases.
	maintainers []string

	// contributors is a list of individuals who have contributed to the project
	// through code, documentation, or other improvements.
	contributors []string

	// specialContributors is a list of individuals who have made significant or
	// notable contributions to the project, such as funding, mentorship, or key features.
	specialContributors []string
}

var (
	authorsParameter authorsCmdParameter
	authorsCmd       = &cobra.Command{
		Use:   "authors",
		Short: "Generate an AUTHORS file with maintainers and contributors.",
		Long: `The 'authors' command generates an AUTHORS file that includes 
		maintainers, contributors, and special contributors based on the provided details.`,
		Example: `  docwiz authors -o AUTHORS.md authors -m 'name=Alice,duty="Lead Developer"'
    -c 'name=Bob,duty="Contributor"' -s 'name=Charlie,duty="Special Contributor"'`,
		Run: func(cmd *cobra.Command, args []string) {
			var maintainers, contributors, specialContributors []user
			log.Info("parsing users")
			for _, m := range authorsParameter.maintainers {
				user, err := unmarshalUser(m)
				if err != nil {
					log.WithError(err).Fatal("parsing maintainers")
				}
				log.IncreasePadding()
				log.Infof("%s, maintainer", user.Name)
				log.DecreasePadding()
				maintainers = append(maintainers, user)
			}

			for _, c := range authorsParameter.contributors {
				user, err := unmarshalUser(c)
				if err != nil {
					log.WithError(err).Fatal("parsing contributors")
				}
				log.IncreasePadding()
				log.Infof("%s, contributors", user.Name)
				log.DecreasePadding()
				contributors = append(contributors, user)
			}

			for _, s := range authorsParameter.specialContributors {
				user, err := unmarshalUser(s)
				if err != nil {
					log.WithError(err).Fatal("parsing special contributors")
				}
				log.IncreasePadding()
				log.Infof("%s, special contributors", user.Name)
				log.DecreasePadding()
				specialContributors = append(specialContributors, user)
			}

			authrosPath := filepath.Join(os.TemplatePath, "AUTHORS")
			if authorsParameter.language != defaultLanguage {
				authrosPath = filepath.Join(os.TemplatePath, "AUTHORS", authorsParameter.language)
			}

			tpl := filepath.Join(authrosPath, fmt.Sprintf("%s.tpl", authorsParameter.theme))

			log.Infof("creating %s", authorsParameter.output)
			output, err := io.NewSafeFile(authorsParameter.output)
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

			log.WithField("target", tpl).Info("loading template")
			tmpl, err := template.Default(tpl)
			if err != nil {
				log.WithError(err).Fatal("fail to load template")
			}

			log.Info("executing template")
			log.IncreasePadding()
			log.WithField("Maintainers", authorsParameter.maintainers).
				WithField("Contributors", authorsParameter.contributors).
				WithField("SpecialContributors", authorsParameter.specialContributors).
				WithField("License", authorsParameter.license).Info("parameters")
			log.DecreasePadding()
			err = tmpl.Execute(output, map[string]any{
				"Maintainers":         maintainers,
				"Contributors":        contributors,
				"SpecialContributors": specialContributors,
				"License":             authorsParameter.license,
			})
			if err != nil {
				log.WithError(err).Fatal("fail to execute template")
			}

			if !authorsParameter.disableCopyright {
				output.Write(COPYRIGHT)
			}
			log.Infof("generating %s", style.Bold(authorsParameter.output))
			log.Info("thanks for using docwiz!")
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
	authorsCmd.PersistentFlags().StringVarP(&authorsParameter.license, "license", "L", "", "License type to include in the AUTHORS file (e.g., MIT, Apache)")
	authorsCmd.PersistentFlags().StringArrayVarP(&authorsParameter.maintainers, "maintainers", "m", []string{}, "List of maintainers in JSON format")
	authorsCmd.PersistentFlags().StringArrayVarP(&authorsParameter.contributors, "contributors", "c", []string{}, "List of contributors in JSON format")
	authorsCmd.PersistentFlags().StringArrayVarP(&authorsParameter.specialContributors, "special-contributors", "s", []string{}, "List of special contributors in JSON format")
	authorsCmd.PersistentFlags().BoolVarP(&authorsParameter.disableCopyright, "disable-copyright", "d", false, "Disable copyright information in the authors")
	authorsCmd.PersistentFlags().StringVarP(&authorsParameter.language, "language", "l", "en_us", "Set the language for contributing file (e.g. zh_cn)")
	authorsCmd.PersistentFlags().BoolVarP(&authorsParameter.verbose, "verbose", "v", false, "")
}

// parse key=value to map
func unmarshal(v string) (map[string]string, error) {
	result := make(map[string]string)
	pairs := strings.Split(v, ",")

	for _, pair := range pairs {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid format: %s", pair)
		}

		key := strings.TrimSpace(kv[0])
		value := strings.TrimSpace(kv[1])

		if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
			value = strings.Trim(value, "\"")
		}

		result[key] = value
	}

	return result, nil
}

func unmarshalUser(v string) (user, error) {
	u := user{IsIndividual: false, Others: make(map[string]string)}
	data, err := unmarshal(v)
	if err != nil {
		return u, err
	}

	for k, v := range data {
		switch k {
		case "name":
			u.Name = v
		case "duty":
			u.Duty = v
		case "homepage":
			u.HomePage = v
		case "profile":
			u.Profile = v
		case "isIndividual":
			if v == "true" {
				u.IsIndividual = true
			}
		default:
			u.Others[k] = v
		}
	}
	return u, nil
}

// user represents an individual or an entity that is a contributor or maintainer in the project.
// It includes personal information such as name, duty, and profile, and allows for customization
// with additional fields like homepage and other metadata.
type user struct {
	// Name holds the name of the user (contributor or maintainer).
	Name string

	// Duty represents the role or responsibility of the user in the project (e.g., "Lead Developer").
	Duty string

	// HomePage is a link to the user's personal or project website.
	HomePage string

	// Profile contains the link to the user's profile, like a GitHub or LinkedIn profile.
	Profile string

	// IsIndividual indicates whether the user is an individual (true) or an organization/entity (false).
	IsIndividual bool

	// Others holds any additional metadata about the user that doesn't fit into the predefined fields.
	Others map[string]string
}
