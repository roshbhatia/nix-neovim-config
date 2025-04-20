vim.g.vscode = (vim.fn.exists("g:vscode") == 1) or (vim.env.VSCODE_GIT_IPC_HANDLE ~= nil)
-- Ensure the Go remote plugin binary is discoverable for registration
local config_bin = vim.fn.stdpath("config") .. "/bin"
vim.env.PATH = config_bin .. ":" .. vim.env.PATH

local lazypath = vim.fn.stdpath("data") .. "/lazy/lazy.nvim"
if not vim.loop.fs_stat(lazypath) then
  vim.fn.system({
    "git",
    "clone",
    "--filter=blob:none",
    "https://github.com/folke/lazy.nvim.git",
    "--branch=stable",
    lazypath,
  })
end
vim.opt.rtp:prepend(lazypath)

local init_dir = vim.fn.fnamemodify(vim.fn.expand("$MYVIMRC"), ":p:h")
local lua_dir = init_dir .. "/lua"
vim.opt.rtp:prepend(lua_dir)

vim.keymap.set({ "n", "v" }, "<Space>", "<Nop>", { noremap = true, silent = true })

vim.g.mapleader = " "
vim.g.maplocalleader = " "

vim.opt.clipboard = "unnamedplus"

vim.keymap.set("n", ":", ":", { noremap = true, desc = "Command mode" })


local function setup_neovim_settings()
  vim.opt.number = true
  vim.opt.cursorline = true
  vim.opt.signcolumn = "yes"
  vim.opt.termguicolors = true
  vim.opt.showmode = false -- Hide mode since we use lualine
  vim.opt.lazyredraw = true

  vim.opt.foldmethod = "expr"
  vim.opt.foldexpr = "nvim_treesitter#foldexpr()"
  vim.opt.foldenable = false
  vim.opt.foldlevel = 99

  vim.opt.pumheight = 10 -- Limit completion menu height
  vim.opt.cmdheight = 1 -- More space for displaying messages
  vim.opt.hidden = true -- Enable background buffers
  vim.opt.showtabline = 2 -- Always show tabline
  vim.opt.shortmess:append("c") -- Don't show completion messages
  vim.opt.completeopt = { "menuone", "noselect" }
end

local function setup_vscode_settings()
  vim.notify("SysInit -- VSCode Neovim integration detected", vim.log.levels.INFO)
end

-- Neovim vs VSCode specific settings
if vim.g.vscode then
  setup_vscode_settings()
else
  setup_neovim_settings()
end

local module_loader = require("core.module_loader")

-- determine module loading system
local module_system

if not vim.g.vscode then
  module_system = {
    -- UI-related modules (load first)
    ui = {
      "wezterm",
      "devicons",
      "nvimtree",
      "dropbar",
      "lualine",
      "neominimap",
      "barbar",
      "themify",
    },
    -- Core editor functionality
    editor = {
      "wilder",
    },
    -- Tool modules
    tools = {
      "comment",
      "hop",
      "neoscroll",
      "treesitter",
      "cmp",
      "conform",
      "lazygit",
      "lsp-zero",
      "nvim-lint",
      "copilot",
      "copilot-chat",
      "copilot-cmp",
      "trouble",
      "alpha",
      "autosession",
    },
  }
else
  -- VSCode-Neovim modules (minimal set)
  module_system = {
    -- Core functionality
    editor = {
      "vscode",
    },
    -- No UI modules needed
    ui = {
      "devicons",
      "lualine",
    },
    -- Minimal tool modules
    tools = {
      "comment",
      "treesitter",
      "hop",
      "alpha",      
    },
  }
end
local function collect_plugin_specs()
  local specs = module_loader.get_plugin_specs(module_system)
  -- Always include which-key plugin spec (Go plugin config will handle setup)
  table.insert(specs, {
    "folke/which-key.nvim",
    lazy = false,
    dependencies = { "echasnovski/mini.icons" },
  })
  -- Autopairs plugin
  table.insert(specs, {
    "windwp/nvim-autopairs",
    event = "InsertEnter",
    dependencies = { "nvim-treesitter/nvim-treesitter" },
  })

  if not vim.g.vscode then
    -- Git signs
    table.insert(specs, {
      "lewis6991/gitsigns.nvim",
      config = function()
        require("gitsigns").setup()
      end,
    })
    -- Oil: file explorer plugin
    table.insert(specs, {
      "stevearc/oil.nvim",
      lazy = false,
    })
  end

  return specs
end

require("lazy").setup(collect_plugin_specs())

-- Configuration via Go plugin; skip Lua module loader
-- module_loader.setup_modules(module_system)

if vim.g.vscode then
  local vscode_module = require("modules.editor.vscode")
  if vscode_module and vscode_module.setup_layer then
    vscode_module.setup_layer()
  end
end
