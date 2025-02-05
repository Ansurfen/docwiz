// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type NimWalker struct {
	BaseWalker
}

func (*NimWalker) SubscribeExt() []string {
	return []string{".nim", ".nims"}
}

func (*NimWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Nim", upgradeBadge("Nim", badge.ShieldNim))
	return nil
}
