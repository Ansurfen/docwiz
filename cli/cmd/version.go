// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const logo = `
██████╗  ██████╗  ██████╗██╗    ██╗██╗███████╗
██╔══██╗██╔═══██╗██╔════╝██║    ██║██║╚══███╔╝
██║  ██║██║   ██║██║     ██║ █╗ ██║██║  ███╔╝ 
██║  ██║██║   ██║██║     ██║███╗██║██║ ███╔╝  
██████╔╝╚██████╔╝╚██████╗╚███╔███╔╝██║███████╗
╚═════╝  ╚═════╝  ╚═════╝ ╚══╝╚══╝ ╚═╝╚══════╝`

type versionCmdParameter struct {
	pure bool
}

var (
	versionParameter versionCmdParameter
	versionCmd       = &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			if versionParameter.pure {
				fmt.Print(version)
			} else {
				fmt.Printf("%s\n🚀 CLI that generates beautiful git files\nVersion: %s\nHomePage: https://github.com/ansurfen/docwiz", logo, version)
			}
		},
	}
)

func init() {
	docwizCmd.AddCommand(versionCmd)
	versionCmd.PersistentFlags().BoolVarP(&versionParameter.pure, "pure", "p", false, "")
}
