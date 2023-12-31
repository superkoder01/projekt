/*
BMSFES. 
Copyright (C) 2022-2035 

This file is part of the BMSFES solution. 
BMSFES is free software: you can redistribute it and/or modify 
it under the terms of the GNU Affero General Public License as 
published by the Free Software Foundation, either version 3 of the 
License, or (at your option) any later version.
 
BMSFES solution is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the 
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License 
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	ports "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	worker "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/messagebroker/worker"
)

// InvoiceEventSubscriberFactory is an autogenerated mock type for the InvoiceEventSubscriberFactory type
type InvoiceEventSubscriberFactory struct {
	mock.Mock
}

// MakeSubscriber provides a mock function with given fields: _a0
func (_m *InvoiceEventSubscriberFactory) MakeSubscriber(_a0 worker.Worker) ports.InvoiceEventSubscriber {
	ret := _m.Called(_a0)

	var r0 ports.InvoiceEventSubscriber
	if rf, ok := ret.Get(0).(func(worker.Worker) ports.InvoiceEventSubscriber); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ports.InvoiceEventSubscriber)
		}
	}

	return r0
}

type NewInvoiceEventSubscriberFactoryT interface {
	mock.TestingT
	Cleanup(func())
}

// NewInvoiceEventSubscriberFactory creates a new instance of InvoiceEventSubscriberFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInvoiceEventSubscriberFactory(t NewInvoiceEventSubscriberFactoryT) *InvoiceEventSubscriberFactory {
	mock := &InvoiceEventSubscriberFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
