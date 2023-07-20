package api_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlQueryToSQL(t *testing.T) {
	url := "https://example.com/api/management/customerAccounts?" +
		"&limit=20&offset=5&sort=asc:name&filterFields=name,lastName,pesel&filterValues=Jan,Kowalski"
	query := ParseQuery("CUSTOMER_ACCOUNT", url)
	sql, err := query.BuildSQLFromApiQuery()
	assert.Nil(t, err)
	assert.Equal(t, 20, sql.Limit)
	assert.Equal(t, 5, sql.Offset)
	assert.Equal(t, "NAME asc", sql.Order)

	assert.Equal(t,
		"(`CUSTOMER_ACCOUNT`.`NAME` like '%Jan%' or `CUSTOMER_ACCOUNT`.`LAST_NAME` like '%Jan%' or `CUSTOMER_ACCOUNT`.`PESEL` like '%Jan%') "+
			"or (`CUSTOMER_ACCOUNT`.`NAME` like '%Kowalski%' or `CUSTOMER_ACCOUNT`.`LAST_NAME` like '%Kowalski%' or `CUSTOMER_ACCOUNT`.`PESEL` like '%Kowalski%')",
		sql.Filter)
}

func TestUrlQueryToSQLNoSort(t *testing.T) {
	url := "https://example.com/api/management/customerAccounts?" +
		"&limit=20&offset=5&filterFields=name,lastName,pesel&filterValues=Jan,Kowalski,90070602141"
	query := ParseQuery("CUSTOMER_ACCOUNT", url)
	sql, err := query.BuildSQLFromApiQuery()
	assert.Nil(t, err)
	assert.Equal(t, 20, sql.Limit)
	assert.Equal(t, 5, sql.Offset)
	assert.Equal(t, "", sql.Order)

	assert.Equal(t,
		"(`CUSTOMER_ACCOUNT`.`NAME` like '%Jan%' or `CUSTOMER_ACCOUNT`.`LAST_NAME` like '%Jan%' or `CUSTOMER_ACCOUNT`.`PESEL` like '%Jan%') "+
			"or (`CUSTOMER_ACCOUNT`.`NAME` like '%Kowalski%' or `CUSTOMER_ACCOUNT`.`LAST_NAME` like '%Kowalski%' or `CUSTOMER_ACCOUNT`.`PESEL` like '%Kowalski%') "+
			"or (`CUSTOMER_ACCOUNT`.`NAME` like '%90070602141%' or `CUSTOMER_ACCOUNT`.`LAST_NAME` like '%90070602141%' or `CUSTOMER_ACCOUNT`.`PESEL` like '%90070602141%')",
		sql.Filter)
}
