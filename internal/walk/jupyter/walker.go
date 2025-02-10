// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package jupyterwalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/walk"
)

type Walker struct {
	walk.BaseWalker
}

func (*Walker) SubscribeExt() []string {
	return []string{".ipynb"}
}

func (*Walker) ParseFile(fullpath string, file string, ctx *walk.Context) error {
	ctx.Set("Jupyter", walk.UpgradeBadge("Jupyter", badge.ShieldJupyterNotebook))
	return nil
}
