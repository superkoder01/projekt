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
	context "context"

	billing "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"

	mock "github.com/stretchr/testify/mock"
)

// RepurchaseInvoiceGenerator is an autogenerated mock type for the RepurchaseInvoiceGenerator type
type RepurchaseInvoiceGenerator struct {
	mock.Mock
}

// GenerateRepurchaseInvoice provides a mock function with given fields: ctx, number
func (_m *RepurchaseInvoiceGenerator) GenerateRepurchaseInvoice(ctx context.Context, number string) (*billing.InvoiceProsumentRepurchase, *billing.InvoiceProsumentRepurchaseDetails, error) {
	ret := _m.Called(ctx, number)

	var r0 *billing.InvoiceProsumentRepurchase
	if rf, ok := ret.Get(0).(func(context.Context, string) *billing.InvoiceProsumentRepurchase); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billing.InvoiceProsumentRepurchase)
		}
	}

	var r1 *billing.InvoiceProsumentRepurchaseDetails
	if rf, ok := ret.Get(1).(func(context.Context, string) *billing.InvoiceProsumentRepurchaseDetails); ok {
		r1 = rf(ctx, number)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*billing.InvoiceProsumentRepurchaseDetails)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, number)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type NewRepurchaseInvoiceGeneratorT interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepurchaseInvoiceGenerator creates a new instance of RepurchaseInvoiceGenerator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepurchaseInvoiceGenerator(t NewRepurchaseInvoiceGeneratorT) *RepurchaseInvoiceGenerator {
	mock := &RepurchaseInvoiceGenerator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
