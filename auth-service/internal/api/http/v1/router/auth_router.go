package router

import (
	handler "auth-service/internal/api/http/v1/handlers"

	"github.com/go-chi/chi/v5"
)

func NewAuthRouter(authHandler *handler.AuthHandler) chi.Router {
	r := chi.NewRouter()
	r.Post("/register", authHandler.Register)
	r.Post("/login", authHandler.Login)
	r.Post("/refresh", authHandler.Refresh)

	return r
}
