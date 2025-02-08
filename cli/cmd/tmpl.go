// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"docwiz/internal/os"
	"docwiz/internal/tui"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

type tmplCmdParameter struct {
}

var (
	tmplParameter tmplCmdParameter
	tmplCmd       = &cobra.Command{
		Use: "tmpl",
		Run: func(cmd *cobra.Command, args []string) {
			panic("WIP")
			m, err := tui.NewTmpl(filepath.Join(os.TemplatePath, "TMPL", "vue", "index.yaml"))
			if err != nil {
				panic(err)
			}
			err = m.Run()
			if err != nil {
				panic(err)
			}
			fmt.Println(m.Value())
		},
	}
)

func init() {
	docwizCmd.AddCommand(tmplCmd)
}
