package entities

import "time"

type Submission struct {
	ID          uint       `json:"id"`
	QuizID      uint       `json:"quiz_id"`
	UserID      *uint      `json:"user_id"`
	StartedAt   time.Time  `json:"started_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}
