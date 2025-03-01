// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cfg

import (
	"encoding/xml"
	"io"
	"os"
	"strings"
)

type POM struct {
	XMLName      xml.Name          `xml:"project"`
	ModelVersion string            `xml:"modelVersion"`
	ArtifactId   string            `xml:"artifactId"`
	Version      string            `xml:"version"`
	Dependencies []MavenDependency `xml:"dependencies>dependency"`
}

type MavenDependency struct {
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	Version    string `xml:"version"`
}

func (p POM) ProjectName() string {
	return p.ArtifactId
}

func (p POM) ProjectVersion() string {
	return p.Version
}

func (p POM) ProjectAuthor() string { return "" }

func (p POM) ProjectDescription() string { return "" }

func (p POM) ProjectLicense() string { return "" }

func (p POM) ProjectDependencies() []Dependency {
	var deps []Dependency
	for _, dep := range p.Dependencies {
		deps = append(deps, BaseDependency{name: dep.ArtifactId, version: dep.Version})
	}
	return deps
}

func (p POM) ProjectDevDependencies() []Dependency {
	return nil
}

func (p POM) Environments() []Environment { return nil }

func LoadPOM(r io.Reader) (Configure, error) {
	var pom POM
	err := xml.NewDecoder(r).Decode(&pom)
	if err != nil {
		return nil, err
	}
	return pom, nil
}

func LoadPOMFromFile(filename string) (Configure, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return LoadPOM(file)
}

func LoadPOMFromString(str string) (Configure, error) {
	return LoadPOM(strings.NewReader(str))
}
