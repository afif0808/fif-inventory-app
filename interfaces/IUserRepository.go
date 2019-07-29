package interfaces

import "inventory/models"

type IUserRepository interface {
	GetUserById(userId int) (user *models.UserModel, err error)
}
