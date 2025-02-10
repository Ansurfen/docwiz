// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const goModContent = `
module github.com/example/project

go 1.18

require (
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/mod v0.5.1
)
`

func TestLoadGoModFromString(t *testing.T) {
	config, err := LoadGoModFromString(goModContent)
	assert.NoError(t, err, "Loading go.mod from string should not return an error")

	gm, ok := config.(GoMod)
	assert.True(t, ok, "Config should be of type GoMod")

	assert.Equal(t, "github.com/example/project", gm.ProjectName(), "Module name should match")
	assert.Equal(t, "1.18", gm.Environments()[0].Version(), "Go version should be 1.18")

	expected := []Dependency{
		BaseDependency{name: "github.com/sirupsen/logrus", version: "v1.8.1"},
		BaseDependency{name: "golang.org/x/mod", version: "v0.5.1"},
	}

	actual := gm.ProjectDependencies()
	assert.Len(t, actual, len(expected), "Dependency count should match expected")

	for i, dep := range actual {
		assert.Equal(t, expected[i].Name(), dep.Name(), "Dependency name should match expected")
		assert.Equal(t, expected[i].Version(), dep.Version(), "Dependency version should match expected")
	}
}
