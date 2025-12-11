package handler

import (
	"backend-survey-app/internal/service"
	responses "backend-survey-app/pkg/errors"
	httpResponser "backend-survey-app/pkg/http"
	"encoding/json"
	"net/http"
	"strconv"
)

type CreateQuizDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreateQuiz(w http.ResponseWriter, r *http.Request) {
	var resp httpResponser.Response
	defer resp.Concerter(w)

	var dto CreateQuizDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		resp.Message = responses.ErrBadRequest.Error()
		return
	}
	defer r.Body.Close()

	uid := r.Context().Value("user_id")
	if uid == nil {
		resp.Message = responses.ErrUnauthorized.Error()
		return
	}
	userId := uid.(int)

	err = service.CreateQuiz(userId, dto.Title, dto.Description)
	if err != nil {
		resp.Message = err.Error()
		return
	}

	resp.Message = responses.ErrSuccess.Error()
}

func GetQuiz(w http.ResponseWriter, r *http.Request) {
	var resp httpResponser.Response
	defer resp.Concerter(w)

	idStr := r.URL.Query().Get("quiz_id")
	quizId, err := strconv.Atoi(idStr)
	if err != nil {
		resp.Message = responses.ErrBadRequest.Error()
		return
	}

	result, err := service.GetQuiz(quizId)
	if err != nil {
		resp.Message = err.Error()
		return
	}

	resp.Payload = result
	resp.Message = responses.ErrSuccess.Error()
}

func UpdateQuiz(w http.ResponseWriter, r *http.Request) {
	var resp httpResponser.Response
	defer resp.Concerter(w)

	idStr := r.URL.Query().Get("quiz_id")
	quizId, err := strconv.Atoi(idStr)
	if err != nil {
		resp.Message = responses.ErrBadRequest.Error()
		return
	}

	var dto CreateQuizDTO
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		resp.Message = responses.ErrBadRequest.Error()
		return
	}
	defer r.Body.Close()

	err = service.UpdateQuiz(quizId, dto.Title, dto.Description)
	if err != nil {
		resp.Message = err.Error()
		return
	}

	resp.Message = responses.ErrSuccess.Error()
}

func DeleteQuiz(w http.ResponseWriter, r *http.Request) {
	var resp httpResponser.Response
	defer resp.Concerter(w)

	idStr := r.URL.Query().Get("quiz_id")
	quizId, err := strconv.Atoi(idStr)
	if err != nil {
		resp.Message = responses.ErrBadRequest.Error()
		return
	}

	err = service.DeleteQuiz(quizId)
	if err != nil {
		resp.Message = err.Error()
		return
	}

	resp.Message = responses.ErrSuccess.Error()
}
