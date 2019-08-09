package iServices

import "inventory/iModels"

type IUserService interface {
	GetUserByIdentity(userIdentity string) (iModels.IUserModel, error)
	GetUserByUsernameAndPassword(username, password string) (iModels.IUserModel, error)
}
