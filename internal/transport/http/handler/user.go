package handler

import (
	"backend-survey-app/internal/service"
	responses "backend-survey-app/pkg/errors"
	httpResponser "backend-survey-app/pkg/http"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var resp httpResponser.Response
	defer resp.Concerter(w)

	var newUser CreateUserDTO

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		resp.Message = responses.ErrBadRequest.Error()
	}

	err = service.CreateUser(newUser.Login, newUser.Email, newUser.Password)
	if err != nil {
		resp.Message = err.Error()
		return
	}
	defer r.Body.Close()

	resp.Message = responses.ErrSuccess.Error()
}
