package handler

import (
	"auth-service/internal/entity"
	"context"
)

type AuthUsecase interface {
	Register(context.Context, string, string, string, string) error
	Login(context.Context, string, string) (string, string, int64, error)
	Refresh(string) (string, int64, error)
}

type UserUsecase interface {
	GetUsersList(context.Context) ([]entity.User, error)
	GetUserInfo(context.Context, string) (entity.User, error)
	ChangeUserRole(context.Context, string, string) error
	DeleteUser(context.Context, string) error
	GetRoleName(context.Context, int64) (string, error)
	CreateNewUser(context.Context, string, string, string, string) error
	ChangeUserName(context.Context, string, string, string) error
}
