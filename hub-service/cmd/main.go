package main

import (
	"fmt"
	"hub-service/internal/api/http/v1/handlers"
	"hub-service/internal/api/http/v1/router"
	"hub-service/internal/config"
	"hub-service/internal/repository/postgres"
	"hub-service/internal/usecase"
	"hub-service/pkg/token"
	"log"
	"net/http"
)

// @title Hub API
// @version 1.0
// @description Это документация для сервиса хаба
// @host localhost:8081
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

	hubRep, err := postgres.NewPostgresHubRepository(connectionString)

	if err != nil {
		log.Fatalf("FATAL ERROR connect to DB: %v", err)
	}

	deviceRep, err := postgres.NewPostgresDeviceRepository(connectionString)
	if err != nil {
		log.Fatalf("FATAL Error connect to DB: %v", err)
	}

	tokenSer := token.NewJWTService(cfg.Token.AccessSecret)

	hubUsecase := usecase.NewHubUsecase(hubRep, deviceRep)

	hubHandler := handlers.NewAuthHandler(hubUsecase)

	r := router.NewHubRouter(hubHandler, tokenSer)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	log.Printf("Hub-service is working")
}
