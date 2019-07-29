package mocks

import (
	"inventory/iModels"
	"inventory/models"

	"github.com/stretchr/testify/mock"
)

type IAuthenticationRepository struct {
	mock.Mock
}

func (_m *IAuthenticationRepository) GetAuthenticationByToken(token string) (authentication iModels.IAuthenticationModel, err error) {
	ret := _m.Called(token)

	var r0 models.AuthenticationModel
	if rf, ok := ret.Get(0).(func(token string) (authentication models.AuthenticationModel)); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(models.AuthenticationModel)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(token string) (err error)); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}
	authentication, err = r0, r1
	return

}
