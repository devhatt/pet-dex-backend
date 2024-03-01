package middlewares

import (
	"pet-dex-backend/v2/infra/config"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

var jwtSecret = config.GetEnvConfig().JWT_SECRET

func NewAccessToken(claims UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken.SignedString([]byte(jwtSecret))
}

func ParseAccessToken(accessToken string) *UserClaims {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	return parsedAccessToken.Claims.(*UserClaims)
}
