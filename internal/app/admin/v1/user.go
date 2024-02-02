package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"proj/internal/domain/model"
	"proj/lifecycle"
	"proj/public/httplib"
	"proj/public/utils"
	"strconv"
	"unicode/utf8"
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

func (ua *UserApp) UpdateUserName(c *gin.Context) {
	uid, err := strconv.ParseInt(c.Param("uid"), 10, 0)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	req := struct {
		Name  string `json:"name,omitempty"`
		Alias string `json:"alias,omitempty"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		httplib.Failure(c, err)
		return
	}
	if req.Name == "" || req.Alias == "" {
		httplib.Failure(c, errors.New("invalid user name"))
		return
	}
	_, err = ua.userService.UpdateUserByID(c, uid, utils.Any2Map(req))
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

func (ua *UserApp) UpdateUserDetail(c *gin.Context) {
	uid, err := strconv.ParseInt(c.Param("uid"), 10, 0)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	req := struct {
		Email      string `json:"email,omitempty"`
		Preference string `json:"preference,omitempty"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		httplib.Failure(c, err)
		return
	}
	_, err = ua.userService.UpdateUserDetailByUserID(c, uid, utils.Any2Map(req))
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

func (ua *UserApp) UpdateUserPassword(c *gin.Context) {
	uid, err := strconv.ParseInt(c.Param("uid"), 10, 0)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	req := struct {
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		httplib.Failure(c, err)
		return
	}
	if utf8.RuneCountInString(req.NewPassword) < 5 {
		httplib.Failure(c, errors.New("new password is too short"))
		return
	}
	_, userDetail, err := ua.userService.QueryByID(c, uid)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	if ua.hash.Sha256([]byte(req.OldPassword)) != userDetail.Password {
		httplib.Failure(c, errors.New("password error"))
		return
	}
	_, err = ua.userService.UpdateUserDetailByUserID(c, uid, map[string]interface{}{
		"password": ua.hash.Sha256([]byte(req.NewPassword)),
	})
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, nil)
}

func (ua *UserApp) UpdateUserRole(c *gin.Context) {
	uid, err := strconv.ParseInt(c.Param("uid"), 10, 0)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	req := struct {
		RoleID []int64 `json:"roleID"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		httplib.Failure(c, err)
		return
	}
	err = ua.userService.UpdateUserRoleByID(c, uid, req.RoleID)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, nil)
}

func (ua *UserApp) UpdateUserEnable(c *gin.Context) {
	uid, err := strconv.ParseInt(c.Param("uid"), 10, 0)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	enable, err := strconv.ParseBool(c.Query("enable"))
	if err != nil {
		httplib.Failure(c, errors.New("invalid enable"))
		return
	}
	req := struct {
		Enable bool `json:"enable"`
	}{
		Enable: enable,
	}
	_, err = ua.userService.UpdateUserDetailByUserID(c, uid, utils.Any2Map(req))
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, nil)
}

func (ua *UserApp) DeleteUser(c *gin.Context) {
	uid, err := strconv.ParseInt(c.Param("uid"), 10, 0)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	if err := ua.userService.DeleteByID(c, uid); err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, nil)
}
