// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type AndroidWalker struct {
	BaseWalker
}

func (*AndroidWalker) SubscribeFile() []string {
	return []string{"AndroidManifest.xml"}
}

func (*AndroidWalker) ParseFile(fullpath string, file string, ctx *Context) error {
	ctx.Set("Android", upgradeBadge("Android", badge.ShieldAndroid))
	return nil
}
