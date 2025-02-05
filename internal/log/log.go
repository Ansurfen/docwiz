// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package log

import (
	. "docwiz/internal/os"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var (
	errorStyle = lipgloss.NewStyle().
			SetString("ERROR").
			Padding(0, 1, 0, 1).
			Background(lipgloss.Color("#FF5555")).
			Foreground(lipgloss.Color("#000000"))

	warnStyle = lipgloss.NewStyle().
			SetString("WARN").
			Padding(0, 1, 0, 1).
			Background(lipgloss.Color("#FFCC00")).
			Foreground(lipgloss.Color("#000000"))

	infoStyle = lipgloss.NewStyle().
			SetString("INFO").
			Padding(0, 1, 0, 1).
			Background(lipgloss.Color("#0096FF")).
			Foreground(lipgloss.Color("#FFFFFF"))

	debugStyle = lipgloss.NewStyle().
			SetString("DEBUG").
			Padding(0, 1, 0, 1).
			Background(lipgloss.Color("#33CC33")).
			Foreground(lipgloss.Color("#000000"))

	fatalStyle = lipgloss.NewStyle().
			SetString("FATA").
			Padding(0, 1, 0, 1).
			Background(lipgloss.Color("#990000")).
			Foreground(lipgloss.Color("#FFFFFF"))
)

func Errorf(fmt string, a ...any) {
	printf(errorStyle, fmt, a...)
}

func Infof(fmt string, a ...any) {
	printf(infoStyle, fmt, a...)
}

func Warnf(fmt string, a ...any) {
	printf(warnStyle, fmt, a...)
}

func Debugf(fmt string, a ...any) {
	printf(debugStyle, fmt, a...)
}

func Fataf(fmt string, a ...any) {
	printf(fatalStyle, fmt, a...)
	os.Exit(1)
}

func printf(style lipgloss.Style, format string, a ...any) {
	fmt.Fprintf(logWriter, style.String()+" "+format, a...)
}

var logWriter io.Writer = os.Stdout

func init() {
	logFileName := filepath.Join(LogPath, time.Now().Format("2006-01-02_15-04-05")+".log")

	logFile, err := os.Create(logFileName)
	if err != nil {
		Panic("fail to create log file, err: %s", err.Error())
		return
	}
	logWriter = logFile
}
