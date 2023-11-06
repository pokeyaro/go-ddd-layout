package user

import (
	"fmt"
	"log"
	"net/http"

	"server/infrastructure/common/context"
	"server/infrastructure/common/response"
	dto "server/interfaces/dto/user"

	"github.com/gin-gonic/gin"
)

// @Summary View user information
// @Description View personal user information
// @Tags Authentication Required
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.BaseResponse[dto.Userinfo] "Success"
// @Failure 400 {object} response.ErrorResponse "BadRequest"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Router /user/info [get]
func (ctl *EndpointCtl) Userinfo(c *gin.Context) {
	username, err := context.GetUsername(c.Request.Context())
	roles, err := context.GetRoleList(c.Request.Context())

	log.Printf("The platform role permissions for the current logged-in user (%s) are: %v", username, roles)

	user, err := ctl.Srv.GetUserinfo(username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(fmt.Errorf("username: %s ??? %v", username, err)))
		return
	}

	userinfo := dto.Userinfo{
		Name:         user.NickName,
		EmpID:        user.EmployeeID,
		Username:     user.Username,
		Avatar:       user.Avatar,
		Email:        user.Email,
		Organization: user.Department,
		Company:      user.Company,
		Region:       user.WorkCountry,
		City:         user.WorkCity,
		Roles:        roles,
	}

	c.JSON(http.StatusOK, response.Data(userinfo))
}
