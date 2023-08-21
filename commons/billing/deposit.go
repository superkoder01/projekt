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
package billing

import "time"

type DepositItemValue struct {
	Amount int `json:"amount" bson:"amount"`
	Value  int `json:"value" bson:"value"`
}

type DepositItem struct {
	In       DepositItemValue `json:"in" bson:"in"`
	Out      DepositItemValue `json:"out" bson:"out"`
	Has      DepositItemValue `json:"has" bson:"has"`
	Used     DepositItemValue `json:"used" bson:"used"`
	Residual DepositItemValue `json:"residual" bson:"residual"`
}

type Timestamps struct {
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type DepositRecord struct {
	DepositId  int                 `json:"depositId" bson:"depositId"`
	CustomerId string              `json:"customerId" bson:"customerId"`
	FromDt     time.Time           `json:"fromDt" bson:"fromDt"`
	ToDt       time.Time           `json:"toDt" bson:"toDt"`
	Timestamps Timestamps          `json:"timestamps" bson:"timestamps"`
	TimeZone   string              `json:"timeZone,omitempty" bson:"timeZone,omitempty"`
	Deposit    map[int]DepositItem `json:"deposit" bson:"deposit"`
}
