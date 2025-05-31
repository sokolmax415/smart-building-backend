package response

import (
	"time"
)

type LocationTreeResponse struct {
	LocationId   int64                   `json:"location_id"`
	ParentId     *int64                  `json:"parent_id"`
	LocationType string                  `json:"location_type"`
	LocationName string                  `json:"location_name"`
	CreatedAt    time.Time               `json:"created_at"`
	Children     []*LocationTreeResponse `json:"children"`
}
