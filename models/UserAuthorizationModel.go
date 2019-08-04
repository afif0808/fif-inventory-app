package models

type UserAuthorizationModel struct {
	UserId int
}

func (model UserAuthorizationModel) GetUserAuthorizationUserId() (userId int) {
	userId = model.UserId
	return
}
