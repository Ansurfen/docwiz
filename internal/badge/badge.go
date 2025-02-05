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

type BadgeUnion struct {
	BadgenBadge TypedBadge
	ShieldBadge TypedBadge
}

func (*BadgeUnion) Name() string { return "" }

func (*BadgeUnion) SetVersion(string) {}

func (*BadgeUnion) URL() string { return "" }

func (*BadgeUnion) Markdown() string { return "" }

func (*BadgeUnion) RSt() string { return "" }

func (*BadgeUnion) AsciiDoc() string { return "" }

func (*BadgeUnion) HTML() string { return "" }

type TypedBadge struct {
	Type string
	Badge
}
