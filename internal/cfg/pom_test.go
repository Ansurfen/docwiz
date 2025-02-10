package cfg

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Sample POM string for testing
const testPOM = `
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/POM/4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>
    <artifactId>MyAwesomeApp</artifactId>
    <version>1.0.0</version>
    <dependencies>
        <dependency>
            <groupId>org.springframework</groupId>
            <artifactId>spring-core</artifactId>
            <version>5.3.9</version>
        </dependency>
        <dependency>
            <groupId>org.apache.commons</groupId>
            <artifactId>commons-io</artifactId>
            <version>2.8.0</version>
        </dependency>
    </dependencies>
</project>
`

func TestLoadPOMFromString(t *testing.T) {
	// Parse the POM string
	pom, err := LoadPOMFromString(testPOM)
	assert.NoError(t, err, "Failed to parse POM string")

	// Validate Project Name and Version
	assert.Equal(t, "MyAwesomeApp", pom.ProjectName(), "Project name mismatch")
	assert.Equal(t, "1.0.0", pom.ProjectVersion(), "Project version mismatch")

	// Prepare expected dependencies
	expectedDeps := []Dependency{
		BaseDependency{name: "spring-core", version: "5.3.9"},
		BaseDependency{name: "commons-io", version: "2.8.0"},
	}

	// Validate dependencies using expected values
	deps := pom.ProjectDependencies()
	assert.Len(t, deps, len(expectedDeps), "Incorrect number of dependencies")

	// Iterate and compare each expected dependency
	for i, expectedDep := range expectedDeps {
		assert.Equal(t, expectedDep.Name(), deps[i].Name(), "Dependency name mismatch at index %d", i)
		assert.Equal(t, expectedDep.Version(), deps[i].Version(), "Dependency version mismatch at index %d", i)
	}
}

func TestLoadPOMFromFile(t *testing.T) {
	// Create a temporary POM file for testing
	tempFile := "test_pom.xml"
	err := os.WriteFile(tempFile, []byte(testPOM), 0644)
	assert.NoError(t, err, "Failed to write test POM file")

	// Ensure the file is cleaned up after the test
	defer func() {
		err := os.Remove(tempFile)
		if err != nil {
			t.Errorf("Failed to delete temporary POM file: %v", err)
		}
	}()

	// Load POM from the file
	pom, err := LoadPOMFromFile(tempFile)
	assert.NoError(t, err, "Failed to parse POM from file")

	// Validate Project Name and Version
	assert.Equal(t, "MyAwesomeApp", pom.ProjectName(), "Project name mismatch")
	assert.Equal(t, "1.0.0", pom.ProjectVersion(), "Project version mismatch")

	// Prepare expected dependencies
	expectedDeps := []Dependency{
		BaseDependency{name: "spring-core", version: "5.3.9"},
		BaseDependency{name: "commons-io", version: "2.8.0"},
	}

	// Validate dependencies using expected values
	deps := pom.ProjectDependencies()
	assert.Len(t, deps, len(expectedDeps), "Incorrect number of dependencies")

	// Iterate and compare each expected dependency
	for i, expectedDep := range expectedDeps {
		assert.Equal(t, expectedDep.Name(), deps[i].Name(), "Dependency name mismatch at index %d", i)
		assert.Equal(t, expectedDep.Version(), deps[i].Version(), "Dependency version mismatch at index %d", i)
	}
}
