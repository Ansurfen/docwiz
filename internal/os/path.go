// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package os

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	// RootPath to where the executable is stored
	RootPath     string
	TemplatePath string
	LogPath      string
)

func Panic(format string, a ...any) {
	if len(LogPath) != 0 {

	} else {
		fmt.Println("\x1b[48;5;167m\x1b[38;5;0m ERROR \x1b[0m", fmt.Sprintf(format, a...))
	}
}

func init() {
	var err error

	RootPath, err = os.Executable()
	if err != nil {
		Panic("fail to get root path, err: %s", err.Error())
	}

	TemplatePath = filepath.Join(RootPath, "../template")

	LogPath = filepath.Join(RootPath, "../log")
	os.Mkdir(LogPath, 0755)
}
