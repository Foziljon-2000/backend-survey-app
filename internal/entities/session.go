package entities

import "time"

type Session struct {
	SessionID    uint      `json:"session_id"`
	UserID       uint      `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}
