package aush

import (
	responses "backend-survey-app/pkg/errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userId int) (newToken string, err error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		err = responses.ErrInternalServer
		return
	}

	newToken, err = token.SignedString([]byte(secret))
	if err != nil {
		err = responses.ErrInternalServer
		return
	}

	return
}
