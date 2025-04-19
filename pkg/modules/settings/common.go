package settings

import "github.com/roshbhatia/nix-neovim-config/pkg/core"

// CommonModule applies general Neovim settings.
type CommonModule struct{}

// NewCommonModule constructs a new CommonModule.
func NewCommonModule() *CommonModule { return &CommonModule{} }

// Name returns the module name.
func (m *CommonModule) Name() string { return "settings.common" }

// Setup applies the common settings.
func (m *CommonModule) Setup(ctx *core.Context) error {
    cmds := []string{
        "set number",
        "set cursorline",
        "set signcolumn=yes",
        "set termguicolors",
        "set noshowmode",
        "set lazyredraw",
        "set foldmethod=expr",
        "set foldexpr=nvim_treesitter#foldexpr()",
        "set nofoldenable",
        "set foldlevel=99",
        "set pumheight=10",
        "set cmdheight=1",
        "set hidden",
        "set showtabline=2",
        "set shortmess+=c",
        "set completeopt=menuone,noselect",
        "set hlsearch",
        "set incsearch",
        "set ignorecase",
        "set smartcase",
        "set expandtab",
        "set shiftwidth=2",
        "set tabstop=2",
        "set smartindent",
        "set nowrap",
        "set linebreak",
        "set breakindent",
        "set splitbelow",
        "set splitright",
        "set updatetime=100",
        "set timeoutlen=300",
        "set scrolloff=8",
        "set sidescrolloff=8",
        "set mouse=a",
        "set clipboard=unnamedplus",
    }
    for _, cmd := range cmds {
        if err := ctx.Command(cmd); err != nil {
            return err
        }
    }
    return nil
}