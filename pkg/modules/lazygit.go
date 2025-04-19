package modules

import (
  "github.com/roshbhatia/nix-neovim-config/pkg/core"
)

// LazygitModule configures kdheepak/lazygit.nvim and key mappings.
type LazygitModule struct{}

// NewLazygitModule constructs a new LazygitModule.
func NewLazygitModule() *LazygitModule {
  return &LazygitModule{}
}

// Name returns the module name.
func (m *LazygitModule) Name() string {
  return "lazygit"
}

// Setup configures lazygit plugin settings and which-key mappings.
func (m *LazygitModule) Setup(ctx *core.Context) error {
  // Configure floating window settings and use plenary
  cfgCmd := `lua << EOF
vim.g.lazygit_floating_window_winblend = 0
vim.g.lazygit_floating_window_scaling_factor = 0.9
vim.g.lazygit_floating_window_border_chars = { '╭', '─', '╮', '│', '╯', '─', '╰', '│' }
vim.g.lazygit_floating_window_use_plenary = 0
vim.g.lazygit_use_neovim_remote = 1
-- Load Telescope extension if available
pcall(function() require('telescope').load_extension('lazygit') end)
EOF`
  if err := ctx.Command(cfgCmd); err != nil {
    return err
  }
  // Which-key mappings under <leader>g
  if err := ctx.Command(`lua require('which-key').register({ ['<leader>g'] = { name = 'Git' } }, { mode = 'n' })`); err != nil {
    return err
  }
  opts := map[string]bool{"noremap": true, "silent": true}
  mappings := []struct{ lhs, rhs, desc string }{
    {"<leader>gg", "<cmd>LazyGit<CR>", "Open LazyGit"},
    {"<leader>gc", "<cmd>LazyGitCurrentFile<CR>", "LazyGit (Current File)"},
    {"<leader>gf", "<cmd>LazyGitFilter<CR>", "LazyGit (Filter)"},
    {"<leader>gF", "<cmd>LazyGitFilterCurrentFile<CR>", "LazyGit (Filter File)"},
    {"<leader>gC", "<cmd>LazyGitConfig<CR>", "Edit LazyGit Config"},
  }
  for _, m2 := range mappings {
    if err := ctx.Map("n", m2.lhs, m2.rhs, opts); err != nil {
      return err
    }
  }
  return nil
}