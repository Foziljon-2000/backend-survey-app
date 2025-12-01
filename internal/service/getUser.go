package service

import (
	"backend-survey-app/internal/storage"
	responses "backend-survey-app/pkg/errors"
)

func GetUser(userId int) (result map[string]interface{}, err error) {
	// valid
	if userId <= 0 {
		err = responses.ErrBadRequest
		return
	}

	user, err := storage.GetUserByUserId(userId)
	if err != nil {
		return
	}

	result = make(map[string]interface{})

	result["user_id"] = user.UserID
	result["login"] = user.Login
	result["email"] = user.Email

	return
}
