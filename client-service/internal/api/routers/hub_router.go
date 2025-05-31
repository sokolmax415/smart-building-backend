package router

import (
	handler "client-service/internal/api/handlers"
	"client-service/internal/api/middleware"

	"github.com/go-chi/chi/v5"
)

func NewHubRouter(hubHandler *handler.HubHandler, tokenServ middleware.TokenService) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.AuthMiddleware(tokenServ))
	r.Get("/{hub_sn}", hubHandler.GetHubInfo)
	r.Get("/{hub_sn}/devices", hubHandler.GetDevicesForHub)
	r.Get("/{hub_sn}/devices/count", hubHandler.GetDeviceCountForHub)

	r.Group(func(r chi.Router) {
		r.Use(middleware.AdminOnlyMiddleware)
		r.Delete("/{hub_sn}", hubHandler.DeleteHub)
	})
	return r
}
