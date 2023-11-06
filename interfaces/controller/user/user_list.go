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

// @Summary User list
// @Description Get user list (This API is only available to administrators)
// @Tags Authentication Required
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.BaseResponse[any] "Success"
// @Failure 400 {object} response.ErrorResponse "BadRequest"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Router /user [get]
func (ctl *EndpointCtl) UserList(c *gin.Context) {
	isAdmin, err := context.GetIsAdminPerm(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(err))
		return
	}

	if !isAdmin {
		c.AbortWithStatusJSON(http.StatusForbidden, response.Fail(errors.New("unauthorized access")))
		return
	}

	userList, err := ctl.Srv.GetUserList()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(err))
		return
	}

	dtoUsers := assembler.ListEntityToDTO(userList)
	c.JSON(http.StatusOK, response.List[dto.User](dtoUsers))
}
