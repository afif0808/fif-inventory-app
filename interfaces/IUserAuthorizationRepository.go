package interfaces

import "inventory/iModels"

type IUserAuthorizationRepository interface {
	GetUserAuthorizationByTokenFromDB(token string) (iModels.IUserAuthorizationModel, error)
}
