package service

import (
	"backend-survey-app/internal/service/aush"
	"backend-survey-app/internal/storage"
	responses "backend-survey-app/pkg/errors"
	"backend-survey-app/pkg/utils"
	"time"
)

func Login(email string, password string) (access string, refresh string, err error) {
	user, err := storage.GetUserByEmail(email)
	if err != nil {
		return "", "", responses.ErrUnauthorized
	}

	err = utils.ComparePassword(user.Password, password)
	if err != nil {
		return "", "", responses.ErrUnauthorized
	}

	access, err = aush.GenerateAccessToken(int(user.UserID))
	if err != nil {
		return
	}

	refresh, err = aush.GenerateRefreshToken(int(user.UserID))
	if err != nil {
		return
	}

	expires := time.Now().Add(720 * time.Hour)

	err = storage.CreateRefreshToken(int(user.UserID), refresh, expires)
	if err != nil {
		return
	}

	return
}
