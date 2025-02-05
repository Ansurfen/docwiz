// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type TypeScriptWalker struct {
	JavaScriptWalker
}

func (*TypeScriptWalker) SubscribeExt() []string {
	return []string{".ts", ".tsx"}
}

func (*TypeScriptWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("TypeScript", upgradeBadge("TypeScript", badge.ShieldTypeScript))
	return nil
}
