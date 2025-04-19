package modules

import (
   "fmt"
   "github.com/roshbhatia/nix-neovim-config/pkg/core"
)

// OilModule configures the stevearc/oil.nvim plugin and its key mappings.
type OilModule struct{}

// NewOilModule constructs a new OilModule.
func NewOilModule() *OilModule {
   return &OilModule{}
}

// Name returns the module name.
func (m *OilModule) Name() string {
   return "oil"
}

// Setup registers oil.nvim setup and which-key mappings.
func (m *OilModule) Setup(ctx *core.Context) error {
   // oil.nvim configuration
   setupCmd := `lua << EOF
require("oil").setup({
  default_file_explorer = true,
  columns = {"icon", "size", "mtime"},
  view_options = {
    show_hidden = true,
    is_hidden_file = function(name)
      return vim.startswith(name, ".")
    end,
  },
  float = {
    border = "rounded",
    max_width = 80,
    max_height = 30,
  },
})
EOF`
   if err := ctx.Command(setupCmd); err != nil {
       return err
   }
   // Map '-' to open parent directory
   if err := ctx.Map("n", "-", "<CMD>Oil<CR>", map[string]bool{"noremap": true, "silent": true}); err != nil {
       return err
   }
   // Which-key group and mappings
   // Register group <leader>o
   if err := ctx.Command(
       `lua require("which-key").register({ ["<leader>o"] = { name = "Oil" } }, { mode = "n" })`,
   ); err != nil {
       return err
   }
   // Register <leader>of and <leader>oh mappings
   mappings := []struct{ lhs, rhs, desc string }{
       {"<leader>of", "<cmd>Oil --float<cr>", "Open Oil in Float"},
       {"<leader>oh", "lua require('oil').open(vim.fn.expand('~'))", "Open Home Directory"},
   }
   opts := map[string]bool{"noremap": true, "silent": true}
   for _, m2 := range mappings {
       // Register which-key
       if err := ctx.Command(
           fmt.Sprintf(`lua require("which-key").register({ ["%s"] = "%s" }, { mode = "n" })`, m2.lhs, m2.desc),
       ); err != nil {
           return err
       }
       // Map key
       if err := ctx.Map("n", m2.lhs, m2.rhs, opts); err != nil {
           return err
       }
   }
   return nil
}