package modules

import (
  "github.com/roshbhatia/nix-neovim-config/pkg/core"
)

// NvimLintModule configures mfussenegger/nvim-lint for on-save linting.
type NvimLintModule struct{}

// NewNvimLintModule constructs a new NvimLintModule.
func NewNvimLintModule() *NvimLintModule {
  return &NvimLintModule{}
}

// Name returns the module name.
func (m *NvimLintModule) Name() string {
  return "nvim-lint"
}

// Setup configures linting commands, autocmds, and toggles.
func (m *NvimLintModule) Setup(ctx *core.Context) error {
  cmd := `lua << EOF
local lint = require("lint")
local linters_by_ft = {
  python = {"pylint", "ruff"},
  javascript = {"eslint"},
  typescript = {"eslint"},
  javascriptreact = {"eslint"},
  typescriptreact = {"eslint"},
  go = {"golangci-lint"},
  sh = {"shellcheck"},
  bash = {"shellcheck"},
  zsh = {"shellcheck"},
  json = {"jsonlint"},
  markdown = {"markdownlint"},
  lua = {"luacheck"},
  terraform = {"tflint"},
}
if vim.fn.executable("yamllint") == 1 then
  linters_by_ft.yaml = {"yamllint"}
else
  vim.notify("yamllint not found: YAML linting disabled", vim.log.levels.WARN)
end
lint.linters_by_ft = linters_by_ft
lint.linters.pylint.args = {"--output-format=text","--score=no","--msg-template='{line}:{column}:{category}:{msg} ({symbol})'"}
lint.linters.shellcheck.args = {"--format=gcc","--external-sources","--shell=bash"}
vim.api.nvim_create_autocmd({"BufWritePost","BufEnter"}, { callback = function() require("lint").try_lint() end })
vim.api.nvim_create_user_command("Lint", function() require("lint").try_lint() end, {})
local auto_lint = true
vim.api.nvim_create_user_command("ToggleAutoLint", function()
  auto_lint = not auto_lint
  if auto_lint then
    vim.api.nvim_clear_autocmds({ group = "nvim_lint_augroup" })
    vim.api.nvim_create_autocmd({"BufWritePost","BufEnter"}, { group = vim.api.nvim_create_augroup("nvim_lint_augroup", { clear = true }), callback = function() require("lint").try_lint() end })
    vim.notify("Automatic linting enabled", vim.log.levels.INFO)
  else
    vim.api.nvim_clear_autocmds({ group = "nvim_lint_augroup" })
    vim.notify("Automatic linting disabled", vim.log.levels.INFO)
  end
end, {})
EOF`
  return ctx.Command(cmd)
}