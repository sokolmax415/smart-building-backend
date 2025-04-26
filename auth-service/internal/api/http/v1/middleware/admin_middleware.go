package middleware

import (
	response "auth-service/internal/api/http/v1/types/generic_response"
	"net/http"
)

func AdminOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		roleVal := r.Context().Value("role")
		role, ok := roleVal.(string)

		if !ok || role != "admin" {
			response.WriteErrorResponse(w, http.StatusForbidden, "Admin access required")
			return
		}

		next.ServeHTTP(w, r)
	})
}
