// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cfg

import (
	"os"

	"golang.org/x/mod/modfile"
)

type GoMod struct {
	name         string
	version      string
	dependencies []Dependency
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

func ParseGoMod(path string) (Configure, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	modFile, err := modfile.Parse(path, data, nil)
	if err != nil {
		return nil, err
	}

	mod := GoMod{
		name:    modFile.Module.Mod.Path,
		// TODO project version / environment version
		version: modFile.Go.Version,
	}

	for _, req := range modFile.Require {
		mod.dependencies = append(mod.dependencies, BaseDependecy{name: req.Mod.Path, version: req.Mod.Version})
	}
	return mod, nil
}
