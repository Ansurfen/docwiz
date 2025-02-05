// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package tui

import (
	"fmt"
	"strings"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type IndexFile struct {
	Questions []*Question `json:"questions"`
}

type Question struct {
	textinput.Model
	Suggest  string `json:"suggest"`
	Required bool   `json:"required"`
	Default  string `json:"default"`
	Binding  string `json:"binding"`
}

func (t *Question) init() {
	t.Model = textinput.New()
	if len(t.Default) != 0 {
		t.Model.Placeholder = t.Default
	}
}

func (q *Question) SetDefault() {
	q.Model.SetValue(q.Default)
}

func (q *Question) Focus() tea.Cmd {
	q.Model.Prompt = focusing.String() + " " + textfocusedStyle.SetString(q.Suggest).String() + " "
	q.Model.PromptStyle = focusedStyle
	q.Model.TextStyle = focusedStyle
	return q.Model.Focus()
}

var (
	textfocusedStyle    = lipgloss.NewStyle().Bold(true)
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Bold(true)
	textStyle           = lipgloss.NewStyle().Foreground(lipgloss.Color("#299aba"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
	redStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
)

var (
	pengding  = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575")).SetString("?")
	focusing  = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575")).Bold(true).SetString(">")
	finishing = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575")).SetString("✔")
)

var DefaultQuestion = []*Question{
	{Suggest: "Project name", Required: true, Default: "example project", Binding: "Name"},
	{Suggest: "Project version", Required: true, Default: "1.0.0", Binding: "Version"},
	{Suggest: "Project license", Required: false, Default: "MIT", Binding: "License"},
	{Suggest: "Project homepage", Required: true, Default: "https://github.com/yourname/repository", Binding: "HomePage"},
	{Suggest: "README.md content (use ';' to divide different sections)", Required: false, Default: "📦 Install;🚀 Usage;✅ Test", Binding: "Content"},
}

type ReadmeModel struct {
	inputs     []*Question
	focusIndex int
	cursorMode cursor.Mode
	isContinue bool
	showError  bool
	form       *huh.Form
	v          map[string]any
}

func NewReadmeModel(question ...*Question) *ReadmeModel {
	m := &ReadmeModel{
		inputs:     question,
		isContinue: true,
		showError:  true,
		v:          map[string]any{},
	}

	m.form = huh.NewForm(huh.NewGroup(huh.NewConfirm().
		Title("Continue to improve the README.md?").
		Description("Select continue to complete the content of README.md").
		Affirmative("Continue").
		Negative("Done").
		Value(&m.isContinue)))

	for i, t := range m.inputs {
		t.init()
		// t.PromptStyle = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("#04B575"))
		t.Cursor.Style = cursorStyle

		switch i {
		case 0:
			t.Focus()
		default:
			t.Prompt = pengding.String() + " " + t.Suggest + " "
			t.TextStyle = textStyle
		}

	}
	return m
}

func (m *ReadmeModel) current() *Question {
	return m.inputs[m.focusIndex]
}

func (m *ReadmeModel) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.form.Init())
}

func (m *ReadmeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		// Change cursor mode
		case "ctrl+r":
			m.cursorMode++
			if m.cursorMode > cursor.CursorHide {
				m.cursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(m.inputs))
			for i := range m.inputs {
				cmds[i] = m.inputs[i].Cursor.SetMode(m.cursorMode)
			}
			return m, tea.Batch(cmds...)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				ok := true
				for _, t := range m.inputs {
					if t.Required && len(t.Value()) == 0 {
						ok = false
					}
				}
				if !ok {
					m.showError = true
				} else {
					return m, tea.Quit
				}
			}

			if s == "tab" && m.focusIndex != len(m.inputs) {
				m.current().SetDefault()
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				if len(m.inputs[i].Value()) > 0 {
					m.inputs[i].Prompt = finishing.String() + " " + m.inputs[i].Suggest + " "
				} else {
					m.inputs[i].Prompt = pengding.String() + " " + m.inputs[i].Suggest + " "
				}
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = textStyle

			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	var cmds []tea.Cmd
	cmds = append(cmds, m.updateInputs(msg))

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m *ReadmeModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i].Model, cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m *ReadmeModel) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	if m.showError {
		b.WriteString(redStyle.SetString("\n[ERROR] You have an unfilled field\n").String())
		m.showError = false
	}

	b.WriteString("\n\n" + m.form.View())

	b.WriteString(helpStyle.Render("cursor mode is "))
	b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String()))
	b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

	return b.String()
}

func (m *ReadmeModel) Run() error {
	_, err := tea.NewProgram(m).Run()
	if err != nil {
		return err
	}

	for _, q := range m.inputs {
		m.v[q.Binding] = q.Value()
	}

	if m.isContinue {
		rm := newREADMEContentModel(strings.Split(m.v["Content"].(string), ";"))
		if _, err := tea.NewProgram(rm).Run(); err != nil {
			return err
		}
		content := map[string]string{}
		for i, item := range rm.items {
			content[rm.titles[i]] = item.Value()
		}
		m.v["Content"] = content
	}
	return err
}

func (m *ReadmeModel) Value() map[string]any {
	return m.v
}

type READMEContentModel struct {
	titles    []string
	items     []textarea.Model
	paginator paginator.Model
}

const TextareaPlaceholder = "Describe it here..."

func newREADMEContentModel(titles []string) READMEContentModel {
	var items []textarea.Model

	for range titles {
		ta := textarea.New()
		ta.Focus()
		ta.Placeholder = TextareaPlaceholder
		items = append(items, ta)
	}

	p := paginator.New()
	p.KeyMap = paginator.KeyMap{
		PrevPage: key.NewBinding(key.WithKeys("ctrl+left")),
		NextPage: key.NewBinding(key.WithKeys("ctrl+right")),
	}
	p.Type = paginator.Dots
	p.PerPage = 1
	p.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•")
	p.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")
	p.SetTotalPages(len(items))

	return READMEContentModel{
		paginator: p,
		items:     items,
		titles:    titles,
	}
}

func (m READMEContentModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m READMEContentModel) index() int {
	start, _ := m.paginator.GetSliceBounds(len(m.items))
	return start
}

func (m READMEContentModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	item := m.items[m.index()]

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlU:
			return m, tea.Quit
		case tea.KeyEsc:
			if item.Focused() {
				item.Blur()
			}
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			if !item.Focused() {
				cmd = item.Focus()
				cmds = append(cmds, cmd)
			}
		}
	}

	m.items[m.index()], cmd = item.Update(msg)
	cmds = append(cmds, cmd)

	m.paginator, cmd = m.paginator.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m READMEContentModel) View() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("\n  %s (use ctrl+u to submit)\n\n", m.titles[m.index()]))

	textarea := m.items[m.index()]
	b.WriteString(textarea.View())

	b.WriteString("\n  " + m.paginator.View())
	b.WriteString("\n\n  ctrl + ←/→ page\n")
	return b.String()
}
