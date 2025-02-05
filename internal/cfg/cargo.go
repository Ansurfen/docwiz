// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cfg

import "github.com/BurntSushi/toml"

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
	if _, ok := ct.Package.Authors.(map[string]any); ok {
		return ct.Workspace.Package.Authors.(string)
	} else {
		return ct.Package.Authors.(string)
	}
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
				deps = append(deps, BaseDependecy{name: name, version: vv})
			} else {
				deps = append(deps, BaseDependecy{name: name})
			}
		} else {
			deps = append(deps, BaseDependecy{name: name, version: value.(string)})
		}
	}
	return deps
}

func (ct CargoToml) ProjectDevDependencies() []Dependency {
	var deps []Dependency
	for name, value := range ct.DevDependencies {
		if v, ok := value.(map[string]any); ok {
			if vv, ok := v["version"].(string); ok {
				deps = append(deps, BaseDependecy{name: name, version: vv})
			} else {
				deps = append(deps, BaseDependecy{name: name})
			}
		} else {
			deps = append(deps, BaseDependecy{name: name, version: value.(string)})
		}
	}
	return deps
}

func ParseCargoToml(path string) (Configure, error) {
	var cargo CargoToml
	_, err := toml.DecodeFile(path, &cargo)
	if err != nil {
		return nil, err
	}
	
	return cargo, nil
}