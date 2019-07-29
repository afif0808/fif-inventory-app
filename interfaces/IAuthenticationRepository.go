package interfaces

import "inventory/iModels"

type IAuthenticationRepository interface {
	GetAuthenticationByToken(token string) (iAuthentication iModels.IAuthenticationModel, err error)
}
