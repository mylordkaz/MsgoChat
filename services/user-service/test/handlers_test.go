package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/handlers"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/routes"
)

func TestHandlerUser(t *testing.T) {
    router := mux.NewRouter()
    routes.UserRoutes(router)
    router.HandleFunc("/", handlers.HandleUser).Methods("GET")

    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.Handler(router)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    expected := "Hello, you've reached the user service\n"
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}