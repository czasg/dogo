package server

import (
	"github.com/spf13/cobra"
	"proj/cmd"
)

var (
	serverCmd = &cobra.Command{
		Use:   "webserver",
		Short: "web server",
		Long:  `start a web server`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
)

func init() {
	cmd.App.AddCommand(serverCmd)
}
