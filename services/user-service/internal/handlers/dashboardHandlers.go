package handlers

import (
	"fmt"
	"net/http"

	"github.com/mylordkaz/MsgoChat/services/user-service/contextkeys"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
    email := r.Context().Value(contextkeys.UserContextKey).(string)
    w.Write([]byte(fmt.Sprintf("Welcome to the Dashboard, %s!", email)))
}