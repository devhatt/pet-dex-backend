package middlewares

import (
	"context"
	"net/http"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/pkg/encoder"
	"strings"
	"time"
)

type ContextKey string

func AuthMiddleware(next http.Handler) http.Handler {
	const UserClaimsContextKey ContextKey = "userClaims"

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := encoder.NewEncoderAdapter(config.GetEnvConfig().JWT_SECRET)

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		headerSplited := strings.Split(authHeader, " ")
		if len(headerSplited) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		bearerToken := headerSplited[1]
		if bearerToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		userclaims := encoder.ParseAccessToken(bearerToken)
		if userclaims.ExpiresAt != 0 && userclaims.ExpiresAt < time.Now().Unix() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		context := context.WithValue(r.Context(), UserClaimsContextKey, userclaims)
		next.ServeHTTP(w, r.WithContext(context))
	})
}
