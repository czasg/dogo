package v1

import (
	"github.com/gin-gonic/gin"
	"proj/internal/domain/model"
	"proj/lifecycle"
	"proj/public/httplib"
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
	err := app.roleService.Create(c, &model.Role{Name: req.Name, Alias: req.Alias})
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, nil)
}

func (app *RoleApp) UpdateMenus(c *gin.Context) {}

func (app *RoleApp) UpdateApis(c *gin.Context) {}
