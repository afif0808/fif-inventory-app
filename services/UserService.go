package services

import (
	"inventory/interfaces"
	"inventory/iModels"
)

type UserService struct {
	interfaces.IUserRepository
}

func (service *UserService) GetUser(userId int) (iUser iModels.IUserModel, iUserErr error) {
	iUser, iUserErr = service.GetUserById(userId)
	return
}
