package containers

import (
	usercontroller "main/cmd/internals/controllers/UserController"
	UserRepository "main/cmd/internals/repositories"
	UserService "main/cmd/internals/services"
	AuthRoutes "main/cmd/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserContainer(r *gin.Engine, DB *gorm.DB) {

	userRepo := UserRepository.InitUserRepo(DB)
	userService := UserService.InitUserService(userRepo)
	userController := usercontroller.InitUserController(userService)

	AuthRoutes.AuthRoutes(r, userController)
}
