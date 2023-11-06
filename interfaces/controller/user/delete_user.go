package user

import (
	"errors"
	"net/http"
	"strconv"

	"server/infrastructure/common/context"
	"server/infrastructure/common/response"

	"github.com/gin-gonic/gin"
)

// @Summary Delete user
// @Description Delete a user (This API is only available to administrators)
// @Tags Authentication Required
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param uid path int true "User ID"
// @Success 200 {object} response.BaseResponse[any] "Success"
// @Failure 400 {object} response.ErrorResponse "BadRequest"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 403 {object} response.ErrorResponse "Forbidden"
// @Router /user/{id} [delete]
func (ctl *EndpointCtl) DeleteUser(c *gin.Context) {
	isAdmin, err := context.GetIsAdminPerm(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(err))
		return
	}

	if !isAdmin {
		c.AbortWithStatusJSON(http.StatusForbidden, response.Fail(errors.New("unauthorized access")))
		return
	}

	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(errors.New("parameter error")))
		return
	}

	err = ctl.Srv.DeleteUser(uid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(err))
		return
	}

	c.JSON(http.StatusOK, response.Ok())
}
