package iRepositories

import "inventory/iModels"

type IUserRepository interface {
	GetUserByIdentity(userIdentity string) (iModels.IUserModel, error)
	GetUserByUsernameAndPassword(username, password string) (iModels.IUserModel, error)
}
