package auth

import (
	"encoding/base64"
	"strings"

	"github.com/gin-gonic/gin"
)

// ExtractBearerToken extracts the Bearer Token from the request header
func ExtractBearerToken(c *gin.Context) (string, bool) {
	authHeader := c.GetHeader("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer "), true
	}
	return "", false
}

// ExtractBasicToken extracts the Basic Token from the request header
func ExtractBasicToken(c *gin.Context) (string, bool) {
	authHeader := c.GetHeader("Authorization")
	if strings.HasPrefix(authHeader, "Basic ") {
		encodedCreds := strings.TrimPrefix(authHeader, "Basic ")
		decodedCreds, err := base64.StdEncoding.DecodeString(encodedCreds)
		if err != nil {
			return "", false
		}
		return string(decodedCreds), true
	}
	return "", false
}
