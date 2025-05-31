package handler

import (
	response "client-service/internal/api/types/generic_response"
	"client-service/internal/entity"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type TelemetryHandler struct {
	telemUsecase TelemetryUsecase
}

func NewTelemetryHandler(telemUsecase TelemetryUsecase) *TelemetryHandler {
	return &TelemetryHandler{telemUsecase: telemUsecase}
}

func writeJson(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (h *TelemetryHandler) GetLatestTelemetry(w http.ResponseWriter, r *http.Request) {
	deviceSn := chi.URLParam(r, "device_sn")

	telemetry, err := h.telemUsecase.GetLatestTelemetry(r.Context(), deviceSn)

	if err != nil {
		if errors.Is(err, entity.ErrDeviceNotFound) {
			response.WriteErrorResponse(w, http.StatusNotFound, "Device not found")
			return
		}
		if errors.Is(err, entity.ErrTelemetryNotFound) {
			response.WriteErrorResponse(w, http.StatusNoContent, "No telemetry for this device")
			return
		}
		response.WriteErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	writeJson(w, http.StatusOK, telemetry)
}

func (h *TelemetryHandler) GetTelemetryInRange(w http.ResponseWriter, r *http.Request) {
	deviceSn := chi.URLParam(r, "device_sn")
	fromStr := r.URL.Query().Get("from")
	tillStr := r.URL.Query().Get("till")

	from, err := time.Parse(time.RFC3339, fromStr)
	if err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, "Invalid 'from' timestamp format, expected RFC3339")
		return
	}

	till, err := time.Parse(time.RFC3339, tillStr)
	if err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, "Invalid 'till' timestamp format, expected RFC3339")
		return
	}

	telemetries, err := h.telemUsecase.GetTelemetryInRange(r.Context(), deviceSn, from, till)

	if err != nil {
		if errors.Is(err, entity.ErrDeviceNotFound) {
			response.WriteErrorResponse(w, http.StatusNotFound, "Device not found")
			return
		}
		if errors.Is(err, entity.ErrTelemetryNotFound) {
			response.WriteErrorResponse(w, http.StatusNoContent, "No telemetry for this device")
			return
		}
		response.WriteErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	writeJson(w, http.StatusOK, telemetries)

}
