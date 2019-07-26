package services

import (
	"inventory/interfaces"
	"inventory/models"
)

type UserService struct {
	interfaces.IUserRepository
	interfaces.ILoginRepository
}

func (service *UserService) GetUser(token string) (user models.UserModel, userErr error) {
	login, loginErr := service.GetLoginByToken(token)

	if loginErr != nil {
		return
	}
	userId := login.GetLoginUserId()
	user, userErr = service.GetUserById(userId)
	if userErr != nil {
		return
	}
	return
}
