package usecase

import (
	"auth-service/internal/entity"
	"context"
	"fmt"
	"log"
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
		log.Printf("User with login=%s already exists", login)
		return entity.ErrUserAlreadyExists
	}

	hashedPassword, err := usecase.hashingSer.HashPassword(password)
	if err != nil {
		return err
	}

	role := "user"
	roleId, err := usecase.roleRep.GetIdByRole(ctx, role)
	if err != nil {
		return err
	}

	user := entity.User{Firstname: firstname, Lastname: lastname, Login: login, PasswordHash: hashedPassword, RoleId: roleId, RegistrationTime: time.Now()}
	err = usecase.userRep.CreateNewUser(ctx, user)

	if err != nil {
		return err
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

	if err != nil {
		return "", "", 0, entity.ErrInvalidPassword
	}

	role, err := usecase.roleRep.GetRoleById(ctx, user.RoleId)
	if err != nil {
		return "", "", 0, err
	}

	accessToken, expiresIn, err := usecase.tokenSer.CreateAccessToken(user.Id, role)

	if err != nil {
		return "", "", 0, err
	}

	refreshToken, err := usecase.tokenSer.CreateRefreshToken(user.Id, role)
	if err != nil {
		return "", "", 0, err
	}
	return accessToken, refreshToken, expiresIn, nil
}

// Refresh creates new access token and expires_in by refresh token.
func (usecase *AuthUsecase) Refresh(refreshToken string) (string, int64, error) {
	id, role, err := usecase.tokenSer.ParseRefreshToken(refreshToken)

	if err != nil {
		return "", 0, err
	}

	accessToken, expiresIn, err := usecase.tokenSer.CreateAccessToken(id, role)

	if err != nil {
		return "", 0, err
	}
	return accessToken, expiresIn, nil
}
