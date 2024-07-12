package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mylordkaz/MsgoChat/services/auth-service/internal/config"
	"github.com/mylordkaz/MsgoChat/services/auth-service/internal/handlers"
	"github.com/mylordkaz/MsgoChat/services/auth-service/internal/repository"
	"github.com/mylordkaz/MsgoChat/services/auth-service/internal/service"
	"github.com/mylordkaz/MsgoChat/services/auth-service/pkg/database"
)

func main(){
	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialsize database
	db, err := database.Newdatabase(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize repository
	authRepo := repository.NewAuthRepository(db.DB)

	// Initialize service
	authService := service.NewAuthService(authRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)

	// Initialize router
	r := mux.NewRouter()

	// Setup routes
	r.HandleFunc("/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/login", authHandler.Login).Methods("POST")

	// start server
	log.Printf("Starting server on %s", cfg.ServerAdress)
    log.Fatal(http.ListenAndServe(cfg.ServerAdress, r))

}