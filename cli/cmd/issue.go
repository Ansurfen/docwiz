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
	"gopkg.in/yaml.v3"
)

// issueCmdParameter stores the parameters for the "issue" command.
type issueCmdParameter struct {
	// issueName defines the name of the issue (e.g., "Bug report", "Feature request").
	// If not provided, the name defaults based on the issue kind.
	issueName string

	// issueDescription contains the description of the issue, providing details on the problem or feature.
	// If not provided, a default description is set based on the issue kind.
	issueDescription string

	// issueAssigness is a comma-separated list of assignees for the issue.
	// It specifies who is responsible for addressing the issue.
	issueAssigness string

	// output defines the output file where the generated issue template will be saved.
	// Defaults to "ISSUE.md" for Markdown or "ISSUE.yaml" for YAML, depending on the format.
	output string

	// theme specifies the theme for the issue template rendering, which influences the layout and structure of the generated file.
	// The default theme is "default".
	theme string

	// kind defines the type of the issue (e.g., "bug" or "feature").
	// It determines the structure and labels of the issue template.
	kind string

	// format specifies the format of the generated issue template.
	// It can be either "md" for Markdown or "yaml" for YAML.
	format string

	// repoPath specifies the path to the Git repository, from which information like tags will be gathered.
	// The default value is the current directory ("./").
	repoPath string
}

// Constants for the issue format and type.
const (
	// issueFormatMarkdown represents the Markdown format for the issue template.
	issueFormatMarkdown = "md"

	// issueForamtYAML represents the YAML format for the issue template.
	issueForamtYAML = "yaml"

	// issueKindBug represents a bug report issue.
	issueKindBug = "bug"

	// issueKindFeature represents a feature request issue.
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

			var (
				output *io.SafeFile
				err    error
			)
			log.Infof("creating %s", issueParameter.output)
			output, err = io.NewSafeFile(issueParameter.output)
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

			if issueParameter.format == issueForamtYAML {
				if len(issueParameter.output) == 0 {
					issueParameter.output = "ISSUE.yaml"
				}

				log.WithField("path", issueParameter.repoPath).Info("parsing .git directory")
				repo, err := git.New(issueParameter.repoPath)
				if err != nil {
					log.WithError(err).Fatal("fail to read git repository")
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
				issueTmpl := IssueTemplate{
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
					log.WithError(err).Fatal("marshaling template")
				}

				output.Write(data)
			} else {
				if len(issueParameter.output) == 0 {
					issueParameter.output = "ISSUE.md"
				}
				issuePath := filepath.Join(os.TemplatePath, "ISSUE")
				tpl := filepath.Join(issuePath, fmt.Sprintf("%s.%s.tpl", issueParameter.theme, issueParameter.kind))

				log.WithField("target", tpl).Info("loading template")
				tmpl, err := template.Default(tpl)
				if err != nil {
					log.WithError(err).Fatal("fail to load template")
				}

				log.Info("executing template")
				log.IncreasePadding()
				log.WithField("IssueName", issueParameter.issueName).
					WithField("IssueDescription", issueParameter.issueDescription).
					WithField("IssueAssigness", issueParameter.issueAssigness).Info("parameters")
				log.DecreasePadding()
				err = tmpl.Execute(output, map[string]any{
					"IssueName":        issueParameter.issueName,
					"IssueDescription": issueParameter.issueDescription,
					"IssueAssigness":   issueParameter.issueAssigness,
				})

				if err != nil {
					log.WithError(err).Fatal("fail to execute template")
				}
			}

			log.Info("thanks for using docwiz!")
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

type IssueTemplate struct {
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
