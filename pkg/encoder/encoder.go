package encoder

import (
	"pet-dex-backend/v2/interfaces"

	"github.com/golang-jwt/jwt"
)

type EncoderAdapter struct {
	secret string
}

func NewEncoderAdapter(secret string) *EncoderAdapter {
	return &EncoderAdapter{
		secret: secret,
	}
}

func (e *EncoderAdapter) NewAccessToken(claims interfaces.UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken.SignedString([]byte(e.secret))
}

func (e *EncoderAdapter) ParseAccessToken(accessToken string) *interfaces.UserClaims {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &interfaces.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(e.secret), nil
	})
	return parsedAccessToken.Claims.(*interfaces.UserClaims)
}
