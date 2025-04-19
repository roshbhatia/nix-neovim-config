package modules

import "github.com/roshbhatia/nix-neovim-config/pkg/core"

// LspZeroModule sets up basic LSP functionality via lsp-zero.
type LspZeroModule struct{}

// NewLspZeroModule constructs a new LspZeroModule.
func NewLspZeroModule() *LspZeroModule { return &LspZeroModule{} }

// Name returns the module name.
func (m *LspZeroModule) Name() string { return "lsp-zero" }

// Setup is a placeholder for LSP zero configuration.
// TODO: add full LSP, mason, and keybinding setup.
func (m *LspZeroModule) Setup(ctx *core.Context) error {
   // Basic LSP-zero preset
   return ctx.Command(`lua require('lsp-zero').preset({ name = 'recommended', set_lsp_keymaps = true, manage_nvim_cmp = false })`)
}