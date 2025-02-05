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
		b.SetVersion(cargoToml.Package.RustVersion.(string))
	}
	return nil
}
