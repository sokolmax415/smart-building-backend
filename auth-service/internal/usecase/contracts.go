package usecase

import (
	"auth-service/internal/entity"
	"context"
)

type UserRepository interface {
	CreateNewUser(context.Context, entity.User) error
	GetUserByLogin(context.Context, string) (entity.User, error)
	IsUserExists(context.Context, string) (bool, error)
	ChangeRoleByLogin(context.Context, string, int64) error
	DeleteUser(context.Context, string) error
	GetAllUsers(context.Context) ([]entity.User, error)
	ChangeUserName(context.Context, string, string, string) error
}

type RoleRepository interface {
	GetIdByRole(context.Context, string) (int64, error)
	GetRoleById(context.Context, int64) (string, error)
	GetAllRoles(context.Context) ([]entity.Role, error)
	CreateNewRole(context.Context, string, string) error
}

type HashingService interface {
	HashPassword(string) (string, error)
	CompareHashAndPassword(string, string) error
}

type TokenService interface {
	CreateAccessToken(int64, string) (string, int64, error)
	CreateRefreshToken(int64, string) (string, error)
	ParseAccessToken(string) (int64, string, error)
	ParseRefreshToken(string) (int64, string, error)
}
