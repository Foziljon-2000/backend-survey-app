package aush

import (
	"os"
	"time"

	responses "backend-survey-app/pkg/errors"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateRefreshToken(userId int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(720 * time.Hour).Unix(), // 30 дней
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", responses.ErrInternalServer
	}

	return token.SignedString([]byte(secret))
}
