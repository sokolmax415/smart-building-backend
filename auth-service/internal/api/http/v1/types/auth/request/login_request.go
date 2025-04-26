package request

type LoginRequest struct {
	Login    string `json:"login" example:"ivanov@228"`
	Password string `json:"password" example:"12345"`
}
