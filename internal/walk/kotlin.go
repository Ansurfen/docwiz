// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type KotlinWalker struct {
	BaseWalker
}

func (*KotlinWalker) SubscribeExt() []string {
	return []string{".kt", ".kts"}
}

func (*KotlinWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Kotlin", upgradeBadge("Kotlin", badge.ShieldKotlin))
	return nil
}
