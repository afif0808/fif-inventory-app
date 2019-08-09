package services

import (
	"errors"
	"fmt"
	"inventory/iModels"
	"inventory/iRepositories"
	"inventory/iServices"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type UserAccessService struct {
	iServices.IUserService
	iServices.IJWTService
	iRepositories.IUserAccessRepositories
}

func (service *UserAccessService) AuthenticateUser(username, password string) (string, error) {
	var user iModels.IUserModel
	var userId int
	var authenticateUserErr error
	var token string
	var claims JWTClaims
	user, authenticateUserErr = service.GetUserByUsernameAndPassword(username, password)
	if authenticateUserErr != nil {
		return "", authenticateUserErr
	}

	userId = user.GetUserId()

	service.UserWhiteList(fmt.Sprint(userId))

	claims = JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "FIF_INVENTORY",
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
		UserIdentity: fmt.Sprint(userId),
	}
	token, authenticateUserErr = service.JWTCreate(claims)
	if authenticateUserErr != nil {
		return "", authenticateUserErr
	}

	return token, nil
}

func (service *UserAccessService) ValidateUser(token string) (iModels.IUserAccess, error) {
	var userAccess iModels.IUserAccess
	var validateUserErr error
	var isUserBlackListed bool
	userAccess, validateUserErr = service.JWTAuthenticate(token)
	isUserBlackListed, validateUserErr = service.UserIsBlackListed(userAccess.GetUserIdentity())
	if isUserBlackListed {
		// NOTE: return "has no access"
		return nil, errors.New("text")
	}
	if validateUserErr != nil {
		return nil, validateUserErr
	}
	return userAccess, nil
}

func (service *UserAccessService) InvalidateUser(userIdentity string) error {
	var invalidateUserErr error
	invalidateUserErr = service.UserBlackList(userIdentity)
	if invalidateUserErr != nil {
		return invalidateUserErr
	}
	return nil
}
