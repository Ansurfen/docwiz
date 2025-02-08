// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type TextFrame struct {
	content string
}

func NewTextFrame(content string) *TextFrame {
	return &TextFrame{content: content}
}

func (f *TextFrame) Run() error {
	var sb strings.Builder
	fmt.Fprint(&sb, f.content)

	fmt.Println(
		lipgloss.NewStyle().
			Width(40).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1, 2).
			Render(sb.String()),
	)
	return nil
}
