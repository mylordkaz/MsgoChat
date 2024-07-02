package handlers

import (
	"fmt"
	"net/http"

	"github.com/mylordkaz/MsgoChat/services/user-service/contextkeys"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
    email, ok := r.Context().Value(contextkeys.UserContextKey).(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
    w.Write([]byte(fmt.Sprintf("Welcome to the Dashboard, %s!", email)))
}