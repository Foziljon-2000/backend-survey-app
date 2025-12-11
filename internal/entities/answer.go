package entities

type IntArray []int

type Answer struct {
	ID               uint     `json:"id"`
	SubmissionID     uint     `json:"submission_id"`
	QuestionID       uint     `json:"question_id"`
	TextAnswer       *string  `json:"text_answer,omitempty"`
	SelectedChoiceID *uint    `json:"selected_choice_id,omitempty"`
	SelectedChoices  IntArray `json:"selected_choices,omitempty"`
}
