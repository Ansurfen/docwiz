// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import "docwiz/internal/badge"

type PHPWalker struct {
	BaseWalker
}

func (*PHPWalker) SubscribeExt() []string {
	return []string{".php", ".phtml", ".php3", ".php4", ".php5"}
}

func (*PHPWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("PHP", upgradeBadge("PHP", badge.ShieldPHP))
	return nil
}

var phpLib = &DependencyManager{
	fullMatches: map[string]badge.Badge{
		"codeigniter4/framework": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldCodeIgniter},
		},
		"getgrav/grav": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldGrav},
		},
	},
	fuzzyMatches: map[string]badge.Badge{
		"joomla": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldJoomla},
		},
	},
}
