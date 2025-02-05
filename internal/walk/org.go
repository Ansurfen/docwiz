// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type OrgModeWalker struct {
	BaseWalker
}

func (*OrgModeWalker) SubscribeExt() []string {
	return []string{".org"}
}

func (*OrgModeWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Org Mode", upgradeBadge("Org Mode", badge.ShieldOrgMode))
	return nil
}
