// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type PowerShellWalker struct {
	BaseWalker
}

func (*PowerShellWalker) SubscribeExt() []string {
	return []string{".ps1", ".psm1", ".psd1"}
}

func (*PowerShellWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("PowerShell", upgradeBadge("PowerShell", badge.ShieldPowerShell))
	return nil
}
