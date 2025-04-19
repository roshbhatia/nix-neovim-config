package modules

import "github.com/roshbhatia/nix-neovim-config/pkg/core"

// CopilotModule configures github/copilot.vim and its mappings.
type CopilotModule struct{}

// NewCopilotModule constructs a new CopilotModule.
func NewCopilotModule() *CopilotModule {
   return &CopilotModule{}
}

// Name returns the module name.
func (m *CopilotModule) Name() string {
   return "copilot"
}

// Setup applies Copilot settings, keymaps, and which-key integration.
func (m *CopilotModule) Setup(ctx *core.Context) error {
   cmd := `lua << EOF
vim.g.copilot_no_tab_map = true
vim.g.copilot_assume_mapped = true
vim.g.copilot_tab_fallback = ""
vim.g.copilot_filetypes = {
  ["*"] = true,
  ["markdown"] = true,
  ["yaml"] = true,
  ["help"] = false,
  ["gitcommit"] = false,
  ["gitrebase"] = false,
  ["hgcommit"] = false,
  ["svn"] = false,
  ["cvs"] = false,
}
vim.keymap.set("i", "<C-j>", 'copilot#Accept("<CR>")', { expr = true, silent = true, replace_keycodes = false })
vim.keymap.set("i", "<C-k>", "<Plug>(copilot-previous)")
vim.keymap.set("i", "<C-l>", "<Plug>(copilot-next)")
vim.keymap.set("i", "<C-\\>", "<Plug>(copilot-dismiss)")
local function toggle_copilot()
  if vim.g.copilot_enabled == 0 then
    vim.cmd("Copilot enable")
    vim.notify("Copilot enabled", vim.log.levels.INFO)
  else
    vim.cmd("Copilot disable")
    vim.notify("Copilot disabled", vim.log.levels.INFO)
  end
end
local wk = require("which-key")
wk.register({ ["<leader>a"] = { name = "AI" } }, { mode = "n" })
wk.register({ ["<leader>at"] = "Toggle Copilot" }, { mode = "n" })
vim.keymap.set("n", "<leader>at", function() toggle_copilot() end, { noremap = true, silent = true })
wk.register({ ["<leader>as"] = "Copilot Status" }, { mode = "n" })
vim.keymap.set("n", "<leader>as", "<cmd>Copilot status<CR>", { noremap = true, silent = true })
wk.register({ ["<leader>ap"] = "Copilot Panel" }, { mode = "n" })
vim.keymap.set("n", "<leader>ap", "<cmd>Copilot panel<CR>", { noremap = true, silent = true })
vim.api.nvim_create_autocmd("ColorScheme", {
  pattern = "*",
  callback = function()
    vim.api.nvim_set_hl(0, "CopilotSuggestion", { fg = "#888888", ctermfg = 8, force = true })
  end,
})
EOF`
   return ctx.Command(cmd)
}