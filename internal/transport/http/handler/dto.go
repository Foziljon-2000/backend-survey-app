package handler

type CreateUserDTO struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `Json:"password"`
}

type LoginRespons struct {
	Email    string `Json:"email"`
	Password string `json:"password"`
}
