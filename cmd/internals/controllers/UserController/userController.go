package usercontroller

import (
	UserService "main/cmd/internals/services"
	UserModel "main/cmd/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *UserService.UserService
}

func InitUserController(service *UserService.UserService) *UserController {
	return &UserController{
		service: service,
	}
}
func (s UserController) RegisterUserHandler(ctx *gin.Context) {
	var reqBody *UserModel.UserModel
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Insufficent details",
		})
		return
	}
	responseDto := s.service.RegisterUser(reqBody)
	ctx.JSON(responseDto.StatusCode, gin.H{
		"message": responseDto.Message,
	})
}
