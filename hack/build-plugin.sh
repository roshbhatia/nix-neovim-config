#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."
mkdir -p bin
echo "Tidying modules..."
go mod tidy
echo "Building Go plugin (sysinit-nvim-core)..."
go build -o bin/sysinit-nvim-core ./cmd/plugin
echo "Build complete: $(pwd)/bin/sysinit-nvim-core"
echo "Generating remote plugin manifest..."
bin/sysinit-nvim-core -manifest=sysinit-nvim-core > plugin/manifest.tmp

# Inject manifest into stub between markers
awk '\
  /" MANIFEST-BEGIN/ {\
    print;\
    while ((getline line < "plugin/manifest.tmp") > 0) print "  " line;\
    close("plugin/manifest.tmp"); skip=1; next;\
  }\
  /" MANIFEST-END/ { print; skip=0; next; }\
  skip { next; }\
  { print; }' \
  plugin/sysinit-nvim-core.vim > plugin/sysinit-nvim-core.vim.tmp
mv plugin/sysinit-nvim-core.vim.tmp plugin/sysinit-nvim-core.vim
rm plugin/manifest.tmp
echo "Manifest updated in plugin/sysinit-nvim-core.vim"
