// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import (
	"docwiz/internal/badge"
	"docwiz/internal/cfg"
)

type JavaScriptWalker struct {
	BaseWalker
}

func (*JavaScriptWalker) SubscribeExt() []string {
	return []string{".js", ".mjs", ".cjs"}
}

func (*JavaScriptWalker) SubscribeFile() []string {
	return []string{"package.json", "bun.lockb", "package-lock.json", "deno.json", "deno.jsonc"}
}

func (*JavaScriptWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("JavaScript", upgradeBadge("JavaScript", badge.ShieldJavaScript))
	return nil
}

func (*JavaScriptWalker) ParseFile(fullpath string, file string, ctx *Context) error {
	ctx.Set("JavaScript", upgradeBadge("JavaScript", badge.ShieldJavaScript))
	switch file {
	case "deno.json", "deno.jsonc":
		ctx.Set("Deno", upgradeBadge("JavaScript", badge.ShieldDenoJS))
	case "package.json":
		pkg, err := cfg.ParsePackageJSON(fullpath)
		if err != nil {
			return err
		}
		pkgJson := pkg.(cfg.PackageJSON)
		if len(pkgJson.Engines.NPM) != 0 {
			ctx.Set("NPM", upgradeBadge("JavaScript", badge.ShieldNPM)).Badge.SetVersion(pkgJson.Engines.NPM)
		}
		if len(pkgJson.Engines.Node) != 0 {
			ctx.Set("NodeJS", upgradeBadge("JavaScript", badge.ShieldNodeJS)).Badge.SetVersion(pkgJson.Engines.Node)
		}

		for _, dep := range pkg.ProjectDependencies() {
			b := jsLib.Match(dep.Name(), ctx.stackKind)
			if b.Badge == nil {
				continue
			}
			if b.Type == useLibVersion {
				b.Badge.SetVersion(dep.Version())
			}
			ctx.Set(b.Name(), upgradeBadge("JavaScript", b))
		}

		for _, dep := range pkg.ProjectDevDependencies() {
			b := jsLib.Match(dep.Name(), ctx.stackKind)
			if b.Badge == nil {
				continue
			}
			if b.Type == useLibVersion {
				b.Badge.SetVersion(dep.Version())
			}
			ctx.Set(b.Name(), upgradeBadge("JavaScript", b))
		}
	case "bun.lockb":
		ctx.Set("Bun", upgradeBadge("JavaScript", badge.ShieldBun))
	}
	return nil
}

var jsLib = &DependencyManager{
	fullMatches: map[string]badge.Badge{
		"bun": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldBun},
		},
		"chart.js": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldChartJS},
		},
		"ejs": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldEJS},
		},
		"vue": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldVueJS},
		},
		"vite": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldVite},
		},
		"sass": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldSASS},
		},
		"antd": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldAntDesign},
		},
		"@adonisjs/core": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldAdonisJS},
		},
		"alpinejs": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldAlpineJS},
		},
		"@apollo/client": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldApolloGraphQL},
		},
		"apollo-server": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldApolloGraphQL},
		},
		"astro": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldAstro},
		},
		"bootstrap": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldBootstrap},
		},
		"Buefy": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldBuefy},
		},
		"bulma": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldBulma},
		},
		"@chakra-ui/react": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldChakraUI},
		},
		"daisyui": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldDaisyUI},
		},
		"directus": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldDirectus},
		},
		"electron": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldElectronJS},
		},
		"ember-source": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldEmber},
		},
		"esbuild": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldEsbuild},
		},
		"expo": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldExpo},
		},
		"express": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldExpressJS},
		},
		"fastify": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldFastify},
		},
		"gatsby": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldGatsby},
		},
		"gsap": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldGreenSock},
		},
		"gulp": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldGulp},
		},
		"@wordpress/block-editor": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldGutenberg},
		},
		"handlebars": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldHandlebars},
		},
		"jasmine": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldJasmine},
		},
		"jquery": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldJQuery},
		},
		"zod": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldZod},
		},
	},
	fuzzyMatches: map[string]badge.Badge{
		"@angular": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldAngular},
		},
		"aurelia": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldAurelia},
		},
		"elasticsearch": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useMiddleVersion, Badge: badge.ShieldElasticSearch},
		},
		"Framework7": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldFramework7},
		},
		"@ionic": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldIonic},
		},
	},
}

// Context API
// filament
