// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cfg

import (
	"io"
	"strings"

	"github.com/BurntSushi/toml"
)

type Poetry struct {
	Tool struct {
		Poetry struct {
			Name            string         `toml:"name"`
			Version         string         `toml:"version"`
			Description     string         `toml:"description"`
			Authors         []string       `toml:"authors"`
			License         string         `toml:"license"`
			Dependencies    map[string]any `toml:"dependencies"`
			DevDependencies map[string]any `toml:"dev-dependencies"`
		} `toml:"poetry"`
	} `toml:"tool"`
}

func (p Poetry) ProjectName() string {
	return p.Tool.Poetry.Name
}

func (p Poetry) ProjectVersion() string {
	return p.Tool.Poetry.Version
}

func (p Poetry) ProjectDescription() string {
	return p.Tool.Poetry.Description
}

func (p Poetry) ProjectAuthor() string {
	return strings.Join(p.Tool.Poetry.Authors, ";")
}

func (p Poetry) ProjectLicense() string {
	return p.Tool.Poetry.License
}

func (p Poetry) ProjectDependencies() []Dependency {
	var deps []Dependency
	for name, value := range p.Tool.Poetry.Dependencies {
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

func (p Poetry) ProjectDevDependencies() []Dependency {
	var deps []Dependency
	for name, value := range p.Tool.Poetry.DevDependencies {
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

func (p Poetry) Environments() []Environment {
	return nil
}

func LoadPoetry(r io.Reader) (Configure, error) {
	var poetry Poetry
	_, err := toml.NewDecoder(r).Decode(&poetry)
	if err != nil {
		return nil, err
	}

	return poetry, nil
}

func LoadPoetryFromFile(filename string) (Configure, error) {
	var poetry Poetry
	_, err := toml.DecodeFile(filename, &poetry)
	if err != nil {
		return nil, err
	}

	return poetry, nil
}

func LoadPoetryFromString(str string) (Configure, error) {
	return LoadPoetry(strings.NewReader(str))
}
