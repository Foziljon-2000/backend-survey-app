package storage

import (
	"backend-survey-app/internal/entities"
	responses "backend-survey-app/pkg/errors"
	"database/sql"
)

func GetUserByUserId(userId int) (user entities.User, err error) {
	row := database.QueryRow("select user_id, login, email, password, created_at, updated_at from users where user_id = $1", userId)

	err = row.Scan(&user.UserID, &user.Login, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		err = responses.ErrUserDoesNotExist
		return
	} else if err != nil {
		err = responses.ErrInternalServer
		return
	}

	err = nil
	return
}
