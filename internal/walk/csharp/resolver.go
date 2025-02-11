// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package csharpwalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/walk"
)

var shieldCSharpResolver = &walk.DependencyResolver{
	Partial: walk.ResolverPattern{
		"OpenCvSharp4": walk.SystemVersionBadge{Badge: badge.ShieldOpenCV},
		"Xamarin":      walk.DependencyVersionBadge{Badge: badge.ShieldXamarin},
	},
	Full: walk.ResolverPattern{
		"Microsoft.AspNetCore.Blazor": walk.DependencyVersionBadge{Badge: badge.ShieldBlazor},
		"OpenTK":                      walk.SystemVersionBadge{Badge: badge.ShieldOpenGL},
		"SharpGL":                     walk.SystemVersionBadge{Badge: badge.ShieldOpenGL},
	},
}
