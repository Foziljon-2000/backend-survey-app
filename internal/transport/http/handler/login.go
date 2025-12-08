package handler

import (
	responses "backend-survey-app/pkg/errors"
	httpResponser "backend-survey-app/pkg/http"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var resp httpResponser.Response

	var user LoginRespons

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		resp.Message = responses.ErrBadRequest.Error()
		return
	}

	defer r.Body.Close()

	resp.Message = responses.ErrSuccess.Error()

}
