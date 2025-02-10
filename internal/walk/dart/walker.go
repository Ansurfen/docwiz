package dartwalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/cfg"
	"docwiz/internal/walk"
)

type Walker struct {
	walk.BaseWalker
}

func (*Walker) SubscribeExt() []string {
	return []string{".dart", ".dart.js"}
}

func (*Walker) SubscribeFile() []string {
	return []string{"pubspec.yaml"}
}

func (*Walker) ParseExt(fullpath string, ext string, ctx *walk.Context) error {
	ctx.Set("Dart", walk.UpgradeBadge("Dart", badge.ShieldDart))
	return nil
}

func (*Walker) ParseFile(fullpath string, file string, ctx *walk.Context) error {
	ctx.Set("Dart", walk.UpgradeBadge("Dart", badge.ShieldDart))
	pubspec, err := cfg.LoadPubSpecFromFile(fullpath)
	if err != nil {
		return err
	}

	for _, env := range pubspec.Environments() {
		if env.Name() == "sdk" {
			ctx.Get("Dart").Badge.SetVersion(env.Version())
		}
	}

	return walk.ResolveDependency(ctx,
		map[walk.BadgeKind]*walk.DependencyResolver{
			walk.BadgeKindShield: shiledDartResolver,
		},
		pubspec, "Dart")
}

var shiledDartResolver = &walk.DependencyResolver{
	Full: walk.ResolverPattern{
		"flutter": walk.DependencyVersionBadge{Badge: badge.ShieldFlutter},
	},
}
