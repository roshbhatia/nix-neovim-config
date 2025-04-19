package modules

import (
   "github.com/roshbhatia/nix-neovim-config/pkg/core"
)

// ThemifyModule sets up which-key mappings for the themify theme picker.
type ThemifyModule struct{}

// NewThemifyModule constructs a new ThemifyModule.
func NewThemifyModule() *ThemifyModule {
   return &ThemifyModule{}
}

// Name returns the module name.
func (m *ThemifyModule) Name() string {
   return "themify"
}

// Setup registers which-key group and mappings for the themify theme switcher.
func (m *ThemifyModule) Setup(ctx *core.Context) error {
   // Register themify keybindings via which-key
   setup := `lua << EOF
local wk = require("which-key")
wk.add({
  { "<leader>t", group = "🎨 Theme", icon = { icon = "🎨" } },
  { "<leader>tt", "<cmd>Themify<CR>", desc = "Open Theme Switcher" },
  { "<leader>tr", "<cmd>ThemifyReload<CR>", desc = "Reload Themes" },
  { "<leader>ti", "<cmd>ThemifyInstall<CR>", desc = "Install Missing Themes" },
})
EOF`
   if err := ctx.Command(setup); err != nil {
       return err
   }
   return nil
}