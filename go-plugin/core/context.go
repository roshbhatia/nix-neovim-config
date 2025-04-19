package core

import (
    "github.com/neovim/go-client/nvim"
    nvimplugin "github.com/neovim/go-client/nvim/plugin"
)

// Context wraps the Neovim plugin and API client.
type Context struct {
    Plugin *nvimplugin.Plugin
    Nvim   *nvim.Nvim
}

// NewContext creates a new Context from the plugin instance.
func NewContext(p *nvimplugin.Plugin) *Context {
    return &Context{Plugin: p, Nvim: p.Nvim}
}

// Module defines a pluggable component.
type Module interface {
    Name() string
    Setup(ctx *Context) error
}

// Map sets a keymap for the given mode.
func (c *Context) Map(mode, lhs, rhs string, opts map[string]interface{}) error {
    return c.Nvim.SetKeymap(mode, lhs, rhs, opts)
}

// Command runs a Vim command.
func (c *Context) Command(cmd string) error {
    return c.Plugin.Command(cmd)
}