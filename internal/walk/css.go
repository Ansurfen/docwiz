// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type CSSWalker struct {
	BaseWalker
}

func (*CSSWalker) SubscribeExt() []string {
	return []string{".css", ".scss", ".sass", ".less", ".styl"}
}

func (*CSSWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("CSS3", upgradeBadge("CSS3", badge.ShieldCSS3))
	return nil
}
