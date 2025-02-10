// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cfg

import (
	"io"
	"os"
	"strings"

	"golang.org/x/mod/modfile"
)

type GoMod struct {
	name         string
	version      string
	dependencies []Dependency
	file         *modfile.File
}

func (gm GoMod) ProjectName() string {
	return gm.name
}

func (gm GoMod) ProjectVersion() string {
	return gm.version
}

func (gm GoMod) ProjectAuthor() string { return "" }

func (gm GoMod) ProjectDescription() string { return "" }

func (gm GoMod) ProjectLicense() string { return "" }

func (gm GoMod) ProjectDependencies() []Dependency {
	return gm.dependencies
}

func (gm GoMod) ProjectDevDependencies() []Dependency {
	return gm.dependencies
}

func (gm GoMod) Environments() []Environment {
	return []Environment{BaseEnvironment{name: "Go", version: gm.file.Go.Version}}
}

func LoadGoMod(r io.Reader) (Configure, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	modFile, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return nil, err
	}

	mod := GoMod{
		name: modFile.Module.Mod.Path,
		file: modFile,
	}

	for _, req := range modFile.Require {
		mod.dependencies = append(mod.dependencies, BaseDependency{name: req.Mod.Path, version: req.Mod.Version})
	}
	return mod, nil
}

func LoadGoModFromFile(filename string) (Configure, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return LoadGoMod(file)
}

func LoadGoModFromString(str string) (Configure, error) {
	return LoadGoMod(strings.NewReader(str))
}
