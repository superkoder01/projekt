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
package domain

import (
	"fmt"
	flowcontrol_util "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/pkg/utils/flowcontrol"
)

type Summary struct {
	VatRate    int
	NetValue   float64
	VatValue   float64
	GrossValue float64
}

func (s Summary) String() string {
	return fmt.Sprintf("{VatRate: %v, NetValue: %v, VatValue: %v, GrossValue: %v}", s.VatRate, s.NetValue, s.VatValue, s.GrossValue)
}

func (s Summary) Negate() *Summary {
	negate := new(Summary)

	negate.VatRate = s.VatRate
	negate.NetValue = flowcontrol_util.Ternary(s.NetValue <= 0, float64(0), -s.NetValue).(float64)
	negate.VatValue = flowcontrol_util.Ternary(s.VatValue <= 0, float64(0), -s.VatValue).(float64)
	negate.GrossValue = flowcontrol_util.Ternary(s.GrossValue <= 0, float64(0), -s.GrossValue).(float64)

	return negate
}

type PaymentSummary struct {
	EnergySell         map[int]*Summary
	EnergyRepurchase   map[int]*Summary
	TradeFee           map[int]*Summary
	EnergyDistribution map[int]*Summary
	DepositIncluded    map[int]*Summary
	Excise             float64
}

func (p PaymentSummary) String() string {
	return fmt.Sprintf("EnergySell: %v, EnergyRepurchase: %v, TradeFee: %v, EnergyDistribution: %v, DepositIncluded: %v, Excise: %v", p.EnergySell, p.EnergyRepurchase, p.TradeFee, p.EnergyDistribution, p.DepositIncluded, p.Excise)
}

type SummaryMap map[int]*Summary

func (sm SummaryMap) Sum() *Summary {
	summary := new(Summary)
	for _, v := range map[int]*Summary(sm) {
		summary.NetValue += v.NetValue
		summary.VatValue += v.VatValue
		summary.GrossValue += v.GrossValue
	}

	return summary
}

func (sm SummaryMap) Negate() map[int]*Summary {
	negate := make(map[int]*Summary)

	for k, v := range map[int]*Summary(sm) {
		negate[k] = v.Negate()
	}

	return negate
}
