// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	model "go-issues-api/model"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: userId, _a1
func (_m *Usecase) Create(userId int, title string, description string) error {
	ret := _m.Called(userId, title, description)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string, string) error); ok {
		r0 = rf(userId, title, description)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Usecase) GetAll() ([]*model.Issue, error) {
	ret := _m.Called()

	var r0 []*model.Issue
	if rf, ok := ret.Get(0).(func() []*model.Issue); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Issue)
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

type mockConstructorTestingTNewUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsecase creates a new instance of Usecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsecase(t mockConstructorTestingTNewUsecase) *Usecase {
	mock := &Usecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
