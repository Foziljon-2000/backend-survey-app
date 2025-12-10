package service

import (
	"backend-survey-app/internal/entities"
	"backend-survey-app/internal/storage"
	responses "backend-survey-app/pkg/errors"
	"backend-survey-app/pkg/utils"
	"net/mail"
	"time"
)

func CreateUser(login, email, password string) (err error) {
	// valid
	_, err = mail.ParseAddress(email)
	if err != nil {
		err = responses.ErrInvalidEmail
		return
	}

	user, err := storage.GetUserByEmail(email)
	if err != nil && err != responses.ErrUserDoesNotExist {

		return
	}

	if user.Email == email {
		err = responses.ErrThisEmailIsAlreadyTaken
		return
	}

	if login == "" || password == "" {
		err = responses.ErrBadRequest
		return
	}

	// logic

	now := time.Now()
	hashPass, err := utils.CreateHashPassword(password)
	if err != nil {
		err = responses.ErrInternalServer
		return
	}

	newUser := entities.User{
		Login:     login,
		Email:     email,
		Password:  hashPass,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = storage.CreateUser(newUser)
	if err != nil {
		return
	}

	return
}
