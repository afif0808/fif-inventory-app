package models

type AuthenticationModel struct {
	UserId int
}

func (model AuthenticationModel) GetAuthenticationUserId() (userId int) {
	userId = model.UserId
	return
}
