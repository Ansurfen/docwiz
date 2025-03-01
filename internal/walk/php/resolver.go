// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package phpwalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/walk"
)

var shieldPHPResolver = &walk.DependencyResolver{
	Full: walk.ResolverPattern{
		"codeigniter4/framework": walk.DependencyVersionBadge{Badge: badge.ShieldCodeIgniter},
		"getgrav/grav":           walk.DependencyVersionBadge{Badge: badge.ShieldGrav},
		"laravel/framework":      walk.DependencyVersionBadge{Badge: badge.ShieldLaravel},
		"livewire/livewire":      walk.DependencyVersionBadge{Badge: badge.ShieldLivewire},
		"opencv":                 walk.SystemVersionBadge{Badge: badge.ShieldOpenCV},
	},
	Fuzzy: walk.ResolverPattern{
		"joomla":  walk.DependencyVersionBadge{Badge: badge.ShieldJoomla},
		"symfony": walk.SystemVersionBadge{Badge: badge.ShieldSymfony},
	},
}
