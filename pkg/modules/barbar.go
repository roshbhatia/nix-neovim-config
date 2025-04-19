package modules

import (
   "github.com/roshbhatia/nix-neovim-config/pkg/core"
)

// BarbarModule configures the barbar.nvim bufferline and key mappings.
type BarbarModule struct{}

// NewBarbarModule constructs a new BarbarModule.
func NewBarbarModule() *BarbarModule {
   return &BarbarModule{}
}

// Name returns the module name.
func (m *BarbarModule) Name() string {
   return "barbar"
}

// Setup runs the barbar setup and registers which-key mappings.
func (m *BarbarModule) Setup(ctx *core.Context) error {
   // Disable auto setup
   if err := ctx.Command(`lua vim.g.barbar_auto_setup = false`); err != nil {
       return err
   }
   // Configure barbar
   cfg := `lua << EOF
require("barbar").setup({
  animation = true,
  tabpages = true,
  clickable = true,
  auto_hide = false,
  hide = {extensions = true, inactive = false},
  highlight_visible = true,
  highlight_alternate = true,
  highlight_inactive_file_icons = true,
  icons = {
    buffer_index = false,
    buffer_number = false,
    button = '',
    diagnostics = {
      [vim.diagnostic.severity.ERROR] = {enabled = true, icon = ' '},
      [vim.diagnostic.severity.WARN] = {enabled = true, icon = ' '},
      [vim.diagnostic.severity.INFO] = {enabled = false},
      [vim.diagnostic.severity.HINT] = {enabled = false},
    },
    gitsigns = { added = {enabled = true, icon = '+'}, changed = {enabled = true, icon = '~'}, deleted = {enabled = true, icon = '-'} },
    filetype = {custom_colors = false, enabled = true},
    separator = {left = '▎', right = ''},
    separator_at_end = true,
    modified = {button = '●'},
    pinned = {button = '', filename = true},
    preset = 'powerline',
    alternate = {filetype = {enabled = false}},
    current = {buffer_index = true},
    inactive = {button = '×'},
    visible = {modified = {buffer_number = false}},
  },
  insert_at_end = false,
  insert_at_start = false,
  maximum_padding = 1,
  minimum_padding = 1,
  maximum_length = 30,
  minimum_length = 0,
  semantic_letters = true,
  sidebar_filetypes = {
    NvimTree = true,
    undotree = { text = 'undotree' },
    ['neo-tree'] = { event = 'BufWipeout' },
    Outline = { event = 'BufWinLeave', text = 'symbols-outline' },
    alpha = { event = 'BufWinLeave' },
  },
  letters = 'asdfjkl;ghnmxcvbziowerutyqpASDFJKLGHNMXCVBZIOWERUTYQP',
  sort = {ignore_case = true},
  focus_on_close = 'left',
})
local bufferline_api = require('barbar.api')
vim.api.nvim_create_autocmd("User", {
  pattern = "AlphaReady",
  callback = function() vim.g.barbar_auto_setup_events = false end,
})
vim.api.nvim_create_autocmd("BufUnload", {
  pattern = "alpha",
  callback = function()
    vim.g.barbar_auto_setup_events = true
    bufferline_api.update()
  end,
})
EOF`
   if err := ctx.Command(cfg); err != nil {
       return err
   }
   // Register which-key mappings
   wk := `lua << EOF
local wk = require("which-key")
wk.add({
  { "<leader>b", group = "📑 Buffer", icon = { icon = "📑" } },
  { "<leader>bc", "<cmd>BufferClose<CR>", desc = "Close Buffer" },
  { "<leader>ba", "<cmd>BufferCloseAllButCurrent<CR>", desc = "Close All But Current" },
  { "<leader>bv", "<cmd>BufferCloseAllButVisible<CR>", desc = "Close All But Visible" },
  { "<leader>b,", "<cmd>BufferPrevious<CR>", desc = "Previous Buffer" },
  { "<leader>b.", "<cmd>BufferNext<CR>", desc = "Next Buffer" },
  { "<leader>b<", "<cmd>BufferMovePrevious<CR>", desc = "Move Buffer Left" },
  { "<leader>b>", "<cmd>BufferMoveNext<CR>", desc = "Move Buffer Right" },
})
EOF`
   if err := ctx.Command(wk); err != nil {
       return err
   }
   return nil
}