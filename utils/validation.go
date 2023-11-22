package utils

import (
	"easygoals/model"
)

func ValidateUserData(myUSer *model.User) string {

	if myUSer.Name == "" {
		return "Name is Empty"
	}
	if myUSer.Email == "" {
		return "Email not be null"
	}
	if myUSer.Password == "" {
		return "Password not be null"
	}

	return ""
}
