package entity

import (
	"encoding/json"
	"time"
)

type Telemetry struct {
	DeviceSn string          `json:"device_sn"`
	Data     json.RawMessage `json:"data"`
	SendTime time.Time       `json:"send_time"`
}
