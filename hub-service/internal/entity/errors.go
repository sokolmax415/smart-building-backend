package entity

import "errors"

//Errors for repository
var (
	ErrOpenDb    error = errors.New("failed to open buildingDB.hubs")
	ErrConnectDb error = errors.New("failed to connect to buildingDB.hubs")

	ErrCheckHubExistence      = errors.New("failed to check hub existence")
	ErrCheckLocationExistence = errors.New("failed to check location existence")
	ErrCheckDeviceExistence   = errors.New("failed to check device existence")
	ErrRegisterOrUpdateHub    = errors.New("failed to update or register hub")
	ErrUpdateHubUptime        = errors.New("failed to update uptime")
	ErrRegisterDevice         = errors.New("failed to register device")
	ErrSaveTelemetry          = errors.New("failed to save telemetry")
	ErrCheckdeviceExistence   = errors.New("failed to check device existence")
)

//Errors for usecase
var (
	ErrHubNotFound      = errors.New("hub not found")
	ErrDeviceNotFound   = errors.New("device not found")
	ErrLocationNotFound = errors.New("location not found")
)

//Errors for parsing requests
var (
	ErrParseRegisterHubRequest    = errors.New("failed to parse register hub request")
	ErrParsePingHubRequest        = errors.New("failed to parse ping hub request")
	ErrParseRegisterDeviceRequest = errors.New("failed to parse register device request")
	ErrParseSaveTelemetryRequest  = errors.New("failed to parse save telemetry request")
)

//Errors for validation
var (
	ErrValidateHubSn      = errors.New("failed to validate hub_sn")
	ErrValidateLocationId = errors.New("failed to validate location_id")
	ErrValidateUptime     = errors.New("failed to validate uptime")
	ErrValidateFwVersion  = errors.New("failed to validate fw_version")
	ErrValidateDeviceSn   = errors.New("failed to validate device_sn")
	ErrValidateDeviceType = errors.New("failed to validate device_type")
	ErrValidateDeviceName = errors.New("failed to validate device_name")
	ErrValidateSendTime   = errors.New("failed to validate send_time")
	ErrValidateData       = errors.New("failed to validate data")
)

//Errors for token
var (
	ErrParseAccessToken    error = errors.New("failed to parse accesToken")
	ErrTokenSub            error = errors.New("unexpected token subject")
	ErrSigningMethod       error = errors.New("unexpected signing method")
	ErrValidateAccessToken error = errors.New("failed to validate accesstoken")
)

func IsNotFound(err error) bool {
	return errors.Is(err, ErrHubNotFound) || errors.Is(err, ErrDeviceNotFound) || errors.Is(err, ErrLocationNotFound)
}

func IsBadRequest(err error) bool {
	return errors.Is(err, ErrParseRegisterHubRequest) || errors.Is(err, ErrParseRegisterDeviceRequest) ||
		errors.Is(err, ErrParsePingHubRequest) || errors.Is(err, ErrParseSaveTelemetryRequest)
}

func IsBadValidateRequest(err error) bool {
	return errors.Is(err, ErrValidateHubSn) || errors.Is(err, ErrValidateLocationId) || errors.Is(err, ErrValidateUptime) ||
		errors.Is(err, ErrValidateFwVersion) || errors.Is(err, ErrValidateDeviceSn) || errors.Is(err, ErrValidateDeviceType) || errors.Is(err, ErrValidateDeviceName) ||
		errors.Is(err, ErrValidateSendTime) || errors.Is(err, ErrValidateData)
}
