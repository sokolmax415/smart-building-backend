package router

import (
	"hub-service/internal/api/http/v1/handlers"
	"hub-service/internal/api/http/v1/middleware"

	"github.com/go-chi/chi/v5"
)

func NewHubRouter(hubHandler *handlers.HubHandler, tokenServ middleware.TokenService) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.AuthMiddleware(tokenServ))
	r.Post("/smartbuilding/v1/hubs/register", hubHandler.RegisterHub)
	r.Post("/smartbuilding/v1/hubs/ping", hubHandler.PingHub)
	r.Post("/smartbuilding/v1/hubs/devices/register", hubHandler.RegisterDevice)
	r.Post("/smartbuilding/v1/hubs/devices/telemetry", hubHandler.SaveTelemetry)
	return r
}
