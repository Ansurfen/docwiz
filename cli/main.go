// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"docwiz/cli/cmd"

	"github.com/caarlos0/log"
)

func main() {
	log.SetLevel(log.DebugLevel)
	cmd.Execute()
}
