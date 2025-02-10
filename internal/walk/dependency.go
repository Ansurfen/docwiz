// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import (
	"docwiz/internal/cfg"
	"errors"
	"strings"
)

type ResolverPattern map[string]ExtendedBadge

type DependencyResolver struct {
	Partial ResolverPattern
	Full    ResolverPattern
	Fuzzy   ResolverPattern
}

func (w *DependencyResolver) Match(name string) ExtendedBadge {
	if v, ok := w.Full[name]; ok {
		return v
	}
	for k, v := range w.Partial {
		if strings.HasPrefix(name, k) {
			return v
		}
	}
	for k, v := range w.Fuzzy {
		if strings.Contains(name, k) {
			return v
		}
	}
	return nil
}

func ResolveDependency(ctx *Context, resolvers map[BadgeKind]*DependencyResolver, cfg cfg.Configure, tag string) error {
	if resolver, ok := resolvers[ctx.StackBadgeKind()]; ok {
		for _, dep := range cfg.ProjectDependencies() {
			eb := resolver.Match(dep.Name())
			if eb != nil {
				b := eb.Unwrap()
				if eb.Kind() == ExtraInfoUseUseDependencyVersion {
					b.SetVersion(dep.Version())
				}
				ctx.Set(b.Name(), UpgradeBadge(tag, b))
			}
		}

		for _, dep := range cfg.ProjectDevDependencies() {
			eb := resolver.Match(dep.Name())
			if eb != nil {
				b := eb.Unwrap()
				if eb.Kind() == ExtraInfoUseUseDependencyVersion {
					b.SetVersion(dep.Version())
				}
				ctx.Set(b.Name(), UpgradeBadge(tag, b))
			}
		}
		return nil
	}

	return errors.New("invalid resolver")
}
