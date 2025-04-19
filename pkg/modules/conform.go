package modules

import "github.com/roshbhatia/nix-neovim-config/pkg/core"

// ConformModule configures stevearc/conform.nvim for code formatting.
type ConformModule struct{}

// NewConformModule constructs a new ConformModule.
func NewConformModule() *ConformModule {
   return &ConformModule{}
}

// Name returns the module name.
func (m *ConformModule) Name() string {
   return "conform"
}

// Setup configures conform.nvim: setup, commands, keymaps, and which-key.
func (m *ConformModule) Setup(ctx *core.Context) error {
   // Setup conform with formatters and hooks
   setupCmd := `lua << EOF
require("conform").setup({
  formatters_by_ft = {
    lua = { "stylua" },
    python = { "isort", "black" },
    javascript = { { "prettierd", "prettier" } },
    typescript = { { "prettierd", "prettier" } },
    javascriptreact = { { "prettierd", "prettier" } },
    typescriptreact = { { "prettierd", "prettier" } },
    go = { "goimports", "gofmt" },
    json = { { "prettierd", "prettier" } },
    jsonc = { { "prettierd", "prettier" } },
    yaml = { { "prettierd", "prettier" } },
    markdown = { { "prettierd", "prettier" } },
    ["markdown.mdx"] = { { "prettierd", "prettier" } },
    html = { { "prettierd", "prettier" } },
    css = { { "prettierd", "prettier" } },
    scss = { { "prettierd", "prettier" } },
    sh = { "shfmt" },
    bash = { "shfmt" },
    zsh = { "shfmt" },
    terraform = { "terraform_fmt" },
    rust = { "rustfmt" },
    toml = { "taplo" },
    nix = { "nixfmt" },
  },
  format_on_save = function(bufnr)
    if not vim.g.format_on_save then return end
    local ft = vim.bo[bufnr].filetype
    local fmt = conform.formatters_by_ft[ft]
    if not fmt or #fmt == 0 then return end
    local ok, st = pcall(vim.loop.fs_stat, vim.api.nvim_buf_get_name(bufnr))
    if ok && st && st.size > 100*1024 then
      vim.notify("File too large for formatting", vim.log.levels.WARN)
      return
    end
    return { timeout_ms = 500, lsp_fallback = true }
  end,
  formatters = {
    stylua = { prepend_args = { "--indent-type", "spaces", "--indent-width", "2" } },
    shfmt = { prepend_args = { "-i", "2", "-ci" } },
  },
})
EOF`
   if err := ctx.Command(setupCmd); err != nil {
      return err
   }
   // Create Format and ToggleFormatOnSave commands and keymaps
   cmdCmd := `lua << EOF
vim.api.nvim_create_user_command("Format", function(args)
  local range = nil
  if args.count ~= -1 then
    range = { start = { args.line1, 0 }, ["end"] = { args.line2, 999999 } }
  end
  require("conform").format({ async = true, lsp_fallback = true, range = range })
end, { range = true })
vim.api.nvim_create_user_command("ToggleFormatOnSave", function()
  vim.g.format_on_save = not vim.g.format_on_save
  vim.notify("Format on save " .. (vim.g.format_on_save and "enabled" or "disabled"), vim.log.levels.INFO)
end, {})
EOF`
   if err := ctx.Command(cmdCmd); err != nil {
      return err
   }
   // Map <leader>cf in normal and visual with which-key
   opts := map[string]bool{"noremap": true, "silent": true}
   // Normal
   if err := ctx.Command("lua require('which-key').add({ {'<leader>cf', 'Format', desc='Format Document'} })"); err != nil {
      return err
   }
   if err := ctx.Map("n", "<leader>cf", "<cmd>Format<CR>", opts); err != nil {
      return err
   }
   // Visual
   if err := ctx.Command("lua require('which-key').add({ {'<leader>cf', 'Format Selection'} }, { mode='v' })"); err != nil {
      return err
   }
   if err := ctx.Map("v", "<leader>cf", "<cmd>Format<CR>", opts); err != nil {
      return err
   }
   return nil
}