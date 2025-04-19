package modules

import (
   "github.com/roshbhatia/nix-neovim-config/pkg/core"
)

// NeominimapModule configures Isrothy/neominimap.nvim with rightmost-window logic.
type NeominimapModule struct{}

// NewNeominimapModule constructs a new NeominimapModule.
func NewNeominimapModule() *NeominimapModule { return &NeominimapModule{} }

// Name returns the module name.
func (m *NeominimapModule) Name() string { return "neominimap" }

// Setup configures the minimap, highlight groups, and toggling logic.
func (m *NeominimapModule) Setup(ctx *core.Context) error {
   cmd := `lua << EOF
local neominimap = require("neominimap")
---@type Neominimap.UserConfig
local config = {
  auto_enable = false,
  log_level = vim.log.levels.OFF,
  notification_level = vim.log.levels.INFO,
  exclude_filetypes = {"help","bigfile","NvimTree","neo-tree","lazy","mason","oil","alpha","dashboard","TelescopePrompt","toggleterm"},
  exclude_buftypes = {"nofile","nowrite","quickfix","terminal","prompt"},
  x_multiplier = 4,
  y_multiplier = 1,
  layout = "float",
  split = { minimap_width = 20, fix_width = true, direction = "right", close_if_last_window = true },
  float = { minimap_width = 20, max_minimap_height = nil, margin = { right = 0, top = 0, bottom = 0 }, z_index = 10, window_border = "single" },
  delay = 200,
  sync_cursor = true,
  click = { enabled = true, auto_switch_focus = false },
  diagnostic = { enabled = true, severity = vim.diagnostic.severity.WARN, mode = "line" },
  git = { enabled = true, mode = "sign" },
  treesitter = { enabled = true },
  search = { enabled = true, mode = "line" },
  mark = { enabled = true, mode = "icon", show_builtins = false },
  fold = { enabled = true },
  winopt = function(opt, winid) opt.winblend = 10 end,
  bufopt = function(opt, bufnr) end,
}
vim.g.neominimap = config
-- lualine extension
local ok, _ = pcall(require, "lualine")
if ok then
  local ext = require("neominimap.statusline").lualine_default
  require('lualine').setup({ extensions = vim.list_extend(require('lualine').extensions or {}, { ext }) })
end
-- which-key mappings
local wk = require("which-key")
wk.add({
  { "<leader>m", group = "Minimap", icon = { icon = "🗺️", hl = "WhichKeyIconBlue" } },
  { "<leader>mm", "<cmd>Neominimap toggle<CR>", desc = "Toggle Minimap", mode = "n" },
  { "<leader>mo", "<cmd>Neominimap on<CR>", desc = "Enable Minimap", mode = "n" },
  { "<leader>mc", "<cmd>Neominimap off<CR>", desc = "Disable Minimap", mode = "n" },
  { "<leader>mr", "<cmd>Neominimap refresh<CR>", desc = "Refresh Minimap", mode = "n" },
  { "<leader>mf", "<cmd>Neominimap focus<CR>", desc = "Focus Minimap", mode = "n" },
  { "<leader>mu", "<cmd>Neominimap unfocus<CR>", desc = "Unfocus Minimap", mode = "n" },
  { "<leader>ms", "<cmd>Neominimap toggleFocus<CR>", desc = "Switch Minimap Focus", mode = "n" },
})
-- highlight groups
vim.api.nvim_create_autocmd("ColorScheme", { pattern = "*", callback = function()
  vim.api.nvim_set_hl(0, "NeominimapBackground", { link = "Normal" })
  vim.api.nvim_set_hl(0, "NeominimapBorder", { link = "FloatBorder" })
  vim.api.nvim_set_hl(0, "NeominimapCursorLine", { link = "CursorLine" })
  vim.api.nvim_set_hl(0, "NeominimapErrorLine", { link = "DiagnosticVirtualTextError" })
  vim.api.nvim_set_hl(0, "NeominimapWarnLine", { link = "DiagnosticVirtualTextWarn" })
  vim.api.nvim_set_hl(0, "NeominimapInfoLine", { link = "DiagnosticVirtualTextInfo" })
  vim.api.nvim_set_hl(0, "NeominimapHintLine", { link = "DiagnosticVirtualTextHint" })
  vim.api.nvim_set_hl(0, "NeominimapGitAddLine", { link = "DiffAdd" })
  vim.api.nvim_set_hl(0, "NeominimapGitChangeLine", { link = "DiffChange" })
  vim.api.nvim_set_hl(0, "NeominimapGitDeleteLine", { link = "DiffDelete" })
  vim.api.nvim_set_hl(0, "NeominimapSearchLine", { link = "Search" })
end })
-- rightmost window logic
local function is_rightmost_window()
  local mux = require('smart-splits.mux').get()
  if mux and mux.current_pane_at_edge then
    return mux.current_pane_at_edge('right')
  else
    local win = vim.api.nvim_get_current_win()
    local wins = vim.api.nvim_tabpage_list_wins(0)
    local pos = {}
    for _, id in ipairs(wins) do
      local p = vim.api.nvim_win_get_position(id)
      local w = vim.api.nvim_win_get_width(id)
      pos[id] = p[2] + w
    end
    for id, edge in pairs(pos) do
      if id ~= win and edge > pos[win] then return false end
    end
    return true
  end
end
vim.api.nvim_create_augroup("NeominimapRightmost", { clear = true })
vim.api.nvim_create_autocmd({"WinEnter","BufEnter","VimResized"}, {
  group = "NeominimapRightmost",
  callback = function()
    if is_rightmost_window() and vim.g.minimap_enabled then
      vim.cmd("Neominimap winOn")
    else
      vim.cmd("Neominimap winOff")
    end
  end,
})
vim.g.minimap_enabled = false
-- custom commands
vim.api.nvim_create_user_command("RightmostMinimapToggle", function()
  vim.g.minimap_enabled = not vim.g.minimap_enabled
  if is_rightmost_window() and vim.g.minimap_enabled then vim.cmd("Neominimap winOn") else vim.cmd("Neominimap winOff") end
end, {})
vim.api.nvim_create_user_command("RightmostMinimapOn", function()
  vim.g.minimap_enabled = true
  if is_rightmost_window() then vim.cmd("Neominimap winOn") end
end, {})
vim.api.nvim_create_user_command("RightmostMinimapOff", function()
  vim.g.minimap_enabled = false
  vim.cmd("Neominimap winOff")
end, {})
-- toggle keymaps
vim.keymap.set("n", "<leader>mm", "<cmd>RightmostMinimapToggle<CR>", { desc = "Toggle Rightmost Minimap" })
vim.keymap.set("n", "<leader>mo", "<cmd>RightmostMinimapOn<CR>", { desc = "Enable Rightmost Minimap" })
vim.keymap.set("n", "<leader>mc", "<cmd>RightmostMinimapOff<CR>", { desc = "Disable Minimap" })
EOF`
   return ctx.Command(cmd)
}