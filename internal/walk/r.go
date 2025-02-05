// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type RWaler struct {
	BaseWalker
}

func (*RWaler) SubscribeExt() []string {
	return []string{".R", ".r", ".Rmd", ".Rprofile"}
}

func (*RWaler) ParseExt(fullpath string, file string, ctx *Context) error {
	ctx.Set("R", upgradeBadge("R", badge.ShieldR))
	return nil
}
