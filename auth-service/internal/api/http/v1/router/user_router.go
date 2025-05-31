package router

import (
	handler "auth-service/internal/api/http/v1/handlers"
	"auth-service/internal/api/http/v1/middleware"

	"github.com/go-chi/chi/v5"
)

func NewUserRouter(userHandler *handler.UserHandler, tokenServ middleware.TokenService) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.AuthMiddleware(tokenServ))
	r.Use(middleware.AdminOnlyMiddleware)

	r.Get("/", userHandler.GetAllUsers)
	r.Get("/{login}", userHandler.GetUserInfo)
	r.Post("/", userHandler.CreateNewUser)
	r.Put("/{login}", userHandler.ChangeUserRole)
	r.Delete("/{login}", userHandler.DeleteUser)
	r.Put("/name/{login}", userHandler.ChangeUserName)

	return r
}
