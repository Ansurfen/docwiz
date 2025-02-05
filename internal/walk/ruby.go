// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type RubyWalker struct {
	BaseWalker
}

func (*RubyWalker) SubscribeExt() []string {
	return []string{".rb", ".ru", ".rake", ".gemspec"}
}

func (*RubyWalker) ParseExt(fullpath string, file string, ctx *Context) error {
	ctx.Set("Ruby", upgradeBadge("Ruby", badge.ShieldRuby))
	return nil
}
