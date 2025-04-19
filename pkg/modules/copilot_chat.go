package modules

import "github.com/roshbhatia/nix-neovim-config/pkg/core"

// CopilotChatModule configures CopilotChat.nvim plugin and mappings.
type CopilotChatModule struct{}

// NewCopilotChatModule constructs a new CopilotChatModule.
func NewCopilotChatModule() *CopilotChatModule {
   return &CopilotChatModule{}
}

// Name returns the module name.
func (m *CopilotChatModule) Name() string {
   return "copilot-chat"
}

// Setup applies CopilotChat.nvim setup and which-key mappings.
func (m *CopilotChatModule) Setup(ctx *core.Context) error {
   cmd := `lua << EOF
require("CopilotChat").setup({
  model = "gpt-4o",
  agent = "none",
  window = { layout = "float", width = 0.8, height = 0.7, border = "rounded", title = "Copilot Chat" },
  highlight_selection = true,
  auto_follow_cursor = true,
  auto_insert_mode = true,
  show_help = true,
  clear_chat_on_new_prompt = false,
  prompts = {
    ImplementFeature = { prompt = "Implement a new feature based on the context and description below.", context = "buffer" },
    DebugThis = { prompt = "Debug the selected code and provide a detailed explanation of the problem and solution.", context = "buffer" },
    ExplainInSimple = { prompt = "Explain the selected code in simple terms as if teaching to a beginner.", context = "buffer" },
  },
})
local wk = require("which-key")
wk.register({ ["<leader>a"] = { name = "AI" } }, { mode = "n" })
wk.register({ ["<leader>ac"] = { name = "Chat" } }, { mode = "n" })
wk.register({ ["<leader>acc"] = "Open Chat" }, { mode = "n" })
vim.keymap.set("n", "<leader>acc", "<cmd>CopilotChat<CR>", { noremap = true, silent = true })
wk.register({ ["<leader>act"] = "Toggle Chat" }, { mode = "n" })
vim.keymap.set("n", "<leader>act", "<cmd>CopilotChatToggle<CR>", { noremap = true, silent = true })
wk.register({ ["<leader>ace"] = "Explain Code" }, { mode = { "n", "v" } })
vim.keymap.set({ "n", "v" }, "<leader>ace", "<cmd>CopilotChatExplain<CR>", { noremap = true, silent = true })
wk.register({ ["<leader>acr"] = "Review Code" }, { mode = { "n", "v" } })
vim.keymap.set({ "n", "v" }, "<leader>acr", "<cmd>CopilotChatReview<CR>", { noremap = true, silent = true })
wk.register({ ["<leader>acf"] = "Fix Code" }, { mode = { "n", "v" } })
vim.keymap.set({ "n", "v" }, "<leader>acf", "<cmd>CopilotChatFix<CR>", { noremap = true, silent = true })
wk.register({ ["<leader>aco"] = "Optimize Code" }, { mode = { "n", "v" } })
vim.keymap.set({ "n", "v" }, "<leader>aco", "<cmd>CopilotChatOptimize<CR>", { noremap = true, silent = true })
wk.register({ ["<leader>acd"] = "Generate Docs" }, { mode = { "n", "v" } })
vim.keymap.set({ "n", "v" }, "<leader>acd", "<cmd>CopilotChatDocs<CR>", { noremap = true, silent = true })
wk.register({ ["<leader>act"] = "Generate Tests" }, { mode = { "n", "v" } })
vim.keymap.set({ "n", "v" }, "<leader>act", "<cmd>CopilotChatTests<CR>", { noremap = true, silent = true })
wk.register({ ["<leader>aci"] = "Implement Feature" }, { mode = { "n", "v" } })
vim.keymap.set({ "n", "v" }, "<leader>aci", "<cmd>CopilotChatImplementFeature<CR>", { noremap = true, silent = true })
wk.register({ ["<leader>acb"] = "Debug Code" }, { mode = { "n", "v" } })
vim.keymap.set({ "n", "v" }, "<leader>acb", "<cmd>CopilotChatDebugThis<CR>", { noremap = true, silent = true })
wk.register({ ["<leader>acs"] = "Explain Simple" }, { mode = { "n", "v" } })
vim.keymap.set({ "n", "v" }, "<leader>acs", "<cmd>CopilotChatExplainInSimple<CR>", { noremap = true, silent = true })
EOF`
   return ctx.Command(cmd)
}