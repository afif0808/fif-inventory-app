package services

import (
	"inventory/iModels"
	"inventory/iServices"
)

type UserAuthorizationService struct {
	iServices.IJWTService
}

func (service *UserAuthorizationService) UserAuthorization(token string) (iModels.IUserAuthorizationModel, error) {
	//Nothing
	return nil, nil
}
