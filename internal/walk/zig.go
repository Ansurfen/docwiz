// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type ZigWalker struct {
	BaseWalker
}

func (*ZigWalker) SubscribeExt() []string {
	return []string{".zig"}
}

func (*ZigWalker) ParseFile(fullpath string, file string, ctx *Context) error {
	ctx.Set("Zig", upgradeBadge("Zig", badge.ShieldZig))
	return nil
}
