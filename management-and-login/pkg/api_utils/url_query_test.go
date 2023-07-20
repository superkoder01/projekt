package api_utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Run test cases
	code := m.Run()

	os.Exit(code)
}

func TestQueryOffsetSortAscFilter(t *testing.T) {
	url := "https://example.com/api/management/customerAccounts?" +
		"&limit=20&offset=5&sort=asc:name&filterFields=name,lastName,pesel&filterValues=Jan,Kowalski"
	query := ParseQuery("CUSTOMER_ACCOUNT", url)
	assert.Equal(t, 20, query.Limit)
	assert.Equal(t, 5, query.Offset)
	assert.Equal(t, ASC, query.Sort.Order)
	assert.Equal(t, "name", query.Sort.Value)

	var filterFields []string
	for _, ff := range query.FilterFields {
		if ff == "name" {
			filterFields = append(filterFields, ff)
		} else if ff == "lastName" {
			filterFields = append(filterFields, ff)
		} else if ff == "pesel" {
			filterFields = append(filterFields, ff)
		} else {
			assert.Fail(t, "parsing error, query filter field not found")
		}
	}
	assert.Equal(t, 3, len(filterFields))

	var filterValues []string
	for _, fv := range query.FilterValues {
		if fv == "Jan" {
			filterValues = append(filterValues, fv)
		} else if fv == "Kowalski" {
			filterValues = append(filterValues, fv)
		} else {
			assert.Fail(t, "parsing error, query filter value not found")
		}
	}
	assert.Equal(t, 2, len(filterValues))
}

func TestQueryNoSortFilter(t *testing.T) {
	url := "https://example.com/api/management/customerAccounts?" +
		"&limit=20&offset=3&filterFields=name,lastName,pesel&filterValues=Jan,Kowalski,90070602141"
	query := ParseQuery("CUSTOMER_ACCOUNT", url)
	assert.Equal(t, 20, query.Limit)
	assert.Equal(t, 3, query.Offset)
	assert.Nil(t, query.Sort)

	var filterFields []string
	for _, ff := range query.FilterFields {
		if ff == "name" {
			filterFields = append(filterFields, ff)
		} else if ff == "lastName" {
			filterFields = append(filterFields, ff)
		} else if ff == "pesel" {
			filterFields = append(filterFields, ff)
		} else {
			assert.Fail(t, "parsing error, query filter field not found")
		}
	}
	assert.Equal(t, 3, len(filterFields))

	var filterValues []string
	for _, fv := range query.FilterValues {
		if fv == "Jan" {
			filterValues = append(filterValues, fv)
		} else if fv == "Kowalski" {
			filterValues = append(filterValues, fv)
		} else if fv == "90070602141" {
			filterValues = append(filterValues, fv)
		} else {
			assert.Fail(t, "parsing error, query filter value not found")
		}
	}
	assert.Equal(t, 3, len(filterValues))
}

func TestQueryDefault(t *testing.T) {
	url := "https://example.com/api/management/customerAccounts"
	query := ParseQuery("CUSTOMER_ACCOUNT", url)
	assert.Equal(t, 10, query.Limit)
	assert.Equal(t, 0, query.Offset)
	assert.Nil(t, query.Sort)
}
