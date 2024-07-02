package sso

import (
	"errors"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"

	fb "github.com/huandu/facebook/v2"
)

type FacebookSSO struct{}

func NewFacebookGateway() *FacebookSSO {
	return &FacebookSSO{}
}

func (f *FacebookSSO) GetUserDetails(accessToken string) (*dto.UserSSODto, error) {
	env := config.GetEnvConfig()
	if env.FACEBOOK_APP_ID == "" || env.FACEBOOK_APP_SECRET == "" {
		return nil, errors.New("facebook app id or secret missing")
	}

	var globalApp = fb.New(env.FACEBOOK_APP_ID, env.FACEBOOK_APP_SECRET)
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

	res.Decode(&userDetails)

	if userDetails.Email == "" {
		return nil, errors.New("email scope not authorized at facebook app")
	}

	return &userDetails, nil
}

func (f *FacebookSSO) Name() string {
	return "facebook"
}
