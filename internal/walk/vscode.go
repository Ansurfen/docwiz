// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type VSCodeWalker struct {
	BaseWalker
}

func (*VSCodeWalker) SubscribeDir() []string {
	return []string{".vscode"}
}

func (*VSCodeWalker) ParseDir(fullpath string, dir string, ctx *Context) error {
	ctx.Set("Visual Studio Code", upgradeBadge("Visual Studio Code", badge.ShieldVisualStudioCode))
	return nil
}
