package modules

import (
   "fmt"
   "github.com/roshbhatia/nix-neovim-config/pkg/core"
)

// TelescopeModule configures Telescope and its key mappings.
type TelescopeModule struct{}

// NewTelescopeModule constructs a new TelescopeModule.
func NewTelescopeModule() *TelescopeModule {
   return &TelescopeModule{}
}

// Name returns the module name.
func (m *TelescopeModule) Name() string {
   return "telescope"
}

// Setup registers Telescope setup and which-key mappings.
func (m *TelescopeModule) Setup(ctx *core.Context) error {
   // Telescope configuration
   setupCmd := `lua << EOF
require("telescope").setup({
  defaults = {
    prompt_prefix = " ",
    selection_caret = " ",
    path_display = { "smart" },
    sorting_strategy = "ascending",
    layout_config = {
      horizontal = {
        prompt_position = "top",
        preview_width = 0.55,
        results_width = 0.8,
      },
      vertical = { mirror = false },
      width = 0.87,
      height = 0.80,
      preview_cutoff = 120,
    },
  },
  pickers = {
    find_files = {
      hidden = true,
      find_command = { "fd", "--type", "f", "--strip-cwd-prefix" },
    },
  },
  extensions = {
    fzf = {
      fuzzy = true,
      override_generic_sorter = true,
      override_file_sorter = true,
      case_mode = "smart_case",
    },
  },
})
EOF`
   if err := ctx.Command(setupCmd); err != nil {
       return err
   }
   // Which-key groups
   groups := map[string]string{
       "<leader>f": "Find",
       "<leader>g": "Git",
   }
   for keys, name := range groups {
       cmd := fmt.Sprintf(
           `lua require("which-key").register({ ["%s"] = { name = "%s" } }, { mode = "n" })`,
           keys, name,
       )
       if err := ctx.Command(cmd); err != nil {
           return err
       }
   }
   // Which-key mappings and keymaps
   mappings := []struct{ lhs, rhs, desc string }{
       {"<leader>ff", "<cmd>Telescope find_files<cr>", "Find Files"},
       {"<leader>fg", "<cmd>Telescope live_grep<cr>", "Find Text"},
       {"<leader>fb", "<cmd>Telescope buffers<cr>", "Find Buffers"},
       {"<leader>fm", "<cmd>Telescope marks<cr>",   "Find Marks"},
       {"<leader>gc", "<cmd>Telescope git_commits<cr>", "Git Commits"},
       {"<leader>gb", "<cmd>Telescope git_branches<cr>","Git Branches"},
   }
   opts := map[string]bool{"noremap": true, "silent": true}
   for _, m2 := range mappings {
       // Map the key
       if err := ctx.Map("n", m2.lhs, m2.rhs, opts); err != nil {
           return err
       }
       // Register in which-key with description
       cmd := fmt.Sprintf(
           `lua require("which-key").register({ ["%s"] = "%s" }, { mode = "n" })`,
           m2.lhs, m2.desc,
       )
       if err := ctx.Command(cmd); err != nil {
           return err
       }
   }
   return nil
}