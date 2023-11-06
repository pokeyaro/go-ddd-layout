package user

import (
	"errors"
	"net/http"
	"strconv"

	"server/infrastructure/common/context"
	"server/infrastructure/common/response"
	dto "server/interfaces/dto/user"

	"github.com/gin-gonic/gin"
)

// @Summary Update user password
// @Description Update user's password
// @Tags Authentication Required
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Success 200 {object} response.BaseResponse[dto.LoginPasswd] "Success"
// @Failure 400 {object} response.ErrorResponse "BadRequest"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Router /user/{id}/password [post]
func (ctl *EndpointCtl) UpdatePassword(c *gin.Context) {
	username, err := context.GetUsername(c.Request.Context())
	userID, err := context.GetUID(c.Request.Context())
	isAdmin, err := context.GetIsAdminPerm(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(err))
		return
	}

	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(errors.New("parameter error")))
		return
	}

	if !isAdmin && uid != userID {
		c.AbortWithStatusJSON(http.StatusForbidden, response.Fail(errors.New("unauthorized access")))
		return
	}

	newPassword, err := ctl.Srv.UpdatePassword(uid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(err))
		return
	}

	c.JSON(http.StatusOK, response.Data(dto.LoginPasswd{
		Username: username,
		Password: newPassword,
	}))
}
