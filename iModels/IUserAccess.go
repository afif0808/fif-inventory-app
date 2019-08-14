package iModels

type IUserAccessModel interface {
	GetUserIdentity() string
	SetUserIdentity(interface{}) error
}
