package storage

func DeleteRefreshToken(userId int) error {
	_, err := database.Exec("delete from sessions where user_id = $1", userId)
	return err
}
