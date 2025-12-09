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

	access, err = aush.GenerateAccessToken(int(user.UserID))
	if err != nil {
		return
	}

	refresh, err = aush.GenerateRefreshToken(int(user.UserID))
	if err != nil {
		return
	}

	// refresh действует 30 суток
	expires := time.Now().Add(720 * time.Hour)

	err = storage.CreateRefreshToken(int(user.UserID), refresh, expires)
	if err != nil {
		return
	}

	return
}
