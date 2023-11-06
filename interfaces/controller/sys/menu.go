package sys

import (
	"net/http"

	"server/infrastructure/common/response"
	dto "server/interfaces/dto/sys"

	"github.com/gin-gonic/gin"
)

// @Summary Menu bar
// @Description Get user menu bar
// @Tags Authentication Required
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.BaseResponse[dto.Menus] "Success"
// @Router /sys/menu [get]
func (ctl *EndpointCtl) Menu(c *gin.Context) {
	menus := dto.Menus{
		{
			Path: "workplace",
			Name: "Workplace",
			Meta: dto.MenuMeta{
				Locale:       "menu.dashboard",
				RequiresAuth: true,
				SvgIcon:      "menu-overview",
				TopLevel:     true,
				Order:        0,
				Roles:        []string{"*"},
			},
		},
		{
			Path: "/user",
			Name: "User",
			Meta: dto.MenuMeta{
				Locale:       "menu.user",
				RequiresAuth: true,
				Icon:         "icon-user",
				Order:        18,
				HideInMenu:   true,
			},
			Children: []dto.MenuChild{
				{
					Path: "create",
					Name: "UserCreate",
					Meta: dto.MenuMeta{
						Locale:       "menu.user.create",
						RequiresAuth: true,
						Roles:        []string{"*"},
						HideInMenu:   true,
					},
				},
				{
					Path: "list",
					Name: "UserList",
					Meta: dto.MenuMeta{
						Locale:       "menu.user.list",
						RequiresAuth: true,
						Roles:        []string{"*"},
						HideInMenu:   true,
					},
				},
				{
					Path: "info",
					Name: "Info",
					Meta: dto.MenuMeta{
						Locale:       "menu.user.info",
						RequiresAuth: true,
						Roles:        []string{"*"},
						HideInMenu:   true,
					},
				},
			},
		},
	}

	c.JSON(http.StatusOK, response.Data(menus))
}
