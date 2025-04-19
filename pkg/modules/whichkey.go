package modules

import (
   "fmt"
   "github.com/roshbhatia/nix-neovim-config/pkg/core"
)

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
   // Configure which-key plugin
   if err := ctx.SetOption("timeout", true); err != nil {
       return err
   }
   if err := ctx.SetOption("timeoutlen", 300); err != nil {
       return err
   }
   // Setup plugin via Lua
   setupCmd := `lua << EOF
require("which-key").setup({
  plugins = {
    marks = true,
    registers = true,
    spelling = { enabled = true, suggestions = 20 },
    presets = { operators=true, motions=true, text_objects=true, windows=true, nav=true, z=true, g=true },
  },
  win = { border = "rounded", padding = {2,2,2,2} },
  layout = { spacing = 3 },
  icons = { breadcrumb = "»", separator = "➜", group = "+" },
  show_help = true,
  show_keys = true,
  triggers = {{ "<auto>", mode = "nxsotc" }},
})
EOF`
   if err := ctx.Command(setupCmd); err != nil {
       return err
   }
   // Register which-key groups
   groups := map[string]string{
       "<leader>f": "Find/Files",
       "<leader>b": "Buffers",
       "<leader>w": "Windows",
       "<leader>g": "Git",
       "<leader>l": "LSP",
       "<leader>s": "Session",
       "<leader>c": "Comment",
       "<leader>o": "Oil",
       "<leader>e": "Explorer",
       "<leader>m": "Minimap",
   }
   for keys, name := range groups {
       cmd := fmt.Sprintf(`lua require("which-key").register({ ["%s"] = { name = "%s" } }, { mode = "n" })`, keys, name)
       if err := ctx.Command(cmd); err != nil {
           return err
       }
   }
   // Register buffer and command mode mappings
   mappings := []struct{ mode, lhs, rhs string }{
       {"n", "<leader>;", ":"},
       {"n", "<leader>bd", "<cmd>bdelete<cr>"},
       {"n", "<leader>bn", "<cmd>bnext<cr>"},
       {"n", "<leader>bp", "<cmd>bprevious<cr>"},
       {"n", "<leader>bN", "<cmd>enew<cr>"},
       {"n", "<leader>bl", "<cmd>Telescope buffers<cr>"},
   }
   opts := map[string]bool{"noremap": true, "silent": true}
   for _, m := range mappings {
       if err := ctx.Map(m.mode, m.lhs, m.rhs, opts); err != nil {
           return err
       }
   }
   // Window Hydra mapping
   hydraCmd := `lua vim.keymap.set("n", "<leader>W", function() require("which-key").show({ keys = "<c-w>", mode = "n", loop = true }) end, { desc = "Window Hydra Mode" })`
   if err := ctx.Command(hydraCmd); err != nil {
       return err
   }
   return nil
}
