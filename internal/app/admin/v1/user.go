package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"proj/internal/domain/middleware/auth"
	"proj/internal/domain/model"
	"proj/internal/service"
	"proj/lifecycle"
	"proj/public/httplib"
	"proj/public/utils"
	"strconv"
	"unicode/utf8"
)

func DefaultUserApp() *UserApp {
	return &UserApp{
		userService:     model.UserService{DB: lifecycle.MySQL},
		roleMenuService: model.RoleMenuService{DB: lifecycle.MySQL},
		jwtService:      auth.JwtService{Cache: lifecycle.Redis},
		hash:            utils.NewHash(nil),
	}
}

type UserApp struct {
	userService     model.UserService
	roleMenuService model.RoleMenuService
	jwtService      auth.JwtService
	hash            utils.Hash
}

func (app *UserApp) Login(c *gin.Context) {
	req := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		httplib.Failure(c, err)
		return
	}
	user, userDetail, err := app.userService.QueryByName(c, req.Username)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	if app.hash.Sha256([]byte(req.Password)) != userDetail.Password {
		httplib.Failure(c, fmt.Errorf("password error"))
		return
	}
	jwt := c.MustGet(auth.JwtKey).(*auth.Jwt)
	token, err := jwt.Encrypt(user.Name, user.ID, user.Admin)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	c.SetCookie(jwt.Config.SessionKey, token, jwt.Config.SessionKeyExpire, "/", "", false, false)
	httplib.Success(c, token)
}

func (app *UserApp) Logout(c *gin.Context) {
	jwt, payload, ok := app.jwtService.Enable(c)
	if !ok {
		httplib.Failure(c, fmt.Errorf("server error"))
		return
	}
	err := app.jwtService.Logout(c, jwt, payload)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, nil)
}

func (app *UserApp) UserList(c *gin.Context) {
	query, err := (httplib.QueryMapping{
		StringMap: httplib.QueryStringMap{
			Lk: map[string]string{
				"name": "name",
			},
		},
	}).Parse(c)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	users, err := app.userService.Query(c, query)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, users)
}

func (app *UserApp) UserDetails(c *gin.Context) {
	uid, err := strconv.ParseInt(c.Param("uid"), 10, 0)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	user, userDetail, err := app.userService.QueryByID(c, uid)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, map[string]interface{}{
		"user":   user,
		"detail": userDetail,
	})
}

func (app *UserApp) CreateUser(c *gin.Context) {
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
		_, _, err := us.QueryByName(c, req.Name)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if err == nil {
			return fmt.Errorf("user name[%s] has already exists.", req.Name)
		}
		user := model.User{
			Name:   req.Name,
			Alias:  req.Alias,
			Enable: true,
		}
		userDetail := model.UserDetail{
			Email:    req.Email,
			Password: app.hash.Sha256([]byte(req.Password)),
		}
		return us.Create(c, &user, &userDetail)
	})
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, nil)
}

func (app *UserApp) UpdateUserName(c *gin.Context) {
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
	_, err = app.userService.UpdateUserByID(c, uid, utils.Any2Map(req))
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	user, userDetail, err := app.userService.QueryByID(c, uid)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, map[string]interface{}{
		"user":   user,
		"detail": userDetail,
	})
}

func (app *UserApp) UpdateUserDetail(c *gin.Context) {
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
	_, err = app.userService.UpdateUserDetailByUserID(c, uid, utils.Any2Map(req))
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	user, userDetail, err := app.userService.QueryByID(c, uid)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, map[string]interface{}{
		"user":   user,
		"detail": userDetail,
	})
}

func (app *UserApp) UpdateUserPassword(c *gin.Context) {
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
	_, userDetail, err := app.userService.QueryByID(c, uid)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	if app.hash.Sha256([]byte(req.OldPassword)) != userDetail.Password {
		httplib.Failure(c, errors.New("password error"))
		return
	}
	_, err = app.userService.UpdateUserDetailByUserID(c, uid, map[string]interface{}{
		"password": app.hash.Sha256([]byte(req.NewPassword)),
	})
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, nil)
}

func (app *UserApp) UserRoleList(c *gin.Context) {
	uid, err := strconv.ParseInt(c.Param("uid"), 10, 0)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	roles, err := app.userService.QueryUserRoleByID(c, uid)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, roles)
}

func (app *UserApp) UpdateUserRole(c *gin.Context) {
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
	err = app.userService.UpdateUserRoleByID(c, uid, req.RoleID)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, nil)
}

func (app *UserApp) GetUserMenu(c *gin.Context) {
	uid, err := strconv.ParseInt(c.Param("uid"), 10, 0)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	svc := &service.MenuService{
		UserService:     app.userService,
		RoleMenuService: app.roleMenuService,
	}
	menus, err := svc.MenusByUserID(c, uid)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, menus)
}

func (app *UserApp) UpdateUserEnable(c *gin.Context) {
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
	_, err = app.userService.UpdateUserByID(c, uid, utils.Any2Map(req))
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, nil)
}

func (app *UserApp) DeleteUser(c *gin.Context) {
	uid, err := strconv.ParseInt(c.Param("uid"), 10, 0)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	if err := app.userService.DeleteByID(c, uid); err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, nil)
}
