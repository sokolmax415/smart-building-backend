package handlers

import (
	"errors"
	"hub-service/internal/api/http/v1/types"
	resp "hub-service/internal/api/http/v1/types/generic_repsponse"
	"hub-service/internal/entity"
	"net/http"
)

type HubHandler struct {
	hubUsecase HubUsecase
}

func NewAuthHandler(hubUsecase HubUsecase) *HubHandler {
	return &HubHandler{hubUsecase: hubUsecase}
}

func (handler *HubHandler) RegisterHub(w http.ResponseWriter, r *http.Request) {
	hub, err := types.ParseRegisterHubRequest(r)

	if entity.IsBadRequest(err) {
		resp.WriteErrorResponse(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	if entity.IsBadValidateRequest(err) {
		resp.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Validation failed")
		return
	}

	err = handler.hubUsecase.RegisterHub(r.Context(), hub)

	if err != nil {
		if errors.Is(err, entity.ErrLocationNotFound) {
			resp.WriteErrorResponse(w, http.StatusNotFound, "Location does not exist")
			return
		}
		resp.WriteErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	resp.WriteSuccessResponse(w, http.StatusCreated, "Hub was successfully created")
}

func (handler *HubHandler) RegisterDevice(w http.ResponseWriter, r *http.Request) {
	device, err := types.ParseRegisterDeviceRequest(r)

	if entity.IsBadRequest(err) {
		resp.WriteErrorResponse(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	if entity.IsBadValidateRequest(err) {
		resp.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Validation failed")
		return
	}

	err = handler.hubUsecase.RegisterDevice(r.Context(), device)

	if err != nil {
		if errors.Is(err, entity.ErrHubNotFound) {
			resp.WriteErrorResponse(w, http.StatusNotFound, "Hub does not exist")
			return
		}
		resp.WriteErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	resp.WriteSuccessResponse(w, http.StatusCreated, "Device was successfully created")
}

func (handler *HubHandler) PingHub(w http.ResponseWriter, r *http.Request) {
	pingReq, err := types.ParsePingHubRequest(r)

	if entity.IsBadRequest(err) {
		resp.WriteErrorResponse(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	if entity.IsBadValidateRequest(err) {
		resp.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Validation failed")
		return
	}

	err = handler.hubUsecase.PingHub(r.Context(), pingReq.HubSn, pingReq.Uptime)

	if err != nil {
		if errors.Is(err, entity.ErrHubNotFound) {
			resp.WriteErrorResponse(w, http.StatusNotFound, "Hub does not exist")
			return
		}
		resp.WriteErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	resp.WriteSuccessResponse(w, http.StatusOK, "Hub was pinged successfully")
}

func (handler *HubHandler) SaveTelemetry(w http.ResponseWriter, r *http.Request) {
	telemetry, err := types.ParseSaveTelemetryRequest(r)

	if entity.IsBadRequest(err) {
		resp.WriteErrorResponse(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	if entity.IsBadValidateRequest(err) {
		resp.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Validation failed")
		return
	}

	err = handler.hubUsecase.SaveTelemetry(r.Context(), telemetry)

	if err != nil {
		if errors.Is(err, entity.ErrDeviceNotFound) {
			resp.WriteErrorResponse(w, http.StatusNotFound, "Device does not exist")
			return
		}
		resp.WriteErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	resp.WriteSuccessResponse(w, http.StatusCreated, "Telemetry was successfully saved")
}
