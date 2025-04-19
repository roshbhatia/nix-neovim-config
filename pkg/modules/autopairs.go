package modules

import "github.com/roshbhatia/nix-neovim-config/pkg/core"

// AutopairsModule configures windwp/nvim-autopairs plugin and integrates with nvim-cmp.
type AutopairsModule struct{}

// NewAutopairsModule constructs a new AutopairsModule.
func NewAutopairsModule() *AutopairsModule {
    return &AutopairsModule{}
}

// Name returns the module name.
func (m *AutopairsModule) Name() string {
    return "autopairs"
}

// Setup sets up nvim-autopairs and binds confirm_done event for nvim-cmp.
func (m *AutopairsModule) Setup(ctx *core.Context) error {
    // Configure autopairs
    setupCmd := `lua << EOF
require("nvim-autopairs").setup({
  check_ts = true,
  ts_config = {
    lua = {"string"},
    javascript = {"template_string"},
    typescript = {"template_string"},
    java = false,
  },
  disable_filetype = {"TelescopePrompt", "vim"},
  enable_check_bracket_line = true,
  fast_wrap = {
    map = "<M-e>",
    chars = {"{","[","(","\"","'"},
    pattern = [=[[%'%"%)%>%]%)%}%,]]=],
    offset = 0,
    end_key = "$",
    keys = "qwertyuiopzxcvbnmasdfghjkl",
    check_comma = true,
    highlight = "Search",
    highlight_grey = "Comment",
  },
})
EOF`
    if err := ctx.Command(setupCmd); err != nil {
        return err
    }
    // Integrate with nvim-cmp confirm_done
    cmpCmd := `lua << EOF
local cmp = require("cmp")
local cmp_autopairs = require("nvim-autopairs.completion.cmp")
cmp.event:on("confirm_done", cmp_autopairs.on_confirm_done())
EOF`
    if err := ctx.Command(cmpCmd); err != nil {
        return err
    }
    return nil
}