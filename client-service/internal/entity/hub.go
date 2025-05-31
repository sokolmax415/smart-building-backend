package entity

import "time"

type Hub struct {
	HubSn        string    `json:"hub_sn"`
	LocationId   int64     `json:"location_id"`
	Uptime       int64     `json:"uptime"`
	LastPingTime time.Time `json:"last_ping_time"`
	FwVersion    string    `json:"fw_version"`
	CreatedAt    time.Time `json:"created_at"`
}
