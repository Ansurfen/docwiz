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

// versionCmdParameter stores parameters for the "version" command.
type versionCmdParameter struct {
	// pure indicates whether to print just the version number without any extra information (logo, description, etc.)
	pure bool
}

var (
	versionParameter versionCmdParameter
	versionCmd       = &cobra.Command{
		Use:   "version",
		Short: "Displays the current version of the CLI tool",
		Long: `The 'version' command outputs the version of the CLI tool,
	including its logo and description. The 'pure' flag can be used to display just the version number.`,
		Example: "docwiz version\n  docwiz version -p",
		Run: func(cmd *cobra.Command, args []string) {
			if versionParameter.pure {
				fmt.Print(Version)
			} else {
				fmt.Printf("%s\n🚀 CLI that generates beautiful git files\nVersion: %s\nHomePage: https://github.com/ansurfen/docwiz", logo, Version)
			}
		},
	}
)

func init() {
	docwizCmd.AddCommand(versionCmd)
	versionCmd.PersistentFlags().BoolVarP(&versionParameter.pure, "pure", "p", false, "Only display the version number, without any additional info")
}
