package modules

import (
   "fmt"
   "github.com/roshbhatia/nix-neovim-config/pkg/core"
)

// NvimTreeModule configures the nvim-tree file explorer and its key mappings.
type NvimTreeModule struct{}

// NewNvimTreeModule constructs a new NvimTreeModule.
func NewNvimTreeModule() *NvimTreeModule {
   return &NvimTreeModule{}
}

// Name returns the module name.
func (m *NvimTreeModule) Name() string {
   return "nvim-tree"
}

// Setup runs the nvim-tree setup and registers key mappings via which-key.
func (m *NvimTreeModule) Setup(ctx *core.Context) error {
   // Disable netrw
   if err := ctx.Command(`lua vim.g.loaded_netrw = 1`); err != nil {
       return err
   }
   if err := ctx.Command(`lua vim.g.loaded_netrwPlugin = 1`); err != nil {
       return err
   }
   // Plugin setup
   setup := `lua << EOF
require("nvim-tree").setup({
  sync_root_with_cwd = true,
  respect_buf_cwd = true,
  update_focused_file = {
    enable = true,
    update_root = true,
  },
  view = { width = 30 },
  renderer = { group_empty = true },
  filters = { dotfiles = true },
})
EOF`
   if err := ctx.Command(setup); err != nil {
       return err
   }
   // Which-key group
   if err := ctx.Command(
       `lua require("which-key").register({ ["<leader>e"] = { name = "Explorer", icon = { icon = "🌲", hl = "WhichKeyIconGreen" } } }, { mode = "n" })`,
   ); err != nil {
       return err
   }
   // Mappings: Toggle, Find, Refresh, Collapse, Copy Path
   mappings := []struct{ lhs, rhs, desc string }{
       {"<leader>ee", "<cmd>NvimTreeToggle<CR>", "Toggle NvimTree"},
       {"<leader>ef", "<cmd>NvimTreeFindFile<CR>", "Find Current File"},
       {"<leader>er", "<cmd>NvimTreeRefresh<CR>", "Refresh NvimTree"},
       {"<leader>ec", "<cmd>NvimTreeCollapse<CR>", "Collapse NvimTree"},
   }
   opts := map[string]bool{"noremap": true, "silent": true}
   for _, m2 := range mappings {
       // which-key registration
       reg := fmt.Sprintf(
           `lua require("which-key").register({ ["%s"] = "%s" }, { mode = "n" })`,
           m2.lhs, m2.desc,
       )
       if err := ctx.Command(reg); err != nil {
           return err
       }
       // key mapping
       if err := ctx.Map("n", m2.lhs, m2.rhs, opts); err != nil {
           return err
       }
   }
   // Copy node path mapping separately for Lua function
   copyCmd := `lua vim.keymap.set("n", "<leader>ep", function()
  local api = require("nvim-tree.api")
  local node = api.tree.get_node_under_cursor()
  require("nvim-tree.actions.clipboard.clipboard").copy(node)
end, { noremap = true, silent = true, desc = "Copy Node Path" })`
   if err := ctx.Command(copyCmd); err != nil {
       return err
   }
   return nil
}