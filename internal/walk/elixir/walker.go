// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package elixirwalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/walk"
)

type Walker struct {
	walk.BaseWalker
}

func (*Walker) SubscribeExt() []string {
	return []string{".ex", ".exs", ".eex", ".leex"}
}

func (*Walker) SubscribeFile() []string {
	return []string{"mix.exs"}
}

func (*Walker) ParseExt(fullpath string, ext string, ctx *walk.Context) error {
	ctx.Set("Elixir", walk.UpgradeBadge("Elixir", badge.ShieldElixir))
	return nil
}

// Phoenix Framework
