// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package gowalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/cfg"
	"docwiz/internal/walk"
)

type Walker struct {
	walk.BaseWalker
}

func (*Walker) SubscribeExt() []string {
	return []string{".go", ".sum"}
}

func (*Walker) SubscribeFile() []string {
	return []string{"go.mod", "go.sum"}
}

func (*Walker) ParseExt(fullpath, ext string, ctx *walk.Context) error {
	ctx.Set("Go", walk.UpgradeBadge("Go", badge.ShieldGo))
	return nil
}

func (*Walker) ParseFile(fullpath, file string, ctx *walk.Context) error {
	switch file {
	case "go.mod":
		goBadge := ctx.Set("Go", walk.UpgradeBadge("Go", badge.ShieldGo))
		mod, err := cfg.LoadGoModFromFile(fullpath)
		if err != nil {
			return err
		}

		if envs := mod.Environments(); len(envs) > 0 {
			goBadge.Badge.SetVersion(envs[0].Version())
		}

		err = walk.ResolveDependency(ctx,
			map[walk.BadgeKind]*walk.DependencyResolver{
				walk.BadgeKindShield: shiledGoResolver,
				walk.BadgeKindBadgen: badgenGoResolver,
			},
			mod, "Go")
		if err != nil {
			return err
		}

	case "go.sum":
		ctx.Set("Go", walk.UpgradeBadge("Go", badge.ShieldGo))
	}
	return nil
}
