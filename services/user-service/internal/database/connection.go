package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func ConnectDB() *sql.DB {
	err := godotenv.Load("../.env")
	if err != nil{
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("postgresql://%s@%s:%s/%s?sslmode=disable", 
	dbUser, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil{
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.Ping()
    if err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }

    log.Println("Successfully connected to database")
    return db

}
