// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type LuaWalker struct {
	BaseWalker
}

func (*LuaWalker) SubscribeExt() []string {
	return []string{".lua", ".luac"}
}

func (*LuaWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Lua", upgradeBadge("Lua", badge.ShieldLua))
	return nil
}
