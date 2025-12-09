package handler

import (
	"backend-survey-app/internal/service"
	responses "backend-survey-app/pkg/errors"
	httpResponser "backend-survey-app/pkg/http"
	"net/http"
	"strconv"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	var resp httpResponser.Response
	defer resp.Concerter(w)

	userIdStr := r.Header.Get("X-User-ID")
	userId, _ := strconv.Atoi(userIdStr)

	err := service.Logout(userId)
	if err != nil {
		resp.Message = err.Error()
		return
	}

	resp.Message = responses.ErrSuccess.Error()
}
