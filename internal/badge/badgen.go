// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package badge

const BadgenBaseURL = "https://badgen.net/"

type BadgenBadge struct {
	raw string
}

func (b *BadgenBadge) Name() string {
	return ""
}

func (*BadgenBadge) SetVersion(string) {}

func (b *BadgenBadge) URL() string {
	if len(b.raw) != 0 {
		return BadgenBaseURL + b.raw
	}
	return ""
}

func (*BadgenBadge) Markdown() string { return "" }

func (*BadgenBadge) RSt() string { return "" }

func (*BadgenBadge) AsciiDoc() string { return "" }

func (*BadgenBadge) HTML() string { return "" }

var (
	BadgenAtom    = &BadgenBadge{raw: "/badge/icon/atom?icon=atom&label"}
	BadgenAwesome = &BadgenBadge{raw: "/badge/icon/awesome?icon=awesome&label"}
	BadgenDocker  = &BadgenBadge{raw: "/badge/icon/docker?icon=docker&label"}
	BadgenCodecov = &BadgenBadge{raw: "/badge/icon/codecov?icon=codecov&label"}
)
