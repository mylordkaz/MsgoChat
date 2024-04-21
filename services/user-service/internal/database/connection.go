package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func ConnectDB() *sql.DB {

	dbHost := os.Getenv("BD_HOST")
	dbPort := os.Getenv("BD_PORT")
	dbUser := os.Getenv("BD_USER")
	dbPass := os.Getenv("BD_PASS")
	dbName := os.Getenv("BD_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s name=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

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
