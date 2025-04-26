package entity

import "time"

type User struct {
	Id               int64     `json:"id"`
	Firstname        string    `json:"firstname"`
	Lastname         string    `json:"lastname"`
	Login            string    `json:"login"`
	PasswordHash     string    `json:"-"`
	RoleId           int64     `json:"role_id"`
	RegistrationTime time.Time `json:"registration_time"`
}
