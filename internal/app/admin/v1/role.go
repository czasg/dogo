package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"proj/internal/domain/model"
	"proj/lifecycle"
	"proj/public/httplib"
	"proj/public/utils"
	"strconv"
)

func DefaultRoleApp() *RoleApp {
	return &RoleApp{
		roleService: model.RoleService{DB: lifecycle.MySQL},
	}
}

type RoleApp struct {
	roleService model.RoleService
}

func (app *RoleApp) List(c *gin.Context) {
	query := httplib.QueryParams{}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	fmt.Println(query)
	users, err := app.roleService.Query(c, &query)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, users)
}

func (app *RoleApp) Create(c *gin.Context) {
	req := struct {
		Name  string `json:"name"`
		Alias string `json:"alias"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		httplib.Failure(c, err)
		return
	}
	_, err := app.roleService.QueryByName(c, req.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Error(err)
		httplib.Failure(c, errors.New("server unknown error."))
		return
	}
	if err == nil {
		httplib.Failure(c, fmt.Errorf("role[%s] has already exists.", req.Name))
		return
	}
	err = app.roleService.Create(c, &model.Role{Name: req.Name, Alias: req.Alias})
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, nil)
}

func (app *RoleApp) UpdateRoleDetails(c *gin.Context) {
	rid, err := strconv.ParseInt(c.Param("rid"), 10, 0)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	req := struct {
		Alias string `json:"alias,omitempty"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		httplib.Failure(c, err)
		return
	}
	role, err := app.roleService.UpdateRoleByID(c, rid, utils.Any2Map(req))
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, role)
}

func (app *RoleApp) UpdateMenus(c *gin.Context) {}

func (app *RoleApp) UpdateApis(c *gin.Context) {}
