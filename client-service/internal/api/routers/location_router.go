package router

import (
	handler "client-service/internal/api/handlers"
	"client-service/internal/api/middleware"

	"github.com/go-chi/chi/v5"
)

func NewLocationRouter(locationHandler *handler.LocationHandler, tokenServ middleware.TokenService) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.AuthMiddleware(tokenServ))

	r.Group(func(r chi.Router) {
		r.Use(middleware.AdminOnlyMiddleware)
		r.Post("/", locationHandler.CreateNewLocation)
		r.Delete("/{location_id}", locationHandler.DeleteLocation)
		r.Patch("/{location_id}", locationHandler.UpdateLocation)
	})

	r.Get("/tree", locationHandler.GetLocationTree)
	r.Get("/{location_id}", locationHandler.GetLocation)
	r.Get("/{location_id}/details", locationHandler.GetFullLocationInfo)
	r.Get("/root", locationHandler.GetRootLocations)
	r.Get("/{location_id}/children", locationHandler.GetChildrenList)
	r.Get("/{location_id}/parent", locationHandler.GetLocationParent)
	r.Get("/{location_id}/path", locationHandler.GetLocationPath)
	return r
}
