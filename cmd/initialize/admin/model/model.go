package model

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"proj/internal/domain/model"
	"proj/lifecycle"
	"proj/thirdparty/store"
)

var (
	ModelCmd = &cobra.Command{
		Use:   "model",
		Short: "init admin model",
		Long:  "init admin model",
		RunE: func(cmd *cobra.Command, args []string) error {
			return lifecycle.MySQL.AutoMigrate(
				&model.User{},
				&model.UserDetail{},
				&model.Role{},
				&model.UserRole{},
				&model.AccessControl{},
				&model.Menu{},
				&model.RoleMenu{},
			)
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
	)
}
