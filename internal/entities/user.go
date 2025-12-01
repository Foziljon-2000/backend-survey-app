package entities

import "time"

type User struct {
	UserID    uint      `json:"user_id"`
	Login     string    `json:"login"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
