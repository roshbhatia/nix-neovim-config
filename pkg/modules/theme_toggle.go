package modules

import (
   "fmt"
   "github.com/roshbhatia/nix-neovim-config/pkg/core"
)

// ThemeToggleModule provides light/dark theme mappings for both VSCode and Neovim
type ThemeToggleModule struct{}

// NewThemeToggleModule constructs a new ThemeToggleModule.
func NewThemeToggleModule() *ThemeToggleModule {
   return &ThemeToggleModule{}
}

// Name returns the module name.
func (m *ThemeToggleModule) Name() string {
   return "theme-toggle"
}

// Setup registers which-key group and key mappings for theme toggling.
func (m *ThemeToggleModule) Setup(ctx *core.Context) error {
   // Register a which-key group for theme
   if err := ctx.Command(fmt.Sprintf(
       `lua require('which-key').register({ ['%s'] = { name = '%s' } }, { mode = 'n' })`,
       "<leader>t", "🎨 Theme")); err != nil {
       return err
   }
   // Define light level mappings
   type mappingDef struct {
       suffix    string
       desc      string
       vscodeAct string
       neovimAct string
   }
   mappings := []mappingDef{
       {"l", "Light", "workbench.action.selectColorTheme", "require('themify.api').set_current(nil, 'tokyonight-day')"},
       {"d", "Dark",  "workbench.action.selectColorTheme", "require('themify.api').set_current(nil, 'tokyonight-night')"},
   }
   // Register each mapping
   for _, def := range mappings {
       lhs := "<leader>t" + def.suffix
       // Register mapping in which-key
       if err := ctx.Command(fmt.Sprintf(
           `lua require('which-key').register({ ['%s'] = '%s' }, { mode = 'n' })`, lhs, def.desc)); err != nil {
           return err
       }
       // Map key to VSCode action or Neovim themify call
       rhsLua := fmt.Sprintf(
           "if vim.g.vscode then require('vscode').action('%s') else %s end",
           def.vscodeAct, def.neovimAct,
       )
       rhs := fmt.Sprintf("<cmd>lua %s<CR>", rhsLua)
       if err := ctx.Map("n", lhs, rhs, map[string]bool{"noremap": true, "silent": true}); err != nil {
           return err
       }
   }
   return nil
}