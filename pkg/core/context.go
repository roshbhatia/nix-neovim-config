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

// SetOption sets a global Vim option.
func (c *Context) SetOption(name string, value interface{}) error {
   if c.Nvim == nil {
       return nil
   }
   return c.Nvim.SetOption(name, value)
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
// Map sets a keymap for the given mode.
func (c *Context) Map(mode, lhs, rhs string, opts map[string]bool) error {
   if c.Nvim == nil {
       return nil
   }
   return c.Nvim.SetKeyMap(mode, lhs, rhs, opts)
}

// Command runs a Vim command.
func (c *Context) Command(cmd string) error {
   if c.Nvim == nil {
       return nil
   }
   return c.Nvim.Command(cmd)
}