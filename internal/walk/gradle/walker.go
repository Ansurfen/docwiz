// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package gradlewalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/walk"
)

type Walker struct {
	walk.BaseWalker
}

func (*Walker) SubscribeExt() []string {
	return []string{".gradle"}
}

func (*Walker) SubscribeFile() []string {
	return []string{"gradlew", "gradle.bat", "gradle-wrapper.properties"}
}

func (*Walker) ParseExt(fullpath string, ext string, ctx *walk.Context) error {
	ctx.Set("Gradle", walk.UpgradeBadge("Gradle", badge.ShieldGradle))
	return nil
}

func (*Walker) ParseFile(fullpath string, file string, ctx *walk.Context) error {
	ctx.Set("Gradle", walk.UpgradeBadge("Gradle", badge.ShieldGradle))
	return nil
}
