// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gopkg.in/yaml.v3"
)

type TmplQuestion struct {
	Type        string `yaml:"type"`
	Prompt      string `yaml:"prompt"`
	Placeholder string `yaml:"placeholder"`
	Binding     string `yaml:"binding"`
}

type TmplModel struct {
	Question      []TmplQuestion `yaml:"question"`
	BeforeExecute string         `yaml:"beforeExecute"`
	AfterExecute  string         `yaml:"afterExecute"`

	data map[string]string
}

func NewTmpl(tmpl string) (*TmplModel, error) {
	m := &TmplModel{data: make(map[string]string)}
	data, err := os.ReadFile(tmpl)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (mod *TmplModel) Run() error {
	for _, q := range mod.Question {
		switch q.Type {
		case "input":
			m := newTmplInputModel(q.Prompt, q.Placeholder)
			if _, err := tea.NewProgram(m).Run(); err != nil {
				panic(err)
			}
			mod.data[q.Binding] = m.textInput.Value()
			// TODO
		case "confirm":
			m := newTmplInputModel(q.Prompt, q.Placeholder)
			if _, err := tea.NewProgram(m).Run(); err != nil {
				panic(err)
			}
			mod.data[q.Binding] = m.textInput.Value()
		}
	}
	return nil
}

func (m *TmplModel) Value() map[string]string {
	return m.data
}

type tmplInputModel struct {
	textInput    textinput.Model
	prompt       string
	defaultValue string
	done         bool
}

func newTmplInputModel(prompt, value string) *tmplInputModel {
	ti := textinput.New()
	ti.Prompt = pengding.String() + " " + prompt + ": "
	ti.PromptStyle = lipgloss.NewStyle().Bold(true)
	ti.Placeholder = value
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 40

	return &tmplInputModel{textInput: ti, prompt: prompt, defaultValue: value}
}

func (m *tmplInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *tmplInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			m.textInput.SetValue(m.defaultValue)
		case "enter":
			m.done = true
			return m, tea.Quit
		case "ctrl+c", "esc":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m *tmplInputModel) View() string {
	if m.done {
		return fmt.Sprintf(finishing.String()+" %s: %s\n", m.prompt, m.textInput.Value())
	}
	return fmt.Sprintf("%s\n%s", m.textInput.View(), "(Press Tab to auto-complete, Enter to submit)")
}
