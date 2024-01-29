package webserver

import (
	"github.com/gin-gonic/gin"
	"proj/internal/app/admin"
	"proj/internal/app/health"
	"proj/internal/middleware/cors"
	"proj/internal/middleware/jwt"
	"proj/internal/middleware/trace"
)

func NewApp() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.Use(
		gin.Recovery(),
		trace.NewTraceHandler(),
		cors.NewCorsHandler(),
		jwt.NewJwtHandler(
			jwt.IgnorePrefix([]string{"/readiness", "/liveness"}),
			jwt.IgnoreSuffix([]string{"/login"}),
		),
	)
	health.Bind(app)
	admin.Bind(app)
	return app
}
