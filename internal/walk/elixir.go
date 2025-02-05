// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type ElixirWalker struct {
	BaseWalker
}

func (*ElixirWalker) SubscribeExt() []string {
	return []string{".ex", ".exs", ".eex", ".leex"}
}

func (*ElixirWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Elixir", upgradeBadge("Elixir", badge.ShieldElixir))
	return nil
}
