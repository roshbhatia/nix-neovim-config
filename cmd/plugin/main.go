package main

import (
   "log"
   nvimplugin "github.com/neovim/go-client/nvim/plugin"
   "github.com/roshbhatia/nix-neovim-config/pkg/core"
   "github.com/roshbhatia/nix-neovim-config/pkg/modules"
   settings "github.com/roshbhatia/nix-neovim-config/pkg/modules/settings"
)

func main() {
   nvimplugin.Main(func(p *nvimplugin.Plugin) error {
       ctx := core.NewContext(p)
       var mods []core.Module
       // Register modules here in desired order
       mods = append(mods,
           settings.NewCommonModule(),
           modules.NewWhichKeyModule(),
           modules.NewCommentModule(),
           modules.NewHopModule(),
           modules.NewThemeToggleModule(),
           modules.NewNeoscrollModule(),
           modules.NewTelescopeModule(),
           modules.NewTreesitterModule(),
           modules.NewOilModule(),
           modules.NewCmpModule(),
           modules.NewAutopairsModule(),
       )
       for _, m := range mods {
           if err := m.Setup(ctx); err != nil {
               return err
           }
           log.Printf("Loaded module: %s\n", m.Name())
       }
       return nil
   })
}