package modules

import (
   "github.com/roshbhatia/nix-neovim-config/pkg/core"
)

// WeztermModule configures mrjones2014/smart-splits.nvim for terminal splits.
type WeztermModule struct{}

// NewWeztermModule constructs a new WeztermModule.
func NewWeztermModule() *WeztermModule { return &WeztermModule{} }

// Name returns the module name.
func (m *WeztermModule) Name() string { return "wezterm" }

// Setup initializes smart-splits configuration and keybindings.
func (m *WeztermModule) Setup(ctx *core.Context) error {
   cmd := `lua << EOF
local smart_splits = require("smart-splits")
smart_splits.setup({
  ignored_filetypes = { "NvimTree" },
  multiplexer_integration = "wezterm",
  default_amount = 3,
  at_edge = 'wrap',
  move_cursor_same_row = false,
  cursor_follows_swapped_bufs = true,
  resize_mode = {
    quit_key = '<ESC>',
    resize_keys = { 'h', 'j', 'k', 'l' },
    silent = false,
  },
})
vim.keymap.set('n', '<C-h>', function() smart_splits.move_cursor_left() end, {desc = "Move to left split"})
vim.keymap.set('n', '<C-j>', function() smart_splits.move_cursor_down() end, {desc = "Move to split below"})
vim.keymap.set('n', '<C-k>', function() smart_splits.move_cursor_up() end, {desc = "Move to split above"})
vim.keymap.set('n', '<C-l>', function() smart_splits.move_cursor_right() end, {desc = "Move to right split"})
vim.keymap.set('n', '<C-S-h>', function() smart_splits.resize_left() end, {desc = "Resize split left"})
vim.keymap.set('n', '<C-S-j>', function() smart_splits.resize_down() end, {desc = "Resize split down"})
vim.keymap.set('n', '<C-S-k>', function() smart_splits.resize_up() end, {desc = "Resize split up"})
vim.keymap.set('n', '<C-S-l>', function() smart_splits.resize_right() end, {desc = "Resize split right"})
vim.keymap.set('n', '<C-S-s>', vim.cmd.vsplit, {desc = "Split (Vertical)"})
vim.keymap.set('n', '<C-S-v>', vim.cmd.split, {desc = "Split (Horizontal)"})
EOF`
   return ctx.Command(cmd)
}