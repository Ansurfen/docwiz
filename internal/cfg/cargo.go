// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cfg

import (
	"io"
	"strings"

	"github.com/BurntSushi/toml"
)

type CargoToml struct {
	Workspace       CargoTomlWorkspace     `toml:"workspace"`
	Package         CargoTomlPackageConfig `toml:"package"`
	Dependencies    map[string]any         `toml:"dependencies"`
	DevDependencies map[string]any         `toml:"dev-dependencies"`
}

type CargoTomlWorkspace struct {
	Package      CargoTomlPackageConfig `toml:"package"`
	Dependencies map[string]any         `toml:"dependencies"`
}

type CargoTomlPackageConfig struct {
	Name        string `toml:"name"`
	Description string `toml:"description"`
	Version     any    `toml:"version"`
	Authors     any    `toml:"authors"`
	Repository  any    `toml:"repository"`
	License     any    `toml:"license"`
	Edition     any    `toml:"edition"`
	RustVersion any    `toml:"rust-version"`
}

func (ct CargoToml) ProjectName() string {
	return ct.Package.Name
}

func (ct CargoToml) ProjectDescription() string {
	return ct.Package.Description
}

func (ct CargoToml) ProjectVersion() string {
	if _, ok := ct.Package.Version.(map[string]any); ok {
		return ct.Workspace.Package.Version.(string)
	} else {
		return ct.Package.Version.(string)
	}
}

func (ct CargoToml) ProjectAuthor() string {
	switch authors := ct.Package.Authors.(type) {
	case map[string]any:
		return ct.Workspace.Package.Authors.(string)
	case []any:
		ret := []string{}
		for _, author := range authors {
			ret = append(ret, author.(string))
		}
		return strings.Join(ret, ";")
	case string:
		return authors
	}
	return ""
}

func (ct CargoToml) ProjectLicense() string {
	if _, ok := ct.Package.License.(map[string]any); ok {
		return ct.Workspace.Package.License.(string)
	} else {
		return ct.Package.License.(string)
	}
}

func (ct CargoToml) ProjectDependencies() []Dependency {
	var deps []Dependency
	for name, value := range ct.Dependencies {
		if v, ok := value.(map[string]any); ok {
			if vv, ok := v["version"].(string); ok {
				deps = append(deps, BaseDependency{name: name, version: vv})
			} else {
				deps = append(deps, BaseDependency{name: name})
			}
		} else {
			deps = append(deps, BaseDependency{name: name, version: value.(string)})
		}
	}
	return deps
}

func (ct CargoToml) ProjectDevDependencies() []Dependency {
	var deps []Dependency
	for name, value := range ct.DevDependencies {
		if v, ok := value.(map[string]any); ok {
			if vv, ok := v["version"].(string); ok {
				deps = append(deps, BaseDependency{name: name, version: vv})
			} else {
				deps = append(deps, BaseDependency{name: name})
			}
		} else {
			deps = append(deps, BaseDependency{name: name, version: value.(string)})
		}
	}
	return deps
}

func (ct CargoToml) Environments() []Environment {
	var envs []Environment
	if _, ok := ct.Package.RustVersion.(map[string]any); ok {
		envs = append(envs, BaseEnvironment{name: "rust", version: ct.Workspace.Package.RustVersion.(string)})
	} else {
		if v, ok := ct.Package.RustVersion.(string); ok {
			envs = append(envs, BaseEnvironment{name: "rust", version: v})
		}
	}
	return envs
}

func LoadCargo(r io.Reader) (Configure, error) {
	var cargo CargoToml
	_, err := toml.NewDecoder(r).Decode(&cargo)
	if err != nil {
		return nil, err
	}

	return cargo, nil
}

func LoadCargoFromFile(filename string) (Configure, error) {
	var cargo CargoToml
	_, err := toml.DecodeFile(filename, &cargo)
	if err != nil {
		return nil, err
	}

	return cargo, nil
}

func LoadCargoFromString(str string) (Configure, error) {
	return LoadCargo(strings.NewReader(str))
}
