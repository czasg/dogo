package webserver

import (
	"github.com/gin-gonic/gin"
	"proj/internal/app/admin"
	"proj/internal/app/health"
	"proj/internal/middleware/auth"
	"proj/internal/middleware/cors"
	"proj/internal/middleware/trace"
)

func NewApp() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.ContextWithFallback = true
	app.Use(
		gin.Recovery(),
		trace.NewTraceHandler(),
		cors.NewCorsHandler(),
		auth.NewJwtHandler(
			auth.IgnorePrefix([]string{"/readiness", "/liveness"}),
			auth.IgnoreSuffix([]string{"/login"}),
		),
	)
	health.Bind(app)
	admin.Bind(app)
	return app
}
