package v1

import "github.com/gin-gonic/gin"

func DefaultApiApp() *ApiApp {
	return &ApiApp{}
}

type ApiApp struct{}

func (app *ApiApp) List(c *gin.Context) {}
