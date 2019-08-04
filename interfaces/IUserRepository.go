package interfaces

import "inventory/iModels"

type IUserRepository interface {
	GetUserById(userId int) (iModels.IUserModel, error)
	GetUserByUsernameAndPassword(username, password string) (iModels.IUserModel, error)
}
