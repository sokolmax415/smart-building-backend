package usecase

import (
	"auth-service/internal/entity"
	"context"
	"log"
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
		return nil, err
	}

	return users, nil

}

func (usecase *UserUsecase) GetUserInfo(ctx context.Context, login string) (entity.User, error) {
	user, err := usecase.userRep.GetUserByLogin(ctx, login)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (usecase *UserUsecase) ChangeUserRole(ctx context.Context, login, newRole string) error {
	roleId, err := usecase.roleRep.GetIdByRole(ctx, newRole)
	if err != nil {
		return err
	}

	isExist, err := usecase.userRep.IsUserExists(ctx, login)

	if err != nil {
		return err
	}

	if !isExist {
		log.Printf("User not found login=%s", login)
		return entity.ErrUserNotExists
	}

	err = usecase.userRep.ChangeRoleByLogin(ctx, login, roleId)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *UserUsecase) DeleteUser(ctx context.Context, login string) error {
	isExist, err := usecase.userRep.IsUserExists(ctx, login)
	if err != nil {
		return err
	}

	if !isExist {
		log.Printf("User not found login=%s", login)
		return entity.ErrUserNotExists
	}

	err = usecase.userRep.DeleteUser(ctx, login)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *UserUsecase) GetRoleName(ctx context.Context, roleId int64) (string, error) {
	role, err := usecase.roleRep.GetRoleById(ctx, roleId)
	if err != nil {
		return "", err
	}

	return role, nil
}

func (usecase *UserUsecase) CreateNewUser(ctx context.Context, firstname, lastname, login, password string) error {
	isExist, err := usecase.userRep.IsUserExists(ctx, login)
	if err != nil {
		return err
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

func (usecase *UserUsecase) ChangeUserName(ctx context.Context, firstname, lastname, login string) error {
	isExist, err := usecase.userRep.IsUserExists(ctx, login)

	if err != nil {
		return err
	}

	if !isExist {
		log.Printf("User not found login=%s", login)
		return entity.ErrUserNotExists
	}

	err = usecase.userRep.ChangeUserName(ctx, login, firstname, lastname)
	if err != nil {
		return err
	}

	return nil
}
