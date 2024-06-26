package routes

import (
	"github.com/gorilla/mux"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/handlers"
)

func RegisterRoutes(router *mux.Router){
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
}
func AuthRoutes(router *mux.Router){
	router.HandleFunc("/register", handlers.RegisterHandler)
	router.HandleFunc("/login", handlers.LoginHandler)
	router.HandleFunc("/logout", handlers.LogoutHandler)
	router.HandleFunc("/auth/{provider}/callback", handlers.CallbackHandler)
	router.HandleFunc("/auth/{provider}", handlers.AuthHandlers)
}