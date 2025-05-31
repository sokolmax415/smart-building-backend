package types

import (
	types "hub-service/internal/api/http/v1/types/dto"
	"hub-service/internal/entity"
	"log"
	"strings"
)

type DeviceType string

const (
	DeviceTypeTemperature DeviceType = "temperature_sensor"
	DeviceTypeMotion      DeviceType = "motion_sensor"
	DeviceTypeHumidity    DeviceType = "humidity_sensor"
	DeviceTypeSmoke       DeviceType = "smoke_sensor"
)

func ValidateRegisterHubRequest(hub *entity.Hub) error {
	log.Printf("Starting validate Hub Register Request: hub_sn=%s, location_id=%d, uptime=%d, fw_version=%s", hub.HubSn, hub.LocationId, hub.Uptime, hub.FwVersion)

	if hub.HubSn == "" || len(hub.HubSn) < 5 {
		log.Printf("Error to validate hub_sn=%s (too short)", hub.HubSn)
		return entity.ErrValidateHubSn
	}

	if hub.LocationId < 0 {
		log.Printf("Error to validate location_id=%d (negative)", hub.LocationId)
		return entity.ErrValidateLocationId
	}

	if hub.Uptime < 0 {
		log.Printf("Error to validate uptime=%d (negative)", hub.Uptime)
		return entity.ErrValidateUptime
	}

	if !strings.HasPrefix(hub.FwVersion, "v") {
		log.Printf("Error to validate fw_version=%s (without prefix 'v')", hub.FwVersion)
		return entity.ErrValidateFwVersion
	}

	return nil
}

func ValidatePingHubRequest(ping *types.PingRequest) error {
	log.Printf("Starting validate Ping Request: hub_sn=%s, uptime=%d", ping.HubSn, ping.Uptime)

	if ping.HubSn == "" || len(ping.HubSn) < 5 {
		log.Printf("Error to validate hub_sn=%s (too short)", ping.HubSn)
		return entity.ErrValidateHubSn
	}

	if ping.Uptime < 0 {
		log.Printf("Error to validate uptime=%d (negative)", ping.Uptime)
		return entity.ErrValidateUptime
	}

	return nil
}

func ValidateRegisterDeviceRequest(device *entity.Device) error {
	log.Printf("Starting validate Register Device Requst: device_sn=%s, hub_sn=%s, device_type=%s, device_name=%s, fw_version=%s", device.DeviceSn, device.HubSn, device.DeviceType, device.DeviceName, device.FwVersion)

	if device.HubSn == "" || len(device.HubSn) < 5 {
		log.Printf("Error to validate hub_sn=%s (too short)", device.HubSn)
		return entity.ErrValidateHubSn
	}

	if device.DeviceSn == "" || len(device.DeviceSn) < 5 {
		log.Printf("Error to validate device_sn=%s (too short)", device.DeviceSn)
		return entity.ErrValidateDeviceSn
	}

	if device.DeviceName == "" || len(device.DeviceName) < 2 || len(device.DeviceName) > 20 {
		log.Printf("Error to validate device_name=%s (too short or too long)", device.DeviceName)
		return entity.ErrValidateDeviceName
	}

	if !IsTypeValid(DeviceType(device.DeviceType)) {
		log.Printf("Error to validate device_type=%s (should be in ['temperature_sensor','motion_sensor', 'smoke_sensor', 'humidity_sensor'])", device.DeviceType)
		return entity.ErrValidateDeviceType
	}
	if !strings.HasPrefix(device.FwVersion, "v") {
		log.Printf("Error to validate fw_version=%s (without prefix 'v')", device.FwVersion)
		return entity.ErrValidateFwVersion
	}

	return nil
}

func IsTypeValid(deviceType DeviceType) bool {
	switch deviceType {
	case DeviceTypeTemperature, DeviceTypeMotion, DeviceTypeHumidity, DeviceTypeSmoke:
		return true
	}
	return false
}

func ValidateSaveTelemetryRequest(telemetry *entity.Telemetry) error {

	//сейчас от устройств идет время где-то в начале эпохи, а не текущее
	/*maxTime := time.Now().Add(5 * time.Minute)
	minTime := time.Now().Add(-10 * time.Hour)

	if telemetry.SendTime.After(maxTime) || telemetry.SendTime.Before(minTime) {
		log.Printf("Error to validate send_time (too old or too big)")
		return entity.ErrValidateSendTime
	}*/

	if len(telemetry.Data) == 0 {
		log.Printf("Error to validate data (empty)")
		return entity.ErrValidateData
	}

	if telemetry.DeviceSn == "" || len(telemetry.DeviceSn) < 5 {
		log.Printf("Error to validate device_sn=%s (too short)", telemetry.DeviceSn)
		return entity.ErrValidateHubSn
	}

	return nil
}
