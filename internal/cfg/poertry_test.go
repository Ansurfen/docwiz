// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testProjectToml = `[tool.poetry]
name = "my_project"
version = "0.1.0"
description = "A short description of the project"
authors = ["Author Name <author@example.com>"]
license = "MIT"

[tool.poetry.dependencies]
python = "^3.7"
flask = "^2.0"
requests = { version = "^2.25", extras = ["security"] }

[tool.poetry.dev-dependencies]
pytest = "^6.1"
black = "^20.8b1"

[tool.poetry.scripts]
my-command = "my_project.module:main"
`

func TestPoetry(t *testing.T) {
	poetry, err := LoadPoetryFromString(testProjectToml)
	assert.NoError(t, err, "Failed to parse pyproject.toml")

	// Test dependencies extraction
	deps := poetry.ProjectDependencies()

	expectedDeps := map[string]string{
		"requests": "^2.25",
		"flask":    "^2.0",
		"python":   "^3.7",
	}

	for _, dep := range deps {
		expectedVersion, exists := expectedDeps[dep.Name()]
		assert.True(t, exists, "Unexpected dependency name: %s", dep.Name())
		assert.Equal(t, expectedVersion, dep.Version(), "Dependency version mismatch for "+dep.Name())
	}
}
