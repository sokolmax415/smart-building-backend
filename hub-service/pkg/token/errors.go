package token

import "errors"

var (
	ErrParseAccessToken    error = errors.New("failed to parse accesToken")
	ErrTokenSub            error = errors.New("unexpected token subject")
	ErrSigningMethod       error = errors.New("unexpected signing method")
	ErrValidateAccessToken error = errors.New("failed to validate accessToken")
)
