package storage	

import (
	"backend-survey-app/internal/entities"
	responses "backend-survey-app/pkg/errors"
)

func GetTokenByUserId(userId int) (session entities.Session, err error) {
	row := database.QueryRow("select session_id, user_id, refresh_token, expires_at from sessions where user_id = $1", userId)

	err = row.Scan(&session.SessionID, &session.UserID, &session.RefreshToken, &session.ExpiresAt)
	if err != nil {
		err = responses.ErrInternalServer
		return
	}

	return
}
