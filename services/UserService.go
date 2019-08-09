package services

import (
	"inventory/iModels"
	"inventory/iRepositories"
)

type UserService struct {
	iRepositories.IUserRepository
}

func (service *UserService) GetUserByUserIdentity(identity string) (iModels.IUserModel, error) {
	return service.IUserRepository.GetUserByIdentity(identity)
}

func (service *UserService) GetUserByUsernameAndPassword(username, password string) (iModels.IUserModel, error) {
	// NOTE: hashing matters
	return service.IUserRepository.GetUserByUsernameAndPassword(username, password)
}
