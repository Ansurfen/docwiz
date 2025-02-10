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
