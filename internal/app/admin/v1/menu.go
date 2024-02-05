package v1

import "github.com/gin-gonic/gin"

func DefaultMenuApp() *MenuApp {
	return &MenuApp{}
}

type MenuApp struct{}

func (app *MenuApp) List(c *gin.Context) {}
