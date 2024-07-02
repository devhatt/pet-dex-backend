package sso

import (
	"errors"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/interfaces"
)

type Provider struct {
	gateways map[string]interfaces.SingleSignOnGateway
}

var ErrProviderNotFound = errors.New("Provider not found")

func NewProvider(gateways ...interfaces.SingleSignOnGateway) *Provider {
	p := &Provider{
		gateways: make(map[string]interfaces.SingleSignOnGateway),
	}

	for _, g := range gateways {
		p.gateways[g.Name()] = g
	}

	return p
}

func (p *Provider) GetUserDetails(provider, accessToken string) (*dto.UserSSODto, error) {

	g, ok := p.gateways[provider]
	if !ok {
		return nil, ErrProviderNotFound
	}

	return g.GetUserDetails(accessToken)

}
