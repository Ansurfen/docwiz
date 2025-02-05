// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type GradleWalker struct {
	BaseWalker
}

func (*GradleWalker) SubscribeExt() []string {
	return []string{".gradle"}
}

func (*GradleWalker) SubscribeFile() []string {
	return []string{"gradlew", "gradle.bat", "gradle-wrapper.properties"}
}

func (*GradleWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Gradle", upgradeBadge("Gradle", badge.ShieldGradle))
	return nil
}

func (*GradleWalker) ParseFile(fullpath string, file string, ctx *Context) error {
	ctx.Set("Gradle", upgradeBadge("Gradle", badge.ShieldGradle))
	return nil
}
