package modules

import "github.com/roshbhatia/nix-neovim-config/pkg/core"

// TreesitterModule configures nvim-treesitter and its textobjects.
type TreesitterModule struct{}

// NewTreesitterModule constructs a new TreesitterModule.
func NewTreesitterModule() *TreesitterModule {
   return &TreesitterModule{}
}

// Name returns the module name.
func (m *TreesitterModule) Name() string {
   return "treesitter"
}

// Setup runs the Treesitter setup.
func (m *TreesitterModule) Setup(ctx *core.Context) error {
   cmd := `lua << EOF
require("nvim-treesitter.configs").setup({
  ensure_installed = {
    "bash", "c", "cpp", "css", "go", "html", "javascript",
    "json", "lua", "markdown", "markdown_inline", "python",
    "regex", "rust", "toml", "tsx", "typescript", "vim",
    "yaml", "nix", "comment",
  },
  highlight = {
    enable = true,
    additional_vim_regex_highlighting = false,
  },
  indent = {
    enable = true,
  },
  autopairs = {
    enable = true,
  },
  textobjects = {
    select = {
      enable = true,
      lookahead = true,
      keymaps = {
        ["af"] = "@function.outer",
        ["if"] = "@function.inner",
        ["ac"] = "@class.outer",
        ["ic"] = "@class.inner",
        ["ab"] = "@block.outer",
        ["ib"] = "@block.inner",
      },
    },
    move = {
      enable = true,
      set_jumps = true,
      goto_next_start = {
        ["]f"] = "@function.outer",
        ["]c"] = "@class.outer",
      },
      goto_next_end = {
        ["]F"] = "@function.outer",
        ["]C"] = "@class.outer",
      },
      goto_previous_start = {
        ["[f"] = "@function.outer",
        ["[c"] = "@class.outer",
      },
      goto_previous_end = {
        ["[F"] = "@function.outer",
        ["[C"] = "@class.outer",
      },
    },
  },
})
EOF`
   return ctx.Command(cmd)
}