// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type NixWalker struct {
	BaseWalker
}

func (*NixWalker) SubscribeExt() []string {
	return []string{".nix"}
}

func (*NixWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Nix", upgradeBadge("Nix", badge.ShieldNix))
	return nil
}
