package iModels

type IUserModel interface {
	GetUserId() int
	GetUserName() string
	GetUserFullname() string
	GetUserEmail() string
}
