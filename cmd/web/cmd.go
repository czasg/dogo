package web

import (
	"github.com/spf13/cobra"
	"proj/internal/server"
	"proj/internal/server/webserver"
	"proj/lifecycle"
	"proj/public/config"
)

var (
	ServerCmd = &cobra.Command{
		Use:   "webserver",
		Short: "web server",
		Long:  "start a web server",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := lifecycle.RootContext
			cfg := config.Config()
			app := webserver.NewApp()
			return server.Run(ctx, app, cfg.Http)
		},
	}
)
