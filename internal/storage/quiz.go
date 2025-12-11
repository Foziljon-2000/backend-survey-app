package storage

import (
	"backend-survey-app/internal/entities"
	responses "backend-survey-app/pkg/errors"
)

func CreateQuiz(q entities.Quiz) error {
	_, err := database.Exec(
		"insert into quiz(title, description, is_active, created_at, updated_at, creator_id) values($1,$2,$3,$4,$5,$6)",
		q.Title, q.Description, q.IsActive, q.CreatedAt, q.UpdatedAt, q.CreatorID,
	)

	if err != nil {
		return responses.ErrInternalServer
	}
	return nil
}

func GetQuizById(id int) (quiz entities.Quiz, err error) {
	row := database.QueryRow(
		"select id, title, description, is_active, created_at, updated_at, creator_id from quiz where id = $1", id)

	err = row.Scan(
		&quiz.ID,
		&quiz.Title,
		&quiz.Description,
		&quiz.IsActive,
		&quiz.CreatedAt,
		&quiz.UpdatedAt,
		&quiz.CreatorID,
	)

	if err != nil {
		return quiz, responses.ErrInternalServer
	}

	return quiz, nil
}

func UpdateQuiz(id int, title, description string) error {
	_, err := database.Exec(
		"update quiz set title=$1, description=$2, updated_at=now() where id=$3",
		title, description, id,
	)
	if err != nil {
		return responses.ErrInternalServer
	}
	return nil
}

func DeleteQuiz(id int) error {
	_, err := database.Exec("delete from quiz where id=$1", id)
	if err != nil {
		return responses.ErrInternalServer
	}
	return nil
}
