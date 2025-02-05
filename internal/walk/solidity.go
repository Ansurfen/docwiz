// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type SolidityWalker struct {
	BaseWalker
}

func (*SolidityWalker) SubscribeExt() []string {
	return []string{".sol"}
}

func (*SolidityWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Solidity", upgradeBadge("Solidity", badge.ShieldSolidity))
	return nil
}
