package entity

import "time"

type Location struct {
	LocationId   int64     `json:"location_id"`
	ParentId     *int64    `json:"parent_id"`
	LocationType string    `json:"location_type"`
	LocationName string    `json:"location_name"`
	CreatedAt    time.Time `json:"created_at"`
}
