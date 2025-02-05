// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type HTMLWalker struct {
	BaseWalker
}

func (*HTMLWalker) SubscribeExt() []string {
	return []string{".html", ".htm", ".xhtml"}
}

func (*HTMLWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("HTML", upgradeBadge("HTML", badge.ShieldHTML5))
	return nil
}
