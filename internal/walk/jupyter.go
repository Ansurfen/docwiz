// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type JupyterWalker struct {
	BaseWalker
}

func (*JupyterWalker) SubscribeExt() []string {
	return []string{".ipynb"}
}

func (*JupyterWalker) ParseFile(fullpath string, file string, ctx *Context) error {
	ctx.Set("Jupyter", upgradeBadge("Jupyter", badge.ShieldJupyterNotebook))
	return nil
}
