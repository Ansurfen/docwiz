#!/bin/bash
# Copyright 2025 The DocWiz Authors. All rights reserved.
# Use of this source code is governed by a MIT-style
# license that can be found in the LICENSE file.

# Function to build DocWiz
New_DocWiz() {
    local version="$1"
    local os="$2"
    local arch="$3"

    echo "Building for version: $version, OS: $os, Architecture: $arch..."

    # Set environment variables for GOOS and GOARCH
    export GOOS="$os"
    export GOARCH="$arch"

    # Navigate to the cli directory and build the executable
    cd ./cli || exit
    local exePath=""
    if [[ "$os" == "windows" ]]; then
        go build -o "../docwiz.exe" .
        exePath="../docwiz.exe"
    else
        go build -o "../docwiz" .
        exePath="../docwiz"
    fi
    cd ../ || exit

    echo "Build complete for $os/$arch."
    echo "$exePath"
}

# Function to package the build
New_Package() {
    local exePath="$1"
    local version="$2"
    local os="$3"
    local arch="$4"

    echo "Creating package for $os..."
    local packageName="docwiz-$version-$os-$arch"

    if [[ "$os" == "windows" ]]; then
        # Windows: Create zip file
        zip -r "$packageName.zip" "$exePath" "./template" "./License"
    else
        # Linux/Darwin: Create tar file
        tar -cf "$packageName.tar" "$exePath" "./template" "./License"
    fi

    echo "Packaging complete for $os. Package created at: $packageName."
}

# Main Script
if [[ $# -eq 0 ]]; then
    echo "Building for debug version..."
    New_DocWiz "debug" "linux" "amd64"
    exit 0
fi

# Extract release version
version="$1"

# Check if version is empty
if [[ -z "$version" ]]; then
    echo "Release version is required!"
    exit 1
fi

# Platforms to build for (Windows, Linux, Darwin)
declare -A platforms
platforms=(
    [windows]=amd64
    [linux]=amd64
    [darwin]=amd64
)

for os in "${!platforms[@]}"; do
    arch=${platforms[$os]}
    exePath=$(New_DocWiz "$version" "$os" "$arch")
    New_Package "$exePath" "$version" "$os" "$arch"

    # Clean up the built executable
    rm -f "$exePath"
done

echo "All builds and packaging complete for version $version."
