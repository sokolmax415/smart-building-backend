package usecase

import (
	"client-service/internal/entity"
	"context"
	"log"
)

type LocationUsecase struct {
	locRep LocationRepository
}

func NewLocationUsecase(locRep LocationRepository) *LocationUsecase {
	return &LocationUsecase{locRep: locRep}
}

func checkLocationExistence(exist bool, err error, locationId int64) error {
	if err != nil {
		return err
	}

	if !exist {
		log.Printf("MESSAGE IN LocationUsecase: Location with location_id=%d not found", locationId)
		return entity.ErrLocationNotFound
	}
	return nil
}

func (uc *LocationUsecase) CreateNewLocation(ctx context.Context, parentId *int64, locationType string, locationName string) (int64, error) {

	if parentId != nil {
		exist, err := uc.locRep.IsLocationExist(ctx, *parentId)
		if err := checkLocationExistence(exist, err, *parentId); err != nil {
			return 0, err
		}
	}
	return uc.locRep.CreateNewLocation(ctx, &entity.Location{ParentId: parentId, LocationType: locationType, LocationName: locationName})
}

func (uc *LocationUsecase) DeleteLocation(ctx context.Context, locationId int64) error {
	exist, err := uc.locRep.IsLocationExist(ctx, locationId)
	if err := checkLocationExistence(exist, err, locationId); err != nil {
		return err
	}
	return uc.locRep.DeleteLocation(ctx, locationId)
}

func (uc *LocationUsecase) GetLocation(ctx context.Context, locationId int64) (*entity.Location, error) {
	exist, err := uc.locRep.IsLocationExist(ctx, locationId)

	if err := checkLocationExistence(exist, err, locationId); err != nil {
		return nil, err
	}
	location, err := uc.locRep.GetLocationById(ctx, locationId)

	if err != nil {
		return nil, err
	}

	return location, nil
}

func (uc *LocationUsecase) GetRootLocations(ctx context.Context) ([]entity.Location, error) {
	return uc.locRep.GetLocationsListWithoutParent(ctx)
}

func (uc *LocationUsecase) GetChildrenParent(ctx context.Context, locationId int64) (*entity.Location, error) {
	exist, err := uc.locRep.IsLocationExist(ctx, locationId)

	if err := checkLocationExistence(exist, err, locationId); err != nil {
		return nil, err
	}

	location, err := uc.locRep.GetLocationById(ctx, locationId)

	if err != nil {
		return nil, err
	}

	parentId := location.ParentId

	if parentId == nil {
		return nil, entity.ErrParentNull
	}

	return uc.locRep.GetLocationById(ctx, *parentId)
}

func (uc *LocationUsecase) GetPathToLocation(ctx context.Context, locationId int64) ([]entity.Location, error) {
	exist, err := uc.locRep.IsLocationExist(ctx, locationId)
	if err := checkLocationExistence(exist, err, locationId); err != nil {
		return nil, err
	}

	return uc.locRep.GetPathToLocation(ctx, locationId)
}

func (uc *LocationUsecase) GetChildrenList(ctx context.Context, locationId int64) ([]entity.Location, error) {
	exist, err := uc.locRep.IsLocationExist(ctx, locationId)
	if err := checkLocationExistence(exist, err, locationId); err != nil {
		return nil, err
	}

	return uc.locRep.GetLocationChildren(ctx, locationId)
}

func (uc *LocationUsecase) GetLocationList(ctx context.Context) ([]entity.Location, error) {
	return uc.locRep.GetLocationsList(ctx)
}

func (uc *LocationUsecase) UpdateLocation(ctx context.Context, locationId int64, parentId *int64, locationType *string, locationName *string) error {
	exist, err := uc.locRep.IsLocationExist(ctx, locationId)
	if err := checkLocationExistence(exist, err, locationId); err != nil {
		return err
	}

	err = uc.locRep.UpdateLocationParentId(ctx, parentId, locationId)

	if err != nil {
		return err
	}

	if locationType != nil {
		err = uc.locRep.UpdateLocationType(ctx, *locationType, locationId)
		if err != nil {
			return err
		}
	}

	if locationName != nil {
		err = uc.locRep.UpdateLocationName(ctx, *locationName, locationId)
		if err != nil {
			return err
		}
	}
	return nil
}
