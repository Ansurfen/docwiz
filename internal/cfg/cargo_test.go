package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testCargoToml = `
[package]
name = "my_cargo_project"
description = "A sample Cargo project"
version = "1.0.0"
authors = ["John Doe"]
repository = "https://github.com/my_cargo_project"
license = "MIT"
edition = "2018"
rust-version = "1.56.0"

[dependencies]
serde = "1.0"
toml = { version = "0.5", optional = true }

[dev-dependencies]
rustfmt = "1.4.0"
`

func TestLoadCargoFromString(t *testing.T) {
	cargo, err := LoadCargoFromString(testCargoToml)
	assert.NoError(t, err, "Failed to parse Cargo.toml")

	// Test basic fields
	assert.Equal(t, "my_cargo_project", cargo.ProjectName(), "Project name mismatch")
	assert.Equal(t, "A sample Cargo project", cargo.ProjectDescription(), "Project description mismatch")
	assert.Equal(t, "1.0.0", cargo.ProjectVersion(), "Project version mismatch")
	assert.Equal(t, "John Doe", cargo.ProjectAuthor(), "Project author mismatch")
	assert.Equal(t, "MIT", cargo.ProjectLicense(), "Project license mismatch")

	// Expected dependencies
	expectedDeps := map[string]string{
		"serde": "1.0",
		"toml":  "0.5",
	}

	// Test dependencies using expected map
	deps := cargo.ProjectDependencies()
	assert.Len(t, deps, len(expectedDeps), "Incorrect number of dependencies")

	for _, dep := range deps {
		expectedVersion, exists := expectedDeps[dep.Name()]
		assert.True(t, exists, "Unexpected dependency name: %s", dep.Name())
		assert.Equal(t, expectedVersion, dep.Version(), "Dependency version mismatch for "+dep.Name())
	}

	// Test dev dependencies using expected map
	expectedDevDeps := map[string]string{
		"rustfmt": "1.4.0",
	}

	devDeps := cargo.ProjectDevDependencies()
	assert.Len(t, devDeps, len(expectedDevDeps), "Incorrect number of dev dependencies")

	for _, dep := range devDeps {
		expectedVersion, exists := expectedDevDeps[dep.Name()]
		assert.True(t, exists, "Unexpected dev dependency name: %s", dep.Name())
		assert.Equal(t, expectedVersion, dep.Version(), "Dev dependency version mismatch for "+dep.Name())
	}

	// Test environments
	expectedEnv := map[string]string{
		"rust": "1.56.0",
	}

	envs := cargo.Environments()
	assert.Len(t, envs, len(expectedEnv), "Incorrect number of environments")

	for _, env := range envs {
		expectedVersion, exists := expectedEnv[env.Name()]
		assert.True(t, exists, "Unexpected environment name: %s", env.Name())
		assert.Equal(t, expectedVersion, env.Version(), "Environment version mismatch for "+env.Name())
	}
}

func TestLoadCargo_Error(t *testing.T) {
	// Test invalid Cargo.toml content (missing required fields)
	invalidCargoToml := `
	[package]
	name = "my_cargo_project"
	`

	cargo, err := LoadCargoFromString(invalidCargoToml)
	assert.Error(t, err, "Expected error for invalid Cargo.toml")
	assert.Nil(t, cargo, "Expected nil cargo on error")
}
