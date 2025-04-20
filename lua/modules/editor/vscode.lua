-- sysinit.nvim.doc-url="https://raw.githubusercontent.com/vscode-neovim/vscode-neovim/main/README.md"
local M = {}

-- Map of Neovim commands (without <cmd> and <cr>) to VSCode action names
M.cmd_map = {
  w      = 'workbench.action.files.save',
  wa     = 'workbench.action.files.saveAll',
  q      = 'workbench.action.closeActiveEditor',
  qa     = 'workbench.action.quit',
  enew   = 'workbench.action.files.newUntitledFile',
  bdelete= 'workbench.action.closeActiveEditor',
  bn     = 'workbench.action.nextEditor',
  bp     = 'workbench.action.previousEditor',
  split  = 'workbench.action.splitEditorDown',
  vsplit = 'workbench.action.splitEditorRight',
}

-- VSCode plugin spec for Lazy.nvim
M.plugins = {
  {
    "vscode-neovim/vscode-neovim",
    lazy = false,
    cond = function()
      return vim.g.vscode == true
    end,
  },
}

--- Mirror <leader> keymaps from Neovim to VSCode
function M.setup_layer()
  if not vim.g.vscode then return end
  local api = require('vscode')
  local leader = vim.g.mapleader or ' '
  local summary = {}

  for _, m in ipairs(vim.api.nvim_get_keymap('n')) do
    if m.lhs:match('^' .. leader) then
      local lhs = m.lhs
      local rhs = m.rhs
      local opts = { noremap = m.noremap == 1, silent = m.silent == 1 }
      local cmd = rhs:match('^<cmd>(.+)<cr>$')
      if cmd then
        local action = M.cmd_map[cmd]
        if action then
          vim.keymap.set('n', lhs, function() api.action(action) end, opts)
          table.insert(summary, lhs .. ' -> ' .. action)
        else
          vim.keymap.set('n', lhs, '<cmd>' .. cmd .. '<cr>', opts)
          table.insert(summary, lhs .. ' -> ' .. cmd)
        end
      end
    end
  end

  -- Command to list mirrored keymaps
  vim.api.nvim_create_user_command('VsKeymaps', function()
    for _, line in ipairs(summary) do
      vim.api.nvim_echo({{line, 'Normal'}}, false, {})
    end
  end, {})
end

return M