package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/auth"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/handlers"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/routes"
)

func TestRouterSetup(t *testing.T) {
    router := mux.NewRouter()
    
    routes.UserRoutes(router)
    routes.AuthRoutes(router)
    auth.NewAuth(router)
    router.HandleFunc("/", handlers.HandleUser)

    // test cases
    testCases := []struct {
        name            string
        path            string
        method          string
        expectedStatus  int
    }{
        {"Root Path", "/", "GET", http.StatusOK},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T){
            req, err := http.NewRequest(tc.method, tc.path, nil)
            if err != nil {
                t.Fatal(err)
            }

            rr := httptest.NewRecorder()
            router.ServeHTTP(rr, req)

            if status := rr.Code; status != tc.expectedStatus {
                t.Errorf("handler returned wrong status code: got %v want %v", status, tc.expectedStatus)
            }
        })
    }
}