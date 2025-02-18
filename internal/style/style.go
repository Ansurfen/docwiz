// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package style

import "github.com/charmbracelet/lipgloss"

var (
	Bold    = lipgloss.NewStyle().Bold(true).Render
	Keyword = lipgloss.NewStyle().
		Padding(0, 1).
		Foreground(lipgloss.AdaptiveColor{Light: "#FF4672", Dark: "#ED567A"}).
		Background(lipgloss.AdaptiveColor{Light: "#DDDADA", Dark: "#242424"}).
		Render
)
