package initialize

import (
	"net/http"
	"os"

	"server/interfaces/adapter/middleware"
	"server/interfaces/adapter/router"

	"github.com/gin-gonic/gin"
)

func routeEngine() *gin.Engine {
	env := os.Getenv("APP_ENV")
	switch env {
	case "development":
		gin.SetMode(gin.DebugMode)
	case "production":
		gin.SetMode(gin.ReleaseMode)
	default:
		// Set mode to debug by default if no environment is specified
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	r.RedirectTrailingSlash = true

	// Return 404 for unknown routes
	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "route not found"})
	})

	// Load middleware
	r.Use(
		middleware.LoggingMiddleware(),
		gin.Recovery(),
		middleware.ErrorMiddleware(),
		middleware.CorsMiddleware(),
		middleware.AuthMiddleware(),
	)

	// Register sub-routes
	router.ApiRouter(r)

	return r
}
