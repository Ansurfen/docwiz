function New-DocWiz {
    param (
        [string]$version,
        [string]$os,
        [string]$arch
    )

    Write-Host "Building for version: $version, OS: $os, Architecture: $arch..."

    # Set environment variables for GOOS and GOARCH
    $env:GOOS = $os
    $env:GOARCH = $arch

    # Navigate to the cli directory and build the executable
    Set-Location -Path "./cli"
    $exePath = ""
    if ($os -eq "windows") {
        go build -o "../docwiz.exe" .
        $exePath = "./docwiz.exe"
    } else {
        go build -o "../docwiz" .
        $exePath = "./docwiz"
    }

    Set-Location -Path "../"
    Write-Host "Build complete for $os/$arch."

    return $exePath
}

function New-Package {
    param (
        [string]$exePath,
        [string]$version,
        [string]$os,
        [string]$arch
    )

    Write-Host "Creating package for $os..."

    $packageName = "docwiz-$version-$os-$arch"

    if ($os -eq "windows") {
        # Windows: Create zip file using Compress-Archive
        Compress-Archive -Path "$exePath", ".\template", ".\License" -DestinationPath ".\$packageName.zip"
    } else {
        # Linux/Darwin: Create tar file using tar
        $tarPath = ".\$packageName.tar"
        tar -cf $tarPath "$exePath" ".\template" ".\License"
    }

    Write-Host "Packaging complete for $os. Package created at: $packagePath."
}

# Main Script
if ($args.Count -eq 0) {
    Write-Host "Building for debug version..."
    $__ = New-DocWiz -version "debug" -os "windows" -arch "amd64"
    exit
}

# Extract release version
$version = $args[0]

# Check if version is empty
if ([string]::IsNullOrEmpty($version)) {
    Write-Host "Release version is required!"
    exit
}

# Platforms to build for (Windows, Linux, Darwin)
$platforms = @(
    @{os="windows"; arch="amd64"},
    @{os="linux"; arch="amd64"},
    @{os="darwin"; arch="amd64"}
)

foreach ($platform in $platforms) {
    $exePath = New-DocWiz -version $version -os $platform.os -arch $platform.arch
    New-Package -exePath $exePath -version $version -os $platform.os -arch $platform.arch
}

Write-Host "All builds and packaging complete for version $version."
