// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package phpwalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/cfg"
	"docwiz/internal/walk"
)

type Walker struct {
	walk.BaseWalker
}

func (*Walker) SubscribeExt() []string {
	return []string{".php", ".phtml", ".php3", ".php4", ".php5"}
}

func (*Walker) SubscribeFile() []string {
	return []string{"composer.json"}
}

func (*Walker) ParseExt(fullpath string, ext string, ctx *walk.Context) error {
	ctx.Set("PHP", walk.UpgradeBadge("PHP", badge.ShieldPHP))
	return nil
}

func (*Walker) ParseFile(fullpath string, file string, ctx *walk.Context) error {
	ctx.Set("PHP", walk.UpgradeBadge("PHP", badge.ShieldPHP))
	composer, err := cfg.LoadComposerFromFile(fullpath)
	if err != nil {
		return err
	}

	return walk.ResolveDependency(ctx,
		map[walk.BadgeKind]*walk.DependencyResolver{
			walk.BadgeKindShield: shieldPHPResolver,
		}, composer, "PHP")
}
