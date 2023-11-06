package middleware

import (
	"log"
	"net/http"
	"strings"

	"server/infrastructure/common/auth"
	"server/infrastructure/common/context"
	"server/infrastructure/common/jwt"

	"github.com/gin-gonic/gin"
)

var (
	whiteListURLs = map[string]bool{
		"/":                 true,
		"/api/v1":           true,
		"/api/v1/sys/login": true,
	}
	dynamicRoutePrefix = "/swagger"
)

// AuthMiddleware API Token Authentication Middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the requested url path
		uri := c.Request.URL.Path

		// Allow public URLs in the whitelist
		if _, ok := whiteListURLs[uri]; ok || strings.Contains(uri, dynamicRoutePrefix) {
			log.Printf("Accessing whitelisted URL: %s\n", uri)
			c.Next()
			return
		}

		// Verify token for private URLs
		log.Printf("Accessing private URL: %s\n", uri)

		// Extract jwtToken from Bearer Token
		jwtToken, valid := auth.ExtractBearerToken(c)
		if !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":   40000,
				"error":  "Unauthorized",
				"detail": "failed",
			})
			return
		}

		// Validate JWT payload
		claims, err := jwt.ValidateJwt(jwtToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":   40000,
				"error":  "Unauthorized",
				"detail": "failed",
			})
			return
		}

		// Store user-related information in request context
		ctx := context.Store(c.Request.Context(), claims)

		// Update the request context
		c.Request = c.Request.WithContext(ctx)

		// Continue execution
		c.Next()
	}
}
