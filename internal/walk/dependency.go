// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import (
	"docwiz/internal/badge"
	"strings"
)

type DependencyManager struct {
	partialMatches map[string]badge.Badge
	fullMatches    map[string]badge.Badge
	fuzzyMatches   map[string]badge.Badge
}

func (w *DependencyManager) Match(name string, kind BadgeKind) badge.TypedBadge {
	var ret badge.Badge
	if v, ok := w.fullMatches[name]; ok {
		ret = v
		goto handle
	}
	for k, v := range w.partialMatches {
		if strings.HasPrefix(name, k) {
			ret = v
			goto handle
		}
	}
	for k, v := range w.fuzzyMatches {
		if strings.Contains(name, k) {
			ret = v
			goto handle
		}
	}
handle:
	if ret != nil {
		union := ret.(*badge.BadgeUnion)
		if kind == BadgeKindShield {
			return union.ShieldBadge
		}
		return union.BadgenBadge
	}
	return badge.TypedBadge{}
}
