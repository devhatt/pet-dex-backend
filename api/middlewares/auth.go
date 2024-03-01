package middlewares

import (
	"net/http"
	"strings"
	"time"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(401)
			return
		}
		headerSplited := strings.Split(authHeader, " ")
		if len(headerSplited) != 2 {
			w.WriteHeader(401)
			return
		}
		bearerToken := headerSplited[1]
		if bearerToken == "" {
			w.WriteHeader(401)
			return
		}
		userclaims := ParseAccessToken(bearerToken)
		if userclaims.ExpiresAt != 0 && userclaims.ExpiresAt < time.Now().Unix() {
			w.WriteHeader(401)
			return
		}
		next.ServeHTTP(w, r)
	})
}
