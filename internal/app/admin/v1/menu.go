package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"proj/internal/domain/model"
	"proj/internal/service"
	"proj/lifecycle"
	"proj/public/httplib"
	"strconv"
)

func DefaultMenuApp() *MenuApp {
	return &MenuApp{
		menuService: model.MenuService{DB: lifecycle.MySQL},
	}
}

type MenuApp struct {
	menuService model.MenuService
	menuSvc     service.MenuService
}

func (app *MenuApp) MenuList(c *gin.Context) {
	query := httplib.QueryParams{
		PageSize: 100,
		Sort:     "level ASC, order_id ASC",
	}
	menus, err := app.menuService.Query(c, &query)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, app.menuSvc.Menus(menus))
}

func (app *MenuApp) CreateMenu(c *gin.Context) {
	req := struct {
		Name  string `json:"name"`
		Icon  string `json:"icon"`
		Route string `json:"route"`
		Hide  bool   `json:"hide"`
		Order int64  `json:"order"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		httplib.Failure(c, err)
		return
	}
	menu, err := (model.Menu{}).New(req)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	err = app.menuService.Create(c, menu)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, menu)
}

func (app *MenuApp) CreateSecondaryMenu(c *gin.Context) {
	mid, err := strconv.ParseInt(c.Param("mid"), 10, 0)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	parentMenu, err := app.menuService.QueryByID(c, mid)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	if parentMenu.ID < 1 {
		httplib.Failure(c, fmt.Errorf("menu %d not found", mid))
		return
	}
	req := struct {
		Name  string `json:"name"`
		Icon  string `json:"icon"`
		Route string `json:"route"`
		Hide  bool   `json:"hide"`
		Order int64  `json:"order"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		httplib.Failure(c, err)
		return
	}
	menu, err := (model.Menu{}).New(req)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	menu.RootID = mid
	menu.ParentID = mid
	menu.Level = 1
	err = app.menuService.Create(c, menu)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	httplib.Success(c, menu)
}

func (app *MenuApp) MenuDetails(c *gin.Context) {
	mid, err := strconv.ParseInt(c.Param("mid"), 10, 0)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	menu, err := app.menuService.QueryByID(c, mid)
	if err != nil {
		httplib.Failure(c, err)
		return
	}
	if menu.ID < 1 {
		httplib.Failure(c, fmt.Errorf("menu %d not found", mid))
		return
	}
	httplib.Success(c, menu)
}
