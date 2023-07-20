// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	context "context"

	billing "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"

	mock "github.com/stretchr/testify/mock"
)

// InvoicePublisher is an autogenerated mock type for the InvoicePublisher type
type InvoicePublisher struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *InvoicePublisher) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Publish provides a mock function with given fields: ctx, invoice
func (_m *InvoicePublisher) Publish(ctx context.Context, invoice *billing.InvoiceProsument) error {
	ret := _m.Called(ctx, invoice)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *billing.InvoiceProsument) error); ok {
		r0 = rf(ctx, invoice)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewInvoicePublisherT interface {
	mock.TestingT
	Cleanup(func())
}

// NewInvoicePublisher creates a new instance of InvoicePublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInvoicePublisher(t NewInvoicePublisherT) *InvoicePublisher {
	mock := &InvoicePublisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
