// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package jswalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/cfg"
	"docwiz/internal/walk"
)

type Walker struct {
	walk.BaseWalker
}

func (*Walker) SubscribeExt() []string {
	return []string{".js", ".mjs", ".cjs"}
}

func (*Walker) SubscribeFile() []string {
	return []string{"package.json", "bun.lockb", "package-lock.json", "deno.json", "deno.jsonc"}
}

func (*Walker) ParseExt(fullpath string, ext string, ctx *walk.Context) error {
	ctx.Set("JavaScript", walk.UpgradeBadge("JavaScript", badge.ShieldJavaScript))
	return nil
}

func (*Walker) ParseFile(fullpath string, file string, ctx *walk.Context) error {
	ctx.Set("JavaScript", walk.UpgradeBadge("JavaScript", badge.ShieldJavaScript))
	switch file {
	case "deno.json", "deno.jsonc":
		ctx.Set("Deno", walk.UpgradeBadge("JavaScript", badge.ShieldDenoJS))
	case "package.json":
		pkg, err := cfg.LoadCSProjFromFile(fullpath)
		if err != nil {
			return err
		}

		for _, env := range pkg.Environments() {
			var b badge.Badge
			if env.Name() == "NPM" {
				b = badge.ShieldNPM
			} else if env.Name() == "NodeJS" {
				b = badge.ShieldNodeJS
			}
			ctx.Set(env.Name(), walk.UpgradeBadge("JavaScript", b)).Badge.SetVersion(env.Version())
		}

		walk.ResolveDependency(ctx,
			map[walk.BadgeKind]*walk.DependencyResolver{
				walk.BadgeKindShield: shieldJSResolver,
			}, pkg, "JavaScript")
	case "bun.lockb":
		ctx.Set("Bun", walk.UpgradeBadge("JavaScript", badge.ShieldBun))
	}
	return nil
}
