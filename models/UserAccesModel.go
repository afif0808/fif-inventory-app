package models

import (
	"errors"
	"fmt"
)

type UserAcessModel struct {
	UserId int
}

func (userAccess *UserAcessModel) GetUserIdentity() string {
	return fmt.Sprint(userAccess.UserId)
}

func (userAccess *UserAcessModel) SetUserIdentity(userIdentity interface{}) error {
	var isInt bool
	userAccess.UserId, isInt = userIdentity.(int)
	if !isInt {
		return errors.New("SetUserIdentity() ERROR : the assigned identity is not int")
	}
	return nil
}
