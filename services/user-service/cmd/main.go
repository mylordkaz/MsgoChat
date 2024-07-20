package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/config"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/handlers"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/repository"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/service"
	"github.com/mylordkaz/MsgoChat/services/user-service/pkg/database"
)


func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo)
	handler := handlers.NewUserHandler(svc)

	r := mux.NewRouter()
	handler.RegisterRoutes(r)

	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}