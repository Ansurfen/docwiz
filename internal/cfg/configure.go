// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cfg

import "strings"

// Configure defines an interface for project configuration.
// It provides methods to retrieve project metadata, dependencies, and environments.
type Configure interface {
	// ProjectName returns the name of the project.
	ProjectName() string

	// ProjectDescription returns a brief description of the project.
	ProjectDescription() string

	// ProjectAuthor returns the name of the project's author.
	ProjectAuthor() string

	// ProjectVersion returns the current version of the project.
	ProjectVersion() string

	// ProjectLicense returns the license under which the project is distributed.
	ProjectLicense() string

	// ProjectDependencies returns a list of dependencies required for the project.
	ProjectDependencies() []Dependency

	// ProjectDevDependencies returns a list of development dependencies for the project.
	ProjectDevDependencies() []Dependency

	// Environments returns a list of environments in which the project can run.
	Environments() []Environment
}

// Dependency defines an interface for managing project dependencies.
type Dependency interface {
	// Name returns the name of the dependency.
	Name() string

	// Version returns the version of the dependency.
	Version() string

	// Match checks if the given name exactly matches the dependency name.
	Match(string) bool

	// PartialMatch checks if the given name is a prefix of the dependency name.
	PartialMatch(string) bool

	// Contains checks if the dependency name contains the given substring.
	Contains(string) bool
}

// Environment defines an interface for project execution environments.
type Environment interface {
	// Name returns the name of the environment.
	Name() string

	// Version returns the version of the environment.
	Version() string
}

// BaseDependency is a concrete implementation of the Dependency interface.
type BaseDependency struct {
	name    string // The name of the dependency
	version string // The version of the dependency
}

// Name returns the name of the dependency.
func (d BaseDependency) Name() string {
	return d.name
}

// Version returns the version of the dependency.
func (d BaseDependency) Version() string {
	return d.version
}

// Match checks if the provided name exactly matches the dependency name.
func (d BaseDependency) Match(name string) bool {
	return d.name == name
}

// PartialMatch checks if the provided name is a prefix of the dependency name.
func (d BaseDependency) PartialMatch(name string) bool {
	return strings.HasPrefix(d.name, name)
}

// Contains checks if the provided name is a substring of the dependency name.
func (d BaseDependency) Contains(name string) bool {
	return strings.Contains(d.name, name)
}

// BaseEnvironment is a concrete implementation of the Environment interface.
type BaseEnvironment struct {
	name    string // The name of the environment
	version string // The version of the environment
}

// Name returns the name of the environment.
func (e BaseEnvironment) Name() string {
	return e.name
}

// Version returns the version of the environment.
func (e BaseEnvironment) Version() string {
	return e.version
}
