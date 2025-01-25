package authmiddleware

import (
	"log"
	jwtconfigs "main/cmd/pkgs/jwtConfigs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.Request.Header.Get("Access_Token")
		log.Println("Your auth header", accessToken)
		if accessToken == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "auth cookie not found"})
			ctx.Abort()
			return
		}
		isValidToken, err := jwtconfigs.ValidateToken(accessToken)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		if !isValidToken {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "session expired, please signin again"})
			ctx.Abort()
			return
		}
		log.Println("User is upto date.....")
		ctx.Next()
	}
}
