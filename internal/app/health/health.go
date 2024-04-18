package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"proj/lifecycle"
)

func Bind(app gin.IRouter) {
	app.GET("/readiness", noContentHandler)
	app.GET("/liveness", lifecycleCheck)
}

func noContentHandler(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func lifecycleCheck(ctx *gin.Context) {
	if err := lifecycle.MySQL.Exec("SELECT 1").Error; err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if err := lifecycle.Redis.Ping(ctx); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
