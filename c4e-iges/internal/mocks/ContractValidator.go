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

// ContractValidator is an autogenerated mock type for the ContractValidator type
type ContractValidator struct {
	mock.Mock
}

// ValidateContract provides a mock function with given fields: ctx, contract
func (_m *ContractValidator) ValidateContract(ctx context.Context, contract *billing.Contract) error {
	ret := _m.Called(ctx, contract)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *billing.Contract) error); ok {
		r0 = rf(ctx, contract)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewContractValidatorT interface {
	mock.TestingT
	Cleanup(func())
}

// NewContractValidator creates a new instance of ContractValidator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewContractValidator(t NewContractValidatorT) *ContractValidator {
	mock := &ContractValidator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
