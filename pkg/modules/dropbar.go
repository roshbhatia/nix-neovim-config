package modules

import (
   "github.com/roshbhatia/nix-neovim-config/pkg/core"
)

// DropbarModule configures Bekaboo/dropbar.nvim plugin and keybindings.
type DropbarModule struct{}

// NewDropbarModule constructs a new DropbarModule.
func NewDropbarModule() *DropbarModule { return &DropbarModule{} }

// Name returns the module name.
func (m *DropbarModule) Name() string { return "dropbar" }

// Setup configures dropbar and registers keymaps and highlights.
func (m *DropbarModule) Setup(ctx *core.Context) error {
   // Configure dropbar
   setupCmd := `lua << EOF
local dropbar = require("dropbar")
local utils = require("dropbar.utils")
local sources = require("dropbar.sources")

dropbar.setup({
  bar = {
    enable = function(buf, win, _)
      if not vim.api.nvim_buf_is_valid(buf)
        or not vim.api.nvim_win_is_valid(win)
        or vim.fn.win_gettype(win) ~= ''
        or vim.wo[win].winbar ~= ''
        or vim.bo[buf].ft == 'help' then
        return false
      end
      local stat = vim.uv.fs_stat(vim.api.nvim_buf_get_name(buf))
      if stat and stat.size > 1024 * 1024 then
        return false
      end
      return vim.bo[buf].ft == 'markdown'
        or pcall(vim.treesitter.get_parser, buf)
        or not vim.tbl_isempty(vim.lsp.get_clients({ bufnr = buf, method = 'textDocument/documentSymbol' }))
    end,
    update_debounce = 100,
    sources = function(buf, _)
      if vim.bo[buf].ft == 'markdown' then
        return { sources.path, sources.markdown }
      end
      if vim.bo[buf].ft == 'terminal' then
        return { sources.terminal }
      end
      if vim.bo[buf].ft == 'oil' then
        return { sources.path }
      end
      return { sources.path, utils.source.fallback({ sources.lsp, sources.treesitter }) }
    end,
    padding = { left = 1, right = 1 },
    pick = { pivots = 'abcdefghijklmnopqrstuvwxyz' },
    truncate = true,
    hover = true,
  },
  menu = {
    quick_navigation = true,
    entry = { padding = { left = 1, right = 1 } },
    preview = true,
    hover = true,
    win_configs = { border = "rounded", style = "minimal" },
    scrollbar = { enable = true, background = true },
  },
  symbol = {
    preview = { reorient = function(win, range)
      vim.api.nvim_set_current_win(win)
      vim.api.nvim_win_set_cursor(win, { range.start.line + 1, range.start.character })
      vim.cmd("normal! zz")
      vim.api.nvim_win_set_cursor(0, { range.start.line + 1, range.start.character })
    end },
    jump = { reorient = function(win, range)
      vim.api.nvim_set_current_win(win)
      vim.api.nvim_win_set_cursor(win, { range.start.line + 1, range.start.character })
      vim.cmd("normal! zz")
    end },
  },
  sources = {
    path = { max_depth = 3, filter = function(name) return name ~= "." end, modified = function(sym)
      if vim.bo.modified then
        return sym:merge({ name = sym.name .. " [+]", name_hl = "DiffAdd" })
      end
      return sym
    end },
    treesitter = { max_depth = 16 },
    lsp = { max_depth = 16 },
  },
})
EOF`
   if err := ctx.Command(setupCmd); err != nil {
      return err
   }
   // Register which-key mappings and highlights
   wkCmd := `lua << EOF
local wk = require("which-key")
wk.add({
  { "<leader>d", group = "📊 Dropbar" },
  { "<leader>dp", function() require("dropbar.api").pick() end, desc = "Pick Symbols in Winbar", mode = "n" },
  { "<leader>dg", function() require("dropbar.api").goto_context_start() end, desc = "Go to Context Start", mode = "n" },
  { "<leader>dn", function() require("dropbar.api").select_next_context() end, desc = "Select Next Context", mode = "n" },
})
vim.api.nvim_create_autocmd("ColorScheme", {
  pattern = "*",
  callback = function()
    vim.api.nvim_set_hl(0, "DropBarCurrentContext", { link = "Visual" })
    vim.api.nvim_set_hl(0, "DropBarIconKindArray", { link = "Operator" })
    vim.api.nvim_set_hl(0, "DropBarIconKindFunction", { link = "Function" })
    vim.api.nvim_set_hl(0, "DropBarIconKindClass", { link = "Type" })
    vim.api.nvim_set_hl(0, "DropBarMenuNormalFloat", { link = "NormalFloat" })
    vim.api.nvim_set_hl(0, "DropBarMenuFloatBorder", { link = "FloatBorder" })
  end,
})
vim.cmd("doautocmd ColorScheme")
EOF`
   if err := ctx.Command(wkCmd); err != nil {
      return err
   }
   return nil
}