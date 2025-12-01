package router

import (
	"backend-survey-app/internal/transport/http/handler"

	"github.com/gorilla/mux"
)

func NewRouterCompl() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/create-user", handler.CreateUser).Methods("POST")
	router.HandleFunc("/user/me", handler.GetUser).Methods("GET")

	return router
}
