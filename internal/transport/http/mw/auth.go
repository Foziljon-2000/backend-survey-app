package mw

import (
	responses "backend-survey-app/pkg/errors"
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			http.Error(w, responses.ErrUnauthorized.Error(), http.StatusUnauthorized)
			return
		}

		if strings.HasPrefix(strings.ToLower(tokenHeader), "bearer ") {
			tokenHeader = strings.TrimSpace(tokenHeader[7:])
		}

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			http.Error(w, responses.ErrInternalServer.Error(), http.StatusInternalServerError)
			return
		}

		token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, responses.ErrUnauthorized.Error(), http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, responses.ErrUnauthorized.Error(), http.StatusUnauthorized)
			return
		}

		userIdFloat, ok := claims["user_id"].(float64)
		if !ok {
			http.Error(w, responses.ErrUnauthorized.Error(), http.StatusUnauthorized)
			return
		}

		userId := int(userIdFloat)

		ctx := context.WithValue(r.Context(), "user_id", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
