// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package pythonwalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/walk"
)

var shieldPythonResolver = &walk.DependencyResolver{
	Full: walk.ResolverPattern{
		"python":    walk.DependencyVersionBadge{Badge: badge.ShieldPython},
		"fastapi":   walk.DependencyVersionBadge{Badge: shieldFastAPI},
		"Jinja2":    walk.DependencyVersionBadge{Badge: badge.ShieldJinja},
		"odps":      walk.DependencyVersionBadge{Badge: badge.ShieldMaxCompute},
		"django":    walk.DependencyVersionBadge{Badge: shieldDjango},
		"flask":     walk.DependencyVersionBadge{Badge: shieldFlask},
		"prefect":   walk.DependencyVersionBadge{Badge: badge.ShieldPrefect},
		"pug":       walk.DependencyVersionBadge{Badge: badge.ShieldPug},
		"pytest":    walk.DependencyVersionBadge{Badge: badge.ShieldPytest},
		"scrapy":    walk.DependencyVersionBadge{Badge: badge.ShieldScrapy},
		"streamlit": walk.DependencyVersionBadge{Badge: badge.ShieldStreamlit},
	},
}
