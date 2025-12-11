package service

import (
	"backend-survey-app/internal/entities"
	"backend-survey-app/internal/storage"
	responses "backend-survey-app/pkg/errors"
	"time"
)

func CreateQuiz(userId int, title, desc string) (err error) {
	if title == "" {
		return responses.ErrBadRequest
	}

	now := time.Now()

	newQuiz := entities.Quiz{
		Title:       title,
		Description: desc,
		IsActive:    true,
		CreatedAt:   now,
		UpdatedAt:   now,
		CreatorID:   uint(userId),
	}

	return storage.CreateQuiz(newQuiz)
}

func GetQuiz(quizId int) (map[string]interface{}, error) {
	if quizId <= 0 {
		return nil, responses.ErrBadRequest
	}

	quiz, err := storage.GetQuizById(quizId)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{
		"id":          quiz.ID,
		"title":       quiz.Title,
		"description": quiz.Description,
		"created_at":  quiz.CreatedAt,
		"updated_at":  quiz.UpdatedAt,
		"is_active":   quiz.IsActive,
	}

	return result, nil
}

func UpdateQuiz(quizId int, title, desc string) error {
	if quizId <= 0 || title == "" {
		return responses.ErrBadRequest
	}

	return storage.UpdateQuiz(quizId, title, desc)
}

func DeleteQuiz(quizId int) error {
	if quizId <= 0 {
		return responses.ErrBadRequest
	}

	return storage.DeleteQuiz(quizId)
}
