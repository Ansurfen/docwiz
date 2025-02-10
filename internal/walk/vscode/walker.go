// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package vscodewalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/walk"
)

type Walker struct {
	walk.BaseWalker
}

func (*Walker) SubscribeDir() []string {
	return []string{".vscode"}
}

func (*Walker) ParseDir(fullpath string, dir string, ctx *walk.Context) error {
	ctx.Set("Visual Studio Code", walk.UpgradeBadge("Visual Studio Code", badge.ShieldVisualStudioCode))
	return nil
}
