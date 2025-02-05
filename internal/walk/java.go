// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package walk

import (
	"docwiz/internal/badge"
	"docwiz/internal/cfg"
)

type JavaWalker struct {
	BaseWalker
}

func (*JavaWalker) SubscribeExt() []string {
	return []string{".java", ".class", ".jar", ".jmod"}
}

func (*JavaWalker) SubscribeFile() []string {
	return []string{"pom.xml"}
}

func (*JavaWalker) ParseExt(fullpath string, ext string, ctx *Context) error {
	ctx.Set("Java", upgradeBadge("Java", badge.ShieldJava))
	return nil
}

func (*JavaWalker) ParseFile(fullpath string, file string, ctx *Context) error {
	ctx.Set("Maven", upgradeBadge("Java", badge.ShieldApacheMaven))
	pom, err := cfg.ParsePOM(fullpath)
	if err != nil {
		return err
	}
	for _, dep := range pom.ProjectDependencies() {
		b := javaLib.Match(dep.Name(), ctx.stackKind)
		if b.Badge == nil {
			continue
		}
		if b.Type == useLibVersion {
			b.Badge.SetVersion(dep.Version())
		}
		ctx.Set(b.Name(), upgradeBadge("Java", b))
	}
	return nil
}

var javaLib = &DependencyManager{
	fullMatches: map[string]badge.Badge{
		"spring-boot-starter-parent": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldSpring},
		},
		"mysql-connector-java": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useMiddleVersion, Badge: badge.ShieldMySQL},
		},
		"neo4j-ogm-http-driver": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useMiddleVersion, Badge: badge.ShieldNeo4J},
		},
		"spring-kafka": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useMiddleVersion, Badge: badge.ShieldApacheKafka},
		},
		"springfox-swagger-ui": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useMiddleVersion, Badge: badge.ShieldSwagger},
		},
		"jedis": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useMiddleVersion, Badge: badge.ShieldRedis},
		},
		"elasticsearch-rest-client": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useMiddleVersion, Badge: badge.ShieldElasticSearch},
		},
	},
	fuzzyMatches: map[string]badge.Badge{
		"tomcat": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldApacheTomcat},
		},
		"spark-": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldApacheSpark},
		},
		"hadoop-": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldApacheHadoop},
		},
		"hive-": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldApacheHive},
		},
		"apollo-": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldApolloGraphQL},
		},
		"javafx": &badge.BadgeUnion{
			ShieldBadge: badge.TypedBadge{Type: useLibVersion, Badge: badge.ShieldJavaFX},
		},
	},
}
