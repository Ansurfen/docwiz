// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package badge

type Badge interface {
	Name() string
	SetVersion(v string)

	URL() string
	Markdown() string
	RSt() string
	AsciiDoc() string
	HTML() string
}

type SortableBadge struct {
	Badge
	Tag string
}
