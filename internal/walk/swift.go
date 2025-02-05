// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type SwiftWalker struct {
	BaseWalker
}

func (*SwiftWalker) SubscribeExt() []string {
	return []string{".swift"}
}

func (*SwiftWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Swift", upgradeBadge("Swift", badge.ShieldSwift))
	return nil
}
