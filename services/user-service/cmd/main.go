package main

import (
	"log"

	"github.com/mylordkaz/MsgoChat/services/user-service/internal/config"
)


func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
}