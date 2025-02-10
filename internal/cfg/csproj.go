// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cfg

import (
	"encoding/xml"
	"os"
	"strings"
)

type CSProj struct {
	XMLName       xml.Name `xml:"Project"`
	PropertyGroup []struct {
		ProjectName     string `xml:"AssemblyName"`
		Description     string `xml:"Description"`
		Company         string `xml:"Company"`
		Authors         string `xml:"Authors"`
		License         string `xml:"License"`
		TargetFramework string `xml:"TargetFramework"`
	} `xml:"PropertyGroup"`
	ItemGroups []struct {
		PackageReferences []struct {
			PackageName string `xml:"Include,attr"`
			Version     string `xml:"Version,attr"`
		} `xml:"PackageReference"`
	} `xml:"ItemGroup"`
}

func (cp CSProj) ProjectName() string {
	for _, group := range cp.PropertyGroup {
		if len(group.ProjectName) != 0 {
			return group.ProjectName
		}
	}
	return ""
}

func (cp CSProj) ProjectLicense() string {
	for _, group := range cp.PropertyGroup {
		if len(group.License) != 0 {
			return group.License
		}
	}
	return ""
}

func (cp CSProj) ProjectVersion() string { return "" }

func (cp CSProj) ProjectAuthor() string {
	for _, group := range cp.PropertyGroup {
		if len(group.Authors) != 0 {
			return group.Authors
		}
	}
	return ""
}

func (cp CSProj) ProjectDescription() string {
	for _, group := range cp.PropertyGroup {
		if len(group.Description) != 0 {
			return group.Description
		}
	}
	return ""
}

func (cp CSProj) ProjectDependencies() []Dependency {
	var deps []Dependency
	for _, itemGroup := range cp.ItemGroups {
		for _, pkg := range itemGroup.PackageReferences {
			deps = append(deps, BaseDependency{name: pkg.PackageName, version: pkg.Version})
		}
	}
	return deps
}

func (cp CSProj) ProjectDevDependencies() []Dependency {
	return nil
}

func (cp CSProj) Environments() []Environment { return nil }

func ParseCSProj(path string) (Configure, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var csproj CSProj
	err = xml.NewDecoder(file).Decode(&csproj)
	if err != nil {
		return nil, err
	}
	return csproj, nil
}

func ParseCSProjFromString(s string) (Configure, error) {
	var csproj CSProj
	err := xml.NewDecoder(strings.NewReader(s)).Decode(&csproj)
	if err != nil {
		return nil, err
	}
	return csproj, nil
}
