// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package badge

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	ShieldStyleFlat        = "flat"
	ShieldStyleFlatSquare  = "flat-square"
	ShieldStylePlastic     = "plastic"
	ShieldStyleForTheBadge = "for-the-badge"
	ShieldStyleSocial      = "social"

	baseURL = "https://img.shields.io/badge/"
)

var ShieldStyleDefault = ShieldStyleForTheBadge

type ShieldBadge struct {
	ID        string
	Label     string
	message   string
	Color     string
	Style     string
	Logo      string
	LogoColor string
	Href      string
}

func (s *ShieldBadge) Name() string {
	return s.ID
}

func (s *ShieldBadge) SetVersion(v string) {
	s.message = v
}

func (b *ShieldBadge) URL() string {
	var sb strings.Builder

	sb.WriteString(url.PathEscape(b.Label))
	if len(b.message) != 0 {
		sb.WriteString("-")
		sb.WriteString(url.PathEscape(b.message))
	}
	if len(b.Color) != 0 {
		sb.WriteString("-")
		sb.WriteString(url.PathEscape(b.Color))
	}
	sb.WriteString(".svg")

	params := url.Values{}

	if len(b.Style) == 0 {
		b.Style = ShieldStyleDefault
	}

	if b.Style != "" {
		params.Set("style", url.QueryEscape(b.Style))
	}
	if b.Logo != "" {
		params.Set("logo", url.QueryEscape(b.Logo))
	}
	if b.LogoColor != "" {
		params.Set("logoColor", url.QueryEscape(b.LogoColor))
	}

	if len(params) > 0 {
		sb.WriteString("?")
		sb.WriteString(params.Encode())
	}
	return baseURL + sb.String()
}

func (s *ShieldBadge) Markdown() string {
	icon := fmt.Sprintf("![%s](%s)", s.ID, s.URL())
	if len(s.Href) != 0 {
		return fmt.Sprintf("[%s](%s)", icon, s.Href)
	}
	return icon
}

func (s *ShieldBadge) RSt() string {
	return fmt.Sprintf(`.. image:: %s
   :alt: %s
`, s.URL(), s.ID)
}

func (s *ShieldBadge) AsciiDoc() string {
	return fmt.Sprintf("image:%s[%s]", s.URL(), s.ID)
}

func (s *ShieldBadge) HTML() string {
	icon := fmt.Sprintf(`<img alt="%s" src="%s">`, s.ID, s.URL())
	if len(s.Href) != 0 {
		return fmt.Sprintf(`<a href="%s">
   %s
</a>
`, s.Href, icon)
	}
	return icon
}

