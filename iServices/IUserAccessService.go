package iServices

import "inventory/iModels"

type IUserAccessService interface {
	AuthenticateUser(username, password string) (token string, authenticateErr error)
	ValidateUser(token string) (userAccess iModels.IUserAccessModel, validateErr error)
	InvalidateUser(userIdentity string) (invalidateErr error)
}
