package AuthRoutes

import (
	usercontroller "main/cmd/internals/controllers/UserController"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, userController *usercontroller.UserController) {
	routerGroup := r.Group("/auth")
	{
		routerGroup.POST("/register", userController.RegisterUserHandler)
	}
}
