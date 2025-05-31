package request

type CreateLocationRequest struct {
	ParentId     *int64 `json:"parent_id"`
	LocationType string `json:"location_type,omitempty"`
	LocationName string `json:"location_name,omitempty"`
}
