package interfaces

import "inventory/iModels"

type IUserService interface {
	GetUser(userId int) (iUser iModels.IUserModel, iUserErr error)
}
