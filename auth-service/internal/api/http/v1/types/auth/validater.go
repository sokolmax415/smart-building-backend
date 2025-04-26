package auth

import (
	"auth-service/internal/api/http/v1/types/auth/request"
	"auth-service/internal/entity"
)

func ValidateRegisterRequest(registerRequest *request.RegisterRequest) error {
	if len(registerRequest.Login) <= 1 || len(registerRequest.Login) >= 20 {
		return entity.ErrValidateLogin
	}

	if len(registerRequest.Firstname) <= 1 || len(registerRequest.Firstname) >= 15 {
		return entity.ErrValidateFirstName
	}

	if len(registerRequest.Lastname) <= 1 || len(registerRequest.Lastname) >= 20 {
		return entity.ErrValidateLastName
	}

	if len(registerRequest.Password) <= 1 || len(registerRequest.Password) >= 20 {
		return entity.ErrValidatePassword
	}

	return nil
}

func ValidateLoginRequest(loginRequest *request.LoginRequest) error {
	if len(loginRequest.Login) <= 1 || len(loginRequest.Login) >= 20 {
		return entity.ErrValidateLogin
	}

	if len(loginRequest.Password) <= 1 || len(loginRequest.Password) >= 15 {
		return entity.ErrValidateFirstName
	}

	return nil
}
