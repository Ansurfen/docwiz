// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cfg

import (
	"os"

	"gopkg.in/yaml.v3"
)

type PubSpec struct {
	Name            string                 `yaml:"name"`
	Description     string                 `yaml:"description"`
	PublishTo       string                 `yaml:"publish_to"`
	Version         string                 `yaml:"version"`
	Environment     map[string]string      `yaml:"environment"`
	Dependencies    map[string]interface{} `yaml:"dependencies"`
	DevDependencies map[string]interface{} `yaml:"dev_dependencies"`
	Flutter         FlutterConfig          `yaml:"flutter"`
}

type FlutterConfig struct {
	UsesMaterialDesign bool     `yaml:"uses-material-design"`
	Assets             []string `yaml:"assets"`
}

func (ps PubSpec) ProjectName() string {
	return ps.Name
}

func (ps PubSpec) ProjectDescription() string {
	return ps.Description
}

func (ps PubSpec) ProjectVersion() string {
	return ps.Version
}

func (ps PubSpec) ProjectAuthor() string { return "" }

func (ps PubSpec) ProjectLicense() string { return "" }

func (ps PubSpec) ProjectDependencies() []Dependency {
	deps := []Dependency{}
	for name, ver := range ps.Dependencies {
		if v, ok := ver.(string); ok {
			deps = append(deps, BaseDependency{name: name, version: v})
			continue
		}
		if v, ok := ver.(map[string]any); ok {
			if vv, ok := v["version"]; ok {
				deps = append(deps, BaseDependency{name: name, version: vv.(string)})
			} else {
				deps = append(deps, BaseDependency{name: name})
			}
		}
	}
	return deps
}

func (ps PubSpec) ProjectDevDependencies() []Dependency {
	deps := []Dependency{}
	for name, ver := range ps.DevDependencies {
		if v, ok := ver.(string); ok {
			deps = append(deps, BaseDependency{name: name, version: v})
			continue
		}
		if v, ok := ver.(map[string]any); ok {
			if vv, ok := v["version"]; ok {
				deps = append(deps, BaseDependency{name: name, version: vv.(string)})
			} else {
				deps = append(deps, BaseDependency{name: name})
			}
		}
	}
	return deps
}

func (ps PubSpec) Environments() []Environment {
	return []Environment{BaseEnvironment{name: "sdk", version: ps.Environment["sdk"]}}
}

func ParsePubSpec(path string) (Configure, error) {
	var pubspec PubSpec
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	err = yaml.NewDecoder(file).Decode(&pubspec)
	if err != nil {
		return nil, err
	}
	return pubspec, nil
}
