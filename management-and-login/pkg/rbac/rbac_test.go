package rbac

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

const (
	TEST_RBAC_CONFIG = "../../tests/configs/rbac.yaml"
)

var (
	rbac                 RBAC
	apiPrefix            string
	basicEndpointMethods map[string][]string
)

func TestMain(m *testing.M) {
	var err error

	// Load RBAC config
	var r Rbac
	rbacFile, err := ioutil.ReadFile(TEST_RBAC_CONFIG)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(rbacFile, &r)
	if err != nil {
		panic(err)
	}

	rbac = &r

	// Get http api prefix
	apiPrefix = "/api/management"
	// Basic endpoints allowed for every user to use
	basicEndpointMethods = map[string][]string{
		"/authenticate":               {"POST"},
		"/activate/activationCode123": {"PUT"},
		"/users/details":              {"GET"},
		"/providers/details":          {"GET"},
	}

	// Run test cases
	code := m.Run()

	os.Exit(code)
}

func testBasicEndpoints(t *testing.T, tokenRole string) {
	for k, v := range basicEndpointMethods {
		for _, method := range v {
			assert.True(t, endpointAllowed(tokenRole, apiPrefix+k, method))
		}
	}
}

func endpointAllowed(tokenRole string, endpoint string, method string) bool {
	if !rbac.Omit(endpoint) {
		return rbac.IsAllowed(tokenRole, endpoint, method)
	}
	return true
}

func TestSuperAdminRules(t *testing.T) {
	tokenRole := "SUPER_ADMIN"

	testBasicEndpoints(t, tokenRole)

	endpoint := apiPrefix + "/providers"
	method := "POST"

	assert.True(t, endpointAllowed(tokenRole, endpoint, method))
}

func TestAdminFullRules(t *testing.T) {
	tokenRole := "ADMINISTRATOR_FULL"

	testBasicEndpoints(t, tokenRole)

	endpoint := apiPrefix + "/providers/1/administrators"
	method := "GET"

	assert.True(t, endpointAllowed(tokenRole, endpoint, method))

	endpointNotAllowed := apiPrefix + "/providers"
	methodNotAllowed := "POST"

	assert.False(t, endpointAllowed(tokenRole, endpointNotAllowed, methodNotAllowed))
}

func TestAdminBasicRules(t *testing.T) {
	tokenRole := "ADMINISTRATOR_BASIC"

	testBasicEndpoints(t, tokenRole)

	endpoint := apiPrefix + "/users/2"
	method := "GET"

	assert.True(t, endpointAllowed(tokenRole, endpoint, method))

	endpointNotAllowed := apiPrefix + "/providers"
	methodNotAllowed := "POST"

	assert.False(t, endpointAllowed(tokenRole, endpointNotAllowed, methodNotAllowed))
}

func TestTraderRules(t *testing.T) {
	tokenRole := "TRADER"

	testBasicEndpoints(t, tokenRole)

	endpoint := apiPrefix + "/customerUsers"
	method := "POST"

	assert.True(t, endpointAllowed(tokenRole, endpoint, method))

	endpointNotAllowed := apiPrefix + "/providers/1/administrators"
	methodNotAllowed := "GET"

	assert.False(t, endpointAllowed(tokenRole, endpointNotAllowed, methodNotAllowed))
}

func TestSuperAgentRules(t *testing.T) {
	tokenRole := "SUPER_AGENT"

	testBasicEndpoints(t, tokenRole)

	endpoint := apiPrefix + "/customerAccounts/1"
	method := "DELETE"

	assert.True(t, endpointAllowed(tokenRole, endpoint, method))

	endpointNotAllowed := apiPrefix + "/customerAccounts"
	methodNotAllowed := "POST"

	assert.False(t, endpointAllowed(tokenRole, endpointNotAllowed, methodNotAllowed))
}

func TestAgentRules(t *testing.T) {
	tokenRole := "AGENT"

	testBasicEndpoints(t, tokenRole)

	endpoint := apiPrefix + "/customerAccounts/1"
	method := "PUT"

	assert.True(t, endpointAllowed(tokenRole, endpoint, method))

	endpointNotAllowed := apiPrefix + "/customerAccounts/details"
	methodNotAllowed := "GET"

	assert.False(t, endpointAllowed(tokenRole, endpointNotAllowed, methodNotAllowed))
}

func TestProsumerRules(t *testing.T) {
	tokenRole := "PROSUMER"

	testBasicEndpoints(t, tokenRole)

	endpoint := apiPrefix + "/customerAccounts/details"
	method := "GET"

	assert.True(t, endpointAllowed(tokenRole, endpoint, method))

	endpointNotAllowed := apiPrefix + "/serviceAccessPoints/1"
	methodNotAllowed := "PUT"

	assert.False(t, endpointAllowed(tokenRole, endpointNotAllowed, methodNotAllowed))
}
