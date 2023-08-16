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
package tests

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao"
	td "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/tests/db"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	var err error

	// Run MariaDB instance on container
	mariaDbID, mariaDbIP = td.RunMariaDB()

	// Connect to MariaDB
	s, err = mariaDBConnect(mariaDbIP)
	if err != nil {
		panic(err)
	}

	// Initialize database schema
	mariaDBInitSchema(s)

	// Initialize DaoFactory
	df = billing_dao.NewDaoFactory(s)

	// Run test cases
	code := m.Run()

	// Delete MariaDB container
	td.MariaDBStop(mariaDbID)

	os.Exit(code)
}
