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
