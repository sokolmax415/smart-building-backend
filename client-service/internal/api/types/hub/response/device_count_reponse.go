package response

type DeviceCountResponse struct {
	HubSn       string `json:"hub_sn"`
	DeviceCount int64  `json:"device_count"`
}
