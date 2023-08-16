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
package enum

type ContractStatus string

const (
	CS_DRAFT     ContractStatus = "DRAFT"
	CS_FINAL     ContractStatus = "FINAL"
	CS_SENT      ContractStatus = "SENT"
	CS_DELIVERED ContractStatus = "DELIVERED"
	CS_ACCEPTED  ContractStatus = "ACCEPTED"
)

func (s ContractStatus) Name() string {
	switch s {
	case CS_DRAFT:
		return "DRAFT"
	case CS_FINAL:
		return "FINAL"
	case CS_SENT:
		return "SENT"
	case CS_DELIVERED:
		return "DELIVERED"
	case CS_ACCEPTED:
		return "ACCEPTED"
	default:
		return ""
	}
}
