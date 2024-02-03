package v1

import "github.com/gin-gonic/gin"

func DefaultRoleApp() *RoleApp {
	return &RoleApp{}
}

type RoleApp struct{}

func (app *RoleApp) List(c *gin.Context) {}

func (app *RoleApp) Create(c *gin.Context) {}
