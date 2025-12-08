package service

import (
	"backend-survey-app/internal/service/aush"
	"backend-survey-app/internal/storage"
	responses "backend-survey-app/pkg/errors"
	"backend-survey-app/pkg/utils"
	"time"
)

func Login(email string, password string) (err error) {
	// valid
	user, err := storage.GetUserByEmail(email)
	if err != nil {
		err = responses.ErrUnauthorized
		return
	}

	passwordHash, err := utils.CreateHashPassword(password)
	if err != nil {
		err = responses.ErrInternalServer
		return
	}

	if user.Password != passwordHash {
		err = responses.ErrUnauthorized
		return
	}

	now := time.Now()
	refreshToken, err := aush.GenerateAccessToken(int(user.UserID))
	if err != nil {
		return
	}

	err = storage.CreateAccessTokin(int(user.UserID), refreshToken, now)
	if err != nil {
		return
	}

	return
}
