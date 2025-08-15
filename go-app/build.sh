#!/usr/bin/env bash
set -e

# Windows
echo "Building Windows..."
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H=windowsgui" -o build/app-windows.exe

# MacOS (build only works on Mac)
# echo "Building macOS..."
# GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o build/app-macos

# Android (laundry list of things to be installed and added to PATH -> Android SDK, NDK, Gogio, Java JDK)
echo "Building Android..."
gogio -target android -appid com.goapp.quotes -o build/app-android.apk ./...

echo "Builds complete."
