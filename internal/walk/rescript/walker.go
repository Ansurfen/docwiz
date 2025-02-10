// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package rescriptwalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/walk"
)

type Walker struct {
	walk.BaseWalker
}

func (*Walker) SubscribeExt() []string {
	return []string{".res", ".resi", ".re"}
}

func (*Walker) ParseExt(fullpath string, file string, ctx *walk.Context) error {
	ctx.Set("ReScript", walk.UpgradeBadge("ReScript", badge.ShieldReScript))
	return nil
}
