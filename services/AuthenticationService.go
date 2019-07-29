package services

import (
	"inventory/iModels"
	"inventory/interfaces"
)

type AuthenticationService struct {
	interfaces.IAuthenticationRepository
}

func (service *AuthenticationService) GetAuthentication(token string) (iAuthentication iModels.IAuthenticationModel, iAuthenticationErr error) {
	iAuthentication, iAuthenticationErr = service.GetAuthenticationByToken(token)
	return
}
