package jswalk

import (
	"docwiz/internal/badge"
	"docwiz/internal/walk"
)

var shieldJSResolver = &walk.DependencyResolver{
	Full: walk.ResolverPattern{
		"bun":                     walk.DependencyVersionBadge{Badge: badge.ShieldBun},
		"chart.js":                walk.DependencyVersionBadge{Badge: badge.ShieldChartJS},
		"ejs":                     walk.DependencyVersionBadge{Badge: badge.ShieldEJS},
		"vue":                     walk.DependencyVersionBadge{Badge: badge.ShieldVueJS},
		"vite":                    walk.DependencyVersionBadge{Badge: badge.ShieldVite},
		"sass":                    walk.DependencyVersionBadge{Badge: badge.ShieldSASS},
		"antd":                    walk.DependencyVersionBadge{Badge: badge.ShieldAntDesign},
		"@adonisjs/core":          walk.DependencyVersionBadge{Badge: badge.ShieldAdonisJS},
		"alpinejs":                walk.DependencyVersionBadge{Badge: badge.ShieldAlpineJS},
		"@apollo/client":          walk.DependencyVersionBadge{Badge: badge.ShieldApolloGraphQL},
		"apollo-server":           walk.DependencyVersionBadge{Badge: badge.ShieldApolloGraphQL},
		"astro":                   walk.DependencyVersionBadge{Badge: badge.ShieldAstro},
		"bootstrap":               walk.DependencyVersionBadge{Badge: badge.ShieldBootstrap},
		"Buefy":                   walk.DependencyVersionBadge{Badge: badge.ShieldBuefy},
		"bulma":                   walk.DependencyVersionBadge{Badge: badge.ShieldBulma},
		"@chakra-ui/react":        walk.DependencyVersionBadge{Badge: badge.ShieldChakraUI},
		"daisyui":                 walk.DependencyVersionBadge{Badge: badge.ShieldDaisyUI},
		"directus":                walk.DependencyVersionBadge{Badge: badge.ShieldDirectus},
		"electron":                walk.DependencyVersionBadge{Badge: badge.ShieldElectronJS},
		"ember-source":            walk.DependencyVersionBadge{Badge: badge.ShieldEmber},
		"esbuild":                 walk.DependencyVersionBadge{Badge: badge.ShieldEsbuild},
		"expo":                    walk.DependencyVersionBadge{Badge: badge.ShieldExpo},
		"express":                 walk.DependencyVersionBadge{Badge: badge.ShieldExpressJS},
		"fastify":                 walk.DependencyVersionBadge{Badge: badge.ShieldFastify},
		"gatsby":                  walk.DependencyVersionBadge{Badge: badge.ShieldGatsby},
		"gsap":                    walk.DependencyVersionBadge{Badge: badge.ShieldGreenSock},
		"gulp":                    walk.DependencyVersionBadge{Badge: badge.ShieldGulp},
		"@wordpress/block-editor": walk.DependencyVersionBadge{Badge: badge.ShieldGutenberg},
		"handlebars":              walk.DependencyVersionBadge{Badge: badge.ShieldHandlebars},
		"jasmine":                 walk.DependencyVersionBadge{Badge: badge.ShieldJasmine},
		"jquery":                  walk.DependencyVersionBadge{Badge: badge.ShieldJQuery},
		"zod":                     walk.DependencyVersionBadge{Badge: badge.ShieldZod},
	},
	Fuzzy: walk.ResolverPattern{
		"@angular":      walk.DependencyVersionBadge{Badge: badge.ShieldAngular},
		"aurelia":       walk.DependencyVersionBadge{Badge: badge.ShieldAurelia},
		"elasticsearch": walk.SystemVersionBadge{Badge: badge.ShieldElasticSearch},
		"Framework7":    walk.DependencyVersionBadge{Badge: badge.ShieldFramework7},
		"@ionic":        walk.DependencyVersionBadge{Badge: badge.ShieldIonic},
	},
}
