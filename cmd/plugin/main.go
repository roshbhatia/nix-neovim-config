package main

import (
   "log"
   nvimplugin "github.com/neovim/go-client/nvim/plugin"
   "github.com/roshbhatia/nix-neovim-config/pkg/core"
   "github.com/roshbhatia/nix-neovim-config/pkg/modules"
)

func main() {
   nvimplugin.Main(func(p *nvimplugin.Plugin) error {
       ctx := core.NewContext(p)
       var mods []core.Module
       // Register modules here
       mods = append(mods, modules.NewWhichKeyModule())
       for _, m := range mods {
           if err := m.Setup(ctx); err != nil {
               return err
           }
           log.Printf("Loaded module: %s\n", m.Name())
       }
       return nil
   })
}