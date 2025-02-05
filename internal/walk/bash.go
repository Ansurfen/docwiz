// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type BashWalker struct {
	BaseWalker
}

func (*BashWalker) SubscribeExt() []string {
	return []string{".sh", ".bash", ".zsh", ".bashrc", ".profile"}
}

func (*BashWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Bash", upgradeBadge("Bash", badge.ShieldBashScript))
	return nil
}
