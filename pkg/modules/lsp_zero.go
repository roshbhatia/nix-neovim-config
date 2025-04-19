package modules

import "github.com/roshbhatia/nix-neovim-config/pkg/core"

// LspZeroModule sets up basic LSP functionality via lsp-zero.
type LspZeroModule struct{}

// NewLspZeroModule constructs a new LspZeroModule.
func NewLspZeroModule() *LspZeroModule { return &LspZeroModule{} }

// Name returns the module name.
func (m *LspZeroModule) Name() string { return "lsp-zero" }

// Setup configures lsp-zero, mason, and LSP keybindings.
func (m *LspZeroModule) Setup(ctx *core.Context) error {
	// Configure lsp-zero, mason, and LSP keybindings.
	setupCmd := `lua << EOF
local lsp_zero = require('lsp-zero')
local lspconfig = require('lspconfig')
local mason_lspconfig = require('mason-lspconfig')
local neodev = require('neodev')
local wk = require('which-key')

neodev.setup({})

lsp_zero.preset({
  name = 'recommended',
  set_lsp_keymaps = true,
  manage_nvim_cmp = false,
  suggest_lsp_servers = true,
})

require('mason').setup({
  ui = {
    border = 'rounded',
    icons = {
      package_installed = '✓',
      package_pending = '➜',
      package_uninstalled = '✗',
    },
  },
})

do
  local mlsp = mason_lspconfig
  local servers = {
    'lua_ls', 'gopls', 'golangci_lint_ls', 'russ', 'ts_ls',
    'bashls', 'spectral', 'marksman', 'grammarly', 'nil_ls',
    'tflint', 'terraformls', 'helm-ls',
  }
  local to_install = {}
  if mlsp.get_available_servers then
    local available = mlsp.get_available_servers()
    for _, name in ipairs(servers) do
      if vim.tbl_contains(available, name) then
        table.insert(to_install, name)
      else
        vim.notify("LSP server '" .. name .. "' not available in mason-lspconfig, skipping", vim.log.levels.WARN)
      end
    end
  else
    to_install = servers
  end
  mlsp.setup({ ensure_installed = to_install })
end

mason_lspconfig.setup_handlers({
  lsp_zero.default_setup,

  lua_ls = function()
    lspconfig.lua_ls.setup(lsp_zero.nvim_lua_ls())
  end,

  gopls = function()
    lspconfig.gopls.setup({
      settings = {
        gopls = {
          analyses = { unusedparams = true, shadow = true },
          staticcheck = true,
          gofumpt = true,
          usePlaceholders = true,
          completeUnimported = true,
        },
      },
      capabilities = lsp_zero.get_capabilities(),
    })
  end,
})

lsp_zero.on_attach(function(client, bufnr)
  lsp_zero.default_keymaps({ buffer = bufnr })

  local opts = { buffer = bufnr }
  vim.keymap.set('n', 'gd', '<cmd>lua vim.lsp.buf.definition()<cr>', opts)
  vim.keymap.set('n', 'gi', '<cmd>lua vim.lsp.buf.implementation()<cr>', opts)
  vim.keymap.set('n', 'gr', '<cmd>lua vim.lsp.buf.references()<cr>', opts)
  vim.keymap.set('n', 'K', '<cmd>lua vim.lsp.buf.hover()<cr>', opts)
  vim.keymap.set('n', '<leader>rn', '<cmd>lua vim.lsp.buf.rename()<cr>', opts)
  vim.keymap.set('n', '<leader>ca', '<cmd>lua vim.lsp.buf.code_action()<cr>', opts)
  vim.keymap.set('n', '<leader>do', '<cmd>lua vim.diagnostic.open_float()<cr>', opts)
  vim.keymap.set('n', '<leader>dp', '<cmd>lua vim.diagnostic.goto_prev()<cr>', opts)
  vim.keymap.set('n', '<leader>dn', '<cmd>lua vim.diagnostic.goto_next()<cr>', opts)

  wk.add({
    { 'gd', '<cmd>lua vim.lsp.buf.definition()<cr>', desc = 'Go to Definition', buffer = bufnr },
    { 'gr', '<cmd>lua vim.lsp.buf.references()<cr>', desc = 'Go to References', buffer = bufnr },
    { 'gi', '<cmd>lua vim.lsp.buf.implementation()<cr>', desc = 'Go to Implementation', buffer = bufnr },
    { 'K', '<cmd>lua vim.lsp.buf.hover()<cr>', desc = 'Show Hover', buffer = bufnr },
    { '<leader>d', group = 'Diagnostics', icon = { icon = '󰨮', hl = 'WhichKeyIconRed' }, buffer = bufnr },
    { '<leader>do', '<cmd>lua vim.diagnostic.open_float()<cr>', desc = 'Show Diagnostics', buffer = bufnr },
    { '<leader>dp', '<cmd>lua vim.diagnostic.goto_prev()<cr>', desc = 'Previous Diagnostic', buffer = bufnr },
    { '<leader>dn', '<cmd>lua vim.diagnostic.goto_next()<cr>', desc = 'Next Diagnostic', buffer = bufnr },
    { '<leader>r', group = 'Refactor', icon = { icon = '󰕚', hl = 'WhichKeyIconBlue' }, buffer = bufnr },
    { '<leader>rn', '<cmd>lua vim.lsp.buf.rename()<cr>', desc = 'Rename Symbol', buffer = bufnr },
    { '<leader>ca', '<cmd>lua vim.lsp.buf.code_action()<cr>', desc = 'Code Action', buffer = bufnr },
  })
end)

vim.diagnostic.config({
  virtual_text = { prefix = '●', source = 'if_many' },
  float = { border = 'rounded', source = 'always' },
  severity_sort = true,
  update_in_insert = false,
})

if vim.lsp.inlay_hint then
  vim.keymap.set('n', '<leader>ih', function()
    vim.lsp.inlay_hint.enable(0, not vim.lsp.inlay_hint.is_enabled())
  end, { desc = 'Toggle Inlay Hints' })
end
EOF`
   return ctx.Command(setupCmd)
}