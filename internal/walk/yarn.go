// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type YarnWalker struct {
	BaseWalker
}

func (*YarnWalker) SubscribeFile() []string {
	return []string{"yarn.lock"}
}

func (*YarnWalker) ParseFile(fullpath string, file string, ctx *Context) error {
	ctx.Set("Yarn", upgradeBadge("Yarn", badge.ShieldYarn))
	return nil
}
