package response

import (
	"encoding/json"
	"net/http"
)

// swagger:model errorResponse
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, errMessage string) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	errorReponse := ErrorResponse{Status: "error", Message: errMessage}
	err := json.NewEncoder(w).Encode(errorReponse)
	if err != nil {
		//for slog
	}
}
