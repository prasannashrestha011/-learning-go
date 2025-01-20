package main

import (
	databaseconfig "main/cmd/configs/databaseConfig"
	usercontroller "main/cmd/internals/controllers/UserController"
	UserRepository "main/cmd/internals/repositories"
	UserService "main/cmd/internals/services"
	AuthRoutes "main/cmd/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	DB := databaseconfig.Connect()
	userRepo := UserRepository.InitUserRepo(DB)
	userService := UserService.InitUserService(userRepo)
	userController := usercontroller.InitUserController(userService)

	AuthRoutes.AuthRoutes(r, userController)
	r.Run()

}
