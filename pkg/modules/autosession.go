package modules

import (
  "fmt"
  "github.com/roshbhatia/nix-neovim-config/pkg/core"
)

// AutosessionModule configures rmagatti/auto-session and key mappings.
type AutosessionModule struct{}

// NewAutosessionModule constructs a new AutosessionModule.
func NewAutosessionModule() *AutosessionModule {
  return &AutosessionModule{}
}

// Name returns the module name.
func (m *AutosessionModule) Name() string {
  return "autosession"
}

// Setup configures auto-session plugin and mappings.
func (m *AutosessionModule) Setup(ctx *core.Context) error {
  // Plugin setup
  setupCmd := `lua << EOF
require("auto-session").setup({
  log_level = "error",
  auto_session_enable_last_session = false,
  auto_session_root_dir = vim.fn.stdpath("data") .. "/sessions/",
  auto_session_enabled = true,
  auto_save_enabled = true,
  auto_restore_enabled = false,
  continue_restore_on_error = true,
  auto_session_suppress_dirs = {"~/", "~/Projects", "~/Downloads", "/"},
  auto_session_use_git_branch = true,
  auto_restore_last_session = false,
  bypass_save_filetypes = {"NvimTree", "neo-tree", "dashboard", "alpha", "netrw"},
  close_unsupported_windows = true,
  args_allow_single_directory = true,
  cwd_change_handling = {
    enabled = true,
    restore_upcoming_session = true,
    pre_cwd_changed_hook = nil,
    post_cwd_changed_hook = function()
      local ok, lualine = pcall(require, "lualine")
      if ok then lualine.refresh() end
    end,
  },
  session_lens = {
    load_on_setup = true,
    theme_conf = { border = true },
    previewer = false,
    mappings = {
      delete_session = {"i", "<C-d>"},
      alternate_session = {"i", "<C-s>"},
      copy_session = {"i", "<C-y>"},
    },
  },
  pre_save_cmds = {
    function()
      for _, win in ipairs(vim.api.nvim_list_wins()) do
        local cfg = vim.api.nvim_win_get_config(win)
        if cfg.relative ~= "" then vim.api.nvim_win_close(win, false) end
      end
      local tree_ok, tree = pcall(require, "nvim-tree.api")
      if tree_ok then tree.tree.close() end
    end,
  },
  post_restore_cmds = {
    function()
      local tree_ok, tree = pcall(require, "nvim-tree.api")
      if tree_ok and vim.g.nvim_tree_auto_open_on_session_restore then tree.tree.open() end
      local ok, lualine = pcall(require, "lualine")
      if ok then lualine.refresh() end
    end,
  },
})
EOF`
  if err := ctx.Command(setupCmd); err != nil {
    return err
  }
  // Which-key group
  if err := ctx.Command(`lua require("which-key").register({ ["<leader>s"] = { name = "Session" } }, { mode = "n" })`); err != nil {
    return err
  }
  // Key mappings
  opts := map[string]bool{"noremap": true, "silent": true}
  mappings := []struct{ lhs, rhs, desc string }{
    {"<leader>ss", "<cmd>SessionSave<CR>", "Save Session"},
    {"<leader>sl", "<cmd>SessionRestore<CR>", "Load Session"},
    {"<leader>sd", "<cmd>SessionDelete<CR>", "Delete Session"},
    {"<leader>sp", "<cmd>SessionPurgeOrphaned<CR>", "Purge Orphaned Sessions"},
  }
  for _, m2 := range mappings {
    if err := ctx.Command(fmt.Sprintf(`lua require("which-key").register({ ["%s"] = "%s" }, { mode = "n" })`, m2.lhs, m2.desc)); err != nil {
      return err
    }
    if err := ctx.Map("n", m2.lhs, m2.rhs, opts); err != nil {
      return err
    }
  }
  // Session options
  if err := ctx.Command(`lua vim.o.sessionoptions = "blank,buffers,curdir,folds,help,tabpages,winsize,winpos,terminal,localoptions"`); err != nil {
    return err
  }
  // Auto open Alpha on VimEnter
  autocmd := `lua << EOF
vim.api.nvim_create_autocmd("VimEnter", {
  callback = vim.schedule_wrap(function()
    local argc = vim.fn.argc()
    local bufnr = vim.api.nvim_get_current_buf()
    local bufname = vim.api.nvim_buf_get_name(bufnr)
    local buftype = vim.bo[bufnr].ft
    if argc == 0 and bufname == "" and buftype ~= "directory" then
      local ok = pcall(function()
        require("alpha")
        vim.cmd("silent! %bd")
        vim.cmd("Alpha")
      end)
      if not ok then end
    end
  end),
})
EOF`
  if err := ctx.Command(autocmd); err != nil {
    return err
  }
  return nil
}