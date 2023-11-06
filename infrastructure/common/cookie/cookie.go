package cookie

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	cookieTokenPath string = "/"
	cookieTokenName string = "X-ACCESS-TOKEN"
	expirationTime         = 3600
)

// SetCookies sets the cookie.
func SetCookies(jwtToken string) *http.Cookie {
	return &http.Cookie{
		Name:     cookieTokenName,
		Value:    jwtToken,
		HttpOnly: false,
		Path:     cookieTokenPath,
		MaxAge:   expirationTime,
	}
}

// ParseCookieToken parses the cookie value.
func ParseCookieToken(c *gin.Context) (string, error) {
	cookies, err := c.Request.Cookie(cookieTokenName)
	if err != nil {
		return "", err
	}

	return cookies.Value, nil
}
