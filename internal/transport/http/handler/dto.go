package handler

type CreateUserDTO struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `Json:"password"`
}
