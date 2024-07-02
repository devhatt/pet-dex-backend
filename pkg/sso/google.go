package sso

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleUserDetals struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}

type GoogleSSO struct {
}

func NewGoogleGateway() *GoogleSSO {
	return &GoogleSSO{}
}

func (g *GoogleSSO) GetUserDetails(accessToken string) (*dto.UserSSODto, error) {
	env := config.GetEnvConfig()
	if env.GOOGLE_OAUTH_CLIENT_ID == "" || env.GOOGLE_OAUTH_CLIENT_SECRET == "" {
		return nil, errors.New("google client id or secret missing")
	}
	if env.GOOGLE_REDIRECT_URL == "" {
		return nil, errors.New("google redirect url missing")
	}
	conf := oauth2.Config{
		ClientID:     env.GOOGLE_OAUTH_CLIENT_ID,
		ClientSecret: env.GOOGLE_OAUTH_CLIENT_SECRET,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint:    google.Endpoint,
		RedirectURL: env.GOOGLE_REDIRECT_URL,
	}

	ctx := context.Background()

	token, err := conf.Exchange(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	response, err := conf.Client(ctx, token).Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	client, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var googleUserDetails GoogleUserDetals
	err = json.Unmarshal(client, &googleUserDetails)
	if err != nil {
		return nil, err
	}

	userDetails := dto.UserSSODto{
		Name:  googleUserDetails.Name,
		Email: googleUserDetails.Email,
	}
	return &userDetails, nil
}

func (f *GoogleSSO) Name() string {
	return "google"
}
