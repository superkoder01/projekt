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
package tariff_group_parameter

import "fmt"

type TariffGroupParameter struct {
	ID              int     `json:"id,omitempty"`
	TariffGroupID   int     `json:"tariffGroupId"`
	ParameterNameID int     `json:"parameterNameId"`
	Price           float64 `json:"price"`
}

func (tgp *TariffGroupParameter) String() string {
	return fmt.Sprintf("%s", *tgp)
}

func (tgp *TariffGroupParameter) SetTariffGroupID(i int) {
	tgp.TariffGroupID = i
}

func (tgp *TariffGroupParameter) SetParameterNameID(i int) {
	tgp.ParameterNameID = i
}

func (tgp *TariffGroupParameter) SetPrice(f float64) {
	tgp.Price = f
}
