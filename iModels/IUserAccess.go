package iModels

type IUserAccess interface {
	GetUserIdentity() string
	SetUserIdentity(interface{}) error
}
