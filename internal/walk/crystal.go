// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type CrystalWalker struct {
	BaseWalker
}

func (*CrystalWalker) SubscribeExt() []string {
	return []string{".cr"}
}

func (*CrystalWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Crystal", upgradeBadge("Crystal", badge.ShieldCrystal))
	return nil
}
