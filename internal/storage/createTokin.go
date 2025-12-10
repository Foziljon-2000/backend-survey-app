package storage

import (
	responses "backend-survey-app/pkg/errors"
	"time"
)

func CreateRefreshToken(userId int, refreshToken string, expires time.Time) (err error) {
	_, err = database.Exec("insert into sessions(user_id, refresh_token, expires_at) values($1, $2, $3)", userId, refreshToken, expires)
	if err != nil {
		err = responses.ErrInternalServer
		return
	}
	return
}
