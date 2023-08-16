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

type ServiceOfferGroupStatus string

const (
	SOGS_DRAFT       ServiceOfferGroupStatus = "DRAFT"
	SOGS_FINAL       ServiceOfferGroupStatus = "FINAL"
	SOGS_SENT        ServiceOfferGroupStatus = "SENT"
	SOGS_DELIVERED   ServiceOfferGroupStatus = "DELIVERED"
	SOGS_NEGOTIATION ServiceOfferGroupStatus = "NEGOTIATION"
)

func (s ServiceOfferGroupStatus) Name() string {
	switch s {
	case SOGS_DRAFT:
		return "DRAFT"
	case SOGS_FINAL:
		return "FINAL"
	case SOGS_SENT:
		return "SENT"
	case SOGS_DELIVERED:
		return "DELIVERED"
	case SOGS_NEGOTIATION:
		return "NEGOTIATION"
	default:
		return ""
	}
}
