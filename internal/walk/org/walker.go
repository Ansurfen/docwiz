// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package orgwalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/walk"
)

type OrgModeWalker struct {
	walk.BaseWalker
}

func (*OrgModeWalker) SubscribeExt() []string {
	return []string{".org"}
}

func (*OrgModeWalker) ParseExt(fullpath string, ext string, ctx *walk.Context) error {
	ctx.Set("Org Mode", walk.UpgradeBadge("Org Mode", badge.ShieldOrgMode))
	return nil
}
