#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."
mkdir -p bin
echo "Tidying modules..."
go mod tidy
echo "Building Go plugin..."
go build -o bin/go-plugin ./cmd/plugin
echo "Build complete: $(pwd)/bin/go-plugin"
