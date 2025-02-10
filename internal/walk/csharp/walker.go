// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package csharpwalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/walk"
)

type Walker struct {
	walk.BaseWalker
}

func (*Walker) SubscribeExt() []string {
	return []string{".cs", ".csproj", ".sln"}
}

func (*Walker) SubscribeFile() []string {
	return []string{"nuget.config", "packages.config"}
}

func (*Walker) ParseExt(fullpath string, ext string, ctx *walk.Context) error {
	ctx.Set("C#", walk.UpgradeBadge("C#", badge.ShieldCSharp))
	ctx.Set(".NET", walk.UpgradeBadge("C#", badge.ShieldDotNet))
	return nil
}

func (*Walker) ParseFile(fullpath string, file string, ctx *walk.Context) error {
	ctx.Set(".NET", walk.UpgradeBadge("C#", badge.ShieldDotNet))
	return nil
}

// Blazor
// opencv
// opengl