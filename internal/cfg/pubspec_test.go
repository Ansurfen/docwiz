// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testPubSpec = `
name: my_flutter_app
description: A Flutter project
publish_to: "none"
version: "1.0.0"
environment:
  sdk: "2.10.0"
dependencies:
  flutter: "2.5.0"
  cupertino_icons:
    version: "0.1.3"
dev_dependencies:
  test: "any"
flutter:
  uses-material-design: true
  assets:
    - "assets/images"
`

func TestLoadPubSpecFromString(t *testing.T) {
	pubspec, err := LoadPubSpecFromString(testPubSpec)
	assert.NoError(t, err, "Failed to parse PubSpec YAML")

	// Test basic fields
	assert.Equal(t, "my_flutter_app", pubspec.ProjectName(), "Project name mismatch")
	assert.Equal(t, "A Flutter project", pubspec.ProjectDescription(), "Project description mismatch")
	assert.Equal(t, "1.0.0", pubspec.ProjectVersion(), "Project version mismatch")

	// Expected dependencies
	expectedDeps := map[string]string{
		"flutter":         "2.5.0",
		"cupertino_icons": "0.1.3",
	}

	// Test dependencies using expected map
	deps := pubspec.ProjectDependencies()
	assert.Len(t, deps, len(expectedDeps), "Incorrect number of dependencies")

	for _, dep := range deps {
		expectedVersion, exists := expectedDeps[dep.Name()]
		assert.True(t, exists, "Unexpected dependency name: %s", dep.Name())
		assert.Equal(t, expectedVersion, dep.Version(), "Dependency version mismatch for "+dep.Name())
	}

	// Test dev dependencies using expected map
	expectedDevDeps := map[string]string{
		"test": "any",
	}

	devDeps := pubspec.ProjectDevDependencies()
	assert.Len(t, devDeps, len(expectedDevDeps), "Incorrect number of dev dependencies")

	for _, dep := range devDeps {
		expectedVersion, exists := expectedDevDeps[dep.Name()]
		assert.True(t, exists, "Unexpected dev dependency name: %s", dep.Name())
		assert.Equal(t, expectedVersion, dep.Version(), "Dev dependency version mismatch for "+dep.Name())
	}

	// Test environment
	expectedEnvs := map[string]string{
		"sdk": "2.10.0",
	}

	envs := pubspec.Environments()
	assert.Len(t, envs, len(expectedEnvs), "Incorrect number of environments")

	for _, env := range envs {
		expectedVersion, exists := expectedEnvs[env.Name()]
		assert.True(t, exists, "Unexpected environment name: %s", env.Name())
		assert.Equal(t, expectedVersion, env.Version(), "Environment version mismatch for "+env.Name())
	}
}
