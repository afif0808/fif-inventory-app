package models

type LoginModel struct {
	UserId int
}

func (model LoginModel) GetLoginUserId() (userId int) {
	userId = model.UserId
	return
}
