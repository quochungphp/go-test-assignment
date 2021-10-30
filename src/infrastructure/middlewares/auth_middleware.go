package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/quochungphp/go-test-assignment/src/pkgs/token"
)

// AuthMiddleware ...
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := token.ValidToken(r)
		if err != nil {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Unauthorized")
			return
		}
		// TODO: implement user info
		r.Header.Set("user", "123")
		next.ServeHTTP(w, r)
	})
}
