package middleware

import (
	"context"
	"net/http"
	"strings"

	"BlogApp/utils"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &utils.CustomClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return utils.JwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// store user details in context
		ctx := context.WithValue(r.Context(), "email", claims.Email)
		ctx = context.WithValue(ctx, "username", claims.Username)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}