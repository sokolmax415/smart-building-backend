package router

import (
	handler "client-service/internal/api/handlers"
	"client-service/internal/api/middleware"

	"github.com/go-chi/chi/v5"
)

func NewTelemetryRouter(telemetryHandler *handler.TelemetryHandler, tokenServ middleware.TokenService) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.AuthMiddleware(tokenServ))

	r.Get("/{device_sn}/telemetry/latest", telemetryHandler.GetLatestTelemetry)
	r.Get("/{device_sn}/telemetry", telemetryHandler.GetTelemetryInRange)
	return r
}
