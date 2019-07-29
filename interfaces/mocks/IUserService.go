package mocks

import (
	"inventory/iModels"

	"github.com/stretchr/testify/mock"
)

type IUserService struct {
	mock.Mock
}

func (_m *IUserService) GetUser(userId int) (iUser iModels.IUserModel, iUserErr error) {
	ret := _m.Called(userId)
	var r0 iModels.IUserModel
	if rf, ok := ret.Get(0).(func(userId int) (iUser iModels.IUserModel)); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(iModels.IUserModel)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(userId int) (err error)); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Get(1).(error)
	}
	iUser, iUserErr = r0, r1
	return
}
