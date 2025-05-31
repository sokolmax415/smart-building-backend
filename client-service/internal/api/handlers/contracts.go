package handler

import (
	"client-service/internal/entity"
	"context"
	"time"
)

type LocationUsecase interface {
	CreateNewLocation(context.Context, *int64, string, string) (int64, error)
	GetLocation(context.Context, int64) (*entity.Location, error)
	DeleteLocation(context.Context, int64) error
	GetRootLocations(context.Context) ([]entity.Location, error)
	GetChildrenParent(context.Context, int64) (*entity.Location, error)
	GetPathToLocation(context.Context, int64) ([]entity.Location, error)
	GetChildrenList(context.Context, int64) ([]entity.Location, error)
	GetLocationList(context.Context) ([]entity.Location, error)
	UpdateLocation(context.Context, int64, *int64, *string, *string) error
}

type HubUsecase interface {
	GetDevicesForHub(context.Context, string) ([]entity.Device, error)
	GetDeviceCountForHub(context.Context, string) (int64, error)
	GetHubInfo(context.Context, string) (*entity.Hub, error)
	DeleteHub(context.Context, string) error
	GetHubList(context.Context, int64) ([]entity.Hub, error)
}

type TelemetryUsecase interface {
	GetLatestTelemetry(context.Context, string) (*entity.Telemetry, error)
	GetTelemetryInRange(context.Context, string, time.Time, time.Time) ([]entity.Telemetry, error)
}
