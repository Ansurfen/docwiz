package rustwalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/cfg"
	"docwiz/internal/walk"
)

type Walker struct {
	walk.BaseWalker
}

func (*Walker) SubscribeExt() []string {
	return []string{".rs", ".rlib"}
}

func (*Walker) SubscribeFile() []string {
	return []string{"Cargo.toml"}
}

func (*Walker) ParseExt(fullpath string, ext string, ctx *walk.Context) error {
	ctx.Set("Rust", walk.UpgradeBadge("Rust", badge.ShieldRust))
	return nil
}

func (*Walker) ParseFile(fullpath string, file string, ctx *walk.Context) error {
	b := ctx.Set("Rust", walk.UpgradeBadge("Rust", badge.ShieldRust))
	cargo, err := cfg.LoadCargoFromFile(fullpath)
	if err != nil {
		return err
	}

	if envs := cargo.Environments(); len(envs) > 0 {
		b.SetVersion(envs[0].Version())
	}

	return walk.ResolveDependency(ctx,
		map[walk.BadgeKind]*walk.DependencyResolver{
			walk.BadgeKindShield: shiledRustResolver,
		}, cargo, "Rust")
}
