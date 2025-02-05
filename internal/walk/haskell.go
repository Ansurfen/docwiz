// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type HaskellWalker struct {
	BaseWalker
}

func (*HaskellWalker) SubscribeExt() []string {
	return []string{".hs", ".lhs", ".hsc"}
}

func (*HaskellWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Haskell", upgradeBadge("Haskell", badge.ShieldHaskell))
	return nil
}
