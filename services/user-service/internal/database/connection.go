package database

import (
	"database/sql"
	"fmt"
	"log"

	
)



var db *sql.DB

func InitDB(dataSourceName string) {
	var err error

	db, err = sql.Open("postgres", dataSourceName)
	if err != nil{
		log.Panic(err)
	}
	
	fmt.Println("Successfully connected to database")
}
func GetDB() *sql.DB {
	return db
}