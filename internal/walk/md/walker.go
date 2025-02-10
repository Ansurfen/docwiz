// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package mdwalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/walk"
)

type Walker struct {
	walk.BaseWalker
}

func (*Walker) SubscribeExt() []string {
	return []string{".md", ".markdown", ".mdown"}
}

func (*Walker) ParseExt(fullpath string, ext string, ctx *walk.Context) error {
	ctx.Set("Markdown", walk.UpgradeBadge("Markdown", badge.ShieldMarkdown))
	return nil
}
