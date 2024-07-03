package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/routes"
)


func TestHandlerUser(t *testing.T){
	router := mux.NewRouter()
	routes.UserRoutes(router)
	router.HandleFunc("/", handleUser)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil{
		t.Fatal(err)
	}

	// create a ResponseRecorder to record the response 
	rr := httptest.NewRecorder()
	handler := http.Handler(router)

	// dispach request to handler
	handler.ServeHTTP(rr, req)

	//check status code is what expected
	if status := rr.Code; status != http.StatusOK{
		t.Errorf("handler returned wrong status code : got %v want %v", status, http.StatusOK)
	}

	// check response body
	expected := "Hello, you've reached the user service\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body : got %v want %v", rr.Body.String(), expected)
	}
}