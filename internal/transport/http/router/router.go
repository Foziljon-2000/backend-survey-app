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

	secured := r.PathPrefix("/user").Subrouter()
	secured.Use(mw.Auth)
	secured.HandleFunc("/me", handler.GetUser).Methods("GET")
	secured.HandleFunc("/logout", handler.Logout).Methods("POST")

	quiz := r.PathPrefix("/quiz").Subrouter()
	quiz.Use(mw.Auth)

	quiz.HandleFunc("/create", handler.CreateQuiz).Methods("POST")
	quiz.HandleFunc("/get", handler.GetQuiz).Methods("GET")
	quiz.HandleFunc("/update", handler.UpdateQuiz).Methods("PUT")
	quiz.HandleFunc("/delete", handler.DeleteQuiz).Methods("DELETE")

	return r
}
