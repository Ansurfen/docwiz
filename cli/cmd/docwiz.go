// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const COPYRIGHT = `
---

_This Markdown was generated with ❤️ by [docwiz](https://github.com/ansurfen/docwiz)_`

var docwizCmd = &cobra.Command{
	Use:   "docwiz",
	Short: "🚀 A tool for generating project documentation and related files",
	Long: `docwiz is a versatile command-line tool that helps generate various types of project documentation 
like README, LICENSE, ROADMAP, CONTRIBUTORS, and more. It leverages templates 
and user inputs to create customized and professional documentation files.`,
}

func Execute() {
	err := docwizCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
