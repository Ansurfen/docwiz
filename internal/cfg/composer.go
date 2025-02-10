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

type Composer struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Authors     []struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"authors"`
	License      []string          `json:"license"`
	Dependencies map[string]string `json:"require"`
}

func (c Composer) ProjectName() string {
	return c.Name
}

func (c Composer) ProjectAuthor() string {
	authors := []string{}
	for _, a := range c.Authors {
		authors = append(authors, a.Name)
	}
	return strings.Join(authors, ";")
}

func (c Composer) ProjectVersion() string {
	return ""
}

func (c Composer) ProjectDescription() string {
	return c.Description
}

func (c Composer) ProjectLicense() string {
	return strings.Join(c.License, ";")
}

func (c Composer) ProjectDependencies() []Dependency {
	var deps []Dependency
	for name, ver := range c.Dependencies {
		deps = append(deps, BaseDependency{name: name, version: ver})
	}
	return deps
}

func (c Composer) ProjectDevDependencies() []Dependency { return nil }

func (c Composer) Environments() []Environment { return nil }

func LoadComposer(r io.Reader) (Configure, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var composer Composer

	err = json.Unmarshal(data, &composer)
	if err != nil {
		return nil, err
	}
	return composer, nil
}

func LoadComposerFromFile(filename string) (Configure, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return LoadComposer(file)
}

func LoadComposerFromString(str string) (Configure, error) {
	return LoadComposer(strings.NewReader(str))
}
