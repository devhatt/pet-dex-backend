package middlewares

import (
	"net/http"
	"pet-dex-backend/v2/infra/config"
	"strings"
	"time"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := NewEncoderAdapter(config.GetEnvConfig().JWT_SECRET)

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
		userclaims := encoder.ParseAccessToken(bearerToken)
		if userclaims.ExpiresAt != 0 && userclaims.ExpiresAt < time.Now().Unix() {
			w.WriteHeader(401)
			return
		}
		next.ServeHTTP(w, r)
	})
}
