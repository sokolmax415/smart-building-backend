package request

type RegisterRequest struct {
	Firstname string `json:"firstname" example:"Ivan"`
	Lastname  string `json:"lastname" example:"Ivanov"`
	Login     string `json:"login" example"ivanov@228"`
	Password  string `json:"password" example:"12345"`
}
