package usercontroller

import (
	"log"
	UserDTOS "main/cmd/internals/dtos"
	UserService "main/cmd/internals/services"
	jwtconfigs "main/cmd/pkgs/jwtConfigs"
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
	var reqBody *UserDTOS.CreateUserDTO
	if err := ctx.BindJSON(&reqBody); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	responseDto := s.service.RegisterUser(reqBody)
	ctx.JSON(responseDto.StatusCode, gin.H{
		"message": responseDto.Message,
	})
}

func (s UserController) LoginHandler(ctx *gin.Context) {
	var reqBody *UserDTOS.AuthUserDTO
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login information"})
	}
	authDto := s.service.AuthenticatedUser(*reqBody)
	if authDto.Success {
		accessToken, err := jwtconfigs.CreateAccessToken(reqBody.Username)
		if err != nil {
			log.Fatal("error", err.Error())
		}
		refreshToken, err := jwtconfigs.CreateRefreshToken(reqBody.Username)
		if err != nil {
			log.Fatalln("Refresh token error->", err.Error())
		}
		ctx.Header("Access_Token", accessToken)
		ctx.Header("Refresh_Token", refreshToken)
	}
	ctx.JSON(authDto.StatusCode, gin.H{
		"message": authDto,
	})
}
func (s UserController) RenewAccessTokenHandler(ctx *gin.Context) {
	refreshTokenString := ctx.Request.Header.Get("Refresh_Token")
	if refreshTokenString == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Refresh token is not provided",
		})
		return
	}
	accessToken, err := s.service.RenewAccessToken(refreshTokenString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Header("Access_Token", accessToken)
	ctx.JSON(http.StatusOK, gin.H{})
}
func (s UserController) GetUserByID(ctx *gin.Context) {
	userId := ctx.DefaultQuery("userId", "")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "User id is required",
		})
		return
	}
	responseDto := s.service.GetUserByID(userId)
	ctx.JSON(responseDto.StatusCode, gin.H{
		"message": responseDto,
	})
}
func (s UserController) GetAllUsers(ctx *gin.Context) {
	responseDto := s.service.GetAllUsers()
	ctx.JSON(responseDto.StatusCode, gin.H{
		"message": responseDto,
	})
}
