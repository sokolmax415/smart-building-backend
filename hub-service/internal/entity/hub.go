package entity

type Hub struct {
	HubSn      string `json:"hub_sn"`
	LocationId int64  `json:"location_id"`
	Uptime     int64  `json:"uptime"`
	FwVersion  string `json:"fw_version"`
}
