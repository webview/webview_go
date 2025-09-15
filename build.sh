#!/usr/bin/env bash
set -e

echo "[INFO] Installing dependencies..."

# Проверка платформы
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  sudo apt-get update
  sudo apt-get install -y build-essential pkg-config libgtk-3-dev libwebkit2gtk-4.0-dev
  echo "[INFO] Dependencies installed (Linux)"
elif [[ "$OSTYPE" == "darwin"* ]]; then
  echo "[INFO] No dependencies required on macOS"
else
  echo "[ERROR] Unsupported OS: $OSTYPE"
  exit 1
fi

echo "[INFO] Building Go binary..."
go build -o webview_app ./examples/bind

echo "[INFO] Build complete. Run with ./webview_test"