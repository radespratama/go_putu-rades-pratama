package mocks

import (
	auth "restful-api/src/models/auth"

	mock "github.com/stretchr/testify/mock"
)

type LoginModel struct {
	mock.Mock
}

func (_m *LoginModel) AttemptLogin(data auth.Login) (string, error) {
	ret := _m.Called(data)

	var r0 string
	if rf, ok := ret.Get(0).(func(auth.Login) string); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(auth.Login) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
