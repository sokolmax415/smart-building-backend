package router

import (
	handler "auth-service/internal/api/http/v1/handlers"
	"auth-service/internal/api/http/v1/middleware"

	"github.com/go-chi/chi/v5"
)

func NewRouter(authHandler *handler.AuthHandler, userHandler *handler.UserHandler, tokenServ middleware.TokenService) chi.Router {
	r := chi.NewRouter()

	r.Mount("/smartbuilding/v1/auth", NewAuthRouter(authHandler))
	r.Mount("/smartbuilding/v1/users", NewUserRouter(userHandler, tokenServ))

	return r
}