var (
	ShieldAndroid = &ShieldBadge{
		ID:        "Android",
		Label:     "Android",
		Color:     "#3DDC84",
		Style:     ShieldStyleDefault,
		Logo:      "android",
		LogoColor: "white",
		Href:      "https://www.android.com/",
	}

	ShieldAnsible = &ShieldBadge{
		ID:        "Ansible",
		Label:     "Ansible",
		Color:     "#1A1918",
		Style:     ShieldStyleForTheBadge,
		Logo:      "ansible",
		LogoColor: "white",
		Href:      "https://www.ansible.com/",
	}

	ShieldArduino = &ShieldBadge{
		ID:        "Arduino",
		Label:     "Arduino",
		Color:     "#00979D",
		Style:     ShieldStyleForTheBadge,
		Logo:      "Arduino",
		LogoColor: "white",
		Href:      "https://www.arduino.cc/",
	}

	ShieldBabel = &ShieldBadge{
		ID:        "Babel",
		Label:     "Babel",
		Color:     "#F9DC3e",
		Style:     ShieldStyleForTheBadge,
		Logo:      "babel",
		LogoColor: "black",
		Href:      "https://babeljs.io/",
	}

	ShieldCisco = &ShieldBadge{
		ID:        "Cisco",
		Label:     "Cisco",
		Color:     "#049fd9",
		Style:     ShieldStyleForTheBadge,
		Logo:      "cisco",
		LogoColor: "black",
		Href:      "https://www.cisco.com/",
	}

	ShieldCMake = &ShieldBadge{
		ID:        "CMake",
		Label:     "CMake",
		Color:     "#008FBA",
		Style:     ShieldStyleDefault,
		Logo:      "cmake",
		LogoColor: "white",
		Href:      "https://cmake.org/",
	}

	ShieldCodeCov = &ShieldBadge{
		ID:        "CodeCov",
		Label:     "CodeCov",
		Color:     "#ff0077",
		Style:     ShieldStyleForTheBadge,
		Logo:      "codecov",
		LogoColor: "white",
		Href:      "https://codecov.io/",
	}

	ShieldDocker = &ShieldBadge{
		ID:        "Docker",
		Label:     "Docker",
		Color:     "#0db7ed",
		Style:     ShieldStyleDefault,
		Logo:      "docker",
		LogoColor: "white",
		Href:      "https://www.docker.com/",
	}

	ShieldESLint = &ShieldBadge{
		ID:        "ESLint",
		Label:     "ESLint",
		Color:     "#4B3263",
		Style:     ShieldStyleDefault,
		Logo:      "eslint",
		LogoColor: "white",
		Href:      "https://eslint.org/",
	}

	ShieldElasticSearch = &ShieldBadge{
		ID:        "ElasticSearch",
		Label:     "ElasticSearch",
		Color:     "#005571",
		Style:     ShieldStyleDefault,
		Logo:      "elasticsearch",
		LogoColor: "white",
		Href:      "https://www.elastic.co/elasticsearch/",
	}

	ShieldFFmpeg = &ShieldBadge{
		ID:        "FFmpeg",
		Label:     "FFmpeg",
		Color:     "#171717",
		Style:     ShieldStyleDefault,
		Logo:      "ffmpeg",
		LogoColor: "#5cb85c",
		Href:      "https://ffmpeg.org/",
	}

	ShieldGradle = &ShieldBadge{
		ID:        "Gradle",
		Label:     "Gradle",
		Color:     "#02303A",
		Style:     ShieldStyleDefault,
		Logo:      "gradle",
		LogoColor: "white",
		Href:      "https://gradle.org/",
	}

	ShieldGrafana = &ShieldBadge{
		ID:        "Grafana",
		Label:     "Grafana",
		Color:     "#F46800",
		Style:     ShieldStyleForTheBadge,
		Logo:      "grafana",
		LogoColor: "white",
		Href:      "https://grafana.com/",
	}

	ShieldKubernetes = &ShieldBadge{
		ID:        "Kubernetes",
		Label:     "Kubernetes",
		Color:     "#326ce5",
		Style:     ShieldStyleForTheBadge,
		Logo:      "kubernetes",
		LogoColor: "white",
		Href:      "https://kubernetes.io/",
	}

	ShieldJupyterNotebook = &ShieldBadge{
		ID:        "Jupyter Notebook",
		Label:     "jupyter",
		Color:     "#FA0F00",
		Style:     ShieldStyleDefault,
		Logo:      "jupyter",
		LogoColor: "white",
		Href:      "https://jupyter.org/",
	}

	ShieldPrometheus = &ShieldBadge{
		ID:        "Prometheus",
		Label:     "Prometheus",
		Color:     "#E6522C",
		Style:     ShieldStyleDefault,
		Logo:      "prometheus",
		LogoColor: "white",
		Href:      "https://prometheus.io/",
	}

	ShieldOpenTelemetry = &ShieldBadge{
		ID:        "OpenTelemetry",
		Label:     "OpenTelemetry",
		Color:     "#FFFFFF",
		Style:     ShieldStyleDefault,
		Logo:      "opentelemetry",
		LogoColor: "black",
		Href:      "https://opentelemetry.io/",
	}

	ShieldSwagger = &ShieldBadge{
		ID:        "Swagger",
		Label:     "",
		message:   "Swagger",
		Color:     "#Clojure",
		Style:     ShieldStyleDefault,
		Logo:      "swagger",
		LogoColor: "white",
		Href:      "https://swagger.io/",
	}

	// DATABASE
	ShieldAmazonDynamoDB = &ShieldBadge{
		ID:        "AmazonDynamoDB",
		Label:     "Amazon DynamoDB",
		Color:     "#4053D6",
		Style:     ShieldStyleDefault,
		Logo:      "Amazon DynamoDB",
		LogoColor: "white",
		Href:      "https://aws.amazon.com/dynamodb/",
	}

	ShieldAppwrite = &ShieldBadge{
		ID:        "Appwrite",
		Label:     "Appwrite",
		Color:     "#FD366E",
		Style:     ShieldStyleDefault,
		Logo:      "appwrite",
		LogoColor: "white",
		Href:      "https://appwrite.io/",
	}

	ShieldArangoDB = &ShieldBadge{
		ID:        "ArangoDB",
		Label:     "ArangoDB",
		Color:     "#DDE072",
		Style:     ShieldStyleDefault,
		Logo:      "arangodb",
		LogoColor: "white",
		Href:      "https://www.arangodb.com/",
	}

	ShieldApacheCassandra = &ShieldBadge{
		ID:        "ApacheCassandra",
		Label:     "cassandra",
		Color:     "#1287B1",
		Style:     ShieldStyleDefault,
		Logo:      "apache-cassandra",
		LogoColor: "white",
		Href:      "https://cassandra.apache.org/",
	}

	ShieldClickHouse = &ShieldBadge{
		ID:        "ClickHouse",
		Label:     "ClickHouse",
		Color:     "#FFCC01",
		Style:     ShieldStyleDefault,
		Logo:      "clickhouse",
		LogoColor: "white",
		Href:      "https://clickhouse.com/",
	}

	ShieldCockroachLabs = &ShieldBadge{
		ID:        "CockroachLabs",
		Label:     "Cockroach Labs",
		Color:     "#6933FF",
		Style:     ShieldStyleDefault,
		Logo:      "Cockroach Labs",
		LogoColor: "white",
		Href:      "https://www.cockroachlabs.com/",
	}

	ShieldCouchbase = &ShieldBadge{
		ID:        "Couchbase",
		Label:     "Couchbase",
		Color:     "#EA2328",
		Style:     ShieldStyleDefault,
		Logo:      "couchbase",
		LogoColor: "white",
		Href:      "https://www.couchbase.com/",
	}

	ShieldCrateDB = &ShieldBadge{
		ID:        "CrateDB",
		Label:     "CrateDB",
		Color:     "#009DC7",
		Style:     ShieldStyleDefault,
		Logo:      "CrateDB",
		LogoColor: "white",
		Href:      "https://crate.io/",
	}

	ShieldFirebase = &ShieldBadge{
		ID:        "Firebase",
		Label:     "Firebase",
		Color:     "#A08021",
		Style:     ShieldStyleDefault,
		Logo:      "firebase",
		LogoColor: "white",
		Href:      "https://firebase.google.com/",
	}

	ShieldInfluxDB = &ShieldBadge{
		ID:        "InfluxDB",
		Label:     "InfluxDB",
		Color:     "#22ADF6",
		Style:     ShieldStyleDefault,
		Logo:      "InfluxDB",
		LogoColor: "white",
		Href:      "https://www.influxdata.com/",
	}

	ShieldMariaDB = &ShieldBadge{
		ID:        "MariaDB",
		Label:     "MariaDB",
		Color:     "#003545",
		Style:     ShieldStyleDefault,
		Logo:      "mariadb",
		LogoColor: "white",
		Href:      "https://mariadb.org/",
	}

	ShieldMusicBrainz = &ShieldBadge{
		ID:        "MusicBrainz",
		Label:     "Musicbrainz",
		Color:     "#EB743B",
		Style:     ShieldStyleDefault,
		Logo:      "musicbrainz",
		LogoColor: "BA478F",
		Href:      "https://musicbrainz.org/",
	}

	ShieldMicrosoftSQLServer = &ShieldBadge{
		ID:        "MicrosoftSQLServer",
		Label:     "Microsoft SQL Server",
		Color:     "#CC2927",
		Style:     ShieldStyleDefault,
		Logo:      "microsoft sql server",
		LogoColor: "white",
		Href:      "https://www.microsoft.com/en-us/sql-server",
	}

	ShieldMongoDB = &ShieldBadge{
		ID:        "MongoDB",
		Label:     "MongoDB",
		Color:     "#4ea94b",
		Style:     ShieldStyleDefault,
		Logo:      "mongodb",
		LogoColor: "white",
		Href:      "https://www.mongodb.com/",
	}

	ShieldMySQL = &ShieldBadge{
		ID:        "MySQL",
		Label:     "MySQL",
		Color:     "#4479A1",
		Style:     ShieldStyleDefault,
		Logo:      "mysql",
		LogoColor: "white",
		Href:      "https://www.mysql.com/",
	}

	ShieldNeo4J = &ShieldBadge{
		ID:        "Neo4J",
		Label:     "Neo4j",
		Color:     "#008CC1",
		Style:     ShieldStyleDefault,
		Logo:      "neo4j",
		LogoColor: "white",
		Href:      "https://neo4j.com/",
	}

	ShieldPlanetScale = &ShieldBadge{
		ID:        "PlanetScale",
		Label:     "PlanetScale",
		Color:     "#000000",
		Style:     ShieldStyleDefault,
		Logo:      "planetscale",
		LogoColor: "white",
		Href:      "https://planetscale.com/",
	}

	ShieldPocketBase = &ShieldBadge{
		ID:        "PocketBase",
		Label:     "PocketBase",
		Color:     "#b8dbe4",
		Style:     ShieldStyleDefault,
		Logo:      "Pocketbase",
		LogoColor: "black",
		Href:      "https://pocketbase.io/",
	}

	ShieldPostgres = &ShieldBadge{
		ID:        "Postgres",
		Label:     "Postgres",
		Color:     "#316192",
		Style:     ShieldStyleDefault,
		Logo:      "postgresql",
		LogoColor: "white",
		Href:      "https://www.postgresql.org/",
	}

	ShieldRealm = &ShieldBadge{
		ID:        "Realm",
		Label:     "Realm",
		Color:     "#39477F",
		Style:     ShieldStyleDefault,
		Logo:      "realm",
		LogoColor: "white",
		Href:      "https://realm.io/",
	}

	ShieldRedis = &ShieldBadge{
		ID:        "Redis",
		Label:     "redis",
		Color:     "#DD0031",
		Style:     ShieldStyleDefault,
		Logo:      "redis",
		LogoColor: "white",
		Href:      "https://redis.io/",
	}

	ShieldSingleStore = &ShieldBadge{
		ID:        "SingleStore",
		Label:     "Single Store",
		Color:     "#AA00FF",
		Style:     ShieldStyleDefault,
		Logo:      "singlestore",
		LogoColor: "white",
		Href:      "https://www.singlestore.com/",
	}

	ShieldSQLite = &ShieldBadge{
		ID:        "SQLite",
		Label:     "SQLite",
		Color:     "#07405E",
		Style:     ShieldStyleDefault,
		Logo:      "sqlite",
		LogoColor: "white",
		Href:      "https://www.sqlite.org/",
	}

	ShieldSupabase = &ShieldBadge{
		ID:        "Supabase",
		Label:     "Supabase",
		Color:     "#3ECF8E",
		Style:     ShieldStyleDefault,
		Logo:      "supabase",
		LogoColor: "white",
		Href:      "https://supabase.io/",
	}

	ShieldSurrealDB = &ShieldBadge{
		ID:        "SurrealDB",
		Label:     "SurrealDB",
		Color:     "#FF00A0",
		Style:     ShieldStyleDefault,
		Logo:      "surrealdb",
		LogoColor: "white",
		Href:      "https://surrealdb.com/",
	}

	ShieldTeradata = &ShieldBadge{
		ID:        "Teradata",
		Label:     "Teradata",
		Color:     "#F37440",
		Style:     ShieldStyleDefault,
		Logo:      "teradata",
		LogoColor: "white",
		Href:      "https://www.teradata.com/",
	}
	////////////////////////////////////////////

	// FRAMEWORK
	ShieldDotNet = &ShieldBadge{
		ID:        ".NET",
		Label:     ".NET",
		Color:     "#5C2D91",
		Style:     ShieldStyleDefault,
		Logo:      ".net",
		LogoColor: "white",
		Href:      "https://dotnet.microsoft.com/",
	}

	ShieldAdonisJS = &ShieldBadge{
		ID:        "AdonisJS",
		Label:     "AdonisJS",
		Color:     "#220052",
		Style:     ShieldStyleDefault,
		Logo:      "adonisjs",
		LogoColor: "white",
		Href:      "https://adonisjs.com/",
	}

	ShieldAiohttp = &ShieldBadge{
		ID:        "Aiohttp",
		Label:     "Aiohttp",
		Color:     "#2C5BB4",
		Style:     ShieldStyleDefault,
		Logo:      "aiohttp",
		LogoColor: "white",
		Href:      "https://aiohttp.readthedocs.io/en/stable/",
	}

	ShieldAlpineJS = &ShieldBadge{
		ID:        "Alpine.js",
		Label:     "Alpine.js",
		Color:     "white",
		Style:     ShieldStyleDefault,
		Logo:      "alpinedotjs",
		LogoColor: "#8BC0D0",
		Href:      "https://alpinejs.dev/",
	}

	ShieldAnaconda = &ShieldBadge{
		ID:        "Anaconda",
		Label:     "Anaconda",
		Color:     "#44A833",
		Style:     ShieldStyleDefault,
		Logo:      "anaconda",
		LogoColor: "white",
		Href:      "https://www.anaconda.com/",
	}

	ShieldAngular = &ShieldBadge{
		ID:        "Angular",
		Label:     "Angular",
		Color:     "#DD0031",
		Style:     ShieldStyleDefault,
		Logo:      "angular",
		LogoColor: "white",
		Href:      "https://angular.io/",
	}

	ShieldAngularJS = &ShieldBadge{
		ID:        "Angular.js",
		Label:     "Angular.js",
		Color:     "#E23237",
		Style:     ShieldStyleDefault,
		Logo:      "angularjs",
		LogoColor: "white",
		Href:      "https://angularjs.org/",
	}

	ShieldAntDesign = &ShieldBadge{
		ID:        "Ant-Design",
		Label:     "AntDesign",
		Color:     "#0170FE",
		Style:     ShieldStyleDefault,
		Logo:      "ant-design",
		LogoColor: "white",
		Href:      "https://ant.design/",
	}

	ShieldApacheSpark = &ShieldBadge{
		ID:        "Apache Spark",
		Label:     "Apache Spark",
		Color:     "#FDEE21",
		Style:     ShieldStyleFlatSquare,
		Logo:      "apachespark",
		LogoColor: "black",
		Href:      "https://spark.apache.org/",
	}

	ShieldApacheKafka = &ShieldBadge{
		ID:        "Apache Kafka",
		Label:     "Apache Kafka",
		Color:     "000000",
		Style:     ShieldStyleDefault,
		Logo:      "apachekafka",
		LogoColor: "white",
		Href:      "https://kafka.apache.org/",
	}

	ShieldApacheHadoop = &ShieldBadge{
		ID:        "Apache Hadoop",
		Label:     "Apache Hadoop",
		Color:     "#66CCFF",
		Style:     ShieldStyleDefault,
		Logo:      "apachehadoop",
		LogoColor: "black",
		Href:      "https://hadoop.apache.org/",
	}

	ShieldApacheHive = &ShieldBadge{
		ID:        "Apache Hive",
		Label:     "Apache Hive",
		Color:     "#FDEE21",
		Style:     ShieldStyleDefault,
		Logo:      "apachehive",
		LogoColor: "black",
		Href:      "https://hive.apache.org/",
	}

	ShieldApolloGraphQL = &ShieldBadge{
		ID:        "Apollo-GraphQL",
		Label:     "Apollo GraphQL",
		Color:     "#311C87",
		Style:     ShieldStyleDefault,
		Logo:      "apollo-graphql",
		LogoColor: "white",
		Href:      "https://www.apollographql.com/",
	}

	ShieldAstro = &ShieldBadge{
		ID:        "Astro",
		Label:     "Astro",
		Color:     "#2C2052",
		Style:     ShieldStyleDefault,
		Logo:      "astro",
		LogoColor: "white",
		Href:      "https://astro.build/",
	}

	ShieldAurelia = &ShieldBadge{
		ID:        "Aurelia",
		Label:     "Aurelia",
		Color:     "#ED2B88",
		Style:     ShieldStyleDefault,
		Logo:      "aurelia",
		LogoColor: "white",
		Href:      "https://aurelia.io/",
	}

	ShieldBlazor = &ShieldBadge{
		ID:        "Blazor",
		Label:     "Blazor",
		Color:     "#5C2D91",
		Style:     ShieldStyleDefault,
		Logo:      "blazor",
		LogoColor: "white",
		Href:      "https://dotnet.microsoft.com/en-us/apps/aspnet/web-apps/blazor",
	}

	ShieldBootstrap = &ShieldBadge{
		ID:        "Bootstrap",
		Label:     "Bootstrap",
		Color:     "#8511FA",
		Style:     ShieldStyleDefault,
		Logo:      "bootstrap",
		LogoColor: "white",
		Href:      "https://getbootstrap.com/",
	}

	ShieldBuefy = &ShieldBadge{
		ID:        "Buefy",
		Label:     "Buefy",
		Color:     "#7957D5",
		Style:     ShieldStyleDefault,
		Logo:      "buefy",
		LogoColor: "#48289E",
		Href:      "https://buefy.org/",
	}

	ShieldBulma = &ShieldBadge{
		ID:        "Bulma",
		Label:     "Bulma",
		Color:     "#00D0B1",
		Style:     ShieldStyleDefault,
		Logo:      "bulma",
		LogoColor: "white",
		Href:      "https://bulma.io/",
	}

	ShieldBun = &ShieldBadge{
		ID:        "Bun",
		Label:     "Bun",
		Color:     "#000000",
		Style:     ShieldStyleDefault,
		Logo:      "bun",
		LogoColor: "white",
		Href:      "https://bun.sh/",
	}

	ShieldCelery = &ShieldBadge{
		ID:        "Celery",
		Label:     "Celery",
		Color:     "#A9CC54",
		Style:     ShieldStyleDefault,
		Logo:      "celery",
		LogoColor: "#DDF4A4",
		Href:      "https://celeryproject.org/",
	}

	ShieldChakraUI = &ShieldBadge{
		ID:        "Chakra",
		Label:     "Chakra UI",
		Color:     "#4ED1C5",
		Style:     ShieldStyleDefault,
		Logo:      "chakraui",
		LogoColor: "white",
		Href:      "https://chakra-ui.com/",
	}

	ShieldChartJS = &ShieldBadge{
		ID:        "Chart.js",
		Label:     "Chart.js",
		Color:     "#F5788D",
		Style:     ShieldStyleDefault,
		Logo:      "chart.js",
		LogoColor: "white",
		Href:      "https://www.chartjs.org/",
	}

	ShieldCodeIgniter = &ShieldBadge{
		ID:        "Code-Igniter",
		Label:     "Code Igniter",
		Color:     "#EF4223",
		Style:     ShieldStyleDefault,
		Logo:      "codeIgniter",
		LogoColor: "white",
		Href:      "https://codeigniter.com/",
	}

	ShieldContextAPI = &ShieldBadge{
		ID:        "Context-API",
		Label:     "Context API",
		Color:     "#000000",
		Style:     ShieldStyleDefault,
		Logo:      "react",
		LogoColor: "white",
		Href:      "https://reactjs.org/docs/context.html",
	}

	ShieldCUDA = &ShieldBadge{
		ID:        "nVIDIA",
		Label:     "CUDA",
		Color:     "#000000",
		Style:     ShieldStyleDefault,
		Logo:      "nvidia",
		LogoColor: "green",
		Href:      "https://developer.nvidia.com/cuda-zone",
	}

	ShieldDaisyUI = &ShieldBadge{
		ID:        "DaisyUI",
		Label:     "DaisyUI",
		Color:     "#5A0EF8",
		Style:     ShieldStyleDefault,
		Logo:      "daisyui",
		LogoColor: "white",
		Href:      "https://daisyui.com/",
	}

	ShieldDenoJS = &ShieldBadge{
		ID:        "Deno JS",
		Label:     "Deno JS",
		Color:     "000000",
		Style:     ShieldStyleDefault,
		Logo:      "deno",
		LogoColor: "white",
		Href:      "https://deno.land/",
	}

	ShieldDirectus = &ShieldBadge{
		ID:        "Directus",
		Label:     "Directus",
		Color:     "#64F",
		Style:     ShieldStyleDefault,
		Logo:      "directus",
		LogoColor: "white",
		Href:      "https://directus.io/",
	}

	ShieldDjango = &ShieldBadge{
		ID:        "Django",
		Label:     "Django",
		Color:     "#092E20",
		Style:     ShieldStyleDefault,
		Logo:      "django",
		LogoColor: "white",
		Href:      "https://www.djangoproject.com/",
	}

	ShieldDjangoREST = &ShieldBadge{
		ID:        "DjangoREST",
		Label:     "Django REST",
		Color:     "#FF1709",
		Style:     ShieldStyleDefault,
		Logo:      "django",
		LogoColor: "white",
		Href:      "https://www.django-rest-framework.org/",
	}

	ShieldDrupal = &ShieldBadge{
		ID:        "Drupal",
		Label:     "Drupal",
		Color:     "#0678BE",
		Style:     ShieldStyleDefault,
		Logo:      "drupal",
		LogoColor: "white",
		Href:      "https://www.drupal.org/",
	}

	ShieldEJS = &ShieldBadge{
		ID:        "EJS",
		Label:     "EJS",
		Color:     "#B4CA65",
		Style:     ShieldStyleDefault,
		Logo:      "ejs",
		LogoColor: "black",
		Href:      "https://www.embeddedjs.com/",
	}

	ShieldElasticsearch = &ShieldBadge{
		ID:        "Elasticsearch",
		Label:     "Elasticsearch",
		Color:     "#0377CC",
		Style:     ShieldStyleDefault,
		Logo:      "elasticsearch",
		LogoColor: "white",
		Href:      "https://www.elastic.co/elasticsearch/",
	}

	ShieldElectronJS = &ShieldBadge{
		ID:        "Electron.js",
		Label:     "Electron.js",
		Color:     "#191970",
		Style:     ShieldStyleDefault,
		Logo:      "electron",
		LogoColor: "white",
		Href:      "https://www.electronjs.org/",
	}

	ShieldEmber = &ShieldBadge{
		ID:        "Ember",
		Label:     "Ember",
		Color:     "#1C1E24",
		Style:     ShieldStyleDefault,
		Logo:      "ember.js",
		LogoColor: "#D04A37",
		Href:      "https://emberjs.com/",
	}

	ShieldEsbuild = &ShieldBadge{
		ID:        "Esbuild",
		Label:     "Esbuild",
		Color:     "#FFCF00",
		Style:     ShieldStyleDefault,
		Logo:      "esbuild",
		LogoColor: "black",
		Href:      "https://esbuild.github.io/",
	}

	ShieldExpo = &ShieldBadge{
		ID:        "Expo",
		Label:     "Expo",
		Color:     "#1C1E24",
		Style:     ShieldStyleDefault,
		Logo:      "expo",
		LogoColor: "#D04A37",
		Href:      "https://expo.dev/",
	}

	ShieldExpressJS = &ShieldBadge{
		ID:        "Express.js",
		Label:     "Express.js",
		Color:     "#404D59",
		Style:     ShieldStyleDefault,
		Logo:      "express",
		LogoColor: "#61DAFB",
		Href:      "https://expressjs.com/",
	}

	ShieldFastAPI = &ShieldBadge{
		ID:        "FastAPI",
		Label:     "FastAPI",
		Color:     "#005571",
		Style:     ShieldStyleDefault,
		Logo:      "fastapi",
		LogoColor: "white",
		Href:      "https://fastapi.tiangolo.com/",
	}

	ShieldFastify = &ShieldBadge{
		ID:        "Fastify",
		Label:     "Fastify",
		Color:     "#000000",
		Style:     ShieldStyleDefault,
		Logo:      "fastify",
		LogoColor: "white",
		Href:      "https://www.fastify.io/",
	}

	ShieldFilament = &ShieldBadge{
		ID:        "Filament",
		Label:     "Filament",
		Color:     "#FFAA00",
		Style:     ShieldStyleDefault,
		Logo:      "filament",
		LogoColor: "#000000",
		Href:      "https://filamentphp.com/",
	}

	ShieldFlask = &ShieldBadge{
		ID:        "Flask",
		Label:     "Flask",
		Color:     "#000000",
		Style:     ShieldStyleDefault,
		Logo:      "flask",
		LogoColor: "white",
		Href:      "https://flask.palletsprojects.com/",
	}

	ShieldFlutter = &ShieldBadge{
		ID:        "Flutter",
		Label:     "Flutter",
		Color:     "#02569B",
		Style:     ShieldStyleDefault,
		Logo:      "flutter",
		LogoColor: "white",
		Href:      "https://flutter.dev/",
	}

	ShieldFramework7 = &ShieldBadge{
		ID:        "Framework7",
		Label:     "Framework7",
		Color:     "#EE350F",
		Style:     ShieldStyleDefault,
		Logo:      "framework7",
		LogoColor: "white",
		Href:      "https://framework7.io/",
	}

	ShieldGatsby = &ShieldBadge{
		ID:        "Gatsby",
		Label:     "Gatsby",
		Color:     "#663399",
		Style:     ShieldStyleDefault,
		Logo:      "gatsby",
		LogoColor: "white",
		Href:      "https://www.gatsbyjs.com/",
	}

	ShieldGrav = &ShieldBadge{
		ID:        "Grav",
		Label:     "Grav",
		Color:     "#FFFFFF",
		Style:     ShieldStyleDefault,
		Logo:      "grav",
		LogoColor: "#221E1F",
		Href:      "https://getgrav.org/",
	}

	ShieldGreenSock = &ShieldBadge{
		ID:        "Green Sock",
		Label:     "Green Sock",
		Color:     "#88CE02",
		Style:     ShieldStyleDefault,
		Logo:      "greensock",
		LogoColor: "white",
		Href:      "https://greensock.com/",
	}

	ShieldGulp = &ShieldBadge{
		ID:        "Gulp",
		Label:     "Gulp",
		Color:     "#CF4647",
		Style:     ShieldStyleDefault,
		Logo:      "gulp",
		LogoColor: "white",
		Href:      "https://gulpjs.com/",
	}

	ShieldGutenberg = &ShieldBadge{
		ID:        "Gutenberg",
		Label:     "Gutenberg",
		Color:     "#077CB2",
		Style:     ShieldStyleDefault,
		Logo:      "gutenberg",
		LogoColor: "white",
		Href:      "https://wordpress.org/gutenberg/",
	}

	ShieldInsomnia = &ShieldBadge{
		ID:        "Insomnia",
		Label:     "Insomnia",
		Color:     "black",
		Style:     ShieldStyleDefault,
		Logo:      "insomnia",
		LogoColor: "#5849BE",
		Href:      "https://insomnia.rest/",
	}

	ShieldHandlebars = &ShieldBadge{
		ID:        "Handlebars",
		Label:     "Handlebars",
		Color:     "#000000",
		Style:     ShieldStyleDefault,
		Logo:      "Handlebars.js",
		LogoColor: "white",
		Href:      "https://handlebarsjs.com/",
	}

	ShieldHugo = &ShieldBadge{
		ID:        "Hugo",
		Label:     "Hugo",
		Color:     "black",
		Style:     ShieldStyleDefault,
		Logo:      "hugo",
		LogoColor: "white",
		Href:      "https://gohugo.io/",
	}

	ShieldIonic = &ShieldBadge{
		ID:        "Ionic",
		Label:     "Ionic",
		Color:     "#3880FF",
		Style:     ShieldStyleDefault,
		Logo:      "ionic",
		LogoColor: "white",
		Href:      "https://ionicframework.com/",
	}

	ShieldJasmine = &ShieldBadge{
		ID:        "Jasmine",
		Label:     "Jasmine",
		Color:     "#8A4182",
		Style:     ShieldStyleDefault,
		Logo:      "jasmine",
		LogoColor: "white",
		Href:      "https://jasmine.github.io/",
	}

	ShieldJavaFX = &ShieldBadge{
		ID:        "JavaFX",
		Label:     "JavaFX",
		Color:     "#FF0000",
		Style:     ShieldStyleDefault,
		Logo:      "javafx",
		LogoColor: "white",
		Href:      "https://openjfx.io/",
	}

	ShieldJinja = &ShieldBadge{
		ID:        "Jinja",
		Label:     "Jinja",
		Color:     "white",
		Style:     ShieldStyleDefault,
		Logo:      "jinja",
		LogoColor: "black",
		Href:      "https://jinja.palletsprojects.com/",
	}

	ShieldJoomla = &ShieldBadge{
		ID:        "Joomla",
		Label:     "Joomla",
		Color:     "#5091CD",
		Style:     ShieldStyleDefault,
		Logo:      "joomla",
		LogoColor: "white",
		Href:      "https://www.joomla.org/",
	}

	ShieldJQuery = &ShieldBadge{
		ID:        "jQuery",
		Label:     "jQuery",
		Color:     "#0769AD",
		Style:     ShieldStyleDefault,
		Logo:      "jquery",
		LogoColor: "white",
		Href:      "https://jquery.com/",
	}

	ShieldJWT = &ShieldBadge{
		ID:        "JWT",
		Label:     "JWT",
		Color:     "black",
		Style:     ShieldStyleDefault,
		Logo:      "JSON web tokens",
		LogoColor: "white",
		Href:      "https://jwt.io/",
	}

	ShieldLaravel = &ShieldBadge{
		ID:        "Laravel",
		Label:     "Laravel",
		Color:     "#FF2D20",
		Style:     ShieldStyleDefault,
		Logo:      "laravel",
		LogoColor: "white",
		Href:      "https://laravel.com/",
	}

	ShieldLivewire = &ShieldBadge{
		ID:        "Livewire",
		Label:     "Livewire",
		Color:     "#4E56A6",
		Style:     ShieldStyleDefault,
		Logo:      "livewire",
		LogoColor: "white",
		Href:      "https://www.livewire.io/",
	}

	ShieldLess = &ShieldBadge{
		ID:        "Less",
		Label:     "Less",
		Color:     "#2B4C80",
		Style:     ShieldStyleDefault,
		Logo:      "less",
		LogoColor: "white",
		Href:      "https://lesscss.org/",
	}

	ShieldMUI = &ShieldBadge{
		ID:        "MUI",
		Label:     "MUI",
		Color:     "#0081CB",
		Style:     ShieldStyleDefault,
		Logo:      "mui",
		LogoColor: "white",
		Href:      "https://mui.com/",
	}

	ShieldMeteorJS = &ShieldBadge{
		ID:        "Meteor JS",
		Label:     "Meteor JS",
		Color:     "#D74C4C",
		Style:     ShieldStyleDefault,
		Logo:      "meteor",
		LogoColor: "white",
		Href:      "https://www.meteor.com/",
	}

	ShieldMantine = &ShieldBadge{
		ID:        "Mantine",
		Label:     "Mantine",
		Color:     "ffffff",
		Style:     ShieldStyleDefault,
		Logo:      "mantine",
		LogoColor: "#339AF0",
		Href:      "https://mantine.dev/",
	}

	ShieldMaxCompute = &ShieldBadge{
		ID:        "MaxCompute",
		Label:     "MaxCompute",
		Color:     "#FF6701",
		Style:     ShieldStyleDefault,
		Logo:      "alibabacloud",
		LogoColor: "white",
		Href:      "https://www.alibabacloud.com/product/maxcompute",
	}

	ShieldNPM = &ShieldBadge{
		ID:        "NPM",
		Label:     "NPM",
		Color:     "#CB3837",
		Style:     ShieldStyleDefault,
		Logo:      "npm",
		LogoColor: "white",
		Href:      "https://www.npmjs.com/",
	}

	ShieldNestJS = &ShieldBadge{
		ID:        "NestJS",
		Label:     "NestJS",
		Color:     "#E0234E",
		Style:     ShieldStyleDefault,
		Logo:      "nestjs",
		LogoColor: "white",
		Href:      "https://nestjs.com/",
	}

	ShieldNextJS = &ShieldBadge{
		ID:        "Next JS",
		Label:     "Next JS",
		Color:     "black",
		Style:     ShieldStyleDefault,
		Logo:      "next.js",
		LogoColor: "white",
		Href:      "https://nextjs.org/",
	}

	ShieldNodeJS = &ShieldBadge{
		ID:        "Node.js",
		Label:     "NodeJS",
		Color:     "#6DA55F",
		Style:     ShieldStyleDefault,
		Logo:      "node.js",
		LogoColor: "white",
		Href:      "https://nodejs.org/",
	}

	ShieldNodemon = &ShieldBadge{
		ID:        "Nodemon",
		Label:     "Nodemon",
		Color:     "#323330",
		Style:     ShieldStyleDefault,
		Logo:      "nodemon",
		LogoColor: "#BBDEAD",
		Href:      "https://nodemon.io/",
	}

	ShieldNodeRED = &ShieldBadge{
		ID:        "Node-RED",
		Label:     "Node-RED",
		Color:     "#8F0000",
		Style:     ShieldStyleDefault,
		Logo:      "node-red",
		LogoColor: "white",
		Href:      "https://nodered.org/",
	}

	ShieldNuxtJS = &ShieldBadge{
		ID:        "Nuxt JS",
		Label:     "Nuxt JS",
		Color:     "#002E3B",
		Style:     ShieldStyleDefault,
		Logo:      "nuxtdotjs",
		LogoColor: "#00DC82",
		Href:      "https://nuxtjs.org/",
	}

	ShieldNx = &ShieldBadge{
		ID:        "Nx",
		Label:     "Nx",
		Color:     "#143055",
		Style:     ShieldStyleDefault,
		Logo:      "nx",
		LogoColor: "white",
		Href:      "https://nx.dev/",
	}

	ShieldOpenCV = &ShieldBadge{
		ID:        "OpenCV",
		Label:     "OpenCV",
		Color:     "white",
		Style:     ShieldStyleDefault,
		Logo:      "opencv",
		LogoColor: "white",
		Href:      "https://opencv.org/",
	}

	ShieldOpenGL = &ShieldBadge{
		ID:        "OpenGL",
		Label:     "OpenGL",
		Color:     "white",
		Style:     ShieldStyleDefault,
		Logo:      "opengl",
		LogoColor: "white",
		Href:      "https://www.opengl.org/",
	}

	ShieldP5js = &ShieldBadge{
		ID:        "P5js",
		Label:     "P5js",
		Color:     "#ED225D",
		Style:     ShieldStyleDefault,
		Logo:      "p5.js",
		LogoColor: "FFFFFF",
		Href:      "https://p5js.org/",
	}

	ShieldPhoenixFramework = &ShieldBadge{
		ID:        "Phoenix Framework",
		Label:     "Phoenix Framework",
		Color:     "#FD4F00",
		Style:     ShieldStyleDefault,
		Logo:      "phoenixframework",
		LogoColor: "black",
		Href:      "https://www.phoenixframework.org/",
	}

	ShieldPNPM = &ShieldBadge{
		ID:        "PNPM",
		Label:     "PNPM",
		Color:     "#4A4A4A",
		Style:     ShieldStyleDefault,
		Logo:      "pnpm",
		LogoColor: "#F69220",
		Href:      "https://pnpm.io/",
	}

	ShieldPoetry = &ShieldBadge{
		ID:        "Poetry",
		Label:     "Poetry",
		Color:     "#3B82F6",
		Style:     ShieldStyleDefault,
		Logo:      "poetry",
		LogoColor: "#0B3D8D",
		Href:      "https://python-poetry.org/",
	}

	ShieldPrefect = &ShieldBadge{
		ID:        "Prefect",
		Label:     "Prefect",
		Color:     "white",
		Style:     ShieldStyleDefault,
		Logo:      "prefect",
		LogoColor: "white",
		Href:      "https://www.prefect.io/",
	}

	ShieldPug = &ShieldBadge{
		ID:        "Pug",
		Label:     "Pug",
		Color:     "FFF",
		Style:     ShieldStyleDefault,
		Logo:      "pug",
		LogoColor: "#A86454",
		Href:      "https://pugjs.org/",
	}

	ShieldPytest = &ShieldBadge{
		ID:        "Pytest",
		Label:     "Pytest",
		Color:     "white",
		Style:     ShieldStyleDefault,
		Logo:      "pytest",
		LogoColor: "#2F9FE3",
		Href:      "https://pytest.org/",
	}

	ShieldQt = &ShieldBadge{
		ID:        "Qt",
		Label:     "Qt",
		Color:     "#217346",
		Style:     ShieldStyleDefault,
		Logo:      "Qt",
		LogoColor: "white",
		Href:      "https://www.qt.io/",
	}

	ShieldQuarkus = &ShieldBadge{
		ID:        "Quarkus",
		Label:     "Quarkus",
		Color:     "#4794EB",
		Style:     ShieldStyleDefault,
		Logo:      "quarkus",
		LogoColor: "white",
		Href:      "https://quarkus.io/",
	}

	ShieldQuasar = &ShieldBadge{
		ID:        "Quasar",
		Label:     "Quasar",
		Color:     "#16B7FB",
		Style:     ShieldStyleDefault,
		Logo:      "quasar",
		LogoColor: "black",
		Href:      "https://quasar.dev/",
	}

	ShieldROS = &ShieldBadge{
		ID:        "ROS",
		Label:     "ROS",
		Color:     "#0A0FF9",
		Style:     ShieldStyleDefault,
		Logo:      "ros",
		LogoColor: "white",
		Href:      "https://www.ros.org/",
	}

	ShieldRabbitMQ = &ShieldBadge{
		ID:        "RabbitMQ",
		Label:     "RabbitMQ",
		Color:     "#FF6600",
		Style:     ShieldStyleDefault,
		Logo:      "rabbitmq",
		LogoColor: "white",
		Href:      "https://www.rabbitmq.com/",
	}

	ShieldRadixUI = &ShieldBadge{
		ID:        "Radix UI",
		Label:     "Radix UI",
		Color:     "#161618",
		Style:     ShieldStyleDefault,
		Logo:      "radix-ui",
		LogoColor: "white",
		Href:      "https://www.radix-ui.com/",
	}

	ShieldRails = &ShieldBadge{
		ID:        "Rails",
		Label:     "Rails",
		Color:     "#CC0000",
		Style:     ShieldStyleDefault,
		Logo:      "ruby-on-rails",
		LogoColor: "white",
		Href:      "https://rubyonrails.org/",
	}

	ShieldRayLib = &ShieldBadge{
		ID:        "RayLib",
		Label:     "RayLib",
		Color:     "FFFFFF",
		Style:     ShieldStyleDefault,
		Logo:      "raylib",
		LogoColor: "black",
		Href:      "https://www.raylib.com/",
	}

	ShieldReact = &ShieldBadge{
		ID:        "React",
		Label:     "React",
		Color:     "#20232A",
		Style:     ShieldStyleDefault,
		Logo:      "react",
		LogoColor: "#61DAFB",
		Href:      "https://reactjs.org/",
	}

	ShieldReactNative = &ShieldBadge{
		ID:        "React Native",
		Label:     "React Native",
		Color:     "#20232A",
		Style:     ShieldStyleDefault,
		Logo:      "react",
		LogoColor: "#61DAFB",
		Href:      "https://reactnative.dev/",
	}

	ShieldReactQuery = &ShieldBadge{
		ID:        "React Query",
		Label:     "React Query",
		Color:     "#FF4154",
		Style:     ShieldStyleDefault,
		Logo:      "react-query",
		LogoColor: "white",
		Href:      "https://react-query.tanstack.com/",
	}

	ShieldReactRouter = &ShieldBadge{
		ID:        "React Router",
		Label:     "React Router",
		Color:     "#CA4245",
		Style:     ShieldStyleDefault,
		Logo:      "react-router",
		LogoColor: "white",
		Href:      "https://reactrouter.com/",
	}

	ShieldReactHookForm = &ShieldBadge{
		ID:        "React Hook Form",
		Label:     "React Hook Form",
		Color:     "#EC5990",
		Style:     ShieldStyleDefault,
		Logo:      "reacthookform",
		LogoColor: "white",
		Href:      "https://react-hook-form.com/",
	}

	ShieldRedux = &ShieldBadge{
		ID:        "Redux",
		Label:     "Redux",
		Color:     "#593D88",
		Style:     ShieldStyleDefault,
		Logo:      "redux",
		LogoColor: "white",
		Href:      "https://redux.js.org/",
	}

	ShieldRemix = &ShieldBadge{
		ID:        "Remix",
		Label:     "Remix",
		Color:     "black",
		Style:     ShieldStyleDefault,
		Logo:      "remix",
		LogoColor: "white",
		Href:      "https://remix.run/",
	}

	ShieldRollupJS = &ShieldBadge{
		ID:        "RollupJS",
		Label:     "RollupJS",
		Color:     "#EF3335",
		Style:     ShieldStyleDefault,
		Logo:      "rollup.js",
		LogoColor: "white",
		Href:      "https://rollupjs.org/",
	}

	ShieldRxDB = &ShieldBadge{
		ID:        "RxDB",
		Label:     "RxDB",
		Color:     "#B7178C",
		Style:     ShieldStyleDefault,
		Logo:      "reactivex",
		LogoColor: "white",
		Href:      "https://rxdb.info/",
	}

	ShieldRxJS = &ShieldBadge{
		ID:        "RxJS",
		Label:     "RxJS",
		Color:     "#B7178C",
		Style:     ShieldStyleDefault,
		Logo:      "reactivex",
		LogoColor: "white",
		Href:      "https://rxjs.dev/",
	}

	ShieldSASS = &ShieldBadge{
		ID:        "SASS",
		Label:     "SASS",
		Color:     "hotpink",
		Style:     ShieldStyleDefault,
		Logo:      "SASS",
		LogoColor: "white",
		Href:      "https://sass-lang.com/",
	}

	ShieldScrapy = &ShieldBadge{
		ID:        "Scrapy",
		Label:     "Scrapy",
		Color:     "#60A839",
		Style:     ShieldStyleDefault,
		Logo:      "scrapy",
		LogoColor: "#D1D2D3",
		Href:      "https://scrapy.org/",
	}

	ShieldSemanticUIReact = &ShieldBadge{
		ID:        "Semantic UI React",
		Label:     "Semantic UI React",
		Color:     "#35BDB2",
		Style:     ShieldStyleDefault,
		Logo:      "semantic-ui-react",
		LogoColor: "white",
		Href:      "https://react.semantic-ui.com/",
	}

	ShieldSnowflake = &ShieldBadge{
		ID:        "Snowflake",
		Label:     "Snowflake",
		Color:     "#29B5E8",
		Style:     ShieldStyleDefault,
		Logo:      "snowflake",
		LogoColor: "white",
		Href:      "https://www.snowflake.com/",
	}
	ShieldSocketIO = &ShieldBadge{
		ID:        "Socket.io",
		Label:     "Socket.io",
		Color:     "black",
		Style:     ShieldStyleDefault,
		Logo:      "socket.io",
		LogoColor: "white",
		Href:      "https://socket.io/",
	}

	ShieldSolidJS = &ShieldBadge{
		ID:        "SolidJS",
		Label:     "SolidJS",
		Color:     "#2c4f7c",
		Style:     ShieldStyleDefault,
		Logo:      "solid",
		LogoColor: "#c8c9cb",
		Href:      "https://solidjs.com/",
	}

	ShieldSpring = &ShieldBadge{
		ID:        "Spring",
		Label:     "Spring",
		Color:     "#6DB33F",
		Style:     ShieldStyleDefault,
		Logo:      "spring",
		LogoColor: "white",
		Href:      "https://spring.io/",
	}

	ShieldStrapi = &ShieldBadge{
		ID:        "Strapi",
		Label:     "Strapi",
		Color:     "#2E7EEA",
		Style:     ShieldStyleDefault,
		Logo:      "strapi",
		LogoColor: "white",
		Href:      "https://strapi.io/",
	}

	ShieldStreamlit = &ShieldBadge{
		ID:        "Streamlit",
		Label:     "Streamlit",
		Color:     "#FE4B4B",
		Style:     ShieldStyleDefault,
		Logo:      "streamlit",
		LogoColor: "white",
		Href:      "https://streamlit.io/",
	}

	ShieldStyledComponents = &ShieldBadge{
		ID:        "Styled Components",
		Label:     "Styled Components",
		Color:     "#DB7093",
		Style:     ShieldStyleDefault,
		Logo:      "styled-components",
		LogoColor: "white",
		Href:      "https://styled-components.com/",
	}

	ShieldStylus = &ShieldBadge{
		ID:        "Stylus",
		Label:     "Stylus",
		Color:     "#ff6347",
		Style:     ShieldStyleDefault,
		Logo:      "stylus",
		LogoColor: "white",
		Href:      "https://stylus-lang.com/",
	}

	ShieldSvelte = &ShieldBadge{
		ID:        "Svelte",
		Label:     "Svelte",
		Color:     "#f1413d",
		Style:     ShieldStyleDefault,
		Logo:      "svelte",
		LogoColor: "white",
		Href:      "https://svelte.dev/",
	}

	ShieldSvelteKit = &ShieldBadge{
		ID:        "SvelteKit",
		Label:     "SvelteKit",
		Color:     "#f1413d",
		Style:     ShieldStyleDefault,
		Logo:      "svelte",
		LogoColor: "white",
		Href:      "https://kit.svelte.dev/",
	}

	ShieldSymfony = &ShieldBadge{
		ID:        "Symfony",
		Label:     "Symfony",
		Color:     "#000000",
		Style:     ShieldStyleDefault,
		Logo:      "symfony",
		LogoColor: "white",
		Href:      "https://symfony.com/",
	}

	ShieldTailwindCSS = &ShieldBadge{
		ID:        "TailwindCSS",
		Label:     "TailwindCSS",
		Color:     "#38B2AC",
		Style:     ShieldStyleDefault,
		Logo:      "tailwind-css",
		LogoColor: "white",
		Href:      "https://tailwindcss.com/",
	}

	ShieldTauri = &ShieldBadge{
		ID:        "Tauri",
		Label:     "Tauri",
		Color:     "#24C8DB",
		Style:     ShieldStyleDefault,
		Logo:      "tauri",
		LogoColor: "#FFFFFF",
		Href:      "https://tauri.app/",
	}

	ShieldThreeJS = &ShieldBadge{
		ID:        "Three.js",
		Label:     "Three.js",
		Color:     "black",
		Style:     ShieldStyleDefault,
		Logo:      "three.js",
		LogoColor: "white",
		Href:      "https://threejs.org/",
	}

	ShieldThymeleaf = &ShieldBadge{
		ID:        "Thymeleaf",
		Label:     "Thymeleaf",
		Color:     "#005C0F",
		Style:     ShieldStyleDefault,
		Logo:      "thymeleaf",
		LogoColor: "white",
		Href:      "https://www.thymeleaf.org/",
	}

	ShieldTRPC = &ShieldBadge{
		ID:        "tRPC",
		Label:     "tRPC",
		Color:     "#2596BE",
		Style:     ShieldStyleDefault,
		Logo:      "tRPC",
		LogoColor: "white",
		Href:      "https://trpc.io/",
	}

	ShieldTypeGraphQL = &ShieldBadge{
		ID:        "TypeGraphQL",
		Label:     "TypeGraphQL",
		Color:     "#C04392",
		Style:     ShieldStyleDefault,
		Logo:      "type-graphql",
		LogoColor: "white",
		Href:      "https://typegraphql.ml/",
	}

	ShieldUnoCSS = &ShieldBadge{
		ID:        "UnoCSS",
		Label:     "UnoCSS",
		Color:     "#333333",
		Style:     ShieldStyleDefault,
		Logo:      "unocss",
		LogoColor: "white",
		Href:      "https://unocss.dev/",
	}

	ShieldVite = &ShieldBadge{
		ID:        "Vite",
		Label:     "Vite",
		Color:     "#646CFF",
		Style:     ShieldStyleDefault,
		Logo:      "vite",
		LogoColor: "white",
		Href:      "https://vitejs.dev/",
	}

	ShieldVueJS = &ShieldBadge{
		ID:        "Vue.js",
		Label:     "Vue.js",
		Color:     "#35495e",
		Style:     ShieldStyleDefault,
		Logo:      "vuedotjs",
		LogoColor: "#4FC08D",
		Href:      "https://vuejs.org/",
	}

	ShieldVuetify = &ShieldBadge{
		ID:        "Vuetify",
		Label:     "Vuetify",
		Color:     "#1867C0",
		Style:     ShieldStyleDefault,
		Logo:      "vuetify",
		LogoColor: "AEDDFF",
		Href:      "https://vuetifyjs.com/",
	}

	ShieldWebGL = &ShieldBadge{
		ID:        "WebGL",
		Label:     "WebGL",
		Color:     "#990000",
		Style:     ShieldStyleDefault,
		Logo:      "webgl",
		LogoColor: "white",
		Href:      "https://www.khronos.org/webgl/",
	}

	ShieldWebpack = &ShieldBadge{
		ID:        "Webpack",
		Label:     "Webpack",
		Color:     "#8DD6F9",
		Style:     ShieldStyleDefault,
		Logo:      "webpack",
		LogoColor: "black",
		Href:      "https://webpack.js.org/",
	}

	ShieldWeb3JS = &ShieldBadge{
		ID:        "Web3.js",
		Label:     "Web3.js",
		Color:     "#F16822",
		Style:     ShieldStyleDefault,
		Logo:      "web3.js",
		LogoColor: "white",
		Href:      "https://web3js.org/",
	}

	ShieldWindiCSS = &ShieldBadge{
		ID:        "WindiCSS",
		Label:     "WindiCSS",
		Color:     "#48B0F1",
		Style:     ShieldStyleDefault,
		Logo:      "windi-css",
		LogoColor: "white",
		Href:      "https://windicss.org/",
	}

	ShieldWordPress = &ShieldBadge{
		ID:        "WordPress",
		Label:     "WordPress",
		Color:     "#117AC9",
		Style:     ShieldStyleDefault,
		Logo:      "WordPress",
		LogoColor: "white",
		Href:      "https://wordpress.org/",
	}

	ShieldXamarin = &ShieldBadge{
		ID:        "Xamarin",
		Label:     "Xamarin",
		Color:     "#3199DC",
		Style:     ShieldStyleDefault,
		Logo:      "xamarin",
		LogoColor: "white",
		Href:      "https://dotnet.microsoft.com/apps/xamarin",
	}

	ShieldYarn = &ShieldBadge{
		ID:        "Yarn",
		Label:     "Yarn",
		Color:     "#2C8EBB",
		Style:     ShieldStyleDefault,
		Logo:      "yarn",
		LogoColor: "white",
		Href:      "https://yarnpkg.com/",
	}

	ShieldZod = &ShieldBadge{
		ID:        "Zod",
		Label:     "Zod",
		Color:     "#3068b7",
		Style:     ShieldStyleDefault,
		Logo:      "zod",
		LogoColor: "white",
		Href:      "https://zod.dev/",
	}

	ShieldGroovy = &ShieldBadge{
		ID:        "ApacheGroovy",
		Label:     "Apache Groovy",
		Color:     "#4298B8",
		Style:     ShieldStyleDefault,
		Logo:      "Apache Groovy",
		LogoColor: "white",
		Href:      "https://groovy-lang.org/",
	}

	ShieldAssemblyScript = &ShieldBadge{
		ID:        "AssemblyScript",
		Label:     "AssemblyScript",
		Color:     "#000000",
		Style:     ShieldStyleDefault,
		Logo:      "assemblyscript",
		LogoColor: "white",
		Href:      "https://www.assemblyscript.org/",
	}

	ShieldC = &ShieldBadge{
		ID:        "C",
		Label:     "C",
		Color:     "#00599C",
		Style:     ShieldStyleDefault,
		Logo:      "c",
		LogoColor: "white",
		Href:      "https://en.wikipedia.org/wiki/C_(programming_language)",
	}

	ShieldCSharp = &ShieldBadge{
		ID:        "CSharp",
		Label:     "C#",
		Color:     "#239120",
		Style:     ShieldStyleDefault,
		Logo:      "csharp",
		LogoColor: "white",
		Href:      "https://learn.microsoft.com/en-us/dotnet/csharp/",
	}

	ShieldCpp = &ShieldBadge{
		ID:        "C++",
		Label:     "C++",
		Color:     "#00599C",
		Style:     ShieldStyleDefault,
		Logo:      "c++",
		LogoColor: "white",
		Href:      "https://isocpp.org/",
	}

	ShieldClojure = &ShieldBadge{
		ID:        "Clojure",
		Label:     "Clojure",
		Color:     "#Clojure",
		Style:     ShieldStyleDefault,
		Logo:      "clojure",
		LogoColor: "Clojure",
		Href:      "https://clojure.org/",
	}

	ShieldCrystal = &ShieldBadge{
		ID:        "Crystal",
		Label:     "Crystal",
		Color:     "#000000",
		Style:     ShieldStyleDefault,
		Logo:      "crystal",
		LogoColor: "white",
		Href:      "https://crystal-lang.org/",
	}

	ShieldCSS3 = &ShieldBadge{
		ID:        "CSS3",
		Label:     "CSS3",
		Color:     "#1572B6",
		Style:     ShieldStyleDefault,
		Logo:      "css3",
		LogoColor: "white",
		Href:      "https://www.w3.org/Style/CSS/",
	}

	ShieldDart = &ShieldBadge{
		ID:        "Dart",
		Label:     "Dart",
		Color:     "#0175C2",
		Style:     ShieldStyleDefault,
		Logo:      "dart",
		LogoColor: "white",
		Href:      "https://dart.dev/",
	}

	ShieldDgraph = &ShieldBadge{
		ID:        "Dgraph",
		Label:     "Dgraph",
		Color:     "#E50695",
		Style:     ShieldStyleDefault,
		Logo:      "dgraph",
		LogoColor: "white",
		Href:      "https://dgraph.io/",
	}

	ShieldElixir = &ShieldBadge{
		ID:        "Elixir",
		Label:     "Elixir",
		Color:     "#4B275F",
		Style:     ShieldStyleDefault,
		Logo:      "elixir",
		LogoColor: "white",
		Href:      "https://elixir-lang.org/",
	}

	ShieldElm = &ShieldBadge{
		ID:        "Elm",
		Label:     "Elm",
		Color:     "#60B5CC",
		Style:     ShieldStyleDefault,
		Logo:      "elm",
		LogoColor: "white",
		Href:      "https://elm-lang.org/",
	}

	ShieldErlang = &ShieldBadge{
		ID:        "Erlang",
		Label:     "Erlang",
		Color:     "white",
		Style:     ShieldStyleDefault,
		Logo:      "erlang",
		LogoColor: "#A90533",
		Href:      "https://www.erlang.org/",
	}

	ShieldFortran = &ShieldBadge{
		ID:        "Fortran",
		Label:     "Fortran",
		Color:     "#734F96",
		Style:     ShieldStyleDefault,
		Logo:      "fortran",
		LogoColor: "white",
		Href:      "https://fortran-lang.org/",
	}

	ShieldGDScript = &ShieldBadge{
		ID:        "GDScript",
		Label:     "GDScript",
		Color:     "#74267B",
		Style:     ShieldStyleDefault,
		Logo:      "godotengine",
		LogoColor: "white",
		Href:      "https://godotengine.org/",
	}

	ShieldGo = &ShieldBadge{
		ID:        "Go",
		Label:     "Go",
		Color:     "#00ADD8",
		Style:     ShieldStyleDefault,
		Logo:      "go",
		LogoColor: "white",
		Href:      "https://golang.org/",
	}

	ShieldGraphQL = &ShieldBadge{
		ID:        "GraphQL",
		Label:     "GraphQL",
		Color:     "#E10098",
		Style:     ShieldStyleDefault,
		Logo:      "graphql",
		LogoColor: "white",
		Href:      "https://graphql.org/",
	}

	ShieldHaskell = &ShieldBadge{
		ID:        "Haskell",
		Label:     "Haskell",
		Color:     "#5e5086",
		Style:     ShieldStyleDefault,
		Logo:      "haskell",
		LogoColor: "white",
		Href:      "https://www.haskell.org/",
	}

	ShieldHTML5 = &ShieldBadge{
		ID:        "HTML5",
		Label:     "HTML5",
		Color:     "#E34F26",
		Style:     ShieldStyleDefault,
		Logo:      "html5",
		LogoColor: "white",
		Href:      "https://www.w3.org/Style/CSS/",
	}

	ShieldJava = &ShieldBadge{
		ID:        "Java",
		Label:     "Java",
		Color:     "#ED8B00",
		Style:     ShieldStyleDefault,
		Logo:      "openjdk",
		LogoColor: "white",
		Href:      "https://www.java.com/",
	}

	ShieldJavaScript = &ShieldBadge{
		ID:        "JavaScript",
		Label:     "JavaScript",
		Color:     "#323330",
		Style:     ShieldStyleDefault,
		Logo:      "javascript",
		LogoColor: "#F7DF1E",
		Href:      "https://developer.mozilla.org/en-US/docs/Web/JavaScript",
	}

	ShieldJulia = &ShieldBadge{
		ID:        "Julia",
		Label:     "Julia",
		Color:     "#9558B2",
		Style:     ShieldStyleDefault,
		Logo:      "julia",
		LogoColor: "white",
		Href:      "https://julialang.org/",
	}

	ShieldKotlin = &ShieldBadge{
		ID:        "Kotlin",
		Label:     "Kotlin",
		Color:     "#7F52FF",
		Style:     ShieldStyleDefault,
		Logo:      "kotlin",
		LogoColor: "white",
		Href:      "https://kotlinlang.org/",
	}

	ShieldLaTeX = &ShieldBadge{
		ID:        "LaTeX",
		Label:     "LaTeX",
		Color:     "#008080",
		Style:     ShieldStyleDefault,
		Logo:      "latex",
		LogoColor: "white",
		Href:      "https://www.latex-project.org/",
	}

	ShieldLua = &ShieldBadge{
		ID:        "Lua",
		Label:     "Lua",
		Color:     "#2C2D72",
		Style:     ShieldStyleDefault,
		Logo:      "lua",
		LogoColor: "white",
		Href:      "https://www.lua.org/",
	}

	ShieldMarkdown = &ShieldBadge{
		ID:        "Markdown",
		Label:     "Markdown",
		Color:     "#000000",
		Style:     ShieldStyleDefault,
		Logo:      "markdown",
		LogoColor: "white",
		Href:      "https://www.markdownguide.org/",
	}

	ShieldNim = &ShieldBadge{
		ID:        "Nim",
		Label:     "Nim",
		Color:     "#FFE953",
		Style:     ShieldStyleDefault,
		Logo:      "nim",
		LogoColor: "white",
		Href:      "https://nim-lang.org/",
	}

	ShieldNix = &ShieldBadge{
		ID:        "Nix",
		Label:     "Nix",
		Color:     "#5277C3",
		Style:     ShieldStyleDefault,
		Logo:      "nixos",
		LogoColor: "white",
		Href:      "https://nixos.org/",
	}

	ShieldObjectiveC = &ShieldBadge{
		ID:        "ObjectiveC",
		Label:     "Objective-C",
		Color:     "#3A95E3",
		Style:     ShieldStyleDefault,
		Logo:      "apple",
		LogoColor: "white",
		Href:      "https://developer.apple.com/documentation/objectivec",
	}

	ShieldOCaml = &ShieldBadge{
		ID:        "OCaml",
		Label:     "OCaml",
		Color:     "#E98407",
		Style:     ShieldStyleDefault,
		Logo:      "ocaml",
		LogoColor: "white",
		Href:      "https://ocaml.org/",
	}

	ShieldOctave = &ShieldBadge{
		ID:        "Octave",
		Label:     "Octave",
		Color:     "darkblue",
		Style:     ShieldStyleDefault,
		Logo:      "octave",
		LogoColor: "fcd683",
		Href:      "https://www.gnu.org/software/octave/",
	}

	ShieldOrgMode = &ShieldBadge{
		ID:        "OrgMode",
		Label:     "Org Mode",
		Color:     "#77AA99",
		Style:     ShieldStyleDefault,
		Logo:      "org",
		LogoColor: "white",
		Href:      "https://orgmode.org/",
	}

	ShieldPerl = &ShieldBadge{
		ID:        "Perl",
		Label:     "Perl",
		Color:     "#39457E",
		Style:     ShieldStyleDefault,
		Logo:      "perl",
		LogoColor: "white",
		Href:      "https://www.perl.org/",
	}

	ShieldPHP = &ShieldBadge{
		ID:        "PHP",
		Label:     "PHP",
		Color:     "#777BB4",
		Style:     ShieldStyleDefault,
		Logo:      "php",
		LogoColor: "white",
		Href:      "https://www.php.net/",
	}

	ShieldPowerShell = &ShieldBadge{
		ID:        "PowerShell",
		Label:     "PowerShell",
		Color:     "#5391FE",
		Style:     ShieldStyleDefault,
		Logo:      "powershell",
		LogoColor: "white",
		Href:      "https://learn.microsoft.com/en-us/powershell/",
	}

	ShieldPython = &ShieldBadge{
		ID:        "Python",
		Label:     "Python",
		Color:     "#3670A0",
		Style:     ShieldStyleDefault,
		Logo:      "python",
		LogoColor: "#ffdd54",
		Href:      "https://www.python.org/",
	}

	ShieldR = &ShieldBadge{
		ID:        "R",
		Label:     "R",
		Color:     "#276DC3",
		Style:     ShieldStyleDefault,
		Logo:      "r",
		LogoColor: "white",
		Href:      "https://www.r-project.org/",
	}

	ShieldReScript = &ShieldBadge{
		ID:        "ReScript",
		Label:     "ReScript",
		Color:     "#14162c",
		Style:     ShieldStyleDefault,
		Logo:      "rescript",
		LogoColor: "#e34c4c",
		Href:      "https://rescript-lang.org/",
	}

	ShieldRuby = &ShieldBadge{
		ID:        "Ruby",
		Label:     "Ruby",
		Color:     "#CC342D",
		Style:     ShieldStyleDefault,
		Logo:      "ruby",
		LogoColor: "white",
		Href:      "https://www.ruby-lang.org/",
	}

	ShieldRust = &ShieldBadge{
		ID:        "Rust",
		Label:     "Rust",
		Color:     "#000000",
		Style:     ShieldStyleDefault,
		Logo:      "rust",
		LogoColor: "white",
		Href:      "https://www.rust-lang.org/",
	}

	ShieldScala = &ShieldBadge{
		ID:        "Scala",
		Label:     "Scala",
		Color:     "#DC322F",
		Style:     ShieldStyleDefault,
		Logo:      "scala",
		LogoColor: "white",
		Href:      "https://www.scala-lang.org/",
	}

	ShieldBashScript = &ShieldBadge{
		ID:        "BashScript",
		Label:     "Bash Script",
		Color:     "#121011",
		Style:     ShieldStyleDefault,
		Logo:      "gnu-bash",
		LogoColor: "white",
		Href:      "https://www.gnu.org/software/bash/",
	}

	ShieldSolidity = &ShieldBadge{
		ID:        "Solidity",
		Label:     "Solidity",
		Color:     "#363636",
		Style:     ShieldStyleDefault,
		Logo:      "solidity",
		LogoColor: "white",
		Href:      "https://soliditylang.org/",
	}

	ShieldSwift = &ShieldBadge{
		ID:        "Swift",
		Label:     "Swift",
		Color:     "#F54A2A",
		Style:     ShieldStyleDefault,
		Logo:      "swift",
		LogoColor: "white",
		Href:      "https://swift.org/",
	}

	ShieldTypeScript = &ShieldBadge{
		ID:        "TypeScript",
		Label:     "TypeScript",
		Color:     "#007ACC",
		Style:     ShieldStyleDefault,
		Logo:      "typescript",
		LogoColor: "white",
		Href:      "https://www.typescriptlang.org/",
	}

	ShieldWindowsTerminal = &ShieldBadge{
		ID:        "WindowsTerminal",
		Label:     "Windows Terminal",
		Color:     "#4D4D4D",
		Style:     ShieldStyleDefault,
		Logo:      "windows-terminal",
		LogoColor: "white",
		Href:      "https://aka.ms/terminal",
	}

	ShieldYAML = &ShieldBadge{
		ID:        "YAML",
		Label:     "YAML",
		Color:     "#ffffff",
		Style:     ShieldStyleDefault,
		Logo:      "yaml",
		LogoColor: "#151515",
		Href:      "https://yaml.org/",
	}

	ShieldZig = &ShieldBadge{
		ID:        "Zig",
		Label:     "Zig",
		Color:     "#F7A41D",
		Style:     ShieldStyleDefault,
		Logo:      "zig",
		LogoColor: "white",
		Href:      "https://ziglang.org/",
	}

	ShieldKeras = &ShieldBadge{
		ID:        "Keras",
		Label:     "Keras",
		Color:     "#D00000",
		Style:     ShieldStyleDefault,
		Logo:      "keras",
		LogoColor: "white",
		Href:      "https://keras.io/",
	}

	ShieldMatplotlib = &ShieldBadge{
		ID:        "Matplotlib",
		Label:     "Matplotlib",
		Color:     "#ffffff",
		Style:     ShieldStyleDefault,
		Logo:      "matplotlib",
		LogoColor: "black",
		Href:      "https://matplotlib.org/",
	}

	ShieldMlflow = &ShieldBadge{
		ID:        "mlflow",
		Label:     "mlflow",
		Color:     "#d9ead3",
		Style:     ShieldStyleDefault,
		Logo:      "numpy",
		LogoColor: "blue",
		Href:      "https://mlflow.org/",
	}

	ShieldNumPy = &ShieldBadge{
		ID:        "NumPy",
		Label:     "NumPy",
		Color:     "#013243",
		Style:     ShieldStyleDefault,
		Logo:      "numpy",
		LogoColor: "white",
		Href:      "https://numpy.org/",
	}

	ShieldPandas = &ShieldBadge{
		ID:        "Pandas",
		Label:     "Pandas",
		Color:     "#150458",
		Style:     ShieldStyleDefault,
		Logo:      "pandas",
		LogoColor: "white",
		Href:      "https://pandas.pydata.org/",
	}

	ShieldPlotly = &ShieldBadge{
		ID:        "Plotly",
		Label:     "Plotly",
		Color:     "#3F4F75",
		Style:     ShieldStyleDefault,
		Logo:      "plotly",
		LogoColor: "white",
		Href:      "https://plotly.com/",
	}

	ShieldPyTorch = &ShieldBadge{
		ID:        "PyTorch",
		Label:     "PyTorch",
		Color:     "#EE4C2C",
		Style:     ShieldStyleDefault,
		Logo:      "pytorch",
		LogoColor: "white",
		Href:      "https://pytorch.org/",
	}

	ShieldScikitLearn = &ShieldBadge{
		ID:        "scikit-learn",
		Label:     "scikit-learn",
		Color:     "#F7931E",
		Style:     ShieldStyleDefault,
		Logo:      "scikit-learn",
		LogoColor: "white",
		Href:      "https://scikit-learn.org/",
	}

	ShieldSciPy = &ShieldBadge{
		ID:        "SciPy",
		Label:     "SciPy",
		Color:     "#0C55A5",
		Style:     ShieldStyleDefault,
		Logo:      "scipy",
		LogoColor: "white",
		Href:      "https://scipy.org/",
	}

	ShieldTensorFlow = &ShieldBadge{
		ID:        "TensorFlow",
		Label:     "TensorFlow",
		Color:     "#FF6F00",
		Style:     ShieldStyleDefault,
		Logo:      "tensorflow",
		LogoColor: "white",
		Href:      "https://www.tensorflow.org/",
	}

	ShieldHibernate = &ShieldBadge{
		ID:        "Hibernate",
		Label:     "Hibernate",
		Color:     "#59666C",
		Style:     ShieldStyleDefault,
		Logo:      "hibernate",
		LogoColor: "white",
		Href:      "https://hibernate.org/",
	}

	ShieldPrisma = &ShieldBadge{
		ID:        "Prisma",
		Label:     "Prisma",
		Color:     "#3982CE",
		Style:     ShieldStyleDefault,
		Logo:      "prisma",
		LogoColor: "white",
		Href:      "https://www.prisma.io/",
	}

	ShieldSequelize = &ShieldBadge{
		ID:        "Sequelize",
		Label:     "Sequelize",
		Color:     "#52B0E7",
		Style:     ShieldStyleDefault,
		Logo:      "sequelize",
		LogoColor: "white",
		Href:      "https://sequelize.org/",
	}

	ShieldTypeORM = &ShieldBadge{
		ID:        "TypeORM",
		Label:     "TypeORM",
		Color:     "#FE0803",
		Style:     ShieldStyleDefault,
		Logo:      "typeorm",
		LogoColor: "white",
		Href:      "https://typeorm.io/",
	}

	ShieldQuill = &ShieldBadge{
		ID:        "Quill",
		Label:     "Quill",
		Color:     "#52B0E7",
		Style:     ShieldStyleDefault,
		Logo:      "apache",
		LogoColor: "white",
		Href:      "https://quilljs.com/",
	}

	ShieldApache = &ShieldBadge{
		ID:        "Apache",
		Label:     "Apache",
		Color:     "#D42029",
		Style:     ShieldStyleDefault,
		Logo:      "apache",
		LogoColor: "white",
		Href:      "https://apache.org/",
	}

	ShieldApacheAirflow = &ShieldBadge{
		ID:        "Apache Airflow",
		Label:     "Apache Airflow",
		Color:     "#017CEE",
		Style:     ShieldStyleDefault,
		Logo:      "apache-airflow",
		LogoColor: "white",
		Href:      "https://airflow.apache.org/",
	}

	ShieldApacheAnt = &ShieldBadge{
		ID:        "Apache Ant",
		Label:     "Apache Ant",
		Color:     "#A81C7D",
		Style:     ShieldStyleDefault,
		Logo:      "apache-ant",
		LogoColor: "white",
		Href:      "https://ant.apache.org/",
	}

	ShieldApacheFlink = &ShieldBadge{
		ID:        "Apache Flink",
		Label:     "Apache Flink",
		Color:     "#E6526F",
		Style:     ShieldStyleDefault,
		Logo:      "apache-flink",
		LogoColor: "white",
		Href:      "https://flink.apache.org/",
	}

	ShieldApacheMaven = &ShieldBadge{
		ID:        "Apache Maven",
		Label:     "Apache Maven",
		Color:     "#C71A36",
		Style:     ShieldStyleDefault,
		Logo:      "apache-maven",
		LogoColor: "white",
		Href:      "https://maven.apache.org/",
	}

	ShieldApacheTomcat = &ShieldBadge{
		ID:        "Apache Tomcat",
		Label:     "Apache Tomcat",
		Color:     "#F8DC75",
		Style:     ShieldStyleDefault,
		Logo:      "apache-tomcat",
		LogoColor: "black",
		Href:      "https://tomcat.apache.org/",
	}

	ShieldGunicorn = &ShieldBadge{
		ID:        "Gunicorn",
		Label:     "Gunicorn",
		Color:     "#892729",
		Style:     ShieldStyleDefault,
		Logo:      "gunicorn",
		LogoColor: "white",
		Href:      "https://gunicorn.org/",
	}

	ShieldJenkins = &ShieldBadge{
		ID:        "Jenkins",
		Label:     "Jenkins",
		Color:     "#2C5263",
		Style:     ShieldStyleDefault,
		Logo:      "jenkins",
		LogoColor: "white",
		Href:      "https://www.jenkins.io/",
	}

	ShieldNginx = &ShieldBadge{
		ID:        "Nginx",
		Label:     "Nginx",
		Color:     "#009639",
		Style:     ShieldStyleDefault,
		Logo:      "nginx",
		LogoColor: "white",
		Href:      "https://nginx.org/",
	}

	ShieldApacheSubversion = &ShieldBadge{
		ID:        "Apache Subversion",
		Label:     "Apache Subversion",
		Color:     "#809CC9",
		Style:     ShieldStyleDefault,
		Logo:      "subversion",
		LogoColor: "white",
		Href:      "https://subversion.apache.org/",
	}

	ShieldBitbucket = &ShieldBadge{
		ID:        "Bitbucket",
		Label:     "Bitbucket",
		Color:     "#0047B3",
		Style:     ShieldStyleDefault,
		Logo:      "bitbucket",
		LogoColor: "white",
		Href:      "https://bitbucket.org/",
	}

	ShieldForgejo = &ShieldBadge{
		ID:        "Forgejo",
		Label:     "Forgejo",
		Color:     "#FB923C",
		Style:     ShieldStyleDefault,
		Logo:      "forgejo",
		LogoColor: "white",
		Href:      "https://forgejo.org/",
	}

	ShieldGit = &ShieldBadge{
		ID:        "Git",
		Label:     "Git",
		Color:     "#F05033",
		Style:     ShieldStyleDefault,
		Logo:      "git",
		LogoColor: "white",
		Href:      "https://git-scm.com/",
	}

	ShieldGitea = &ShieldBadge{
		ID:        "Gitea",
		Label:     "Gitea",
		Color:     "#34495E",
		Style:     ShieldStyleDefault,
		Logo:      "gitea",
		LogoColor: "#5D9425",
		Href:      "https://gitea.io/",
	}

	ShieldGitee = &ShieldBadge{
		ID:        "Gitee",
		Label:     "Gitee",
		Color:     "#C71D23",
		Style:     ShieldStyleDefault,
		Logo:      "gitee",
		LogoColor: "white",
		Href:      "https://gitee.com/",
	}

	ShieldGitHub = &ShieldBadge{
		ID:        "GitHub",
		Label:     "GitHub",
		Color:     "#121011",
		Style:     ShieldStyleDefault,
		Logo:      "github",
		LogoColor: "white",
		Href:      "https://github.com/",
	}

	ShieldGitLab = &ShieldBadge{
		ID:        "GitLab",
		Label:     "GitLab",
		Color:     "#181717",
		Style:     ShieldStyleDefault,
		Logo:      "gitlab",
		LogoColor: "white",
		Href:      "https://gitlab.com/",
	}

	ShieldGitpod = &ShieldBadge{
		ID:        "Gitpod",
		Label:     "Gitpod",
		Color:     "#F06611",
		Style:     ShieldStyleDefault,
		Logo:      "gitpod",
		LogoColor: "white",
		Href:      "https://www.gitpod.io/",
	}

	ShieldMercurial = &ShieldBadge{
		ID:        "Mercurial",
		Label:     "Mercurial",
		Color:     "#999999",
		Style:     ShieldStyleDefault,
		Logo:      "mercurial",
		LogoColor: "white",
		Href:      "https://www.mercurial-scm.org/",
	}

	ShieldPerforceHelix = &ShieldBadge{
		ID:        "Perforce Helix",
		Label:     "Perforce Helix",
		Color:     "#00AEEF",
		Style:     ShieldStyleDefault,
		Logo:      "perforce",
		LogoColor: "white",
		Href:      "https://www.perforce.com/",
	}

	ShieldVisualStudioCode = &ShieldBadge{
		ID:        "Visual Studio Code",
		Label:     "Visual Studio Code",
		Color:     "#0078d7",
		Style:     ShieldStyleDefault,
		Logo:      "visual-studio-code",
		LogoColor: "white",
		Href:      "https://code.visualstudio.com/",
	}
)
