package token

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secretAccessTokenKey []byte
}

func NewJWTService(secretAccessTokenKey string) *JWTService {
	return &JWTService{secretAccessTokenKey: []byte(secretAccessTokenKey)}

}

type Claims struct {
	UserId int64
	Role   string
	jwt.RegisteredClaims
}

func (ser *JWTService) ParseAccessToken(accessToken string) (int64, string, error) {
	secret := ser.secretAccessTokenKey
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrSigningMethod
		}
		return secret, nil
	})

	if err != nil {
		return 0, "", ErrParseAccessToken
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return 0, "", ErrValidateAccessToken
	}

	if claims.Subject != "access_token" {
		return 0, "", ErrTokenSub
	}

	return claims.UserId, claims.Role, nil
}
