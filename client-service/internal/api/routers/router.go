package router

import (
	handler "client-service/internal/api/handlers"
	"client-service/internal/api/middleware"

	"github.com/go-chi/chi/v5"
)

var baseURL = "/smartbuilding/v1/client/"

func NewRouter(telemetryHandler *handler.TelemetryHandler, hubHandler *handler.HubHandler, locationHandler *handler.LocationHandler, tokenServ middleware.TokenService) chi.Router {
	r := chi.NewRouter()
	r.Mount(baseURL+"devices", NewTelemetryRouter(telemetryHandler, tokenServ))
	r.Mount(baseURL+"hubs", NewHubRouter(hubHandler, tokenServ))
	r.Mount(baseURL+"locations", NewLocationRouter(locationHandler, tokenServ))
	return r
}
