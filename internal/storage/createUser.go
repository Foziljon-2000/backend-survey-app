package storage

import (
	"backend-survey-app/internal/entities"
	responses "backend-survey-app/pkg/errors"
)

func CreateUser(user entities.User) (err error) {
	_, err = database.Exec("insert into users(login, email, password, created_at, updated_at) values($1, $2, $3, $4, $5)", user.Login, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		err = responses.ErrInternalServer
		return
	}

	return
}
