// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const config = `<Project Sdk="Microsoft.NET.Sdk">

  <!--  -->
  <PropertyGroup>
    <AssemblyName>MyAwesomeApp</AssemblyName>
    <Description>This is an awesome app</Description>
    <Company>Awesome Company</Company>
    <Authors>John Doe</Authors>
    <License>MIT</License>
    <TargetFramework>net6.0</TargetFramework>
  </PropertyGroup>

  <!--  -->
  <PropertyGroup Condition="'$(Configuration)' == 'Debug'">
    <DefineConstants>DEBUG</DefineConstants>
    <Optimize>false</Optimize>
  </PropertyGroup>

  <PropertyGroup Condition="'$(Configuration)' == 'Release'">
    <DefineConstants>RELEASE</DefineConstants>
    <Optimize>true</Optimize>
  </PropertyGroup>

  <!--  -->
  <PropertyGroup>
    <TargetFrameworks>net5.0;net6.0</TargetFrameworks>
  </PropertyGroup>
  
  <ItemGroup>
    <PackageReference Include="Newtonsoft.Json" Version="13.0.3" />
    <PackageReference Include="Xamarin.Forms" Version="5.0.0" />
    <PackageReference Include="Serilog" Version="2.10.0" />
  </ItemGroup>
</Project>
`

const testCSProj = `<Project Sdk="Microsoft.NET.Sdk">
  <PropertyGroup>
    <AssemblyName>TestApp</AssemblyName>
    <Description>A test application</Description>
    <Company>TestCorp</Company>
    <Authors>Jane Doe</Authors>
    <License>MIT</License>
    <TargetFramework>net6.0</TargetFramework>
  </PropertyGroup>
  <ItemGroup>
    <PackageReference Include="Newtonsoft.Json" Version="13.0.3" />
    <PackageReference Include="Xunit" Version="2.4.1" />
  </ItemGroup>
</Project>`

func TestParseCSProjFromString(t *testing.T) {
	csproj, err := LoadCSProjFromString(testCSProj)
	assert.NoError(t, err, "Failed to parse CSProj XML")

	assert.Equal(t, "TestApp", csproj.ProjectName(), "Project name mismatch")
	assert.Equal(t, "A test application", csproj.ProjectDescription(), "Project description mismatch")
	assert.Equal(t, "Jane Doe", csproj.ProjectAuthor(), "Project author mismatch")
	assert.Equal(t, "MIT", csproj.ProjectLicense(), "Project license mismatch")

	expectedDeps := []Dependency{
		BaseDependency{name: "Newtonsoft.Json", version: "13.0.3"},
		BaseDependency{name: "Xunit", version: "2.4.1"},
	}

	actualDeps := csproj.ProjectDependencies()
	assert.Len(t, actualDeps, len(expectedDeps), "Incorrect number of dependencies")

	for i, dep := range actualDeps {
		assert.Equal(t, expectedDeps[i].Name(), dep.Name(), "Dependency name mismatch")
		assert.Equal(t, expectedDeps[i].Version(), dep.Version(), "Dependency version mismatch")
	}
}

func TestParseEmptyCSProj(t *testing.T) {
	emptyCSProj := `<Project Sdk="Microsoft.NET.Sdk">
		<PropertyGroup></PropertyGroup>
	</Project>`
	csproj, err := LoadCSProjFromString(emptyCSProj)
	assert.NoError(t, err, "Parsing empty CSProj should not fail")

	assert.Equal(t, "", csproj.ProjectName(), "Project name should be empty")
	assert.Equal(t, "", csproj.ProjectLicense(), "Project license should be empty")
	assert.Equal(t, "", csproj.ProjectAuthor(), "Project author should be empty")
	assert.Empty(t, csproj.ProjectDependencies(), "Dependencies should be empty")
}

func TestMultiplePropertyGroups(t *testing.T) {
	multiGroupCSProj := `<Project Sdk="Microsoft.NET.Sdk">
      <PropertyGroup>
        <Authors></Authors>
      </PropertyGroup>
      <PropertyGroup>
        <Authors>Valid Author</Authors>
      </PropertyGroup>
    </Project>`

	csproj, err := LoadCSProjFromString(multiGroupCSProj)
	assert.NoError(t, err, "Failed to parse CSProj with multiple PropertyGroups")

	assert.Equal(t, "Valid Author", csproj.ProjectAuthor(), "Should get the first non-empty author")
}
