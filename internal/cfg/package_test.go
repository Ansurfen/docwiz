// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testPackageJSON = `
{
	"name": "TestApp",
	"version": "1.0.0",
	"description": "A test application",
	"main": "index.js",
	"bin": {"testApp": "bin/testApp"},
	"dependencies": {
		"lodash": "4.17.21",
		"express": "4.17.1"
	},
	"devDependencies": {
		"jest": "26.6.0"
	},
	"scripts": {
		"start": "node index.js"
	},
	"repository": {
		"type": "git",
		"url": "https://github.com/test/testapp.git"
	},
	"author": "Jane Doe",
	"license": "MIT",
	"homepage": "https://example.com",
	"engines": {
		"npm": "7.x",
		"node": "14.x"
	}
}
`

func TestLoadPackageJSONFromString(t *testing.T) {
	// Test parsing JSON string into PackageJSON struct
	pkg, err := LoadPackageJSONFromString(testPackageJSON)
	assert.NoError(t, err, "Failed to parse Package JSON")

	// Verify the basic project information
	assert.Equal(t, "TestApp", pkg.ProjectName(), "Project name mismatch")
	assert.Equal(t, "1.0.0", pkg.ProjectVersion(), "Project version mismatch")
	assert.Equal(t, "A test application", pkg.ProjectDescription(), "Project description mismatch")
	assert.Equal(t, "Jane Doe", pkg.ProjectAuthor(), "Project author mismatch")
	assert.Equal(t, "MIT", pkg.ProjectLicense(), "Project license mismatch")

	// Verify the dependencies
	expectedDeps := []Dependency{
		BaseDependency{name: "lodash", version: "4.17.21"},
		BaseDependency{name: "express", version: "4.17.1"},
	}

	actualDeps := pkg.ProjectDependencies()
	assert.Len(t, actualDeps, len(expectedDeps), "Incorrect number of dependencies")

	for i, dep := range actualDeps {
		assert.Equal(t, expectedDeps[i].Name(), dep.Name(), "Dependency name mismatch")
		assert.Equal(t, expectedDeps[i].Version(), dep.Version(), "Dependency version mismatch")
	}

	// Verify the dev dependencies
	expectedDevDeps := []Dependency{
		BaseDependency{name: "jest", version: "26.6.0"},
	}
	actualDevDeps := pkg.ProjectDevDependencies()
	assert.Len(t, actualDevDeps, len(expectedDevDeps), "Incorrect number of dev dependencies")

	for i, dep := range actualDevDeps {
		assert.Equal(t, expectedDevDeps[i].Name(), dep.Name(), "Dev Dependency name mismatch")
		assert.Equal(t, expectedDevDeps[i].Version(), dep.Version(), "Dev Dependency version mismatch")
	}

	// Verify the environment information
	expectedEnvs := []Environment{
		BaseEnvironment{name: "NPM", version: "7.x"},
		BaseEnvironment{name: "NodeJS", version: "14.x"},
	}

	actualEnvs := pkg.Environments()
	assert.Len(t, actualEnvs, len(expectedEnvs), "Incorrect number of environments")

	for i, env := range actualEnvs {
		assert.Equal(t, expectedEnvs[i].Name(), env.Name(), "Environment name mismatch")
		assert.Equal(t, expectedEnvs[i].Version(), env.Version(), "Environment version mismatch")
	}
}
