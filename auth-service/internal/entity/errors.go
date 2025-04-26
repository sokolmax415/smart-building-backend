package entity

import "errors"

// Errors for repository
var (
	ErrOpenDb    error = errors.New("can't open database")
	ErrConnectDb error = errors.New("can't connect to database")

	ErrUserNotExists error = errors.New("can't find user in database")
	ErrRoleNotExists error = errors.New("can't find role in database")

	ErrGetRoleId          error = errors.New("can't get the role ID from database")
	ErrGetRole            error = errors.New("can't get a role from database")
	ErrCreateUser         error = errors.New("can't create the user in database")
	ErrGetUserByLogin     error = errors.New("can't get the user by login from database")
	ErrCheckUserExistence error = errors.New("can't check user existence in database")
	ErrChangeName         error = errors.New("can't change user's firstname and lastname")
	ErrChangeRole         error = errors.New("can't change user's role in database")
	ErrDeleteUser         error = errors.New("can't delete user from database")
	ErrCreateRole         error = errors.New("can't create new role in database")
	ErrGetAllUsers        error = errors.New("can't get all users from database")
	ErrGetAllRoles        error = errors.New("can't get all user from database")
	ErrScanUserRow        error = errors.New("can't scan user row from database")
	ErrScanRoleRow        error = errors.New("can't scan role row from database")
)

var (
	ErrUserAlreadyExists error = errors.New("user already exists")
)

// Errors for hash
var (
	ErrHashPassword    error = errors.New("failed to hash password")
	ErrComparePassword error = errors.New("failed to compare passwords")
	ErrInvalidPassword error = errors.New("invalid password")
)

// Errors for token
var (
	ErrSignAccessToken      error = errors.New("failed to sign access token")
	ErrSignRefreshToken     error = errors.New("failed to sign refresh token")
	ErrParseAccessToken     error = errors.New("failed to parse accesToken")
	ErrParseRefreshToken    error = errors.New("failed to parse refreshToken")
	ErrTokenSub             error = errors.New("unexpected token subject")
	ErrSigningMethod        error = errors.New("unexpected signing method")
	ErrValidateAccessToken  error = errors.New("failed to validate accesstoken")
	ErrValidateRefreshToken error = errors.New("failed to validate refreshToken")
)

// Errors fo http
var (
	ErrParseLoginRequest      error = errors.New("failed to parse LoginRequest")
	ErrParseRegisterRequest   error = errors.New("failed to parse RegisterRequest")
	ErrParseRefreshRequest    error = errors.New("failed to parse RefreshRequest")
	ErrParseUserRequest       error = errors.New("failed to parse UserRequest")
	ErrParseNameRequest       error = errors.New("failed to parse NameRequest")
	ErrParseChangeRoleRequest error = errors.New("failed to parse ChangeRoleRequest")
	ErrValidateLogin          error = errors.New("failed to validate Login")
	ErrValidatePassword       error = errors.New("failed to validate Password")
	ErrValidateFirstName      error = errors.New("failed to validate FirstName")
	ErrValidateLastName       error = errors.New("failed to validate LastName")
)

func IsNotFound(err error) bool {
	return errors.Is(err, ErrUserNotExists) || errors.Is(err, ErrRoleNotExists)
}

func IsBadRequest(err error) bool {
	return errors.Is(err, ErrParseLoginRequest) || errors.Is(err, ErrParseRegisterRequest) ||
		errors.Is(err, ErrParseRefreshRequest) || errors.Is(err, ErrParseUserRequest) || errors.Is(err, ErrParseNameRequest) || errors.Is(err, ErrParseChangeRoleRequest)
}

func IsBadValidateRequest(err error) bool {
	return errors.Is(err, ErrValidateLogin) || errors.Is(err, ErrValidatePassword) || errors.Is(err, ErrValidateFirstName) || errors.Is(err, ErrValidateLastName)
}
