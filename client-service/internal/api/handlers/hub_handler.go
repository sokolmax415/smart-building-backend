package handler

import (
	genericresp "client-service/internal/api/types/generic_response"
	resp "client-service/internal/api/types/hub/response"
	"client-service/internal/entity"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HubHandler struct {
	hubUsecase HubUsecase
}

func NewHubHandler(hubUsecase HubUsecase) *HubHandler {
	return &HubHandler{hubUsecase: hubUsecase}
}

func handleHubError(w http.ResponseWriter, err error) {
	if err != nil {
		if errors.Is(err, entity.ErrHubNotFound) {
			genericresp.WriteErrorResponse(w, http.StatusNotFound, "Hub not found")
			return
		}
		genericresp.WriteErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
}

func (h *HubHandler) GetDevicesForHub(w http.ResponseWriter, r *http.Request) {
	hubSn := chi.URLParam(r, "hub_sn")
	devices, err := h.hubUsecase.GetDevicesForHub(r.Context(), hubSn)
	if err != nil {
		handleHubError(w, err)
		return
	}
	writeJson(w, http.StatusOK, devices)
}

func (h *HubHandler) GetDeviceCountForHub(w http.ResponseWriter, r *http.Request) {
	hubSn := chi.URLParam(r, "hub_sn")
	deviceCount, err := h.hubUsecase.GetDeviceCountForHub(r.Context(), hubSn)
	if err != nil {
		handleHubError(w, err)
		return
	}
	writeJson(w, http.StatusOK, resp.DeviceCountResponse{HubSn: hubSn, DeviceCount: deviceCount})
}

func (h *HubHandler) GetHubInfo(w http.ResponseWriter, r *http.Request) {
	hubSn := chi.URLParam(r, "hub_sn")
	hub, err := h.hubUsecase.GetHubInfo(r.Context(), hubSn)
	if err != nil {
		handleHubError(w, err)
		return
	}
	writeJson(w, http.StatusOK, hub)
}

func (h *HubHandler) DeleteHub(w http.ResponseWriter, r *http.Request) {
	hubSn := chi.URLParam(r, "hub_sn")
	err := h.hubUsecase.DeleteHub(r.Context(), hubSn)
	if err != nil {
		handleHubError(w, err)
		return
	}
	genericresp.WriteSuccessResponse(w, http.StatusOK, "Hub was deleted successfully")
}
