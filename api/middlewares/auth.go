package middlewares

import (
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		w.Write([]byte(fmt.Sprintf("%v", claims["user_id"])))
		next.ServeHTTP(w, r)
	})
}
