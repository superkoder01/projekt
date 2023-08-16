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
package distribution_network_operator

import "fmt"

type DistributionNetworkOperator struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name"`
	Nip     string `json:"nip"`
	Address string `json:"address"`
	City    string `json:"city"`
}

func (c *DistributionNetworkOperator) String() string {
	return fmt.Sprintf("%s", *c)
}

func (c *DistributionNetworkOperator) SetCity(s string) {
	c.City = s
}

func (c *DistributionNetworkOperator) SetAddress(s string) {
	c.Address = s
}

func (c *DistributionNetworkOperator) SetNip(s string) {
	c.Nip = s
}

func (c *DistributionNetworkOperator) SetName(s string) {
	c.Name = s
}
