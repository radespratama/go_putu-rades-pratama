package mocks

import (
	book "restful-api/src/models/book"

	mock "github.com/stretchr/testify/mock"
)

type BookModel struct {
	mock.Mock
}

func (_m *BookModel) CreateBook(_a0 book.Book) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(book.Book) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBook provides a mock function with given fields: id
func (_m *BookModel) DeleteBook(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllBooks provides a mock function with given fields:
func (_m *BookModel) GetAllBooks() ([]book.Book, error) {
	ret := _m.Called()

	var r0 []book.Book
	if rf, ok := ret.Get(0).(func() []book.Book); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]book.Book)
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

// GetBook provides a mock function with given fields: id
func (_m *BookModel) GetBook(id string) (book.Book, error) {
	ret := _m.Called(id)

	var r0 book.Book
	if rf, ok := ret.Get(0).(func(string) book.Book); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(book.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBook provides a mock function with given fields: _a0, id
func (_m *BookModel) UpdateBook(_a0 book.Book, id string) error {
	ret := _m.Called(_a0, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(book.Book, string) error); ok {
		r0 = rf(_a0, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
