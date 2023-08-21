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
package configuration

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/rbac"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var r rbac.RBAC

func GetRBACConfig() rbac.RBAC {
	return r
}

func LoadRBACConfig(path string) error {
	var rb rbac.Rbac
	rbacFile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(rbacFile, &rb)
	if err != nil {
		return err
	}

	r = &rb

	return nil
}
