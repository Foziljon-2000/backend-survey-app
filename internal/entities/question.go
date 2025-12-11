package entities

type Question struct {
	ID           uint   `json:"id"`
	QuizID       uint   `json:"quiz_id"`
	Text         string `json:"text"`
	QuestionType string `json:"question_type"`
	Order        int    `json:"order"`
	Required     bool   `json:"required"`
}
