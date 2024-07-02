package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/handlers"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/middleware"
)

func UserRoutes(router *mux.Router){
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
	router.Handle("/dashboard", middleware.AuthMiddleware(http.HandlerFunc(handlers.DashboardHandler))).Methods("GET")
}