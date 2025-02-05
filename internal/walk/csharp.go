// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type CSharpWalker struct {
	BaseWalker
}

func (*CSharpWalker) SubscribeExt() []string {
	return []string{".cs", ".csproj", ".sln"}
}

func (*CSharpWalker) SubscribeFile() []string {
	return []string{"nuget.config", "packages.config"}
}

func (*CSharpWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("C#", upgradeBadge("C#", badge.ShieldCSharp))
	ctx.Set(".NET", upgradeBadge("C#", badge.ShieldDotNet))
	return nil
}

func (*CSharpWalker) ParseFile(fullpath string, file string, ctx *Context) error {
	ctx.Set(".NET", upgradeBadge("C#", badge.ShieldDotNet))
	return nil
}

// Blazor