#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."
echo "Tidying modules..."
go mod tidy
echo "Building Go plugin..."
go build -o go-plugin ./cmd/plugin
echo "Build complete: $(pwd)/go-plugin"
