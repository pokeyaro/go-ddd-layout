package router

import (
	"net/http"
	"os"

	"server/interfaces/controller"
	"server/interfaces/controller/public"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	baseURI = "/api/v1"
)

func ApiRouter(r *gin.Engine) {
	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	/**
	 * Public routing
	 */

	r.GET("/", public.HelloWorld)
	r.GET(baseURI, func(c *gin.Context) { c.Redirect(http.StatusFound, os.Getenv("APP_LOCAL")) })

	/**
	 * Business routing
	 */

	sysGroup := r.Group(baseURI + "/sys")
	{
		sysGroup.POST("/login", controller.APIs.Sys.Login)   // [POST] /api/v1/sys/login - User login
		sysGroup.POST("/logout", controller.APIs.Sys.Logout) // [POST] /api/v1/sys/logout - User logout
		sysGroup.GET("/menu", controller.APIs.Sys.Menu)      // [GET]  /api/v1/sys/menu - Get menu
	}

	userGroup := r.Group(baseURI + "/user")
	{
		userGroup.GET("", controller.APIs.User.UserList)                       // [GET]  /api/v1/user - Get user list
		userGroup.POST("", controller.APIs.User.CreateUser)                    // [POST]  /api/v1/user - Create new user
		userGroup.DELETE("/:uid", controller.APIs.User.DeleteUser)             // [DELETE]  /api/v1/user/:id - Delete user by ID
		userGroup.PATCH("/:uid", controller.APIs.User.UpdateUserPartialFields) // [PATCH]  /api/v1/user/:id - Update user with partial fields
		userGroup.PUT("/:uid", controller.APIs.User.UpdateUserAllFields)       // [PUT]  /api/v1/user/:id - Update user with all fields
		userGroup.GET("/:uid", controller.APIs.User.UserDetail)                // [GET]  /api/v1/user/:id - Get user details by ID
		userGroup.GET("/info", controller.APIs.User.Userinfo)                  // [GET]  /api/v1/user/info - Get user info
		userGroup.POST("/:uid/password", controller.APIs.User.UpdatePassword)  // [POST]  /api/v1/user/:id/password - Update user password
	}
}
