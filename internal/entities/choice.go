package entities

type Choice struct {
	ID         uint   `json:"id"`
	QuestionID uint   `json:"question_id"`
	Text       string `json:"text"`
}
