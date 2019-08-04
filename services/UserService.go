package services

import (
	"inventory/iModels"
	"inventory/interfaces"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type UserService struct {
	interfaces.IUserRepository
	interfaces.IJWTService
}

func (service *UserService) GetUser(userId int) (iModels.IUserModel, error) {
	var iUser iModels.IUserModel
	var iUserErr error

	iUser, iUserErr = service.GetUserById(userId)
	return iUser, iUserErr
}
func (service *UserService) AuthenticateUser(username, password string) (string, bool) {
	var user iModels.IUserModel
	var userId int
	var token string
	var claims JWTClaims
	user, err := service.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		return "", false
	}

	userId = user.GetUserId()

	claims = JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "FIF_INVENTORY",
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
		UserId: userId,
	}
	token, err = service.JWTCreate(claims)
	if err != nil {
		return "", false
	}
	return token, true
}
