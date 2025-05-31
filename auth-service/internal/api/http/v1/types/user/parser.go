package user

import (
	"auth-service/internal/api/http/v1/types/user/request"
	"auth-service/internal/entity"
	"encoding/json"
	"log"
	"net/http"
)

func ParseUserRequest(r *http.Request) (*request.UserRequest, error) {
	var userRequest request.UserRequest

	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		log.Printf("Error to parse UserRequest: %v", err)
		return nil, entity.ErrParseUserRequest
	}

	err = ValidateUserRequest(&userRequest)

	if err != nil {
		log.Printf("Error to validate UserRequest: %v", err)
		return nil, err
	}

	return &userRequest, nil
}

func ParseNameRequest(r *http.Request) (*request.NameRequest, error) {
	var nameRequest request.NameRequest

	err := json.NewDecoder(r.Body).Decode(&nameRequest)
	if err != nil {
		log.Printf("Error to parse NameRequest: %v", err)
		return nil, entity.ErrParseNameRequest
	}

	err = ValidateNameRequest(&nameRequest)

	if err != nil {
		log.Printf("Error to validate NameRequest: %v", err)
		return nil, err
	}

	return &nameRequest, nil
}

func ParseChangeRoleRequest(r *http.Request) (*request.ChangeRoleRequest, error) {
	var changeRoleRequest request.ChangeRoleRequest
	err := json.NewDecoder(r.Body).Decode(&changeRoleRequest)
	if err != nil {
		log.Printf("Error to parse RoleRequest: %v", err)
		return nil, entity.ErrParseChangeRoleRequest
	}
	return &changeRoleRequest, nil
}
