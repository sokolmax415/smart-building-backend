package user

import (
	"auth-service/internal/api/http/v1/types/user/request"
	"auth-service/internal/entity"
)

func ValidateUserRequest(userRequest *request.UserRequest) error {
	if len(userRequest.Login) <= 1 || len(userRequest.Login) >= 20 {
		return entity.ErrValidateLogin
	}

	if len(userRequest.Firstname) <= 1 || len(userRequest.Firstname) >= 15 {
		return entity.ErrValidateFirstName
	}

	if len(userRequest.Lastname) <= 1 || len(userRequest.Lastname) >= 20 {
		return entity.ErrValidateLastName
	}

	if len(userRequest.Password) <= 1 || len(userRequest.Password) >= 20 {
		return entity.ErrValidatePassword
	}

	return nil
}

func ValidateNameRequest(nameRequest *request.NameRequest) error {

	if len(nameRequest.Firstname) <= 1 || len(nameRequest.Firstname) >= 15 {
		return entity.ErrValidateFirstName
	}

	if len(nameRequest.Lastname) <= 1 || len(nameRequest.Lastname) >= 20 {
		return entity.ErrValidateLastName
	}

	return nil
}
