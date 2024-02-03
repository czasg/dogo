package cors

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func NewCorsHandler(configFunc ...ConfigFunc) gin.HandlerFunc {
	config := NewDefaultConfig()
	for _, patch := range configFunc {
		patch(config)
	}
	allowOrigins := strings.Join(config.allowOrigins, ",")
	allowMethods := strings.Join(config.allowMethods, ",")
	allowHeaders := strings.Join(config.allowHeaders, ",")
	var allowCredentials string
	if config.allowCredentials {
		allowCredentials = "true"
	} else {
		allowCredentials = "false"
	}
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", allowOrigins)
		c.Header("Access-Control-Allow-Methods", allowMethods)
		c.Header("Access-Control-Allow-Headers", allowHeaders)
		c.Header("Access-Control-Allow-Credentials", allowCredentials)
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func NewDefaultConfig() *Config {
	return &Config{
		allowOrigins:     []string{"*"},
		allowMethods:     []string{"*"},
		allowHeaders:     []string{"*"},
		allowCredentials: true,
	}
}

type ConfigFunc func(cfg *Config)

type Config struct {
	allowOrigins     []string
	allowMethods     []string
	allowHeaders     []string
	allowCredentials bool
}

func AllowOrigins(allowOrigins []string) ConfigFunc {
	return func(cfg *Config) {
		cfg.allowOrigins = allowOrigins
	}
}

func AllowMethods(allowMethods []string) ConfigFunc {
	return func(cfg *Config) {
		cfg.allowMethods = allowMethods
	}
}

func AllowHeaders(allowHeaders []string) ConfigFunc {
	return func(cfg *Config) {
		cfg.allowHeaders = allowHeaders
	}
}

func AllowCredentials(allowCredentials bool) ConfigFunc {
	return func(cfg *Config) {
		cfg.allowCredentials = allowCredentials
	}
}
