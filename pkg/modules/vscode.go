package modules

import (
   "fmt"
   "regexp"
   "strings"

   "github.com/roshbhatia/nix-neovim-config/pkg/core"
)

// VscodeModule mirrors Neovim keybindings in VSCode via vscode-neovim
type VscodeModule struct{}

// NewVscodeModule constructs a new VscodeModule.
func NewVscodeModule() *VscodeModule {
   return &VscodeModule{}
}

// Name returns the module name.
func (m *VscodeModule) Name() string {
   return "vscode"
}

// Setup registers VSCode keymaps mirroring Neovim leader mappings.
// It only runs when 'vim.g.vscode' is true.
func (m *VscodeModule) Setup(ctx *core.Context) error {
   // Skip setup when running manifest generation (no Nvim client)
   if ctx.Nvim == nil {
       return nil
   }
   // Detect VSCode environment via vim.g.vscode
   var enabled bool
   if err := ctx.Nvim.Eval("vim.g.vscode == true", &enabled); err != nil || !enabled {
       return nil
   }

   // Mapping of Neovim commands to VSCode actions
   cmdMap := map[string]string{
       "w":      "workbench.action.files.save",
       "wa":     "workbench.action.files.saveAll",
       "q":      "workbench.action.closeActiveEditor",
       "qa":     "workbench.action.quit",
       "enew":   "workbench.action.files.newUntitledFile",
       "bdelete": "workbench.action.closeActiveEditor",
       "bn":     "workbench.action.nextEditor",
       "bp":     "workbench.action.previousEditor",
       "split":  "workbench.action.splitEditorDown",
       "vsplit": "workbench.action.splitEditorRight",
   }

   // Determine leader key (default to space)
   leader := " "
   var ldr string
   if err := ctx.Nvim.Eval("vim.g.mapleader", &ldr); err == nil && ldr != "" {
       leader = ldr
   }

   // Retrieve all normal mode keymaps
   var keymaps []map[string]interface{}
   if err := ctx.Nvim.Eval("vim.api.nvim_get_keymap('n')", &keymaps); err != nil {
       return err
   }

   // Regex to extract command from <cmd>...<cr>
   re := regexp.MustCompile(`^<cmd>(.+)<cr>$`)

   // Mirror each leader mapping
   for _, km := range keymaps {
       lhsVal, ok := km["lhs"].(string)
       if !ok || !strings.HasPrefix(lhsVal, leader) {
           continue
       }
       rhsVal, ok := km["rhs"].(string)
       if !ok {
           continue
       }
       // Extract options
       opts := map[string]bool{}
       if v, ok := km["noremap"].(int64); ok {
           opts["noremap"] = v != 0
       }
       if v, ok := km["silent"].(int64); ok {
           opts["silent"] = v != 0
       }
       // Match <cmd>...<cr>
       if m := re.FindStringSubmatch(rhsVal); len(m) == 2 {
           cmdText := m[1]
           if action, found := cmdMap[cmdText]; found {
               // Map to VSCode action via Lua call
               rhs := fmt.Sprintf("<cmd>lua require('vscode').action(%q)<cr>", action)
               if err := ctx.Map("n", lhsVal, rhs, opts); err != nil {
                   return err
               }
           } else {
               // Fallback: original command mapping
               if err := ctx.Map("n", lhsVal, rhsVal, opts); err != nil {
                   return err
               }
           }
       }
   }
   return nil
}