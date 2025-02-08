// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package tui

import (
	"github.com/charmbracelet/huh/spinner"
)

func NewSpinner(action func(), title string) *spinner.Spinner {
	return spinner.New().Title(title).Action(action)
}
