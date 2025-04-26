package token

import "errors"

var (
	ErrSignAccessToken      error = errors.New("failed to sign access token")
	ErrSignRefreshToken     error = errors.New("failed to sign refresh token")
	ErrParseAccessToken     error = errors.New("failed to parse accesToken")
	ErrParseRefreshToken    error = errors.New("failed to parse refreshToken")
	ErrTokenSub             error = errors.New("unexpected token subject")
	ErrSigningMethod        error = errors.New("unexpected signing method")
	ErrValidateAccessToken  error = errors.New("failed to validate accessToken")
	ErrValidateRefreshToken error = errors.New("failed to validate refreshToken")
)
