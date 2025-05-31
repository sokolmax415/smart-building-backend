package usecase

import (
	"context"
	"hub-service/internal/entity"
)

type HubRepository interface {
	CreateOrUpdateHub(context.Context, *entity.Hub) error
	IsHubExist(context.Context, string) (bool, error)
	UpdateHubUptime(context.Context, string, int64) error
	IsLocationExist(context.Context, int64) (bool, error)
}

type DeviceRepository interface {
	RegisterDevice(context.Context, *entity.Device) error
	SaveTelemetry(context.Context, *entity.Telemetry) error
	IsDeviceExist(context.Context, string) (bool, error)
}
