// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cfg

import "strings"

type Configure interface {
	ProjectName() string
	ProjectDescription() string
	ProjectAuthor() string
	ProjectVersion() string
	ProjectLicense() string
	ProjectDependencies() []Dependency
	ProjectDevDependencies() []Dependency
}

type Dependency interface {
	Name() string
	Version() string
	Match(string) bool
	PartialMatch(string) bool
	Contains(string) bool
}

type BaseDependecy struct {
	name    string
	version string
}

func (d BaseDependecy) Name() string {
	return d.name
}

func (d BaseDependecy) Version() string {
	return d.version
}

func (d BaseDependecy) Match(name string) bool {
	return d.name == name
}

func (d BaseDependecy) PartialMatch(name string) bool {
	return strings.HasPrefix(d.name, name)
}

func (d BaseDependecy) Contains(name string) bool {
	return strings.Contains(d.name, name)
}
