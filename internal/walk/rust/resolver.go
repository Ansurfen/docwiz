// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package rustwalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/walk"
)

var shiledRustResolver = &walk.DependencyResolver{
	Full: walk.ResolverPattern{
		"hyperlane": walk.DependencyVersionBadge{Badge: shieldBadgeHyperlane},
		"opencv": walk.SystemVersionBadge{Badge: badge.ShieldOpenCV},
	},
}
