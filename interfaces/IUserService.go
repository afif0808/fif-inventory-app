package interfaces

import "inventory/iModels"

type IUserService interface {
	GetUser(userId int) (iModels.IUserModel, error)
	AuthenticateUser(username, password string) (map[string]interface{}, bool)
}
