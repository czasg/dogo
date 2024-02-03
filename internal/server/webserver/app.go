package webserver

import (
	"github.com/gin-gonic/gin"
	"proj/internal/app/admin"
	"proj/internal/app/health"
	auth2 "proj/internal/domain/middleware/auth"
	"proj/internal/domain/middleware/cors"
	"proj/internal/domain/middleware/trace"
	"proj/lifecycle"
)

func NewApp() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.ContextWithFallback = true
	app.Use(
		gin.Recovery(),
		trace.NewTraceHandler(),
		cors.NewCorsHandler(),
		auth2.NewJwtHandler(
			lifecycle.Redis,
			auth2.IgnorePrefix([]string{"/readiness", "/liveness"}),
			auth2.IgnoreSuffix([]string{"/login"}),
		),
	)
	health.Bind(app)
	admin.Bind(app)
	return app
}
