package main

import (
	"log"

	"github.com/mylordkaz/MsgoChat/services/auth-service/internal/config"
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
	
	// Initialize router

	// Initialize handlers

	// Setup routes

	// start server
}