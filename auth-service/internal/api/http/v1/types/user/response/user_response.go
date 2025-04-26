package response

type UserResponse struct {
	Firstname        string `json:"firstname"`
	LastName         string `json:"lastname"`
	Login            string `json:"login"`
	Role             string `json:"role"`
	RegistrationTime string `json:"registration_time"`
}
