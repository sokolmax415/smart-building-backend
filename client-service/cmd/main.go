package main

import (
	handler "client-service/internal/api/handlers"
	router "client-service/internal/api/routers"
	"client-service/internal/config"
	"client-service/internal/repository/postgres"
	"client-service/internal/usecase"
	token "client-service/pkg"
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

	telRep, err := postgres.NewPostgresTelemetryRepository(connectionString)
	if err != nil {
		log.Fatalf("FATAL ERROR connect to DB: %v", err)
	}

	deviceRep, err := postgres.NewPostgresDeviceRepository(connectionString)
	if err != nil {
		log.Fatalf("FATAL ERROR connect to DB: %v", err)
	}

	hubRep, err := postgres.NewPostgresHubRepository(connectionString)
	if err != nil {
		log.Fatalf("FATAL ERROR connect to DB: %v", err)
	}

	locRep, err := postgres.NewPostgresLocationRepository(connectionString)
	if err != nil {
		log.Fatalf("FATAL ERROR connect to DB: %v", err)
	}

	tokenSer := token.NewJWTService(cfg.Token.AccessSecret)

	telemUsecase := usecase.NewTelemetryUsecase(telRep, deviceRep)
	hubUsecase := usecase.NewHubUsecase(hubRep, deviceRep, locRep)
	locationUsecase := usecase.NewLocationUsecase(locRep)

	telemHandler := handler.NewTelemetryHandler(telemUsecase)
	hubHandler := handler.NewHubHandler(hubUsecase)
	locationHandler := handler.NewLocationHandler(locationUsecase, hubUsecase)
	r := router.NewRouter(telemHandler, hubHandler, locationHandler, tokenSer)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	log.Printf("Client-service is working")
}
