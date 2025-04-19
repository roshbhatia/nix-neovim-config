package modules

import (
   "github.com/roshbhatia/nix-neovim-config/pkg/core"
)

// WilderModule configures gelguy/wilder.nvim and keymap integration.
type WilderModule struct{}

// NewWilderModule constructs a new WilderModule.
func NewWilderModule() *WilderModule { return &WilderModule{} }

// Name returns the module name.
func (m *WilderModule) Name() string { return "wilder" }

// Setup initializes wilder and which-key mapping.
func (m *WilderModule) Setup(ctx *core.Context) error {
   cmd := `lua << EOF
local wilder = require("wilder")
wilder.setup({ modes = { ":", "/", "?" } })
local wk = require("which-key")
wk.add({ { "<leader>:", ":", desc = "Command-line (Wilder)", mode = "n" } })
EOF`
   return ctx.Command(cmd)
}