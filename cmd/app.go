package cmd

import (
	"github.com/spf13/cobra"
	"proj/cmd/initialize"
	"proj/cmd/web"
)

var App = &cobra.Command{
	Use:   "app",
	Short: "dogo app command",
	Long:  "this is dogo app command, please make sure `yourself cmd` is registered before app.Execute",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		panic("not found app command")
	},
}

func init() {
	App.AddCommand(
		web.ServerCmd,
		initialize.InitCmd,
	)
}
