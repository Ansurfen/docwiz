// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"docwiz/internal/git"
	"docwiz/internal/io"
	"docwiz/internal/os"
	. "docwiz/internal/template"
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type issueCmdParameter struct {
	issueName        string
	issueDescription string
	issueAssigness   string
	output           string
	theme            string
	kind             string
	format           string
	repoPath         string
}

const (
	issueFormatMarkdown = "md"
	issueForamtYAML     = "yaml"

	issueKindBug     = "bug"
	issueKindFeature = "feature"
)

var (
	issueParameter issueCmdParameter
	issueCmd       = &cobra.Command{
		Use:   "issue",
		Short: "Generate an issue template for bug reports or feature requests",
		Long: `The 'issue' command allows you to generate an issue template 
for bug reports or feature requests in either YAML or Markdown format. 
You can customize the issue name, description, assignees, output file, 
and other parameters.`,
		Example: `  # Generate a bug report issue template in Markdown format
  docwiz issue --kind bug --format md --output ISSUE.md

  # Generate a feature request issue template in YAML format
  docwiz issue --kind feature --format yaml --output ISSUE.yaml

  # Generate an issue template with a custom name and description
  docwiz issue --name "Crash on startup" --description "The app crashes immediately on launch" --format md`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(issueParameter.issueName) == 0 {
				if issueParameter.kind == issueKindBug {
					issueParameter.issueName = "Bug report"
				} else {
					issueParameter.issueName = "Feature request"
				}
			}
			if len(issueParameter.issueDescription) == 0 {
				if issueParameter.kind == issueKindBug {
					issueParameter.issueDescription = "Create a report to help us improve"
				} else {
					issueParameter.issueDescription = "Suggest an idea for this project"
				}
			}
			if issueParameter.format == issueForamtYAML {
				if len(issueParameter.output) == 0 {
					issueParameter.output = "ISSUE.yaml"
				}
				repo, err := git.New(issueParameter.repoPath)
				if err != nil {
					panic(err)
				}
				title := "[Bug]: "
				labels := []string{"bug", "question"}
				assignees := []string{}
				tags := []string{}
				if issueParameter.kind == issueKindFeature {
					title = "[Feature]: "
					labels = []string{"enhancement"}
				}
				for _, tag := range repo.GetTags() {
					tags = append(tags, tag.Name)
				}
				var content BodyAttributes
				if issueParameter.kind == issueKindFeature {
					content = BodyAttributes{
						Label:       "What feature would you like to see?",
						Description: "Describe the feature you'd like to be added and why it's useful.",
						Placeholder: "Explain the feature idea...",
						Value:       "A clear and concise description of the feature request.",
					}
				} else {
					content = BodyAttributes{
						Label:       "What happened?",
						Description: "Also tell us, what did you expect to happen?",
						Placeholder: "Tell us what you see!",
						Value:       "A clear and concise description of what the bug is.",
					}
				}
				issueTmpl := BugReport{
					Name:        issueParameter.issueName,
					Description: issueParameter.issueDescription,
					Title:       title,
					Labels:      labels,
					Assignees:   assignees,
					Body: []BodyPart{
						{
							Type:       "textarea",
							Attributes: content,
							Validations: &Validations{
								Required: true,
							},
						},
						{
							Type: "dropdown",
							Attributes: BodyAttributes{
								Label:       "Version",
								Description: "What version of our software are you running?",
								Options:     tags,
							},
							Validations: &Validations{
								Required: false,
							},
						},
						{
							Type: "dropdown",
							Attributes: BodyAttributes{
								Label:    "What platform are you seeing the problem on?",
								Multiple: true,
								Options:  []string{"Linux", "Darwin", "Windows"},
							},
						},
					},
				}

				data, err := yaml.Marshal(&issueTmpl)
				if err != nil {
					panic(err)
				}

				output, err := io.NewSafeFile(issueParameter.output)
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

				output.Write(data)
			} else {
				if len(issueParameter.output) == 0 {
					issueParameter.output = "ISSUE.md"
				}
				issuePath := filepath.Join(os.TemplatePath, "ISSUE")
				tpl := filepath.Join(issuePath, fmt.Sprintf("%s.%s.tpl", issueParameter.theme, issueParameter.kind))

				output, err := io.NewSafeFile(issueParameter.output)
				if err != nil {
					panic(err)
				}
				defer output.Close()

				defer func() {
					if err := recover(); err != nil {
						output.Rollback()
					}
				}()

				tmpl, err := template.New(filepath.Base(tpl)).Funcs(DocwizFuncMap(issuePath)).ParseFiles(tpl)
				if err != nil {
					panic(err)
				}

				err = tmpl.Execute(output, map[string]any{
					"IssueName":        issueParameter.issueName,
					"IssueDescription": issueParameter.issueDescription,
					"IssueAssigness":   issueParameter.issueAssigness,
				})

				if err != nil {
					panic(err)
				}
			}
		},
	}
)

func init() {
	docwizCmd.AddCommand(issueCmd)
	issueCmd.PersistentFlags().StringVarP(&issueParameter.issueName, "name", "n", "", "Name of the issue (default: 'Bug report' for bugs, 'Feature request' for features)")
	issueCmd.PersistentFlags().StringVarP(&issueParameter.issueDescription, "description", "d", "", "Description of the issue")
	issueCmd.PersistentFlags().StringVarP(&issueParameter.issueAssigness, "assigness", "a", "", "List of assignees for the issue (comma-separated)")
	issueCmd.PersistentFlags().StringVarP(&issueParameter.output, "output", "o", "", "Output file name for the generated issue template (default: ISSUE.md or ISSUE.yaml)")
	issueCmd.PersistentFlags().StringVarP(&issueParameter.theme, "theme", "t", "default", "Theme for issue template rendering")
	issueCmd.PersistentFlags().StringVarP(&issueParameter.kind, "kind", "k", issueKindBug, "Type of issue to generate (bug or feature)")
	issueCmd.PersistentFlags().StringVarP(&issueParameter.format, "format", "f", issueFormatMarkdown, "Format of the issue template (md or yaml)")
	issueCmd.PersistentFlags().StringVarP(&issueParameter.repoPath, "repo", "r", ".", "Path to the target Git repository (default: current directory)")
}

type BugReport struct {
	Name        string     `yaml:"name"`
	Description string     `yaml:"description"`
	Title       string     `yaml:"title"`
	Labels      []string   `yaml:"labels"`
	Projects    []string   `yaml:"projects"`
	Assignees   []string   `yaml:"assignees"`
	Body        []BodyPart `yaml:"body"`
}

type BodyPart struct {
	Type        string         `yaml:"type"`
	Attributes  BodyAttributes `yaml:"attributes"`
	Validations *Validations   `yaml:"validations,omitempty"`
}

type BodyAttributes struct {
	Value       string   `yaml:"value,omitempty"`
	Label       string   `yaml:"label,omitempty"`
	Description string   `yaml:"description,omitempty"`
	Placeholder string   `yaml:"placeholder,omitempty"`
	Options     []string `yaml:"options,omitempty"`
	Multiple    bool     `yaml:"multiple,omitempty"`
	Render      string   `yaml:"render,omitempty"`
}

type Validations struct {
	Required bool `yaml:"required"`
}
