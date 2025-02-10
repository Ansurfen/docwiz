package rustwalk

import "docwiz/internal/walk"

var shiledRustResolver = &walk.DependencyResolver{
	Full: walk.ResolverPattern{
		"hyperlane": walk.DependencyVersionBadge{Badge: shieldBadgeHyperlane},
	},
}
