#!/bin/bash
set -e

cd "$(dirname "$0")"

echo "Building webview_app for macOS..."
go build -o webview_app main.go

echo "Build complete!"
echo "Run ./webview_app to test the custom User-Agent."
