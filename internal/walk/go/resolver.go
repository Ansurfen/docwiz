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
		"github.com/neo4j/neo4j-go-driver":    walk.SystemVersionBadge{Badge: badge.ShieldNeo4J},
	},
	Full: walk.ResolverPattern{
		"github.com/aws/aws-sdk-go-v2":        walk.SystemVersionBadge{Badge: badge.ShieldAmazonDynamoDB},
		"github.com/appwrite/sdk-for-go":      walk.SystemVersionBadge{Badge: badge.ShieldAppwrite},
		"github.com/arangodb/go-driver":       walk.SystemVersionBadge{Badge: badge.ShieldArangoDB},
		"github.com/gocql/gocql":              walk.SystemVersionBadge{Badge: badge.ShieldApacheCassandra},
		"github.com/ClickHouse/clickhouse-go": walk.SystemVersionBadge{Badge: badge.ShieldClickHouse},
		"github.com/couchbase/gocb":           walk.SystemVersionBadge{Badge: badge.ShieldCouchbase},
		// crateDB
		"firebase.google.com/go":                   walk.SystemVersionBadge{Badge: badge.ShieldFirebase},
		"github.com/influxdata/influxdb-client-go": walk.SystemVersionBadge{Badge: badge.ShieldInfluxDB},
		// MariaDB
		// MusicBrainz
		"github.com/denisenkom/go-mssqldb":      walk.SystemVersionBadge{Badge: badge.ShieldMicrosoftSQLServer},
		"go.mongodb.org/mongo-driver":           walk.SystemVersionBadge{Badge: badge.ShieldMongoDB},
		"github.com/planetscale/planetscale-go": walk.SystemVersionBadge{Badge: badge.ShieldPlanetScale},
		"github.com/go-sql-driver/mysql":        walk.SystemVersionBadge{Badge: badge.ShieldMySQL},
		"github.com/pocketbase/pocketbase":      walk.SystemVersionBadge{Badge: badge.ShieldPocketBase},
		"github.com/lib/pq":                     walk.SystemVersionBadge{Badge: badge.ShieldPostgres},
		// Realm
		// SingleStore
		"github.com/mattn/go-sqlite3":      walk.SystemVersionBadge{Badge: badge.ShieldSQLite},
		"github.com/supabase/postgrest-go": walk.SystemVersionBadge{Badge: badge.ShieldSupabase},
		// Teradata
		"github.com/gin-gonic/gin":     walk.DependencyVersionBadge{Badge: shieldBadgeGin},
		"github.com/gofiber/fiber":     walk.DependencyVersionBadge{Badge: shieldBadgeFiber},
		"github.com/labstack/echo":     walk.DependencyVersionBadge{Badge: shieldBadgeEcho},
		"github.com/beego/beego":       walk.DependencyVersionBadge{Badge: shieldBadgeBeego},
		"github.com/kataras/iris":      walk.DependencyVersionBadge{Badge: shieldBadgeIris},
		"github.com/go-chi/chi":        walk.DependencyVersionBadge{Badge: shieldBadgeChi},
		"github.com/revel/revel":       walk.DependencyVersionBadge{Badge: shieldBadgeRevel},
		"github.com/gobuffalo/buffalo": walk.DependencyVersionBadge{Badge: shieldBadgeBuffalo},
		"gocv.io/x/gocv":               walk.SystemVersionBadge{Badge: badge.ShieldOpenCV},

		"github.com/go-gorm/gorm":           walk.DependencyVersionBadge{Badge: shieldBadgeGorm},
		"github.com/jmoiron/sqlx":           walk.DependencyVersionBadge{Badge: shieldBadgeSqlx},
		"github.com/go-xorm/xorm":           walk.DependencyVersionBadge{Badge: shieldBadgeXorm},
		"github.com/ent/ent":                walk.DependencyVersionBadge{Badge: shieldBadgeEnt},
		"github.com/beego/beego/v2/orm":     walk.DependencyVersionBadge{Badge: shieldBadgeBeegoOrm},
		"github.com/asdine/storm":           walk.DependencyVersionBadge{Badge: shieldBadgeStorm},
		"github.com/volatiletech/sqlboiler": walk.DependencyVersionBadge{Badge: shieldBadgeSqlboiler},
	},
	Fuzzy: walk.ResolverPattern{
		"prometheus": walk.SystemVersionBadge{Badge: badge.ShieldPrometheus},
		"swggo":      walk.SystemVersionBadge{Badge: badge.ShieldSwagger},
	},
}

var badgenGoResolver = &walk.DependencyResolver{}
