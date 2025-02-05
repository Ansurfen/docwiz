// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type GDScriptWalker struct {
	BaseWalker
}

func (*GDScriptWalker) SubscribeExt() []string {
	return []string{".gd"}
}

func (*GDScriptWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("GDScript", upgradeBadge("GDScript", badge.ShieldGDScript))
	return nil
}
