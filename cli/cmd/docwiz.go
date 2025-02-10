// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"docwiz/internal/tui"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const defaultLanguage = "en_us"

// Version will be set by the build process using -ldflags.
// Do not modify this variable directly. It is populated at build time with the desired version string.
var Version = ""

var (
	COPYRIGHT = []byte("\n---\n\n_This Markdown was generated with ❤️ by [docwiz](https://github.com/ansurfen/docwiz)_")
	docwizCmd = &cobra.Command{
		Use:   "docwiz",
		Short: "🚀 A tool for generating project documentation and related files",
		Long: `docwiz is a versatile command-line tool that helps generate various types of project documentation 
like README, LICENSE, ROADMAP, CONTRIBUTORS, and more. It leverages templates 
and user inputs to create customized and professional documentation files.`,
	}
)

func Execute() {
	err := docwizCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// baseParameter contains shared parameters across multiple commands.
type baseParameter struct {
	// output specifies the path and filename of the generated output file
	output string

	// theme defines the theme for the document or UI
	theme string

	// language sets the language to be used (e.g., "en_us", "zh_cn")
	//
	// default: en_us
	language string

	// disableCopyright determines whether to disable copyright information
	//
	// default: false
	disableCopyright bool

	verbose bool
}

type generator struct {
	output string
	action func()
}

func (g *generator) run() error {
	err := tui.NewSpinner(g.action, fmt.Sprintf("Generating %s...", g.output)).Run()
	if err != nil {
		return err
	}
	return tui.NewTextFrame(fmt.Sprintf("%s was successfully generated.\n\nThanks for using docwiz!", g.output)).Run()
}

// todo
// func assert(err any, verbose bool) {
// 	if err != nil {
// 		if verbose {
// 			panic(err)
// 		}
// 		log.Fata(err)
// 	}
// }
