package entity

import "errors"

//Errors for repository
var (
	ErrOpenDb    error = errors.New("failed to open buildingDB.hubs")
	ErrConnectDb error = errors.New("failed to connect to buildingDB.hubs")

	ErrCreateLocation                error = errors.New("failed to create location")
	ErrGetLocationById               error = errors.New("failed to get location by id")
	ErrScanLocationRow               error = errors.New("failed to scane location row")
	ErrGetAllLocations               error = errors.New("failed to get all locations for user")
	ErrDeleteLocation                error = errors.New("failed to delete location")
	ErrGetLocationsListWithoutParent error = errors.New("failed to get locations without parents")
	ErrGetLocationChildren           error = errors.New("failed to get location's children")
	ErrCheckLocationForUser          error = errors.New("failed to check location for user")
	ErrUpdateLocationType            error = errors.New("failed to update location type")
	ErrUpdateLocationName            error = errors.New("failed to update location name")
	ErrUpdateLocationParentId        error = errors.New("failed to update location parent id")
	ErrGetHubBySn                    error = errors.New("failed to get hub by hub_sn")
	ErrGetDeviceCount                error = errors.New("failed to get device count")
	ErrDeleteHub                     error = errors.New("failed to delete hub")
	ErrGetDeviceList                 error = errors.New("failed to get device list")
	ErrScanDeviceRow                 error = errors.New("failed to scane device row")
	ErrGetAllDevices                 error = errors.New("failed to get all devices for hub")
	ErrScanHubRow                    error = errors.New("failed to scan hub row")
	ErrGetAllHubs                    error = errors.New("failed to get all hubs")
	ErrGetLatestTelemetry            error = errors.New("failed to get latest telemetry")
	ErrScanTelemetryRow              error = errors.New("failed to scan telemetry row")
	ErrGetTelemetryInRange           error = errors.New("failed to get telemtry in range")
	ErrParentNull                    error = errors.New("location not have parent")
	ErrCheckHubExistence                   = errors.New("failed to check hub existence")
	ErrCheckLocationExistence              = errors.New("failed to check location existence")
	ErrCheckUserExistence                  = errors.New("failed to check user existence")
	ErrCheckDeviceExistence                = errors.New("failed to check device existence")
	ErrRegisterOrUpdateHub                 = errors.New("failed to update or register hub")
	ErrUpdateHubUptime                     = errors.New("failed to update uptime")
	ErrRegisterDevice                      = errors.New("failed to register device")
	ErrSaveTelemetry                       = errors.New("failed to save telemetry")
	ErrCheckdeviceExistence                = errors.New("failed to check device existence")
)

var (
	ErrLocationNotFound  error = errors.New("location not found")
	ErrHubNotFound       error = errors.New("hub not found")
	ErrDeviceNotFound    error = errors.New("device not found")
	ErrTelemetryNotFound error = errors.New("telemetry not found")
)

var (
	ErrParseLocationRequest error = errors.New("failed to parse location request")
	ErrValidateLocationName error = errors.New("failed to validate locationName")
	ErrValidateLocationType error = errors.New("failed to validate locationType")
)
