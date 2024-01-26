package webserver

import (
	"github.com/gin-gonic/gin"
	"proj/internal/api/health"
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
		jwt.NewJwtHandler(jwt.IgnorePrefix([]string{"/"})),
	)
	health.Register(app)
	return app
}
