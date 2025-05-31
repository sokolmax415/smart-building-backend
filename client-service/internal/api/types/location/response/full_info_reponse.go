package response

import "client-service/internal/entity"

type FullInfoResponse struct {
	Location *entity.Location `json:"location"`
	Hubs     []entity.Hub     `json:"hubs"`
	Devices  []entity.Device  `json:"devices"`
}
