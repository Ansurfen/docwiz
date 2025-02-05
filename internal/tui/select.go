// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package tui

import (
	"strings"

	"github.com/charmbracelet/huh"
)

type SelectModel struct {
	form       *huh.Form
	in         string
	license    string
	candicates []string
}

type SelectModelConfigure struct {
	Title       string
	Description string
	Placeholder string
	SelectTitle string
	Candicates  []string
}

func NewSelectModel(cfg SelectModelConfigure) *SelectModel {
	m := &SelectModel{
		candicates: cfg.Candicates,
	}
	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Value(&m.in).
				Title(cfg.Title).
				Description(cfg.Description).
				Placeholder(cfg.Placeholder),

			huh.NewSelect[string]().Title(cfg.SelectTitle).Value(&m.license).
				OptionsFunc(func() []huh.Option[string] {
					var options []huh.Option[string]
					for _, key := range m.candicates {
						if strings.Contains(key, m.in) {
							options = append(options, huh.NewOption(key, key))
						}
					}
					return options
				}, &m.in),
		),
	)

	return m
}

func (m *SelectModel) Run() error {
	return m.form.Run()
}

func (m *SelectModel) Value() string {
	return m.license
}
