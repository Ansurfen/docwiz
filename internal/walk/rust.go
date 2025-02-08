// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import (
	"docwiz/internal/badge"
	"docwiz/internal/cfg"
)

type RustWalker struct {
	BaseWalker
}

func (*RustWalker) SubscribeExt() []string {
	return []string{".rs", ".rlib"}
}

func (*RustWalker) SubscribeFile() []string {
	return []string{"Cargo.toml"}
}

func (*RustWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Rust", upgradeBadge("Rust", badge.ShieldRust))
	return nil
}

func (*RustWalker) ParseFile(fullpath string, file string, ctx *Context) error {
	b := ctx.Set("Rust", upgradeBadge("Rust", badge.ShieldRust))
	cargo, err := cfg.ParseCargoToml(fullpath)
	if err != nil {
		return err
	}
	cargoToml := cargo.(cfg.CargoToml)
	if _, ok := cargoToml.Package.RustVersion.(map[string]any); ok {
		b.SetVersion(cargoToml.Workspace.Package.RustVersion.(string))
	} else {
		if v, ok := cargoToml.Package.RustVersion.(string); ok {
			b.SetVersion(v)
		}
	}

	for _, dep := range cargoToml.ProjectDependencies() {
		b := rustLib.Match(dep.Name(), ctx.stackKind)
		if b.Badge == nil {
			continue
		}
		if b.Type == useLibVersion {
			b.Badge.SetVersion(dep.Version())
		}
		ctx.Set(b.Name(), upgradeBadge("rust", b))
	}
	return nil
}

var rustLib = &DependencyManager{
	fullMatches: map[string]badge.Badge{
		"hyperlane": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: shieldBadgeHyperlane},
		},
	},
}

var (
	shieldBadgeHyperlane = &badge.ShieldBadge{
		ID:        "hyperlane",
		Label:     "hyperlane",
		Color:     "#dea584",
		Style:     badge.ShieldStyleDefault,
		Logo:      "rust",
		LogoColor: "white",
		Href:      "https://github.com/ltpp-universe/hyperlane",
	}
)
