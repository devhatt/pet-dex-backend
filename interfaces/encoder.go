package interfaces

import "github.com/golang-jwt/jwt"

type UserClaims struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type Encoder interface {
	NewAccessToken(claims UserClaims) (string, error)
	ParseAccessToken(accessToken string) *UserClaims
}
