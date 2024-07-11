package sso

import (
	"errors"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"

	fb "github.com/huandu/facebook/v2"
)

type FacebookSSO struct {
	name   string
	id     string
	secret string
}

func NewFacebookGateway(env *config.Envconfig) *FacebookSSO {
	return &FacebookSSO{
		name:   "facebook",
		id:     env.FACEBOOK_APP_ID,
		secret: env.FACEBOOK_APP_SECRET,
	}
}

func (f *FacebookSSO) GetUserDetails(accessToken string) (*dto.UserSSODto, error) {
	if f.id == "" || f.secret == "" {
		return nil, errors.New("facebook app id or secret missing")
	}

	var globalApp = fb.New(f.id, f.secret)
	session := globalApp.Session(accessToken)
	err := session.Validate()

	if err != nil {
		return nil, err
	}

	res, err := session.Get("/me?fields=name,email", nil)

	if err != nil {
		return nil, err
	}

	var userDetails dto.UserSSODto

	err = res.Decode(&userDetails)

	if err != nil {
		return nil, err
	}

	if userDetails.Email == "" {
		return nil, errors.New("email scope not authorized at facebook app")
	}

	return &userDetails, nil
}

func (f *FacebookSSO) Name() string {
	return f.name
}
