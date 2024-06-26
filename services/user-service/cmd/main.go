package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/auth"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/database"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/routes"
)

func main () {
	database.ConnectDB()
	
	router := mux.NewRouter()

	routes.UserRoutes(router)
	routes.AuthRoutes(router)
	auth.NewAuth(router)

	router.HandleFunc("/", handleUser)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, you've reached the user service")
}