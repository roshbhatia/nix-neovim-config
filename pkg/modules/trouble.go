package modules

import (
  "github.com/roshbhatia/nix-neovim-config/pkg/core"
)

// TroubleModule configures folke/trouble.nvim diagnostics list plugin.
type TroubleModule struct{}

// NewTroubleModule constructs a new TroubleModule.
func NewTroubleModule() *TroubleModule {
  return &TroubleModule{}
}

// Name returns the module name.
func (m *TroubleModule) Name() string {
  return "trouble"
}

// Setup runs trouble.nvim setup.
func (m *TroubleModule) Setup(ctx *core.Context) error {
  cmd := `lua << EOF
require("trouble").setup({
  icons = true,
  fold_open = "",
  fold_closed = "",
  signs = { error = "E", warning = "W", hint = "H", information = "I", other = "O" },
  use_diagnostic_signs = true,
})
EOF`
  return ctx.Command(cmd)
}