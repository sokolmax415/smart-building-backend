package resp

import (
	"encoding/json"
	"log"
	"net/http"
)

type SuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func WriteSuccessResponse(w http.ResponseWriter, statusCode int, successMessage string) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	successResponse := SuccessResponse{Status: "success", Message: successMessage}
	err := json.NewEncoder(w).Encode(successResponse)

	if err != nil {
		log.Printf("Err to write success response with staus=%d, message=%s", statusCode, successMessage)
	}
}
