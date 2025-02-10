// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cfg

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)

type PackageJSON struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Description     string            `json:"description"`
	Main            string            `json:"main"`
	Bin             map[string]string `json:"bin"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
	Scripts         map[string]string `json:"scripts"`
	Repository      struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"repository"`
	Author   string `json:"author"`
	License  string `json:"license"`
	Homepage string `json:"homepage"`
	Engines  struct {
		NPM  string `json:"npm"`
		Node string `json:"node"`
	} `json:"engines"`
}

func (c PackageJSON) ProjectName() string {
	return c.Name
}

func (c PackageJSON) ProjectDescription() string {
	return c.Description
}

func (c PackageJSON) ProjectAuthor() string {
	return c.Author
}

func (c PackageJSON) ProjectLicense() string {
	return c.License
}

func (c PackageJSON) ProjectVersion() string {
	return c.Version
}

func (c PackageJSON) ProjectDependencies() []Dependency {
	deps := []Dependency{}
	for depName, depVersion := range c.Dependencies {
		deps = append(deps, BaseDependency{name: depName, version: depVersion})
	}
	return deps
}

func (c PackageJSON) ProjectDevDependencies() []Dependency {
	deps := []Dependency{}
	for depName, depVersion := range c.DevDependencies {
		deps = append(deps, BaseDependency{name: depName, version: depVersion})
	}
	return deps
}

func (c PackageJSON) Environments() []Environment {
	var envs []Environment

	if len(c.Engines.NPM) != 0 {
		envs = append(envs, BaseEnvironment{name: "NPM", version: c.Engines.NPM})
	}

	if len(c.Engines.Node) != 0 {
		envs = append(envs, BaseEnvironment{name: "NodeJS", version: c.Engines.Node})
	}

	return envs
}

func LoadPackageJSON(r io.Reader) (Configure, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var pkg PackageJSON

	err = json.Unmarshal(data, &pkg)
	if err != nil {
		return nil, err
	}
	return pkg, nil
}

func LoadPackageJSONFromFile(filename string) (Configure, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return LoadPackageJSON(file)
}

func LoadPackageJSONFromString(str string) (Configure, error) {
	return LoadPackageJSON(strings.NewReader(str))
}
