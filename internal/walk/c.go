// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type CWalker struct {
	BaseWalker
}

func (*CWalker) SubscribeExt() []string {
	return []string{".c", ".h"}
}

func (*CWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("C", upgradeBadge("C", badge.ShieldC))
	return nil
}
