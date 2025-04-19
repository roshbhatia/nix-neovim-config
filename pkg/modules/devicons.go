package modules

import "github.com/roshbhatia/nix-neovim-config/pkg/core"

// DeviconsModule configures nvim-web-devicons plugin.
type DeviconsModule struct{}

// NewDeviconsModule constructs a new DeviconsModule.
func NewDeviconsModule() *DeviconsModule { return &DeviconsModule{} }

// Name returns the module name.
func (m *DeviconsModule) Name() string { return "devicons" }

// Setup runs the web-devicons setup and creates a user command.
func (m *DeviconsModule) Setup(ctx *core.Context) error {
   setup := `lua << EOF
require("nvim-web-devicons").setup({
  override = {
    default_icon = { icon = "", color = "#6d8086", name = "Default" },
    nix = { icon = "", color = "#7ebae4", name = "Nix" },
  },
  default = true,
  strict = true,
  color_icons = true,
})
EOF`
   if err := ctx.Command(setup); err != nil {
      return err
   }
   cmd := `lua << EOF
vim.api.nvim_create_user_command("ShowDevIcons", function()
  local ok, icons = pcall(require, "nvim-web-devicons")
  if not ok then
    vim.notify("nvim-web-devicons not available", vim.log.levels.ERROR)
    return
  end
  local all = icons.get_icons()
  if all then
    vim.api.nvim_echo({{vim.inspect(all), "Normal"}}, true, {})
  else
    vim.notify("No icons available", vim.log.levels.WARN)
  end
end, {})
EOF`
   return ctx.Command(cmd)
}