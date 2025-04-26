package usecase

import (
	"auth-service/internal/entity"
	"context"
	"fmt"
	"time"
)

type UserUsecase struct {
	userRep    UserRepository
	roleRep    RoleRepository
	hashingSer HashingService
}

func NewUserUsecase(userRep UserRepository, roleRep RoleRepository, hashingSer HashingService) *UserUsecase {
	return &UserUsecase{userRep: userRep, roleRep: roleRep, hashingSer: hashingSer}
}

func (usecase *UserUsecase) GetUsersList(ctx context.Context) ([]entity.User, error) {
	users, err := usecase.userRep.GetAllUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}

	return users, nil

}

func (usecase *UserUsecase) GetUserInfo(ctx context.Context, login string) (entity.User, error) {
	user, err := usecase.userRep.GetUserByLogin(ctx, login)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to get info about user %q: %w", login, err)
	}

	return user, nil
}

func (usecase *UserUsecase) ChangeUserRole(ctx context.Context, login, newRole string) error {
	roleId, err := usecase.roleRep.GetIdByRole(ctx, newRole)
	if err != nil {
		return fmt.Errorf("failed to get id by role: %w", err)
	}

	isExist, err := usecase.userRep.IsUserExists(ctx, login)

	if err != nil {
		return fmt.Errorf("fail to check user existence %q: %w", login, err)
	}

	if !isExist {
		return fmt.Errorf("user not found %q: %w", login, entity.ErrUserNotExists)
	}

	err = usecase.userRep.ChangeRoleByLogin(ctx, login, roleId)
	if err != nil {
		return fmt.Errorf("failed to change user's role %q: %w", login, err)
	}

	return nil
}

func (usecase *UserUsecase) DeleteUser(ctx context.Context, login string) error {
	isExist, err := usecase.userRep.IsUserExists(ctx, login)
	if err != nil {
		return fmt.Errorf("failed to check user %q existence: %w", login, err)
	}

	if !isExist {
		return fmt.Errorf("failed to delete user %q: %w", login, entity.ErrUserNotExists)
	}

	err = usecase.userRep.DeleteUser(ctx, login)
	if err != nil {
		return fmt.Errorf("failed to delete user by login %q: %w", login, err)
	}

	return nil
}

func (usecase *UserUsecase) GetRoleName(ctx context.Context, roleId int64) (string, error) {
	role, err := usecase.roleRep.GetRoleById(ctx, roleId)
	if err != nil {
		return "", fmt.Errorf("failed to get role by id: %w", err)
	}

	return role, nil
}

func (usecase *UserUsecase) CreateNewUser(ctx context.Context, firstname, lastname, login, password string) error {
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

func (usecase *UserUsecase) ChangeUserName(ctx context.Context, firstname, lastname, login string) error {
	isExist, err := usecase.userRep.IsUserExists(ctx, login)

	if err != nil {
		return fmt.Errorf("failed to check user %q existence: %w", login, err)
	}

	if !isExist {
		return fmt.Errorf("failed to change user(user %q not exist): %w", login, err)
	}

	err = usecase.userRep.ChangeUserName(ctx, firstname, lastname, login)
	if err != nil {
		return fmt.Errorf("failed to change user name %q: %w", login, err)
	}

	return nil
}
