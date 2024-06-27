package sso

import (
	"context"
	"encoding/json"
	"io"
	"pet-dex-backend/v2/infra/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type UserDetails struct {
	Name     string
	LastName string
	Email    string
}

type GoogleUserDetals struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}

func GetGoogleUserDetails(accessToken string) (*UserDetails, error) {
	env := config.GetEnvConfig()
	conf := &oauth2.Config{
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

	userDetails := UserDetails{
		Name:     googleUserDetails.GivenName,
		LastName: googleUserDetails.FamilyName,
		Email:    googleUserDetails.Email,
	}
	return &userDetails, nil
}
