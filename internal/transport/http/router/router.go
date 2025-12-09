package router

import (
	"backend-survey-app/internal/transport/http/handler"
	"backend-survey-app/internal/transport/http/mw"

	"github.com/gorilla/mux"
)

func NewRouterCompl() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/create-user", handler.CreateUser).Methods("POST")
	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.HandleFunc("/logout", handler.Logout).Methods("POST")

	// защищённые роуты
	secured := r.PathPrefix("/user").Subrouter()
	secured.Use(mw.Auth)
	secured.HandleFunc("/me", handler.GetUser).Methods("GET")

	return r
}
