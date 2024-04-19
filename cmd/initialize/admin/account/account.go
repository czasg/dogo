package account

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"proj/internal/domain/model"
	"proj/lifecycle"
	"proj/public/utils"
	"proj/thirdparty/store"
)

var (
	AccountCmd = &cobra.Command{
		Use:   "account",
		Short: "init admin account",
		Long:  "init admin account",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			var user model.User
			err := lifecycle.MySQL.WithContext(ctx).Where("name = 'admin'").First(&user).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
			if user.ID > 0 {
				return nil
			}
			var adminSecret string
			fmt.Println("请初始化admin密码：")
			_, err = fmt.Scanln(&adminSecret)
			if err != nil {
				return err
			}
			return lifecycle.MySQL.Transaction(func(tx *gorm.DB) error {
				us := model.UserService{
					DB: tx,
				}
				user := model.User{
					Name:   "admin",
					Alias:  "管理员",
					Enable: true,
					Admin:  true,
				}
				userDetail := model.UserDetail{
					Email:    "",
					Password: utils.NewHash(nil).Sha256([]byte(adminSecret)),
				}
				return us.Create(ctx, &user, &userDetail)
			})
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
