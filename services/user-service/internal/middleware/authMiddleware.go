package middleware

import (
	"net/http"

	"github.com/mylordkaz/MsgoChat/services/user-service/pkg/jwt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)  {
		cookie, err := r.Cookie("Token")
		if err != nil || cookie.Value == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		_, err = jwt.VerifyJWT(cookie.Value)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}