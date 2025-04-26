package usecase

import (
	"auth-service/internal/entity"
	"context"
	"errors"
	"fmt"
	"time"
)

type AuthUsecase struct {
	userRep    UserRepository
	roleRep    RoleRepository
	tokenSer   TokenService
	hashingSer HashingService
}

// NewAuthUsecase creates AuthUsecase struct.
func NewAuthUsecase(userRep UserRepository, roleRep RoleRepository, tokenSer TokenService, hashingSer HashingService) *AuthUsecase {
	return &AuthUsecase{userRep: userRep, roleRep: roleRep, tokenSer: tokenSer, hashingSer: hashingSer}
}

// Register registers a new user in the database.
func (usecase *AuthUsecase) Register(ctx context.Context, firstname, lastname, login, password string) error {
	isExist, err := usecase.userRep.IsUserExists(ctx, login)

	if err != nil {
		return fmt.Errorf("failed to check user %q existence: %w", login, err)
	}

	if isExist {
		return fmt.Errorf("failed to register as %q: %w", login, entity.ErrUserAlreadyExists)
	}

	hashedPassword, err := usecase.hashingSer.HashPassword(password)

	if err != nil {
		return fmt.Errorf("failed to hash password for user '%q': %w", login, err)
	}

	role := "user"
	roleId, err := usecase.roleRep.GetIdByRole(ctx, role)

	if err != nil {
		return fmt.Errorf("failed to get roleID for role %q: %w", role, err)
	}

	user := entity.User{Firstname: firstname, Lastname: lastname, Login: login, PasswordHash: hashedPassword, RoleId: roleId, RegistrationTime: time.Now()}
	err = usecase.userRep.CreateNewUser(ctx, user)

	if err != nil {
		return fmt.Errorf("failed to create user %q: %w", login, err)
	}

	return nil
}

// Login checks a user's existence, password and creates access token, refresh token for the user.
func (usecase *AuthUsecase) Login(ctx context.Context, login, password string) (string, string, int64, error) {
	user, err := usecase.userRep.GetUserByLogin(ctx, login)
	if err != nil {
		return "", "", 0, fmt.Errorf("failed to get user by %q: %w", login, err)
	}

	err = usecase.hashingSer.CompareHashAndPassword(user.PasswordHash, password)

	if errors.Is(err, entity.ErrInvalidPassword) {
		return "", "", 0, fmt.Errorf("invalild user password for %q: %w", login, err)
	}

	if err != nil {
		return "", "", 0, fmt.Errorf("failed to compare hash and password for user %q: %w", login, err)
	}

	role, err := usecase.roleRep.GetRoleById(ctx, user.RoleId)
	if err != nil {
		return "", "", 0, fmt.Errorf("failed to get user role %q: %w", login, err)
	}

	accessToken, expiresIn, err := usecase.tokenSer.CreateAccessToken(user.Id, role)

	if err != nil {
		return "", "", 0, fmt.Errorf("failed to create accessToken for user %q: %w", login, err)
	}

	refreshToken, err := usecase.tokenSer.CreateRefreshToken(user.Id, role)
	if err != nil {
		return "", "", 0, fmt.Errorf("failed to create refreshToken for user %q: %w", login, err)
	}
	return accessToken, refreshToken, expiresIn, nil
}

// Refresh creates new access token and expires_in by refresh token.
func (usecase *AuthUsecase) Refresh(refreshToken string) (string, int64, error) {
	id, role, err := usecase.tokenSer.ParseRefreshToken(refreshToken)

	if err != nil {
		return "", 0, fmt.Errorf("unathorized: %w", err)
	}

	accessToken, expiresIn, err := usecase.tokenSer.CreateAccessToken(id, role)

	if err != nil {
		return "", 0, fmt.Errorf("failed to create newAccessToken for UserId: %w ", err)
	}
	return accessToken, expiresIn, nil
}
