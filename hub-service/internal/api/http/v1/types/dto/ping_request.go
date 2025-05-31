package types

type PingRequest struct {
	HubSn  string `json:"hub_sn"`
	Uptime int64  `json:"uptime"`
}
