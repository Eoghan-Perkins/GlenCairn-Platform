package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretToken = []byte("secret_user_token")

func GenerateToken(id uint) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString(secretToken)
}
