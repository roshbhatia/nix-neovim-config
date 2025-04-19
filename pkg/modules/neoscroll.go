package modules

import "github.com/roshbhatia/nix-neovim-config/pkg/core"

// NeoscrollModule configures karb94/neoscroll.nvim plugin.
type NeoscrollModule struct{}

// NewNeoscrollModule constructs a new NeoscrollModule.
func NewNeoscrollModule() *NeoscrollModule {
   return &NeoscrollModule{}
}

// Name returns the module name.
func (m *NeoscrollModule) Name() string {
   return "neoscroll"
}

// Setup runs the Neoscroll setup function.
func (m *NeoscrollModule) Setup(ctx *core.Context) error {
   setupCmd := `lua << EOF
require("neoscroll").setup()
EOF`
   return ctx.Command(setupCmd)
}