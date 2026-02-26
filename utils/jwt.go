package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("super_secret_key_change_this")

type CustomClaims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string, username string) (string, error) {

	claims := CustomClaims{
		Email:    email,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(JwtKey)
}