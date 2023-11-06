package sys

import (
	"net/http"

	"server/infrastructure/common/response"

	"github.com/gin-gonic/gin"
)

// @Summary Logout from the system
// @Description User logs out of the system
// @Tags Authentication Required
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.BaseResponse[any] "Success"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Router /sys/logout [post]
func (ctl *EndpointCtl) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, response.Ok())
}
