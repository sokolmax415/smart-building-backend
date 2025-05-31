package usecase

import (
	"client-service/internal/entity"
	"context"
	"log"
)

type HubUsecase struct {
	hubRep    HubRepository
	deviceRep DeviceRepository
	locRep    LocationRepository
}

func NewHubUsecase(hubRep HubRepository, deviceRep DeviceRepository, locRep LocationRepository) *HubUsecase {
	return &HubUsecase{hubRep: hubRep, deviceRep: deviceRep, locRep: locRep}
}

func checkHubExistence(err error, exists bool, hubSn string) error {
	if err != nil {
		return err
	}

	if !exists {
		log.Printf("MESSAGE IN HubUsecase: Hub with hubSn=%s not found", hubSn)
		return entity.ErrHubNotFound
	}
	return nil
}

func (uc *HubUsecase) GetDevicesForHub(ctx context.Context, hubSn string) ([]entity.Device, error) {
	exists, err := uc.hubRep.IsHubExist(ctx, hubSn)
	err = checkHubExistence(err, exists, hubSn)
	if err != nil {
		return nil, err
	}

	return uc.deviceRep.GetDeviceList(ctx, hubSn)
}

func (uc *HubUsecase) DeleteHub(ctx context.Context, hubSn string) error {
	exists, err := uc.hubRep.IsHubExist(ctx, hubSn)

	if err := checkHubExistence(err, exists, hubSn); err != nil {
		return err
	}

	return uc.hubRep.DeleteHub(ctx, hubSn)
}

func (uc *HubUsecase) GetHubInfo(ctx context.Context, hubSn string) (*entity.Hub, error) {
	exists, err := uc.hubRep.IsHubExist(ctx, hubSn)
	if err := checkHubExistence(err, exists, hubSn); err != nil {
		return nil, err
	}

	return uc.hubRep.GetHubBySn(ctx, hubSn)
}

func (uc *HubUsecase) GetDeviceCountForHub(ctx context.Context, hubSn string) (int64, error) {
	exists, err := uc.hubRep.IsHubExist(ctx, hubSn)
	if err := checkHubExistence(err, exists, hubSn); err != nil {
		return 0, err
	}

	return uc.deviceRep.GetDeviceCount(ctx, hubSn)
}

func (uc *HubUsecase) GetHubList(ctx context.Context, locationId int64) ([]entity.Hub, error) {
	exist, err := uc.locRep.IsLocationExist(ctx, locationId)
	if err := checkLocationExistence(exist, err, locationId); err != nil {
		return nil, err
	}
	return uc.hubRep.GetHubList(ctx, locationId)
}
