package AuthRoutes

import (
	usercontroller "main/cmd/internals/controllers/UserController"
	authmiddleware "main/cmd/internals/middlewares/auth_middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, userController *usercontroller.UserController) {
	routerGroup := r.Group("/auth")
	{
		routerGroup.POST("/register", userController.RegisterUserHandler)
		routerGroup.GET("/user", authmiddleware.AuthMiddleware(), userController.GetUserByID)
		routerGroup.GET("/all/users", authmiddleware.AuthMiddleware(), userController.GetAllUsers)
		routerGroup.POST("/login", userController.LoginHandler)
	}
}
