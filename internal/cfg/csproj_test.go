// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cfg_test

import (
	"docwiz/internal/cfg"
	"fmt"
	"testing"
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

func TestCSProj(t *testing.T) {
	c, err := cfg.ParseCSProjFromString(config)
	if err != nil {
		panic(err)
	}
	fmt.Println(c.ProjectDependencies())
}
