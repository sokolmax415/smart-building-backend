package usecase

import (
	"client-service/internal/entity"
	"context"
	"log"
	"time"
)

type TelemetryUsecase struct {
	telemRep  TelemetryRepository
	deviceRep DeviceRepository
}

func NewTelemetryUsecase(telemRep TelemetryRepository, deviceRep DeviceRepository) *TelemetryUsecase {
	return &TelemetryUsecase{telemRep: telemRep, deviceRep: deviceRep}
}

func (uc *TelemetryUsecase) GetLatestTelemetry(ctx context.Context, deviceSn string) (*entity.Telemetry, error) {
	exists, err := uc.deviceRep.IsDeviceExist(ctx, deviceSn)

	if err != nil {
		return nil, err
	}

	if !exists {
		log.Printf("MESSAGE IN  GetLatestTelemetry: Device with device_sn=%s not found", deviceSn)
		return nil, entity.ErrDeviceNotFound
	}

	return uc.telemRep.GetLatestTelemetry(ctx, deviceSn)
}

func (uc *TelemetryUsecase) GetTelemetryInRange(ctx context.Context, deviceSn string, from, till time.Time) ([]entity.Telemetry, error) {
	exists, err := uc.deviceRep.IsDeviceExist(ctx, deviceSn)

	if err != nil {
		return nil, err
	}

	if !exists {
		log.Printf("MESSAGE IN GetTelemetryInRange: Device with device_sn=%s not found", deviceSn)
		return nil, entity.ErrDeviceNotFound
	}

	telemetries, err := uc.telemRep.GetTelemetryInRange(ctx, deviceSn, from, till)

	if err != nil {
		return nil, err
	}

	if len(telemetries) == 0 {
		log.Printf("MESSAGE IN GetTelemetryInRange: Telemetry not found for deviceSn=%s", deviceSn)
		return nil, entity.ErrTelemetryNotFound
	}

	return telemetries, nil
}
