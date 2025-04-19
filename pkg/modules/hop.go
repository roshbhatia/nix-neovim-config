package modules

import (
   "fmt"
   "github.com/roshbhatia/nix-neovim-config/pkg/core"
)

// HopModule configures phaazon/hop.nvim and its key mappings.
type HopModule struct{}

// NewHopModule constructs a new HopModule.
func NewHopModule() *HopModule {
   return &HopModule{}
}

// Name returns the module name.
func (m *HopModule) Name() string {
   return "hop"
}

// Setup sets up hop.nvim and registers key mappings.
func (m *HopModule) Setup(ctx *core.Context) error {
   // Setup hop with preferred keys
   setupCmd := `lua << EOF
require("hop").setup({ keys = "etovxqpdygfblzhckisuran" })
EOF`
   if err := ctx.Command(setupCmd); err != nil {
      return err
   }

   // Register which-key group for Hop under <leader>h
   if err := ctx.Command(
      `lua require("which-key").register({ ["<leader>h"] = { name = "Hop" } }, { mode = "n" })`,
   ); err != nil {
      return err
   }

   opts := map[string]bool{"noremap": true, "silent": true}
   // Motion mappings f, F, t, T for normal, visual, and operator-pending modes
   motions := []struct{ mode, lhs, rhs string }{
      {"n", "f", `<cmd>lua require("hop").hint_char1({ direction = require("hop.hint").HintDirection.AFTER_CURSOR, current_line_only = true })<CR>`},
      {"v", "f", `<cmd>lua require("hop").hint_char1({ direction = require("hop.hint").HintDirection.AFTER_CURSOR, current_line_only = true })<CR>`},
      {"o", "f", `<cmd>lua require("hop").hint_char1({ direction = require("hop.hint").HintDirection.AFTER_CURSOR, current_line_only = true })<CR>`},
      {"n", "F", `<cmd>lua require("hop").hint_char1({ direction = require("hop.hint").HintDirection.BEFORE_CURSOR, current_line_only = true })<CR>`},
      {"v", "F", `<cmd>lua require("hop").hint_char1({ direction = require("hop.hint").HintDirection.BEFORE_CURSOR, current_line_only = true })<CR>`},
      {"o", "F", `<cmd>lua require("hop").hint_char1({ direction = require("hop.hint").HintDirection.BEFORE_CURSOR, current_line_only = true })<CR>`},
      {"n", "t", `<cmd>lua require("hop").hint_char1({ direction = require("hop.hint").HintDirection.AFTER_CURSOR, current_line_only = true, hint_offset = -1 })<CR>`},
      {"v", "t", `<cmd>lua require("hop").hint_char1({ direction = require("hop.hint").HintDirection.AFTER_CURSOR, current_line_only = true, hint_offset = -1 })<CR>`},
      {"o", "t", `<cmd>lua require("hop").hint_char1({ direction = require("hop.hint").HintDirection.AFTER_CURSOR, current_line_only = true, hint_offset = -1 })<CR>`},
      {"n", "T", `<cmd>lua require("hop").hint_char1({ direction = require("hop.hint").HintDirection.BEFORE_CURSOR, current_line_only = true, hint_offset = 1 })<CR>`},
      {"v", "T", `<cmd>lua require("hop").hint_char1({ direction = require("hop.hint").HintDirection.BEFORE_CURSOR, current_line_only = true, hint_offset = 1 })<CR>`},
      {"o", "T", `<cmd>lua require("hop").hint_char1({ direction = require("hop.hint").HintDirection.BEFORE_CURSOR, current_line_only = true, hint_offset = 1 })<CR>`},
   }
   for _, m2 := range motions {
      if err := ctx.Map(m2.mode, m2.lhs, m2.rhs, opts); err != nil {
         return err
      }
   }

   // Leader mappings for hop words, lines, and anywhere
   leaderMaps := []struct{ lhs, rhs, desc string }{
      {"<leader>hw", `<cmd>lua require("hop").hint_words()<CR>`, "Hop to Word"},
      {"<leader>hl", `<cmd>lua require("hop").hint_lines()<CR>`, "Hop to Line"},
      {"<leader>ha", `<cmd>lua require("hop").hint_anywhere()<CR>`, "Hop Anywhere"},
   }
   for _, lm := range leaderMaps {
      // Register with which-key
      if err := ctx.Command(
         fmt.Sprintf(`lua require("which-key").register({ ["%s"] = "%s" }, { mode = "n" })`, lm.lhs, lm.desc),
      ); err != nil {
         return err
      }
      // Map key
      if err := ctx.Map("n", lm.lhs, lm.rhs, opts); err != nil {
         return err
      }
   }
   return nil
}