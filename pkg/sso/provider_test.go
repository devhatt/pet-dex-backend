package sso

import (
	"fmt"
	"testing"

	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/interfaces"
	mockInterfaces "pet-dex-backend/v2/mocks/pet-dex-backend/v2/interfaces"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewProvider(t *testing.T) {
	tcases := map[string]struct {
		signInGateways        []interfaces.SingleSignOnGateway
		mock                  func(*testing.T, []interfaces.SingleSignOnGateway)
		expectedCompareReturn []string
	}{
		"New Providers": {
			signInGateways: []interfaces.SingleSignOnGateway{
				mockInterfaces.NewMockSingleSignOnGateway(t),
				mockInterfaces.NewMockSingleSignOnGateway(t),
			},
			mock: func(t *testing.T, m []interfaces.SingleSignOnGateway) {
				m0, ok := m[0].(*mockInterfaces.MockSingleSignOnGateway)
				require.True(t, ok)
				m0.On("Name").Return("gateway1")
				m1, ok := m[1].(*mockInterfaces.MockSingleSignOnGateway)
				require.True(t, ok)
				m1.On("Name").Return("gateway2")
			},
			expectedCompareReturn: []string{
				"gateway1",
				"gateway2",
			},
		},
		"No Providers": {
			signInGateways:        nil,
			mock:                  func(t *testing.T, m []interfaces.SingleSignOnGateway) {},
			expectedCompareReturn: []string{},
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.mock(t, tcase.signInGateways)

			provider := NewProvider(tcase.signInGateways...)

			for _, gateway := range tcase.expectedCompareReturn {
				_, ok := provider.gateways[gateway]
				assert.True(t, ok, fmt.Sprintf("expected gateway \"%s\" not found", gateway))
			}

			expectedSize := len(tcase.expectedCompareReturn)
			returnedSize := len(provider.gateways)

			assert.Equal(t, expectedSize, returnedSize, fmt.Sprintf("unexpected gateway len. expected %d, received %d", expectedSize, returnedSize))
		})
	}
}

func TestGetUserDetails(t *testing.T) {
	tcases := map[string]struct {
		signInGateway interfaces.SingleSignOnGateway
		providerName  string
		mock          func(*testing.T, interfaces.SingleSignOnGateway, *dto.UserSSODto)
		expectedUser  *dto.UserSSODto
		expectedErr   error
	}{
		"Success": {
			signInGateway: mockInterfaces.NewMockSingleSignOnGateway(t),
			providerName:  "gateway1",
			mock: func(t *testing.T, m interfaces.SingleSignOnGateway, getUserDetailsReturn *dto.UserSSODto) {
				m0, ok := m.(*mockInterfaces.MockSingleSignOnGateway)
				require.True(t, ok)
				m0.On("Name").Return("gateway1")
				m0.On("GetUserDetails", "").Return(getUserDetailsReturn, nil)
			},
			expectedUser: &dto.UserSSODto{
				Name:  "g1",
				Email: "abc@gmail.com",
			},
			expectedErr: nil,
		},
		"Error": {
			signInGateway: mockInterfaces.NewMockSingleSignOnGateway(t),
			providerName:  "gateway1",
			mock: func(t *testing.T, m interfaces.SingleSignOnGateway, getUserDetailsReturn *dto.UserSSODto) {
				m0, ok := m.(*mockInterfaces.MockSingleSignOnGateway)
				require.True(t, ok)
				m0.On("Name").Return("gateway2")
			},
			expectedUser: nil,
			expectedErr:  fmt.Errorf("Provider %s not found", "gateway1"),
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.mock(t, tcase.signInGateway, tcase.expectedUser)

			provider := NewProvider(tcase.signInGateway)

			user, err := provider.GetUserDetails(tcase.providerName, "")

			assert.Equal(t, err, tcase.expectedErr, fmt.Sprintf("unexpected Error expected %v, received %v", tcase.expectedErr, err))

			assert.EqualValues(t, user, tcase.expectedUser, fmt.Sprintf("unexpected User expected %v, received %v", tcase.expectedUser, err))
		})
	}
}
