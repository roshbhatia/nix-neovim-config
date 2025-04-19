package main

import (
   "log"
   "os"
   nvimplugin "github.com/neovim/go-client/nvim/plugin"
   "github.com/roshbhatia/nix-neovim-config/go-plugin/pkg/core"
   "github.com/roshbhatia/nix-neovim-config/go-plugin/pkg/modules"
)

func main() {
   opts := &nvimplugin.Options{
       Addr: os.Getenv("NVIM_LISTEN_ADDRESS"),
   }
   nvimplugin.Main(opts, func(p *nvimplugin.Plugin) error {
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