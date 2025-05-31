package middleware

import (
	response "client-service/internal/api/types/generic_response"
	"log"
	"net/http"
)

func AdminOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		roleVal := r.Context().Value("role")
		role, ok := roleVal.(string)

		if !ok || role != "admin" {
			log.Printf("ERROR ACCESS DENIED: Admin access only")
			response.WriteErrorResponse(w, http.StatusForbidden, "Admin access required")
			return
		}

		next.ServeHTTP(w, r)
	})
}
