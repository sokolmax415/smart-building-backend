package usecase

import (
	"context"
	"hub-service/internal/entity"
	"log"
)

type HubUsecase struct {
	hubRep    HubRepository
	deviceRep DeviceRepository
}

func NewHubUsecase(hubRep HubRepository, deviceRep DeviceRepository) *HubUsecase {
	return &HubUsecase{hubRep: hubRep, deviceRep: deviceRep}
}

func (uc *HubUsecase) RegisterHub(ctx context.Context, hub *entity.Hub) error {
	exists, err := uc.hubRep.IsLocationExist(ctx, hub.LocationId)

	if err != nil {
		return err
	}

	if !exists {
		log.Printf("MESSAGE in RegisterHub: location with location_id=%d not exist", hub.LocationId)
		return entity.ErrLocationNotFound
	}
	return uc.hubRep.CreateOrUpdateHub(ctx, hub)
}

func (uc *HubUsecase) PingHub(ctx context.Context, sn string, uptime int64) error {
	exists, err := uc.hubRep.IsHubExist(ctx, sn)
	if err != nil {
		return err
	}

	if !exists {
		log.Printf("MESSAGE in PingHub: hub with hub_sn=%s not exist", sn)
		return entity.ErrHubNotFound
	}

	err = uc.hubRep.UpdateHubUptime(ctx, sn, uptime)
	if err != nil {
		return err
	}

	return nil
}

func (uc *HubUsecase) RegisterDevice(ctx context.Context, device *entity.Device) error {
	exists, err := uc.hubRep.IsHubExist(ctx, device.HubSn)

	if err != nil {
		return err
	}

	if !exists {
		log.Printf("MESSAGE in RegisterDevice: hub with hub_sn=%s not exist", device.HubSn)
		return entity.ErrHubNotFound
	}

	return uc.deviceRep.RegisterDevice(ctx, device)
}

func (uc *HubUsecase) SaveTelemetry(ctx context.Context, telemetry *entity.Telemetry) error {
	exists, err := uc.deviceRep.IsDeviceExist(ctx, telemetry.DeviceSn)
	if err != nil {
		return err
	}

	if !exists {
		log.Printf("MESSAGE in SaveTelemetry: device with device_sn=%s not exist", telemetry.DeviceSn)
		return entity.ErrDeviceNotFound
	}

	err = uc.deviceRep.SaveTelemetry(ctx, telemetry)
	if err != nil {
		return err
	}

	return nil
}
