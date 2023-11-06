package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Hello, World
// @Description Basic example page
// @Tags Example
// @Produce text/plain
// @Success 200 {object} string "Success"
// @Router / [get]
func HelloWorld(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}
