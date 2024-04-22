package main

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/routes"
)


func TestHandlerUser(t *testing.T){
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	router.HandleFunc("/", handleUser)
}