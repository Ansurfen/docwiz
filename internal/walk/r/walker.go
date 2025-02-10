// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package rwalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/walk"
)

type Walker struct {
	walk.BaseWalker
}

func (*Walker) SubscribeExt() []string {
	return []string{".R", ".r", ".Rmd", ".Rprofile"}
}

func (*Walker) ParseExt(fullpath string, file string, ctx *walk.Context) error {
	ctx.Set("R", walk.UpgradeBadge("R", badge.ShieldR))
	return nil
}
