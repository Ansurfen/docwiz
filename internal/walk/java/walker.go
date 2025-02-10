// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package javawalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/cfg"
	"docwiz/internal/walk"
)

type Walker struct {
	walk.BaseWalker
}

func (*Walker) SubscribeExt() []string {
	return []string{".java", ".class", ".jar", ".jmod"}
}

func (*Walker) SubscribeFile() []string {
	return []string{"pom.xml"}
}

func (*Walker) ParseExt(fullpath string, ext string, ctx *walk.Context) error {
	ctx.Set("Java", walk.UpgradeBadge("Java", badge.ShieldJava))
	return nil
}

func (*Walker) ParseFile(fullpath string, file string, ctx *walk.Context) error {
	ctx.Set("Maven", walk.UpgradeBadge("Java", badge.ShieldApacheMaven))
	pom, err := cfg.LoadPOMFromFile(fullpath)
	if err != nil {
		return err
	}

	return walk.ResolveDependency(ctx,
		map[walk.BadgeKind]*walk.DependencyResolver{
			walk.BadgeKindShield: shiledJavaResolver,
		}, pom, "Java")
}
