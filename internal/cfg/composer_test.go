// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectName(t *testing.T) {
	composer := Composer{
		Name: "Test Composer",
	}

	assert.Equal(t, "Test Composer", composer.ProjectName(), "Project name should match")
}

func TestProjectDescription(t *testing.T) {
	composer := Composer{
		Description: "A test composer project",
	}

	assert.Equal(t, "A test composer project", composer.ProjectDescription(), "Project description should match")
}

func TestProjectAuthor(t *testing.T) {
	composer := Composer{
		Authors: []struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}{
			{Name: "Jane Doe", Email: "jane.doe@example.com"},
			{Name: "John Smith", Email: "john.smith@example.com"},
		},
	}

	expected := "Jane Doe;John Smith"
	assert.Equal(t, expected, composer.ProjectAuthor(), "Authors should be concatenated correctly")
}

func TestProjectLicense(t *testing.T) {
	composer := Composer{
		License: []string{"MIT", "GPL-3.0"},
	}

	expected := "MIT;GPL-3.0"
	assert.Equal(t, expected, composer.ProjectLicense(), "Licenses should be concatenated correctly")
}

func TestProjectDependencies(t *testing.T) {
	composer := Composer{
		Dependencies: map[string]string{
			"php":   "7.4",
			"mysql": "5.7",
		},
	}

	deps := composer.ProjectDependencies()
	expectedDeps := []Dependency{
		BaseDependency{name: "php", version: "7.4"},
		BaseDependency{name: "mysql", version: "5.7"},
	}

	// Iterate over expected dependencies and check each one
	for i, dep := range expectedDeps {
		assert.Equal(t, dep.Name(), deps[i].Name(), "Dependency names should match")
		assert.Equal(t, dep.Version(), deps[i].Version(), "Dependency versions should match")
	}
}

func TestLoadComposerFromString(t *testing.T) {
	jsonData := `{
		"name": "Test Composer",
		"description": "A test composer project",
		"license": ["MIT", "GPL-3.0"],
		"require": {
			"php": "7.4",
			"mysql": "5.7"
		},
		"authors": [
			{"name": "Jane Doe", "email": "jane.doe@example.com"},
			{"name": "John Smith", "email": "john.smith@example.com"}
		]
	}`

	composer, err := LoadComposerFromString(jsonData)

	// Assert no error
	assert.NoError(t, err, "Failed to load composer from string")

	// Assert project name
	assert.Equal(t, "Test Composer", composer.ProjectName(), "Project name should match")

	// Assert project description
	assert.Equal(t, "A test composer project", composer.ProjectDescription(), "Project description should match")

	// Assert project author
	assert.Equal(t, "Jane Doe;John Smith", composer.ProjectAuthor(), "Authors should be concatenated correctly")

	// Assert project license
	assert.Equal(t, "MIT;GPL-3.0", composer.ProjectLicense(), "Licenses should be concatenated correctly")

	// Assert project dependencies
	deps := composer.ProjectDependencies()
	expectedDeps := []Dependency{
		BaseDependency{name: "php", version: "7.4"},
		BaseDependency{name: "mysql", version: "5.7"},
	}

	// Iterate over expected dependencies and check each one
	for i, dep := range expectedDeps {
		assert.Equal(t, dep.Name(), deps[i].Name(), "Dependency names should match")
		assert.Equal(t, dep.Version(), deps[i].Version(), "Dependency versions should match")
	}
}

func TestLoadComposer_InvalidJSON(t *testing.T) {
	invalidJSON := `{"name": "Test Composer", "description": "A test composer project",`

	_, err := LoadComposerFromString(invalidJSON)

	// Assert error is returned
	assert.Error(t, err, "An error is expected due to invalid JSON")
}
