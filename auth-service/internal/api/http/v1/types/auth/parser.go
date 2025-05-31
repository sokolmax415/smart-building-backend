package auth

import (
	"auth-service/internal/api/http/v1/types/auth/request"
	"auth-service/internal/entity"
	"encoding/json"
	"log"
	"net/http"
)

func ParseLoginRequest(r *http.Request) (*request.LoginRequest, error) {
	var loginRequest request.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		log.Printf("Error to parse LoginRequest: %v", err)
		return nil, entity.ErrParseLoginRequest
	}

	err = ValidateLoginRequest(&loginRequest)
	if err != nil {
		log.Printf("Error to validate LoginRequest: %v", err)
		return nil, err
	}

	return &loginRequest, nil
}

func ParseRegisterRequest(r *http.Request) (*request.RegisterRequest, error) {
	var registerRequest request.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		log.Printf("Error to parse RegisterRequest: %v", err)
		return nil, entity.ErrParseRegisterRequest
	}

	err = ValidateRegisterRequest(&registerRequest)
	if err != nil {
		log.Printf("Error to validate RegisterRequest: %v", err)
		return nil, err
	}

	return &registerRequest, nil
}

func ParseRefreshRequest(r *http.Request) (*request.RefreshRequest, error) {
	var refreshRequest request.RefreshRequest
	err := json.NewDecoder(r.Body).Decode(&refreshRequest)
	if err != nil {
		log.Printf("Error to parse RefreshRequest: %v", err)
		return nil, entity.ErrParseRefreshRequest
	}

	return &refreshRequest, nil
}
