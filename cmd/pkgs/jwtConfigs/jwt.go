package jwtconfigs

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(username string) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (bool, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return false, fmt.Errorf("JWT_SECRET environment variable not set")
	}

	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, fmt.Errorf("invalid token claims")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return false, fmt.Errorf("expiration claim (exp) missing or invalid in token")
	}

	// Convert float64 to time.Time
	expTime := time.Unix(int64(exp), 0)
	if time.Now().After(expTime) {
		return false, fmt.Errorf("token is expired")
	}

	return true, nil
}
