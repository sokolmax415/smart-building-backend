package location

import (
	"client-service/internal/api/types/location/request"
	"client-service/internal/entity"
)

type LocationType string

const (
	LocationTypeBuilding LocationType = "Building"
	LocationTypeFloor    LocationType = "Floor"
	LocationTypeRoom     LocationType = "Room"
)

func isTypeValid(locationType LocationType) bool {
	switch locationType {
	case LocationTypeBuilding, LocationTypeFloor, LocationTypeRoom:
		return true
	}
	return false
}

func ValidateCreateLocationRequest(locationRequest *request.CreateLocationRequest) error {

	if len(locationRequest.LocationName) < 2 || len(locationRequest.LocationName) > 30 {
		return entity.ErrValidateLocationName
	}

	if !isTypeValid(LocationType(locationRequest.LocationType)) {
		return entity.ErrValidateLocationType
	}
	return nil
}

func ValidateUpdateLocationRequest(locationRequest *request.UpdateLocationRequest) error {

	if locationRequest.LocationName != nil {
		if len(*locationRequest.LocationName) < 2 || len(*locationRequest.LocationName) > 30 {
			return entity.ErrValidateLocationName
		}
	}

	if locationRequest.LocationType != nil {
		if !isTypeValid(LocationType(*locationRequest.LocationType)) {
			return entity.ErrValidateLocationType
		}
	}

	return nil
}
