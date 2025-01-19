package userroute

import (
	usercontroller "main/cmd/internals/controllers/UserController"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine, userController *usercontroller.UserController) {
	userRoutes := route.Group("/user")
	{
		userRoutes.POST("/register", userController.RegisterUserHandler)
	}
}
