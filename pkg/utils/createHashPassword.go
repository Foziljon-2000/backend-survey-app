package utils

import "golang.org/x/crypto/bcrypt"

func CreateHashPassword(pass string) (hashPass string, err error) {
	bytHashPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	hashPass = string(bytHashPass)
	return
}
