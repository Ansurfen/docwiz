// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package androidwalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/walk"
)

type Walker struct {
	walk.BaseWalker
}

func (*Walker) SubscribeFile() []string {
	return []string{"AndroidManifest.xml"}
}

func (*Walker) ParseFile(fullpath string, file string, ctx *walk.Context) error {
	ctx.Set("Android", walk.UpgradeBadge("Android", badge.ShieldAndroid))
	return nil
}
