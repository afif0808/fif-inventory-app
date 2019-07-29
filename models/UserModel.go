package models

type UserModel struct {
	UserId       int
	UserName     string
	UserFullname string
	UserEmail    string
}

func (model *UserModel) GetUserId() (userId int) {
	userId = model.UserId
	return
}
func (model *UserModel) GetUserName() (userName string) {
	userName = model.UserName
	return
}
func (model *UserModel) GetUserFullname() (userFullname string) {
	userFullname = model.UserFullname
	return
}
func (model *UserModel) GetUserEmail() (userEmail string) {
	userEmail = model.UserEmail
	return
}
