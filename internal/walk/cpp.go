// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type CPPWalker struct {
	BaseWalker
}

func (*CPPWalker) SubscribeExt() []string {
	return []string{".cpp", ".hpp", ".cc", ".cxx"}
}

func (*CPPWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("C++", upgradeBadge("C++", badge.ShieldCpp))
	return nil
}
