// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type FortranWalker struct {
	BaseWalker
}

func (*FortranWalker) SubscribeExt() []string {
	return []string{".f", ".for", ".f90", ".f95", ".f03", ".f08"}
}

func (*FortranWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Fortran", upgradeBadge("Fortran", badge.ShieldFortran))
	return nil
}
