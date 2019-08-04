package services

import (
	"inventory/iModels"
	"inventory/interfaces"
)

type UserAuthorizationService struct {
	interfaces.IUserAuthorizationRepository
	interfaces.IJWTService
}

func (service *UserAuthorizationService) UserAuthorization(token string) (iModels.IUserAuthorizationModel, error) {
	var userAuthorization iModels.IUserAuthorizationModel
	var userAuthorizationErr error
	userAuthorization, userAuthorizationErr = service.JWTValidate(token)
	if userAuthorizationErr != nil {
		return nil, userAuthorizationErr
	}
	return userAuthorization, nil
}

func (service *UserAuthorizationService) userSessionAuthorization(token string) (iModels.IUserAuthorizationModel, error) {
	var iUserAuthorization iModels.IUserAuthorizationModel
	var iUserAuthorizationErr error

	return iUserAuthorization, iUserAuthorizationErr
}
