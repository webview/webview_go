#!/usr/bin/env bash
set -euo pipefail

# Echo context
echo "===> Building on $(lsb_release -sd 2>/dev/null || echo Linux) with Go $(go version)"

# Install system dependencies (dev headers for GTK + WebKitGTK 4.1)
sudo apt-get update
sudo apt-get install -y \
  build-essential pkg-config \
  libgtk-3-dev libwebkit2gtk-4.1-dev

# Optional media codecs to reduce WebKit warnings (safe to skip)
# sudo apt-get install -y gstreamer1.0-plugins-base gstreamer1.0-plugins-good \
#     gstreamer1.0-plugins-bad gstreamer1.0-plugins-ugly gstreamer1.0-libav

export CGO_ENABLED=1
export CC=gcc
export CXX=g++
export PKG_CONFIG=pkg-config

# Build the module
go build ./...

# Build the UA example into ./bin
mkdir -p bin
( cd examples/ua && go build -o ../../bin/webview-ua )

echo "===> Built: $(realpath bin/webview-ua 2>/dev/null || echo ./bin/webview-ua)"
echo "===> Run:   ./bin/webview-ua"
