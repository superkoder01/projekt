// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Status is an autogenerated mock type for the Status type
type Status struct {
	mock.Mock
}

// IsAlive provides a mock function with given fields:
func (_m *Status) IsAlive() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewStatusT interface {
	mock.TestingT
	Cleanup(func())
}

// NewStatus creates a new instance of Status. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStatus(t NewStatusT) *Status {
	mock := &Status{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}