package handler

import (
	genericresp "client-service/internal/api/types/generic_response"
	"client-service/internal/api/types/location"
	locationresp "client-service/internal/api/types/location/response"
	"client-service/internal/entity"
	"errors"
	"net/http"
)

type LocationHandler struct {
	locationUcecase LocationUsecase
	hubUsecase      HubUsecase
}

func NewLocationHandler(locationUsecase LocationUsecase, hubUsecase HubUsecase) *LocationHandler {
	return &LocationHandler{locationUcecase: locationUsecase, hubUsecase: hubUsecase}
}

func handleLocationError(w http.ResponseWriter, err error, errMesssage string) {
	if err != nil {
		if errors.Is(err, entity.ErrLocationNotFound) {
			genericresp.WriteErrorResponse(w, http.StatusNotFound, errMesssage)
			return
		}
		genericresp.WriteErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
}

func (h *LocationHandler) CreateNewLocation(w http.ResponseWriter, r *http.Request) {
	locationRequest, err := location.ParseCreateLocationRequest(r)

	if errors.Is(err, entity.ErrParseLocationRequest) {
		genericresp.WriteErrorResponse(w, http.StatusBadRequest, "Ivalid JSON body")
		return
	}

	if errors.Is(err, entity.ErrValidateLocationType) {
		genericresp.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Ivalid location_type")
		return
	}

	if errors.Is(err, entity.ErrValidateLocationName) {
		genericresp.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Ivalid location_name (too short or big)")
		return
	}

	locationId, err := h.locationUcecase.CreateNewLocation(r.Context(), locationRequest.ParentId, locationRequest.LocationType, locationRequest.LocationName)

	if err != nil {
		handleLocationError(w, err, "Location with this parent_id not found")
		return
	}
	writeJson(w, http.StatusOK, locationresp.CreateLocationResponse{LocationId: locationId})
}

func (h *LocationHandler) DeleteLocation(w http.ResponseWriter, r *http.Request) {
	locationId, err := location.ParseLocationId(r)
	if err != nil {
		genericresp.WriteErrorResponse(w, http.StatusBadRequest, "Ivalid location_id format")
		return
	}

	err = h.locationUcecase.DeleteLocation(r.Context(), locationId)
	if err != nil {
		handleLocationError(w, err, "Location not found")
		return
	}
	genericresp.WriteSuccessResponse(w, http.StatusOK, "Location was deleted successfully")
}

func (h *LocationHandler) GetLocation(w http.ResponseWriter, r *http.Request) {
	locationId, err := location.ParseLocationId(r)
	if err != nil {
		genericresp.WriteErrorResponse(w, http.StatusBadRequest, "Ivalid location_id format")
		return
	}

	location, err := h.locationUcecase.GetLocation(r.Context(), locationId)
	if err != nil {
		handleLocationError(w, err, "Location not found")
		return
	}
	writeJson(w, http.StatusOK, location)
}

func (h *LocationHandler) GetRootLocations(w http.ResponseWriter, r *http.Request) {
	locations, err := h.locationUcecase.GetRootLocations(r.Context())
	if err != nil {
		handleLocationError(w, err, "Location not found")
		return
	}
	writeJson(w, http.StatusOK, locations)
}

func (h *LocationHandler) GetChildrenList(w http.ResponseWriter, r *http.Request) {
	locationId, err := location.ParseLocationId(r)
	if err != nil {
		genericresp.WriteErrorResponse(w, http.StatusBadRequest, "Ivalid location_id format")
		return
	}

	locations, err := h.locationUcecase.GetChildrenList(r.Context(), locationId)
	if err != nil {
		handleLocationError(w, err, "Location not found")
		return
	}
	writeJson(w, http.StatusOK, locations)
}

func (h *LocationHandler) GetLocationParent(w http.ResponseWriter, r *http.Request) {
	locationId, err := location.ParseLocationId(r)
	if err != nil {
		genericresp.WriteErrorResponse(w, http.StatusBadRequest, "Ivalid location_id format")
		return
	}

	location, err := h.locationUcecase.GetChildrenParent(r.Context(), locationId)
	if err != nil {
		if errors.Is(err, entity.ErrParentNull) {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		handleLocationError(w, err, "Location not found")
		return
	}
	writeJson(w, http.StatusOK, location)
}

func (h *LocationHandler) GetLocationPath(w http.ResponseWriter, r *http.Request) {
	locationId, err := location.ParseLocationId(r)
	if err != nil {
		genericresp.WriteErrorResponse(w, http.StatusBadRequest, "Ivalid location_id format")
		return
	}
	locations, err := h.locationUcecase.GetPathToLocation(r.Context(), locationId)
	if err != nil {
		handleLocationError(w, err, "Location not found")
		return
	}
	writeJson(w, http.StatusOK, locations)
}

func buildTree(locations []entity.Location) []*locationresp.LocationTreeResponse {
	locationMap := make(map[int64]*locationresp.LocationTreeResponse)
	var roots []*locationresp.LocationTreeResponse
	for _, location := range locations {
		locationMap[location.LocationId] = &locationresp.LocationTreeResponse{
			LocationId: location.LocationId, ParentId: location.ParentId,
			LocationType: location.LocationType, LocationName: location.LocationName,
			CreatedAt: location.CreatedAt, Children: []*locationresp.LocationTreeResponse{}}
	}

	for _, node := range locationMap {
		if node.ParentId == nil {
			roots = append(roots, node)
		}
	}

	for _, node := range locationMap {
		if node.ParentId != nil {
			if parent, exist := locationMap[*node.ParentId]; exist {
				parent.Children = append(parent.Children, node)
			}
		}
	}
	return roots

}

func (h *LocationHandler) GetLocationTree(w http.ResponseWriter, r *http.Request) {
	locations, err := h.locationUcecase.GetLocationList(r.Context())
	if err != nil {
		genericresp.WriteErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	roots := buildTree(locations)
	writeJson(w, http.StatusOK, roots)
}

func (h *LocationHandler) UpdateLocation(w http.ResponseWriter, r *http.Request) {
	locationId, err := location.ParseLocationId(r)
	if err != nil {
		genericresp.WriteErrorResponse(w, http.StatusBadRequest, "Ivalid location_id format")
		return
	}

	locationRequest, err := location.ParseUpdateLocationRequest(r)

	if errors.Is(err, entity.ErrParseLocationRequest) {
		genericresp.WriteErrorResponse(w, http.StatusBadRequest, "Ivalid JSON body")
		return
	}

	if errors.Is(err, entity.ErrValidateLocationType) {
		genericresp.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Ivalid location_type")
		return
	}

	if errors.Is(err, entity.ErrValidateLocationName) {
		genericresp.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Ivalid location_name (too short or big)")
		return
	}

	err = h.locationUcecase.UpdateLocation(r.Context(), locationId, locationRequest.ParentId, locationRequest.LocationType, locationRequest.LocationName)
	if err != nil {
		handleLocationError(w, err, "Location not found")
		return
	}

	genericresp.WriteSuccessResponse(w, http.StatusOK, "Location was successfully updated")
}

func (h *LocationHandler) GetFullLocationInfo(w http.ResponseWriter, r *http.Request) {
	locationId, err := location.ParseLocationId(r)
	if err != nil {
		genericresp.WriteErrorResponse(w, http.StatusBadRequest, "Ivalid location_id format")
		return
	}

	location, err := h.locationUcecase.GetLocation(r.Context(), locationId)
	if err != nil {
		handleLocationError(w, err, "Location not found")
		return
	}

	hubs, err := h.hubUsecase.GetHubList(r.Context(), locationId)
	if err != nil {
		genericresp.WriteErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	devices_slice := make([]entity.Device, 0)
	for _, hub := range hubs {
		devices, err := h.hubUsecase.GetDevicesForHub(r.Context(), hub.HubSn)
		if err != nil {
			genericresp.WriteErrorResponse(w, http.StatusInternalServerError, "Internal error")
			return
		}
		for _, device := range devices {
			devices_slice = append(devices_slice, device)
		}

	}
	writeJson(w, http.StatusOK, locationresp.FullInfoResponse{Location: location, Hubs: hubs, Devices: devices_slice})
}
