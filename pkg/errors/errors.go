package errors

import "errors"

var (
	ErrInternalServer          = errors.New("Внутреняя ошибка сервера")
	ErrBadRequest              = errors.New("Неправильный запрос")
	ErrUserDoesNotExist        = errors.New("Такого пользователя не существует")
	ErrInvalidEmail            = errors.New("неверный формат электронной почты")
	ErrSuccess                 = errors.New("Успешно")
	ErrThisEmailIsAlreadyTaken = errors.New("Этот email уже занят")
)

var StatusCodes map[string]int = map[string]int{
	//200 ...
	ErrSuccess.Error(): 200,
	//400 ...
	ErrUserDoesNotExist.Error():        404,
	ErrBadRequest.Error():              400,
	ErrInvalidEmail.Error():            400,
	ErrThisEmailIsAlreadyTaken.Error(): 409,

	//500 ...
	ErrInternalServer.Error(): 500,
}
