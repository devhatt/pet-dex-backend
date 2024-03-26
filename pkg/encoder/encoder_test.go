package encoder

import (
	"pet-dex-backend/v2/interfaces"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

var FAKE_SECRET = "jwt_secret_key"

func TestCreateEncoderAdapter(t *testing.T) {
	encoder := NewEncoderAdapter(FAKE_SECRET)
	assert.NotNil(t, encoder)
}

func TestCreateNewAccessToken(t *testing.T) {
	encoder := NewEncoderAdapter(FAKE_SECRET)
	claims := interfaces.UserClaims{
		Id:    "any_user_id",
		Name:  "any_user_name",
		Email: "any_user_email",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
	jwt, err := encoder.NewAccessToken(claims)
	assert.Nil(t, err)
	assert.NotNil(t, jwt)
	assert.IsType(t, "string", jwt)
}

func TestParseNewAccessToken(t *testing.T) {
	encoder := NewEncoderAdapter(FAKE_SECRET)
	claims := interfaces.UserClaims{
		Id:    "any_user_id",
		Name:  "any_user_name",
		Email: "any_user_email",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
	jwt, _ := encoder.NewAccessToken(claims)
	claimsParsed := encoder.ParseAccessToken(jwt)
	assert.NotNil(t, claimsParsed)
	assert.Equal(t, claimsParsed.Id, claims.Id)
	assert.Equal(t, claimsParsed.Name, claims.Name)
	assert.Equal(t, claimsParsed.Email, claims.Email)
	assert.Equal(t, claimsParsed.StandardClaims.ExpiresAt, claims.StandardClaims.ExpiresAt)
	assert.NotEqual(t, claimsParsed, claims)
}
