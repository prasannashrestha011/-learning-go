package usercontroller

import (
	"log"
	services "main/cmd/internals/services/UserService"
	"main/cmd/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func InitUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}
func (userService *UserController) RegisterUserHandler(ctx *gin.Context) {
	var reqBody models.UserModel
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	userService.userService.RegisterNewUser(&reqBody)
	log.Println("Request body-> ", reqBody)
	ctx.JSON(http.StatusOK, gin.H{
		"message": reqBody,
	})
}
