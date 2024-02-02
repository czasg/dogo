package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"proj/internal/domain/model"
	"proj/lifecycle"
	"proj/public/httplib"
	"proj/public/utils"
	"strconv"
)

func DefaultUserApp() *UserApp {
	return &UserApp{
		userService: model.UserService{DB: lifecycle.MySQL},
		hash:        utils.NewHash(nil),
	}
}

type UserApp struct {
	userService model.UserService
	hash        utils.Hash
}

func (ua *UserApp) UserList(c *gin.Context) {
	query := httplib.QueryParams{}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	users, err := ua.userService.Query(c, &query)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, users)
}

func (ua *UserApp) UserDetails(c *gin.Context) {
	uid, err := strconv.ParseInt(c.Param("uid"), 10, 0)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	user, userDetail, err := ua.userService.QueryByID(c, uid)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, map[string]interface{}{
		"user":   user,
		"detail": userDetail,
	})
}

func (ua *UserApp) CreateUser(c *gin.Context) {
	req := struct {
		Name     string `json:"name"`
		Alias    string `json:"alias"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		httplib.Failure(c, err)
		return
	}
	err := lifecycle.MySQL.Transaction(func(tx *gorm.DB) error {
		us := model.UserService{
			DB: tx,
		}
		user := model.User{
			Name:   req.Name,
			Alias:  req.Alias,
			Enable: true,
		}
		userDetail := model.UserDetail{
			Email:    req.Email,
			Password: ua.hash.Sha256([]byte(req.Password)),
		}
		return us.Create(c, &user, &userDetail)
	})
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, nil)
}
