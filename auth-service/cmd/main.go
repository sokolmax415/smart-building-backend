package main

import (
	handler "auth-service/internal/api/http/v1/handlers"
	"auth-service/internal/api/http/v1/router"
	"auth-service/internal/config"
	"auth-service/internal/repository/postgres"
	"auth-service/internal/usecase"
	"auth-service/pkg/hash"
	"auth-service/pkg/token"
	"fmt"
	"log"
	"net/http"
)

// @title Auth API
// @version 1.0
// @description Это документация для сервиса аутентификации
// @host localhost:8080
// @BasePath /smartbuilding/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" accessToken.
func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
	)

	userRep, err := postgres.NewPostgresRepository(connectionString)
	fmt.Println(connectionString)
	if err != nil {
		log.Fatalf("Error connect to DB: %v", err)
	}

	roleRep, err := postgres.NewRoleRepository(connectionString)
	if err != nil {
		log.Fatalf("Error connect to DB: %w", err)
	}

	tokenSer := token.NewJWTService(cfg.Token.AccessSecret, cfg.Token.RefreshSecret)
	hashingSer := hash.NewBcryptService()

	authUsecase := usecase.NewAuthUsecase(userRep, roleRep, tokenSer, hashingSer)
	userUsecase := usecase.NewUserUsecase(userRep, roleRep, hashingSer)

	authHandler := handler.NewAuthHandler(authUsecase)
	userHandler := handler.NewUserHandler(userUsecase)
	r := router.NewRouter(authHandler, userHandler, tokenSer)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
