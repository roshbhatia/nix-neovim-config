#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."
mkdir -p bin
echo "Tidying modules..."
go mod tidy
echo "Building Go plugin (sysinit-nvim-core)..."
go build -o bin/sysinit-nvim-core ./cmd/plugin
echo "Build complete: $(pwd)/bin/sysinit-nvim-core"
