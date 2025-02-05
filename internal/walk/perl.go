// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type PerlWalker struct {
	BaseWalker
}

func (*PerlWalker) SubscribeExt() []string {
	return []string{".pl", ".pm"}
}

func (*PerlWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("perl", upgradeBadge("perl", badge.ShieldPerl))
	return nil
}
