package usecase

import (
	"client-service/internal/entity"
	"context"
	"time"
)

type LocationRepository interface {
	CreateNewLocation(context.Context, *entity.Location) (int64, error)
	GetLocationById(context.Context, int64) (*entity.Location, error)
	IsLocationExist(context.Context, int64) (bool, error)
	GetLocationsList(context.Context) ([]entity.Location, error)
	DeleteLocation(context.Context, int64) error
	GetLocationsListWithoutParent(context.Context) ([]entity.Location, error)
	GetLocationChildren(context.Context, int64) ([]entity.Location, error)
	GetPathToLocation(context.Context, int64) ([]entity.Location, error)
	UpdateLocationType(context.Context, string, int64) error
	UpdateLocationName(context.Context, string, int64) error
	UpdateLocationParentId(context.Context, *int64, int64) error
}

type HubRepository interface {
	GetHubBySn(context.Context, string) (*entity.Hub, error)
	IsHubExist(context.Context, string) (bool, error)
	DeleteHub(context.Context, string) error
	GetHubList(context.Context, int64) ([]entity.Hub, error)
}

type TelemetryRepository interface {
	GetLatestTelemetry(context.Context, string) (*entity.Telemetry, error)
	GetTelemetryInRange(context.Context, string, time.Time, time.Time) ([]entity.Telemetry, error)
}

type DeviceRepository interface {
	IsDeviceExist(context.Context, string) (bool, error)
	GetDeviceList(context.Context, string) ([]entity.Device, error)
	GetDeviceCount(context.Context, string) (int64, error)
}
