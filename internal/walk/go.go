// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import (
	"docwiz/internal/badge"
	"docwiz/internal/cfg"
)

type GoWalker struct {
	BaseWalker
}

func (*GoWalker) SubscribeExt() []string {
	return []string{".go", ".sum"}
}

func (*GoWalker) SubscribeFile() []string {
	return []string{"go.mod", "go.sum"}
}

func (*GoWalker) ParseExt(fullpath, ext string, ctx *Context) error {
	ctx.Set("Go", goBadge(ctx))
	return nil
}

func (*GoWalker) ParseFile(fullpath, file string, ctx *Context) error {
	switch file {
	case "go.mod":
		goBadge := ctx.Set("Go", goBadge(ctx))
		mod, err := cfg.ParseGoMod(fullpath)
		if err != nil {
			return err
		}
		goBadge.Badge.SetVersion(mod.ProjectVersion())

		for _, dep := range mod.ProjectDependencies() {
			b := goLib.Match(dep.Name(), ctx.stackKind)
			if b.Badge == nil {
				continue
			}
			if b.Type == useLibVersion {
				b.Badge.SetVersion(dep.Version())
			}
			ctx.Set(b.Name(), upgradeBadge("Go", b))
		}

	case "go.sum":
		ctx.Set("Go", goBadge(ctx))
	}
	return nil
}

func goBadge(ctx *Context) badge.SortableBadge {
	switch ctx.stackKind {
	case BadgeKindShield:
		return badge.SortableBadge{Badge: badge.ShieldGo, Tag: "Go"}
	case BadgeKindBadgen:
	}

	return badge.SortableBadge{}
}

var goLib = &DependencyManager{
	partialMatches: map[string]badge.Badge{
		"github.com/elastic/go-elasticsearch": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useMiddleVersion, Badge: badge.ShieldElasticSearch},
		},
		"github.com/golang-jwt/jwt": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useMiddleVersion, Badge: badge.ShieldJWT},
		},
		"github.com/go-redis/redis": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useMiddleVersion, Badge: badge.ShieldRedis},
		},
	},
	fullMatches: map[string]badge.Badge{
		"go.mongodb.org/mongo-driver": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useMiddleVersion, Badge: badge.ShieldMongoDB},
		},
		"github.com/go-sql-driver/mysql": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useMiddleVersion, Badge: badge.ShieldMySQL},
		},
		"github.com/gin-gonic/gin": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{
				Type:  useLibVersion,
				Badge: shieldBadgeGin,
			},
		},
		"github.com/gofiber/fiber": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{
				Type:  useLibVersion,
				Badge: shieldBadgeFiber,
			},
		},
		"github.com/labstack/echo": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{
				Type:  useLibVersion,
				Badge: shieldBadgeEcho,
			},
		},
		"github.com/beego/beego": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{
				Type:  useLibVersion,
				Badge: shieldBadgeBeego,
			},
		},
		"github.com/kataras/iris": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{
				Type:  useLibVersion,
				Badge: shieldBadgeIris,
			},
		},
		"github.com/go-chi/chi": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{
				Type:  useLibVersion,
				Badge: shieldBadgeChi,
			},
		},
		"github.com/revel/revel": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{
				Type:  useLibVersion,
				Badge: shieldBadgeRevel,
			},
		},
		"github.com/gobuffalo/buffalo": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{
				Type:  useLibVersion,
				Badge: shieldBadgeBuffalo,
			},
		},
	},
	fuzzyMatches: map[string]badge.Badge{
		"prometheus": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{
				Type:  useMiddleVersion,
				Badge: badge.ShieldPrometheus,
			},
		},
		"swaggo": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useMiddleVersion, Badge: badge.ShieldSwagger},
		},
	},
}

const (
	useLibVersion    = "libVersion"
	useMiddleVersion = "middleVersion"
)

var (
	shieldBadgeGin = &badge.ShieldBadge{
		ID:        "Gin",
		Label:     "Gin",
		Color:     "#ffffff",
		Style:     badge.ShieldStyleDefault,
		Logo:      "go",
		LogoColor: "blue",
		Href:      "https://github.com/gin-gonic/gin",
	}
	shieldBadgeFiber = &badge.ShieldBadge{
		ID:        "Fiber",
		Label:     "Fiber",
		Color:     "#1DBA90",
		Style:     badge.ShieldStyleDefault,
		Logo:      "go",
		LogoColor: "white",
		Href:      "https://github.com/gofiber/fiber",
	}
	shieldBadgeEcho = &badge.ShieldBadge{
		ID:        "Echo",
		Label:     "Echo",
		Color:     "#1D9BF0",
		Style:     badge.ShieldStyleDefault,
		Logo:      "go",
		LogoColor: "white",
		Href:      "https://github.com/labstack/echo",
	}
	shieldBadgeBeego = &badge.ShieldBadge{
		ID:        "Beego",
		Label:     "Beego",
		Color:     "#0A74DA",
		Style:     badge.ShieldStyleDefault,
		Logo:      "go",
		LogoColor: "white",
		Href:      "https://github.com/beego/beego",
	}
	shieldBadgeIris = &badge.ShieldBadge{
		ID:        "Iris",
		Label:     "Iris",
		Color:     "#5A4FCF",
		Style:     badge.ShieldStyleDefault,
		Logo:      "go",
		LogoColor: "white",
		Href:      "https://github.com/kataras/iris",
	}
	shieldBadgeChi = &badge.ShieldBadge{
		ID:        "Chi",
		Label:     "Chi",
		Color:     "#CCCCCC",
		Style:     badge.ShieldStyleDefault,
		Logo:      "go",
		LogoColor: "white",
		Href:      "https://github.com/go-chi/chi",
	}
	shieldBadgeRevel = &badge.ShieldBadge{
		ID:        "Revel",
		Label:     "Revel",
		Color:     "#E34F26",
		Style:     badge.ShieldStyleDefault,
		Logo:      "go",
		LogoColor: "white",
		Href:      "https://github.com/revel/revel",
	}
	shieldBadgeBuffalo = &badge.ShieldBadge{
		ID:        "Buffalo",
		Label:     "Buffalo",
		Color:     "#D22B2B",
		Style:     badge.ShieldStyleDefault,
		Logo:      "go",
		LogoColor: "white",
		Href:      "https://github.com/gobuffalo/buffalo",
	}
)
