package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/envy"
)

func CreateToken() (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ExpiresAt": time.Now().Add(time.Minute * 5).Unix(),
	})
	secretKey := []byte(envy.Get("JWT_SECRET", ""))
	tokenString, err = token.SignedString(secretKey)

	if err != nil {
		return "", fmt.Errorf("could not sign token, %v", err)
	}
	return
}
