// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	model "learngo/restapiserver/pkg/movies/models"

	mock "github.com/stretchr/testify/mock"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// CreateMovie provides a mock function with given fields: movie
func (_m *Store) CreateMovie(movie model.MovieTable) error {
	ret := _m.Called(movie)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.MovieTable) error); ok {
		r0 = rf(movie)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMovie provides a mock function with given fields: id
func (_m *Store) DeleteMovie(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListMovies provides a mock function with given fields: page, pagesize, offset
func (_m *Store) ListMovies(page int, pagesize int, offset int) (error, []model.MovieTable) {
	ret := _m.Called(page, pagesize, offset)

	var r0 error
	var r1 []model.MovieTable
	if rf, ok := ret.Get(0).(func(int, int, int) (error, []model.MovieTable)); ok {
		return rf(page, pagesize, offset)
	}
	if rf, ok := ret.Get(0).(func(int, int, int) error); ok {
		r0 = rf(page, pagesize, offset)
	} else {
		r0 = ret.Error(0)
	}

	if rf, ok := ret.Get(1).(func(int, int, int) []model.MovieTable); ok {
		r1 = rf(page, pagesize, offset)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]model.MovieTable)
		}
	}

	return r0, r1
}

// ListOneMovie provides a mock function with given fields: id
func (_m *Store) ListOneMovie(id int) (error, model.MovieTable) {
	ret := _m.Called(id)

	var r0 error
	var r1 model.MovieTable
	if rf, ok := ret.Get(0).(func(int) (error, model.MovieTable)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	if rf, ok := ret.Get(1).(func(int) model.MovieTable); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(model.MovieTable)
	}

	return r0, r1
}

// UpdateMovie provides a mock function with given fields: initialMovie, updatedMovie
func (_m *Store) UpdateMovie(initialMovie model.MovieTable, updatedMovie model.MovieTable) error {
	ret := _m.Called(initialMovie, updatedMovie)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.MovieTable, model.MovieTable) error); ok {
		r0 = rf(initialMovie, updatedMovie)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
