package handler

import (
	"backend-survey-app/internal/service"
	responses "backend-survey-app/pkg/errors"
	httpResponser "backend-survey-app/pkg/http"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var resp httpResponser.Response
	defer resp.Concerter(w)

	var dto LoginRequest

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		resp.Message = responses.ErrBadRequest.Error()
		return
	}
	defer r.Body.Close()

	access, refresh, err := service.Login(dto.Email, dto.Password)
	if err != nil {
		resp.Message = err.Error()
		return
	}

	resp.Message = responses.ErrSuccess.Error()
	resp.Payload = map[string]interface{}{
		"access_token":  access,
		"refresh_token": refresh,
	}
}
