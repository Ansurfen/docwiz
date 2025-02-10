// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package haskellwalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/walk"
)

type Walker struct {
	walk.BaseWalker
}

func (*Walker) SubscribeExt() []string {
	return []string{".hs", ".lhs", ".hsc"}
}

func (*Walker) ParseExt(fullpath string, ext string, ctx *walk.Context) error {
	ctx.Set("Haskell", walk.UpgradeBadge("Haskell", badge.ShieldHaskell))
	return nil
}
