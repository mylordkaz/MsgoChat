package main

import (
	"log"

	"github.com/mylordkaz/MsgoChat/services/auth-service/internal/config"
)

func main(){
	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	
	// Initialsize database

	// Initialize router

	// Initialize handlers

	// Setup routes

	// start server
}