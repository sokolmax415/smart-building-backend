package token

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secretAccessTokenKey  []byte
	secretRefreshTokenKey []byte
}

func NewJWTService(secretAccessTokenKey string, secretRefreshTokenKey string) *JWTService {
	return &JWTService{secretAccessTokenKey: []byte(secretAccessTokenKey), secretRefreshTokenKey: []byte(secretRefreshTokenKey)}

}

type Claims struct {
	UserId int64
	Role   string
	jwt.RegisteredClaims
}

const (
	accessTokenTTL  = time.Minute * 10
	refreshTokenTTL = time.Hour * 24 * 7
)

func (ser *JWTService) CreateAccessToken(id int64, role string) (string, int64, error) {
	now := time.Now()
	expiresAt := now.Add(accessTokenTTL)
	claims := Claims{
		UserId: id,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   "access_token",
		}}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedAccessToken, err := accessToken.SignedString(ser.secretAccessTokenKey)

	if err != nil {
		log.Printf("ERROR IN CreateAccessToken: %v", err)
		return "", 0, ErrSignAccessToken
	}
	expiresIn := int64(expiresAt.Sub(now).Seconds())

	return signedAccessToken, expiresIn, nil
}

func (ser *JWTService) CreateRefreshToken(id int64, role string) (string, error) {
	now := time.Now()
	expiresAt := now.Add(refreshTokenTTL)
	claims := Claims{
		UserId: id,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   "refresh_token",
		}}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedRefreshToken, err := refreshToken.SignedString(ser.secretRefreshTokenKey)

	if err != nil {
		log.Printf("ERROR IN CreateRefreshToken: %v", err)
		return "", ErrSignRefreshToken
	}

	return signedRefreshToken, nil

}

func (ser *JWTService) ParseAccessToken(accessToken string) (int64, string, error) {
	secret := ser.secretAccessTokenKey
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("ERROR IN ParseAccessToken(sign method)")
			return nil, ErrSigningMethod
		}
		return secret, nil
	})

	if err != nil {
		log.Printf("ERROR IN ParseAccessToken: %v", err)
		return 0, "", ErrParseAccessToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		log.Printf("ERROR IN ParseAccessToken(validation)")
		return 0, "", ErrValidateAccessToken
	}

	if claims.Subject != "access_token" {
		log.Printf("ERROR IN ParseAccessToken(sub): %v", ok)
		return 0, "", ErrTokenSub
	}

	return claims.UserId, claims.Role, nil
}

func (ser *JWTService) ParseRefreshToken(accessToken string) (int64, string, error) {
	secret := ser.secretRefreshTokenKey

	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("ERROR IN ParseRefreshToken(sign method)")
			return nil, ErrSigningMethod
		}
		return secret, nil
	})

	if err != nil {
		return 0, "", ErrParseRefreshToken
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		log.Printf("ERROR IN ParseRefreshToken(validation)")
		return 0, "", ErrValidateRefreshToken
	}

	if claims.Subject != "refresh_token" {
		log.Printf("ERROR IN ParseRefreshToken(sub)")
		return 0, "", ErrSigningMethod
	}

	return claims.UserId, claims.Role, nil
}
