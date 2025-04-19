package modules

import "github.com/roshbhatia/nix-neovim-config/pkg/core"

// AlphaModule configures goolord/alpha-nvim startup dashboard.
type AlphaModule struct{}

// NewAlphaModule constructs a new AlphaModule.
func NewAlphaModule() *AlphaModule {
  return &AlphaModule{}
}

// Name returns the module name.
func (m *AlphaModule) Name() string {
  return "alpha"
}

// Setup configures the alpha dashboard, key mappings, and highlights.
func (m *AlphaModule) Setup(ctx *core.Context) error {
  // Dashboard setup
  dashboardCmd := `lua << EOF
local alpha = require("alpha")
local dashboard = require("alpha.themes.dashboard")
local function read_ascii_art()
  local cfg = vim.fn.fnamemodify(vim.fn.expand("$MYVIMRC"), ":p:h")
  for _, path in ipairs({cfg.."/alpha.ascii", cfg.."/../alpha.ascii", cfg.."/../../alpha.ascii"}) do
    if vim.loop.fs_stat(path) then return vim.fn.readfile(path) end
  end
  return require("alpha").default_header()
end
dashboard.section.header.val = read_ascii_art()
dashboard.section.header.opts.hl = "ProfileGreen"
dashboard.section.buttons.val = {
  dashboard.button("f", "  Find Files", ":Telescope find_files<CR>"),
  dashboard.button("r", "  Recent Files", ":Telescope oldfiles<CR>"),
  dashboard.button("g", "  Live Grep", ":Telescope live_grep<CR>"),
  dashboard.button("c", "  Configuration", ":e $MYVIMRC<CR>"),
  dashboard.button("t", "  Change Theme", ":Themify<CR>"),
  dashboard.button("l", "󰒲  Lazy", ":Lazy<CR>"),
  dashboard.button("q", "  Quit", ":qa<CR>"),
}
dashboard.section.footer.val = {"Welcome back, Rosh"}
dashboard.config.layout = {
  { type = "padding", val = 2 },
  dashboard.section.header,
  { type = "padding", val = 2 },
  dashboard.section.buttons,
  { type = "padding", val = 1 },
  dashboard.section.footer,
}
vim.cmd([[autocmd FileType alpha setlocal nofoldenable]])
alpha.setup(dashboard.config)
EOF`
  if err := ctx.Command(dashboardCmd); err != nil {
    return err
  }
  // Map <leader>P to open dashboard
  if err := ctx.Command(`lua require("which-key").register({ ["<leader>P"] = "Open Homepage" }, { mode = "n" })`); err != nil {
    return err
  }
  if err := ctx.Map("n", "<leader>P", "<cmd>Alpha<CR>", map[string]bool{"noremap": true, "silent": true}); err != nil {
    return err
  }
  // Highlight groups on ColorScheme
  hlCmd := `lua << EOF
vim.api.nvim_create_autocmd("ColorScheme", {
  pattern = "*",
  callback = function()
    vim.opt_local.scrolloff = 0
    vim.opt_local.sidescrolloff = 0
    vim.api.nvim_set_hl(0, "ProfileBlue", { fg = "#61afef", bold = true })
    vim.api.nvim_set_hl(0, "ProfileGreen", { fg = "#98c379", bold = true })
    vim.api.nvim_set_hl(0, "ProfileYellow", { fg = "#e5c07b", bold = true })
    vim.api.nvim_set_hl(0, "ProfileRed", { fg = "#e06c75", bold = true })
  end,
})
EOF`
  return ctx.Command(hlCmd)
}