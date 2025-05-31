package entity

import "time"

type Device struct {
	DeviceSn     string    `json:"device_sn"`
	HubSn        string    `json:"hub_sn"`
	DeviceType   string    `json:"device_type"`
	DeviceName   string    `json:"device_name"`
	LastPingTime time.Time `json:"last_ping_time"`
	FwVersion    string    `json:"fw_version"`
	CreatedAt    string    `json:"created_at"`
}
