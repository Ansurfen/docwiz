// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type ElmWalker struct {
	BaseWalker
}

func (*ElmWalker) SubscribeExt() []string {
	return []string{".elm", ".elm.js"}
}

func (*ElmWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Elm", upgradeBadge("Elm", badge.ShieldElm))
	return nil
}
