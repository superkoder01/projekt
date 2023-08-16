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
package fee

import "fmt"

type Fee struct {
	NameID int     `json:"nameId"`
	Price  float64 `json:"price"`
}
type Fees struct {
	Fees []Fee
}

func (f *Fee) String() string {
	return fmt.Sprintf("%s", *f)
}

func (f *Fee) SetNameID(i int) {
	f.NameID = i
}

func (f *Fee) SetPrice(i float64) {
	f.Price = i
}
