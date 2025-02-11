// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
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
		"less":                    walk.DependencyVersionBadge{Badge: badge.ShieldLess},
		"nuxt":                    walk.DependencyVersionBadge{Badge: badge.ShieldNuxtJS},
		"nx":                      walk.DependencyVersionBadge{Badge: badge.ShieldNx},
		"opencv4nodejs":           walk.SystemVersionBadge{Badge: badge.ShieldOpenCV},
		"gl":                      walk.SystemVersionBadge{Badge: badge.ShieldOpenGL},
		"p5":                      walk.DependencyVersionBadge{Badge: badge.ShieldP5js},
		"pnpm":                    walk.DependencyVersionBadge{Badge: badge.ShieldPNPM},
		"quasar":                  walk.DependencyVersionBadge{Badge: badge.ShieldQuasar},
		"react":                   walk.DependencyVersionBadge{Badge: badge.ShieldReact},
		"amqplib":                 walk.SystemVersionBadge{Badge: badge.ShieldRabbitMQ},
		"react-native":            walk.DependencyVersionBadge{Badge: badge.ShieldReactNative},
		"react-query":             walk.DependencyVersionBadge{Badge: badge.ShieldReactQuery},
		"react-router-dom":        walk.DependencyVersionBadge{Badge: badge.ShieldReactRouter},
		"react-hook-form":         walk.DependencyVersionBadge{Badge: badge.ShieldReactHookForm},
		"redux":                   walk.DependencyVersionBadge{Badge: badge.ShieldRedux},
		"remix":                   walk.DependencyVersionBadge{Badge: badge.ShieldRemix},
		"rollup":                  walk.DependencyVersionBadge{Badge: badge.ShieldRollupJS},
		"rxdb":                    walk.DependencyVersionBadge{Badge: badge.ShieldRxDB},
		"rxjs":                    walk.DependencyVersionBadge{Badge: badge.ShieldRxJS},
		"semantic-ui-react":       walk.DependencyVersionBadge{Badge: badge.ShieldSemanticUIReact},
		"snowflake-sdk":           walk.DependencyVersionBadge{Badge: badge.ShieldSnowflake},
		"socket.io":               walk.DependencyVersionBadge{Badge: badge.ShieldSocketIO},
		"styled-components":       walk.DependencyVersionBadge{Badge: badge.ShieldStyledComponents},
		"stylus":                  walk.DependencyVersionBadge{Badge: badge.ShieldStylus},
		"svelte":                  walk.DependencyVersionBadge{Badge: badge.ShieldSvelte},
		"@sveltejs/kit":           walk.DependencyVersionBadge{Badge: badge.ShieldSvelteKit},
		"tailwindcss":             walk.DependencyVersionBadge{Badge: badge.ShieldTailwindCSS},
		"@tauri-apps/api":         walk.DependencyVersionBadge{Badge: badge.ShieldTauri},
		"three":                   walk.DependencyVersionBadge{Badge: badge.ShieldThreeJS},
		"type-graphql":            walk.DependencyVersionBadge{Badge: badge.ShieldTypeGraphQL},
		"unocss":                  walk.DependencyVersionBadge{Badge: badge.ShieldUnoCSS},
		"vuetify":                 walk.DependencyVersionBadge{Badge: badge.ShieldVuetify},
		"webpack":                 walk.DependencyVersionBadge{Badge: badge.ShieldWebpack},
		"web3":                    walk.DependencyVersionBadge{Badge: badge.ShieldWeb3JS},
		"windicss":                walk.DependencyVersionBadge{Badge: badge.ShieldWindiCSS},
	},
	Fuzzy: walk.ResolverPattern{
		"@angular":      walk.DependencyVersionBadge{Badge: badge.ShieldAngular},
		"aurelia":       walk.DependencyVersionBadge{Badge: badge.ShieldAurelia},
		"elasticsearch": walk.SystemVersionBadge{Badge: badge.ShieldElasticSearch},
		"Framework7":    walk.DependencyVersionBadge{Badge: badge.ShieldFramework7},
		"@ionic":        walk.DependencyVersionBadge{Badge: badge.ShieldIonic},
		"@mui":          walk.DependencyVersionBadge{Badge: badge.ShieldMUI},
		"meteor":        walk.DependencyVersionBadge{Badge: badge.ShieldMeteorJS},
		"@mantine":      walk.DependencyVersionBadge{Badge: badge.ShieldMantine},
		"@radix-ui":     walk.SystemVersionBadge{Badge: badge.ShieldRadixUI},
	},
	Partial: walk.ResolverPattern{
		"@trpc": walk.DependencyVersionBadge{Badge: badge.ShieldTRPC},
	},
}

// Nodemon
// Node-RED
// strapi
// WebGL
