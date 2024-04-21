package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mylordkaz/MsgoChat/services/user-service/internal/database"
)



func main () {

	dbHost := os.Getenv("BD_HOST")
	dbPort := os.Getenv("BD_PORT")
	dbUser := os.Getenv("BD_USER")
	dbPass := os.Getenv("BD_PASS")
	dbName := os.Getenv("BD_NAME")
	
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s name=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	database.InitDB(dataSourceName)

	http.HandleFunc("/", handleUser)
	fmt.Println("Starting server at port 8080")
	http.ListenAndServe(":8080", nil)
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, you've reached the user service")
}