package user

import (
	"errors"
	"net/http"

	"server/infrastructure/common/context"
	"server/infrastructure/common/response"
	assembler "server/interfaces/assembler/user"
	dto "server/interfaces/dto/user"

	"github.com/gin-gonic/gin"
)

// @Summary Create user
// @Description Create a new user (This API is only available to administrators)
// @Tags Authentication Required
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user body dto.User true "User information"
// @Success 201 {object} response.BaseResponse[any] "Success"
// @Failure 400 {object} response.ErrorResponse "BadRequest"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 403 {object} response.ErrorResponse "Forbidden"
// @Router /user [post]
func (ctl *EndpointCtl) CreateUser(c *gin.Context) {
	isAdmin, err := context.GetIsAdminPerm(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(err))
		return
	}

	if !isAdmin {
		c.AbortWithStatusJSON(http.StatusForbidden, response.Fail(errors.New("unauthorized access")))
		return
	}

	var dtoUser dto.User
	err = c.ShouldBindJSON(&dtoUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(errors.New("request body error")))
		return
	}

	entityUser := assembler.DTOToEntity(&dtoUser)

	_, err = ctl.Srv.CreateUser(entityUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(err))
		return
	}

	c.JSON(http.StatusCreated, response.Ok())
}
