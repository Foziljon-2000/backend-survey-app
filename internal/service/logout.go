package service

import "backend-survey-app/internal/storage"

func Logout(userId int) error {
	return storage.DeleteRefreshToken(userId)
}
