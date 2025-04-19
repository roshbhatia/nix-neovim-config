package modules

import "github.com/roshbhatia/nix-neovim-config/pkg/core"

// CopilotCmpModule configures zbirenbaum/copilot-cmp and copilot.lua integration.
type CopilotCmpModule struct{}

// NewCopilotCmpModule constructs a new CopilotCmpModule.
func NewCopilotCmpModule() *CopilotCmpModule {
   return &CopilotCmpModule{}
}

// Name returns the module name.
func (m *CopilotCmpModule) Name() string {
   return "copilot-cmp"
}

// Setup applies copilot-cmp and copilot.lua settings.
func (m *CopilotCmpModule) Setup(ctx *core.Context) error {
   cmd := `lua << EOF
require("copilot_cmp").setup({
  event = { "InsertEnter", "LspAttach" },
  fix_pairs = true,
})
require("copilot").setup({
  suggestion = { enabled = false },
  panel = { enabled = false },
  filetypes = {
    ["*"] = true,
    ["help"] = false,
    ["gitcommit"] = false,
    ["gitrebase"] = false,
    ["hgcommit"] = false,
    ["svn"] = false,
    ["cvs"] = false,
  },
})
vim.api.nvim_set_hl(0, "CmpItemKindCopilot", { fg = "#6CC644" })
EOF`
   return ctx.Command(cmd)
}