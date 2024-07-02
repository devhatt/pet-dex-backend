package interfaces

import "pet-dex-backend/v2/entity/dto"

type SingleSignOnGateway interface {
	GetUserDetails(accessToken string) (*dto.UserSSODto, error)
	Name() string
}

type SingleSignOnProvider interface {
	GetUserDetails(provider, accessToken string) (*dto.UserSSODto, error)
}
