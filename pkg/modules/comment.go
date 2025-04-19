package modules

import (
   "fmt"
   "github.com/roshbhatia/nix-neovim-config/pkg/core"
)

// CommentModule configures numToStr/Comment.nvim and integrates with which-key.
type CommentModule struct{}

// NewCommentModule constructs a new CommentModule.
func NewCommentModule() *CommentModule {
   return &CommentModule{}
}

// Name returns the module name.
func (m *CommentModule) Name() string {
   return "comment"
}

// Setup configures the Comment.nvim plugin and key mappings.
func (m *CommentModule) Setup(ctx *core.Context) error {
   // Plugin setup
   setupCmd := `lua << EOF
require("Comment").setup({
  mappings = {
    basic = false,
    extra = false,
  },
  pre_hook = function(ctx)
    local U = require("Comment.utils")
    local location = nil
    if ctx.ctype == U.ctype.block then
      location = require("ts_context_commentstring.utils").get_cursor_location()
    elseif ctx.cmotion == U.cmotion.v or ctx.cmotion == U.cmotion.V then
      location = require("ts_context_commentstring.utils").get_visual_start_location()
    end
    local ok, ts_context_commentstring = pcall(require, "ts_context_commentstring.internal")
    if ok then
      return ts_context_commentstring.calculate_commentstring({ key = ctx.ctype == U.ctype.line and "__default" or "__multiline", location = location })
    end
  end,
})
EOF`
   if err := ctx.Command(setupCmd); err != nil {
      return err
   }
   // Key mappings
   opts := map[string]bool{"noremap": true, "silent": true}
   mappings := []struct{ mode, lhs, rhs, desc string }{
      {"n", "<leader>cc", "<cmd>lua require('Comment.api').toggle.linewise.current()<CR>", "Toggle Line Comment"},
      {"v", "<leader>cc", "<ESC><cmd>lua require('Comment.api').toggle.linewise(vim.fn.visualmode())<CR>", "Toggle Line Comment"},
      {"n", "<leader>cb", "<cmd>lua require('Comment.api').toggle.blockwise.current()<CR>", "Toggle Block Comment"},
      {"v", "<leader>cb", "<ESC><cmd>lua require('Comment.api').toggle.blockwise(vim.fn.visualmode())<CR>", "Toggle Block Comment"},
   }
   for _, mp := range mappings {
      // Register in which-key
      if err := ctx.Command(
         fmt.Sprintf(`lua require("which-key").register({ ["%s"] = "%s" }, { mode = "%s" })`, mp.lhs, mp.desc, mp.mode),
      ); err != nil {
         return err
      }
      // Set keymap
      if err := ctx.Map(mp.mode, mp.lhs, mp.rhs, opts); err != nil {
         return err
      }
   }
   return nil
}