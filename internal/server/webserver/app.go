package webserver

import (
	"github.com/gin-gonic/gin"
	"proj/internal/api/health"
)

func NewApp() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.Use(gin.Recovery())
	health.Register(app)
	return app
}
