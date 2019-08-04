package interfaces

import "inventory/iModels"

type IUserAuthorizationService interface {
	UserAuthorization(token string) (iModels.IUserAuthorizationModel, error)
}
