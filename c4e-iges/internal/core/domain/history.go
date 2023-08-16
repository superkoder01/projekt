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
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing.git/pkg/invoice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
)

type EnergyDepositHistory struct {
	Sap     string
	Deposit invoice.Deposit
	Cfg     *config.AppConfig
}

func (edh *EnergyDepositHistory) GetLast(count int) (*billing.DepositHistory, *billing.DepositHistory) {
	valueHistory := new(billing.DepositHistory)
	amountHistory := new(billing.DepositHistory)

	valueHistory.PpeCode = edh.Sap
	amountHistory.PpeCode = edh.Sap

	history := edh.Deposit.Records[Max(len(edh.Deposit.Records)-count, 0):]

	itemValueHistory := []struct {
		invoiceLineItemCode InvoiceLineItemCode
		invoiceLineItemFunc func() []float64
	}{
		{
			InvoiceLineItemHistoryIncomeValue,
			func() []float64 { return EnergyHistoryDepositRecord(history).GetIncomeValueArr() },
		},
		{
			InvoiceLineItemHistoryOutcomeValue,
			func() []float64 { return EnergyHistoryDepositRecord(history).GetOutcomeValueArr() },
		},
		{
			InvoiceLineItemHistoryDepositValue,
			func() []float64 { return EnergyHistoryDepositRecord(history).GetDepositValueArr() },
		},
		{
			InvoiceLineItemHistoryUsedValue,
			func() []float64 { return EnergyHistoryDepositRecord(history).GetUsedValueValueArr() },
		},
		{
			InvoiceLineItemHistoryResidualValue,
			func() []float64 { return EnergyHistoryDepositRecord(history).GetResidualValueValueArr() },
		},
	}

	itemAmountHistory := []struct {
		invoiceLineItemCode InvoiceLineItemCode
		invoiceLineItemFunc func() []float64
	}{
		{
			InvoiceLineItemHistoryIncomeAmount,
			func() []float64 { return EnergyHistoryDepositRecord(history).GetIncomeAmountArr() },
		},
		{
			InvoiceLineItemHistoryOutcomeAmount,
			func() []float64 { return EnergyHistoryDepositRecord(history).GetOutcomeAmountArr() },
		},
		{
			InvoiceLineItemHistoryDepositAmount,
			func() []float64 { return EnergyHistoryDepositRecord(history).GetDepositAmountArr() },
		},
		{
			InvoiceLineItemHistoryUsedAmount,
			func() []float64 { return EnergyHistoryDepositRecord(history).GetUsedValueAmountArr() },
		},
		{
			InvoiceLineItemHistoryResidualAmount,
			func() []float64 { return EnergyHistoryDepositRecord(history).GetResidualValueAmountArr() },
		},
	}

	for i := 0; i < 5; i++ {
		valueHistory.Items = append(valueHistory.Items, billing.EnergyAnnualBalanceItem{
			ItemName: itemValueHistory[i].invoiceLineItemCode.ToName(edh.Cfg),
			ItemCode: itemValueHistory[i].invoiceLineItemCode.ToCode(),
			Periods:  itemValueHistory[i].invoiceLineItemFunc(),
		})

		amountHistory.Items = append(amountHistory.Items, billing.EnergyAnnualBalanceItem{
			ItemName: itemAmountHistory[i].invoiceLineItemCode.ToName(edh.Cfg),
			ItemCode: itemAmountHistory[i].invoiceLineItemCode.ToCode(),
			Periods:  itemAmountHistory[i].invoiceLineItemFunc(),
		})
	}

	return valueHistory, amountHistory
}

type EnergyHistoryDepositRecord []invoice.DepositRecord

func (ehdr EnergyHistoryDepositRecord) GetIncomeValueArr() []float64 {
	result := make([]float64, 0)

	for _, v := range ehdr {
		result = append(result, v.Value.Income)
	}
	return result
}

func (ehdr EnergyHistoryDepositRecord) GetOutcomeValueArr() []float64 {
	result := make([]float64, 0)

	for _, v := range ehdr {
		result = append(result, v.Value.Outcome)
	}
	return result
}

func (ehdr EnergyHistoryDepositRecord) GetDepositValueArr() []float64 {
	result := make([]float64, 0)

	for _, v := range ehdr {
		result = append(result, v.Value.Deposit)
	}
	return result
}

func (ehdr EnergyHistoryDepositRecord) GetUsedValueValueArr() []float64 {
	result := make([]float64, 0)

	for _, v := range ehdr {
		result = append(result, v.Value.UsedValue)
	}
	return result
}

func (ehdr EnergyHistoryDepositRecord) GetResidualValueValueArr() []float64 {
	result := make([]float64, 0)

	for _, v := range ehdr {
		result = append(result, v.Value.ResidualValue)
	}
	return result
}

func (ehdr EnergyHistoryDepositRecord) GetIncomeAmountArr() []float64 {
	result := make([]float64, 0)

	for _, v := range ehdr {
		result = append(result, v.Amount.Income)
	}
	return result
}

func (ehdr EnergyHistoryDepositRecord) GetOutcomeAmountArr() []float64 {
	result := make([]float64, 0)

	for _, v := range ehdr {
		result = append(result, v.Amount.Outcome)
	}
	return result
}

func (ehdr EnergyHistoryDepositRecord) GetDepositAmountArr() []float64 {
	result := make([]float64, 0)

	for _, v := range ehdr {
		result = append(result, v.Amount.Deposit)
	}
	return result
}

func (ehdr EnergyHistoryDepositRecord) GetUsedValueAmountArr() []float64 {
	result := make([]float64, 0)

	for _, v := range ehdr {
		result = append(result, v.Amount.UsedValue)
	}
	return result
}

func (ehdr EnergyHistoryDepositRecord) GetResidualValueAmountArr() []float64 {
	result := make([]float64, 0)

	for _, v := range ehdr {
		result = append(result, v.Amount.ResidualValue)
	}
	return result
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
