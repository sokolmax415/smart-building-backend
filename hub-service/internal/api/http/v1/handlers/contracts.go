package handlers

import (
	"context"
	"hub-service/internal/entity"
)

type HubUsecase interface {
	RegisterHub(context.Context, *entity.Hub) error
	PingHub(context.Context, string, int64) error
	RegisterDevice(context.Context, *entity.Device) error
	SaveTelemetry(context.Context, *entity.Telemetry) error
}
