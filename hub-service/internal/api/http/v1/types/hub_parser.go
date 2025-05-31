package types

import (
	"encoding/json"
	types "hub-service/internal/api/http/v1/types/dto"
	"hub-service/internal/entity"
	"log"
	"net/http"
)

func ParseRegisterHubRequest(r *http.Request) (*entity.Hub, error) {
	var hub entity.Hub

	err := json.NewDecoder(r.Body).Decode(&hub)
	if err != nil {
		log.Printf("Error to parse RegisterHubRequest: %v", err)
		return nil, entity.ErrParseRegisterHubRequest
	}

	err = ValidateRegisterHubRequest(&hub)
	if err != nil {
		return nil, err
	}

	return &hub, nil
}

func ParsePingHubRequest(r *http.Request) (*types.PingRequest, error) {
	var pingRequest types.PingRequest

	err := json.NewDecoder(r.Body).Decode(&pingRequest)

	if err != nil {
		log.Printf("Error to parse PingHubRequest: %v", err)
		return nil, entity.ErrParsePingHubRequest
	}

	err = ValidatePingHubRequest(&pingRequest)

	if err != nil {
		return nil, err
	}

	return &pingRequest, nil
}

func ParseRegisterDeviceRequest(r *http.Request) (*entity.Device, error) {
	var device entity.Device

	err := json.NewDecoder(r.Body).Decode(&device)
	if err != nil {
		log.Printf("Error to parse RegisterDeviceRequest %v", err)
		return nil, entity.ErrParseRegisterHubRequest
	}

	err = ValidateRegisterDeviceRequest(&device)
	if err != nil {
		log.Printf("Validation err: %v", err)
		return nil, err
	}

	return &device, nil
}

func ParseSaveTelemetryRequest(r *http.Request) (*entity.Telemetry, error) {
	var telemetry entity.Telemetry

	err := json.NewDecoder(r.Body).Decode(&telemetry)

	if err != nil {
		log.Printf("Error to parse SaveTelemetryRequest: %v", err)
		return nil, entity.ErrParseSaveTelemetryRequest
	}

	err = ValidateSaveTelemetryRequest(&telemetry)
	if err != nil {
		log.Printf("Validation err: %v", err)
		return nil, err
	}

	return &telemetry, nil
}
