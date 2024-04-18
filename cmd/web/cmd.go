package web

import (
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"proj/internal/server"
	"proj/internal/server/webserver"
	"proj/lifecycle"
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
			cfg := lifecycle.Config
			app := webserver.NewApp()
			return server.Run(ctx, app, cfg.Http)
		},
	}
)

func init() {
	lifecycle.Inject(
		lifecycle.MySQLCaller(func() *gorm.DB {
			db, err := store.NewMySQL(lifecycle.RootContext, lifecycle.Config.MySQL)
			if err != nil {
				logrus.WithError(err).Panic("init mysql failure")
			}
			return db
		}),
		lifecycle.RedisCaller(func() *redis.Client {
			rds, err := cache.NewRedis(lifecycle.RootContext, lifecycle.Config.Redis)
			if err != nil {
				logrus.WithError(err).Panic("init redis failure")
			}
			return rds
		}),
	)
}
