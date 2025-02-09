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
	"runtime"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var (
	errorStyle = lipgloss.NewStyle().
			SetString("ERROR").
			Padding(0, 1, 0, 1).
			Background(lipgloss.Color("#d251a6")).
			Foreground(lipgloss.Color("#000000"))

	warnStyle = lipgloss.NewStyle().
			SetString("WARN").
			Padding(0, 1, 0, 1).
			Background(lipgloss.Color("#ff7b52")).
			Foreground(lipgloss.Color("#000000"))

	infoStyle = lipgloss.NewStyle().
			SetString("INFO").
			Padding(0, 1, 0, 1).
			Background(lipgloss.Color("#00a1e5")).
			Foreground(lipgloss.Color("#FFFFFF"))

	debugStyle = lipgloss.NewStyle().
			SetString("DEBUG").
			Padding(0, 1, 0, 1).
			Background(lipgloss.Color("#04b9ae")).
			Foreground(lipgloss.Color("#000000"))

	fatalStyle = lipgloss.NewStyle().
			SetString("FATA").
			Padding(0, 1, 0, 1).
			Background(lipgloss.Color("#c53f3f")).
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

func Fata(err any) {
	fmt.Println(fatalStyle, err)
	print(err, "\n", getStackTrace(1))
	os.Exit(1)
}

func print(a ...any) {
	fmt.Fprint(logWriter, a...)
}

func getStackTrace(skip int) string {
	const maxDepth = 32
	var pcs [maxDepth]uintptr

	n := runtime.Callers(skip+2, pcs[:])
	frames := runtime.CallersFrames(pcs[:n])

	stacks := []string{}

	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		stacks = append(stacks, fmt.Sprintf("%s\n  %s:%d\n", frame.Function, frame.File, frame.Line))
	}
	return strings.Join(stacks, "")
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
