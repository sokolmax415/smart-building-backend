package auth

import (
	"auth-service/internal/api/http/v1/types/auth/request"
	"auth-service/internal/entity"
	"encoding/json"
	"net/http"
)

func ParseLoginRequest(r *http.Request) (*request.LoginRequest, error) {
	var loginRequest request.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		return nil, entity.ErrParseLoginRequest
	}

	err = ValidateLoginRequest(&loginRequest)
	if err != nil {
		return nil, err
	}

	return &loginRequest, nil
}

func ParseRegisterRequest(r *http.Request) (*request.RegisterRequest, error) {
	var registerRequest request.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		return nil, entity.ErrParseRegisterRequest
	}

	err = ValidateRegisterRequest(&registerRequest)
	if err != nil {
		return nil, err
	}

	return &registerRequest, nil
}

func ParseRefreshRequest(r *http.Request) (*request.RefreshRequest, error) {
	var refreshRequest request.RefreshRequest
	err := json.NewDecoder(r.Body).Decode(&refreshRequest)
	if err != nil {
		return nil, entity.ErrParseRefreshRequest
	}

	return &refreshRequest, nil
}
