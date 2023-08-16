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
package invoiceutil

import (
	"fmt"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing.git/pkg/invoice"
	"math"
)

func FormatFloat(number float64) string {
	return fmt.Sprintf("%v", number)
}

func FormatFloat2(number float64) string {
	return fmt.Sprintf("%.2f", number)
}

func RoundFloat(value float64) float64 {
	return math.Round(value*100) / 100
}

func FormatReadType(readType invoice.ReadType) string {
	switch readType {
	case invoice.DirectRead:
		return "F"
	case invoice.ProvidedRead:
		return "O"
	case invoice.EstimatedRead:
		return "S"
	case invoice.RemoteRead:
		return "Z"
	default:
		return ""
	}
}
