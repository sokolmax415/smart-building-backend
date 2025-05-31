package middleware

import (
	"context"
	resp "hub-service/internal/api/http/v1/types/generic_repsponse"
	"net/http"
	"strings"
)

type TokenService interface {
	ParseAccessToken(string) (int64, string, error)
}

func AuthMiddleware(verifier TokenService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				resp.WriteErrorResponse(w, http.StatusUnauthorized, "Missing or invalid Authorization header")
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			_, role, err := verifier.ParseAccessToken(token)
			if err != nil {
				resp.WriteErrorResponse(w, http.StatusUnauthorized, "Invalid token")
				return
			}

			ctx := context.WithValue(r.Context(), "role", role)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
