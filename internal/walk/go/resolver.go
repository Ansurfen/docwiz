// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package gowalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/walk"
)

var shiledGoResolver = &walk.DependencyResolver{
	Partial: walk.ResolverPattern{
		"github.com/elastic/go-elasticsearch": walk.SystemVersionBadge{Badge: badge.ShieldElasticSearch},
		"github.com/golang-jwt/jwt":           walk.SystemVersionBadge{Badge: badge.ShieldJWT},
		"github.com/go-redis/redis":           walk.SystemVersionBadge{Badge: badge.ShieldRedis},
		"github.com/go-gl/gl":                 walk.SystemVersionBadge{Badge: badge.ShieldOpenGL},
	},
	Full: walk.ResolverPattern{
		"go.mongodb.org/mongo-driver":    walk.SystemVersionBadge{Badge: badge.ShieldMongoDB},
		"github.com/go-sql-driver/mysql": walk.SystemVersionBadge{Badge: badge.ShieldMySQL},
		"github.com/gin-gonic/gin":       walk.DependencyVersionBadge{Badge: shieldBadgeGin},
		"github.com/gofiber/fiber":       walk.DependencyVersionBadge{Badge: shieldBadgeFiber},
		"github.com/labstack/echo":       walk.DependencyVersionBadge{Badge: shieldBadgeEcho},
		"github.com/beego/beego":         walk.DependencyVersionBadge{Badge: shieldBadgeBeego},
		"github.com/kataras/iris":        walk.DependencyVersionBadge{Badge: shieldBadgeIris},
		"github.com/go-chi/chi":          walk.DependencyVersionBadge{Badge: shieldBadgeChi},
		"github.com/revel/revel":         walk.DependencyVersionBadge{Badge: shieldBadgeRevel},
		"github.com/gobuffalo/buffalo":   walk.DependencyVersionBadge{Badge: shieldBadgeBuffalo},
		"gocv.io/x/gocv":                 walk.SystemVersionBadge{Badge: badge.ShieldOpenCV},
	},
	Fuzzy: walk.ResolverPattern{
		"prometheus": walk.SystemVersionBadge{Badge: badge.ShieldPrometheus},
		"swggo":      walk.SystemVersionBadge{Badge: badge.ShieldSwagger},
	},
}

var badgenGoResolver = &walk.DependencyResolver{}
