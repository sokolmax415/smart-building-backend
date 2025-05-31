package entity

type Device struct {
	DeviceSn   string `json:"device_sn"`
	HubSn      string `json:"hub_sn"`
	DeviceType string `json:"device_type"`
	DeviceName string `json:"device_name"`
	FwVersion  string `json:"fw_version"`
}
