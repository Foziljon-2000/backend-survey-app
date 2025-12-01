package http

import (
	responses "backend-survey-app/pkg/errors"
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"massage"`
	Payload interface{} `json:"payload"`
}

func (r *Response) Concerter(rw http.ResponseWriter) {
	statusCode := responses.StatusCodes[r.Message]

	if statusCode == 0 {
		statusCode = 500
		r.Message = responses.ErrInternalServer.Error()
	}

	res, _ := json.Marshal(r)

	rw.WriteHeader(statusCode)
	rw.Write(res)
}
