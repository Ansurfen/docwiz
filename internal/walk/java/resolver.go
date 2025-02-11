// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package javawalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/walk"
)

var shiledJavaResolver = &walk.DependencyResolver{
	Full: walk.ResolverPattern{
		"spring-boot-starter-parent": walk.DependencyVersionBadge{Badge: badge.ShieldSpring},
		"mysql-connector-java":       walk.SystemVersionBadge{Badge: badge.ShieldMySQL},
		"neo4j-ogm-http-driver":      walk.SystemVersionBadge{Badge: badge.ShieldNeo4J},
		"spring-kafka":               walk.SystemVersionBadge{Badge: badge.ShieldApacheKafka},
		"springfox-swagger-ui":       walk.SystemVersionBadge{Badge: badge.ShieldSwagger},
		"jedis":                      walk.SystemVersionBadge{Badge: badge.ShieldRedis},
		"elasticsearch-rest-client":  walk.SystemVersionBadge{Badge: badge.ShieldElasticSearch},
		"odps-sdk":                   walk.SystemVersionBadge{Badge: badge.ShieldMaxCompute},
		"lwjgl-opengl":               walk.SystemVersionBadge{Badge: badge.ShieldOpenGL},
	},
	Fuzzy: walk.ResolverPattern{
		"tomcat":  walk.DependencyVersionBadge{Badge: badge.ShieldApacheTomcat},
		"spark-":  walk.DependencyVersionBadge{Badge: badge.ShieldApacheSpark},
		"hadoop-": walk.DependencyVersionBadge{Badge: badge.ShieldApacheHadoop},
		"hive-":   walk.DependencyVersionBadge{Badge: badge.ShieldApacheHive},
		"apollo-": walk.DependencyVersionBadge{Badge: badge.ShieldApolloGraphQL},
		"javafx":  walk.DependencyVersionBadge{Badge: badge.ShieldJavaFX},
		"opencv":  walk.SystemVersionBadge{Badge: badge.ShieldOpenCV},
		"quarkus": walk.SystemVersionBadge{Badge: badge.ShieldQuarkus},
	},
}
