package mw

import (
	"context"
	"net/http"
	"os"

	responses "backend-survey-app/pkg/errors"

	"github.com/golang-jwt/jwt/v5"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			http.Error(w, responses.ErrUnauthorized.Error(), http.StatusUnauthorized)
			return
		}

		secret := os.Getenv("JWT_SECRET")
		token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, responses.ErrUnauthorized.Error(), http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userId := claims["user_id"]

		ctx := context.WithValue(r.Context(), "user_id", int(userId.(float64)))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
