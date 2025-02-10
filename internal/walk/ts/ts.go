// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package tswalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/walk"
	jswalk "docwiz/internal/walk/js"
)

type Walker struct {
	jswalk.Walker
}

func (*Walker) SubscribeExt() []string {
	return []string{".ts", ".tsx"}
}

func (*Walker) ParseExt(fullpath string, ext string, ctx *walk.Context) error {
	ctx.Set("TypeScript", walk.UpgradeBadge("TypeScript", badge.ShieldTypeScript))
	return nil
}
