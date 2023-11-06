package user

import (
	"errors"
	"net/http"
	"strconv"

	"server/infrastructure/common/context"
	"server/infrastructure/common/response"
	assembler "server/interfaces/assembler/user"
	dto "server/interfaces/dto/user"

	"github.com/gin-gonic/gin"
)

// @Summary Update user profile
// @Description Update partial profile information of a user (This API is only available to administrators)
// @Tags Authentication Required
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Param user body dto.User true "User information"
// @Success 200 {object} response.BaseResponse[any] "Success"
// @Failure 400 {object} response.ErrorResponse "BadRequest"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 403 {object} response.ErrorResponse "Forbidden"
// @Router /user/{id} [patch]
func (ctl *EndpointCtl) UpdateUserPartialFields(c *gin.Context) {
	updateUserFields(ctl, c, true)
}

// @Summary Update user profile
// @Description Update all profile information of a user (This API is only available to administrators)
// @Tags Authentication Required
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Param user body dto.User true "User information"
// @Success 200 {object} response.BaseResponse[any] "Request successful"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 403 {object} response.ErrorResponse "Forbidden"
// @Router /user/{id} [put]
func (ctl *EndpointCtl) UpdateUserAllFields(c *gin.Context) {
	updateUserFields(ctl, c, false)
}

func updateUserFields(ctl *EndpointCtl, c *gin.Context, partialUpdate bool) {
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

	var dtoUser dto.User
	err = c.ShouldBindJSON(&dtoUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(errors.New("request body error")))
		return
	}

	entityUser := assembler.DTOToEntity(&dtoUser)
	entityUser.ID = user.ID

	_, err = ctl.Srv.UpdateUser(entityUser, partialUpdate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(err))
		return
	}

	c.JSON(http.StatusOK, response.Ok())
}
