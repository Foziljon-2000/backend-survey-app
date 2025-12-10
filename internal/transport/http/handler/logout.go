package handler

import (
	"backend-survey-app/internal/service"
	responses "backend-survey-app/pkg/errors"
	httpResponser "backend-survey-app/pkg/http"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	var resp httpResponser.Response
	defer resp.Concerter(w)

	uid := r.Context().Value("user_id")
	if uid == nil {
		resp.Message = responses.ErrUnauthorized.Error()
		return
	}

	userId, ok := uid.(int)
	if !ok {
		resp.Message = responses.ErrUnauthorized.Error()
		return
	}

	err := service.Logout(userId)
	if err != nil {
		resp.Message = err.Error()
		return
	}

	resp.Message = responses.ErrSuccess.Error()
}
