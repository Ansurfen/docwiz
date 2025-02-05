// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type GroovyWalker struct {
	BaseWalker
}

func (*GroovyWalker) SubscribeExt() []string {
	return []string{".groovy", ".gvy", ".gsh"}
}

func (*GroovyWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Groovy", upgradeBadge("Groovy", badge.ShieldGroovy))
	return nil
}
