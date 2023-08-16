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
package role

type Role int

const (
	SUPER_ADMIN Role = iota + 1
	ADMINISTRATOR_FULL
	ADMINISTRATOR_BASIC
	TRADER
	SUPER_AGENT
	AGENT
	PROSUMER
	NONE
)

func RoleName(name string) Role {
	switch name {
	case "SUPER_ADMIN":
		return SUPER_ADMIN
	case "ADMINISTRATOR_FULL":
		return ADMINISTRATOR_FULL
	case "ADMINISTRATOR_BASIC":
		return ADMINISTRATOR_BASIC
	case "TRADER":
		return TRADER
	case "SUPER_AGENT":
		return SUPER_AGENT
	case "AGENT":
		return AGENT
	case "PROSUMER":
		return PROSUMER
	default:
		return NONE
	}

}

func (r Role) Name() string {
	switch r {
	case SUPER_ADMIN:
		return "SUPER_ADMIN"
	case ADMINISTRATOR_FULL:
		return "ADMINISTRATOR_FULL"
	case ADMINISTRATOR_BASIC:
		return "ADMINISTRATOR_BASIC"
	case TRADER:
		return "TRADER"
	case SUPER_AGENT:
		return "SUPER_AGENT"
	case AGENT:
		return "AGENT"
	case PROSUMER:
		return "PROSUMER"
	default:
		return ""
	}

}
