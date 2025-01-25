package jwtmethods

import (
	"errors"
	jwtconfigs "main/cmd/pkgs/jwtConfigs"

	"github.com/dgrijalva/jwt-go"
)

func RenewAccessToken(refreshTokenStr string) (string, error) {
	isValid, err := jwtconfigs.ValidateToken(refreshTokenStr)
	if err != nil {
		return "", err
	}
	if !isValid {
		return "", errors.New("unable to validate the token")
	}

	refreshToken, err := jwtconfigs.ParseAuthToken(refreshTokenStr)
	if err != nil {
		return "", err
	}

	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to retrive claims from the token")
	}
	subject := claims["sub"].(string)

	newAccessToken, err := jwtconfigs.CreateAccessToken(subject)
	if err != nil {
		return "", err
	}
	return newAccessToken, nil

}
