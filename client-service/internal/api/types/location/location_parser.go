package location

import (
	"client-service/internal/api/types/location/request"
	"client-service/internal/entity"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func ParseCreateLocationRequest(r *http.Request) (*request.CreateLocationRequest, error) {
	var locationRequest request.CreateLocationRequest
	err := json.NewDecoder(r.Body).Decode(&locationRequest)
	if err != nil {
		log.Printf("Error to parse CreateLocationRequest: %v", err)
		return nil, entity.ErrParseLocationRequest
	}

	err = ValidateCreateLocationRequest(&locationRequest)
	if err != nil {
		log.Printf("Error to validate CreateLocationRequest: %v", err)
		return nil, err
	}

	return &locationRequest, nil
}

func ParseUpdateLocationRequest(r *http.Request) (*request.UpdateLocationRequest, error) {
	var locationRequest request.UpdateLocationRequest
	err := json.NewDecoder(r.Body).Decode(&locationRequest)
	if err != nil {
		log.Printf("Error to parse CreateLocationRequest: %v", err)
		return nil, entity.ErrParseLocationRequest
	}

	err = ValidateUpdateLocationRequest(&locationRequest)
	if err != nil {
		log.Printf("Error to validate CreateLocationRequest: %v", err)
		return nil, err
	}

	return &locationRequest, nil
}

func ParseLocationId(r *http.Request) (int64, error) {
	locationIdStr := chi.URLParam(r, "location_id")
	locationId, err := strconv.ParseInt(locationIdStr, 10, 64)
	if err != nil {
		log.Printf("Error to parse location_id: %v", err)
		return 0, err
	}
	return locationId, nil
}
