package interfaces

import "inventory/iModels"

type IAuthenticationService interface {
	GetAuthentication(token string) (iUser iModels.IUserModel, iUserEr error)
}
