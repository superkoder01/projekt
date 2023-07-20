// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	ports "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
)

// StatusFactory is an autogenerated mock type for the StatusFactory type
type StatusFactory struct {
	mock.Mock
}

// MakeService provides a mock function with given fields:
func (_m *StatusFactory) MakeService() ports.Status {
	ret := _m.Called()

	var r0 ports.Status
	if rf, ok := ret.Get(0).(func() ports.Status); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ports.Status)
		}
	}

	return r0
}

type NewStatusFactoryT interface {
	mock.TestingT
	Cleanup(func())
}

// NewStatusFactory creates a new instance of StatusFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStatusFactory(t NewStatusFactoryT) *StatusFactory {
	mock := &StatusFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}