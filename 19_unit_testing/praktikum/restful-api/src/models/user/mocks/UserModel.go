package mocks

import (
	user "restful-api/src/models/user"

	mock "github.com/stretchr/testify/mock"
)

type UserModel struct {
	mock.Mock
}

func (_m *UserModel) CreateUser(_a0 user.User) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(user.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: id
func (_m *UserModel) DeleteUser(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllUsers provides a mock function with given fields:
func (_m *UserModel) GetAllUsers() ([]user.User, error) {
	ret := _m.Called()

	var r0 []user.User
	if rf, ok := ret.Get(0).(func() []user.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: id
func (_m *UserModel) GetUser(id string) (user.User, error) {
	ret := _m.Called(id)

	var r0 user.User
	if rf, ok := ret.Get(0).(func(string) user.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: _a0, id
func (_m *UserModel) UpdateUser(_a0 user.User, id string) error {
	ret := _m.Called(_a0, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(user.User, string) error); ok {
		r0 = rf(_a0, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
