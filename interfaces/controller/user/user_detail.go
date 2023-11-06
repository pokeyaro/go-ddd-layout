package user

import (
	"errors"
	"net/http"
	"strconv"

	"server/infrastructure/common/context"
	"server/infrastructure/common/response"
	assembler "server/interfaces/assembler/user"

	"github.com/gin-gonic/gin"
)

// @Summary View user profile
// @Description View profile information of a specific user (This API is only available to administrators)
// @Tags Authentication Required
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Success 200 {object} response.BaseResponse[user.User] "Success"
// @Failure 400 {object} response.ErrorResponse "BadRequest"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 403 {object} response.ErrorResponse "Forbidden"
// @Router /user/{id} [get]
func (ctl *EndpointCtl) UserDetail(c *gin.Context) {
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

	user, err := ctl.Srv.GetUserDetail(uid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(err))
		return
	}

	dtoUser := assembler.EntityToDTO(user)
	c.JSON(http.StatusOK, response.Data(dtoUser))
}
