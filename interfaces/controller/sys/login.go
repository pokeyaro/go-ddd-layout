package sys

import (
	"fmt"
	"net/http"

	"server/infrastructure/common/response"
	dto "server/interfaces/dto/sys"

	"github.com/gin-gonic/gin"
)

// @Summary Login to the system
// @Description User login to obtain authentication token
// @Tags No Authentication Required
// @Accept x-www-form-urlencoded
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} response.BaseResponse[dto.TokenResp] "Success"
// @Failure 400 {object} response.ErrorResponse "BadRequest"
// @Router /sys/login [post]
func (ctl *EndpointCtl) Login(c *gin.Context) {
	var loginData dto.LoginData
	if err := c.ShouldBind(&loginData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(fmt.Errorf("invalid form data. %v", err)))
		return
	}

	userToken, err := ctl.Srv.LocalLogin(loginData.Username, loginData.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(err))
		return
	}

	c.JSON(http.StatusOK, response.Data(dto.TokenResp{
		AccessToken: userToken,
	}))
}
