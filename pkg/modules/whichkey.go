package modules

import "github.com/roshbhatia/nix-neovim-config/pkg/core"

// WhichKeyModule provides which-key functionality through Go.
type WhichKeyModule struct{}

// NewWhichKeyModule constructs a new WhichKeyModule.
func NewWhichKeyModule() *WhichKeyModule {
    return &WhichKeyModule{}
}

// Name returns the module name.
func (m *WhichKeyModule) Name() string {
    return "which-key"
}

// Setup registers which-key related keybindings and commands.
func (m *WhichKeyModule) Setup(ctx *core.Context) error {
    // TODO: Implement which-key integration using ctx.Map and ctx.Command
    return nil
}
