package web

import (
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"proj/internal/server"
	"proj/internal/server/webserver"
	"proj/lifecycle"
	"proj/public/config"
	"proj/thirdparty/cache"
	"proj/thirdparty/store"
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

func init() {
	lifecycle.InjectMySQL(func() *gorm.DB {
		db, err := store.NewMySQL(lifecycle.RootContext, config.Config().MySQL)
		if err != nil {
			logrus.WithError(err).Panic("init mysql failure")
		}
		return db
	})
}

func init() {
	lifecycle.InjectRedis(func() *redis.Client {
		rds, err := cache.NewRedis(lifecycle.RootContext, config.Config().Redis)
		if err != nil {
			logrus.WithError(err).Panic("init redis failure")
		}
		return rds
	})
}
