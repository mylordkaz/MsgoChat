package routes

import (
	"github.com/gorilla/mux"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/handlers"
)

func RegisterRoutes(router *mux.Router){

	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")

}