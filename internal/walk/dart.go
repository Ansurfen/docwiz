// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import (
	"docwiz/internal/badge"
	"docwiz/internal/cfg"
)

type DartWalker struct {
	BaseWalker
}

func (*DartWalker) SubscribeExt() []string {
	return []string{".dart", ".dart.js"}
}

func (*DartWalker) SubscribeFile() []string {
	return []string{"pubspec.yaml"}
}

func (*DartWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Dart", upgradeBadge("Dart", badge.ShieldDart))
	return nil
}

func (*DartWalker) ParseFile(fullpath string, file string, ctx *Context) error {
	ctx.Set("Dart", upgradeBadge("Dart", badge.ShieldDart))
	pubspec, err := cfg.LoadPubSpecFromFile(fullpath)
	if err != nil {
		return err
	}

	for _, env := range pubspec.Environments() {
		if env.Name() == "sdk" {
			ctx.Get("Dart").Badge.SetVersion(env.Version())
		}
	}

	for _, dep := range pubspec.ProjectDependencies() {
		b := dartLib.Match(dep.Name(), ctx.stackKind)
		if b.Badge == nil {
			continue
		}
		if b.Type == useLibVersion {
			b.Badge.SetVersion(dep.Version())
		}
		ctx.Set(b.Name(), upgradeBadge("Dart", b))
	}

	for _, dep := range pubspec.ProjectDevDependencies() {
		b := dartLib.Match(dep.Name(), ctx.stackKind)
		if b.Badge == nil {
			continue
		}
		if b.Type == useLibVersion {
			b.Badge.SetVersion(dep.Version())
		}
		ctx.Set(b.Name(), upgradeBadge("Dart", b))
	}
	return nil
}

var dartLib = &DependencyManager{
	fullMatches: map[string]badge.Badge{
		"flutter": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldFlutter},
		},
	},
}
