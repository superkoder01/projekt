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
package server

import (
	"github.com/labstack/echo/v4"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/accounting-system-integration-service.git/configuration"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/accounting-system-integration-service.git/error"
	apiUtils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	a "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/auth"
	"net/http"
	"strconv"
)

const (
	ACCESS_CONTROL_ALLOW_ORIGIN  = "Access-Control-Allow-Origin"
	ACCESS_CONTROL_ALLOW_HEADERS = "Access-Control-Allow-Headers"
	ACCESS_CONTROL_ALLOW_METHODS = "Access-Control-Allow-Methods"
	MATCH_ALL                    = "*"
	AUTHORIZATION                = "Authorization"
)

// CORS - Accept request from any address
func attachHeaders(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().Header.Add(ACCESS_CONTROL_ALLOW_ORIGIN, MATCH_ALL)
		c.Request().Header.Add(ACCESS_CONTROL_ALLOW_HEADERS, MATCH_ALL)
		c.Request().Header.Add(ACCESS_CONTROL_ALLOW_METHODS, MATCH_ALL)
		return next(c)
	}
}

func logRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.Debugf("Incoming request: %s %s", c.Request().Method, c.Request().RequestURI)
		return next(c)
	}
}

// Verify RBAC - match configured rules against token content.
// At this point we dont validate token itself.
func verifyRBAC(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.Debugf("Verifying RBAC rules")
		fullPath := c.Request().RequestURI

		rbacConf := configuration.GetRBACConfig()
		auth := a.NewAuth(&a.AuthConfig{})

		if rbacConf.Omit(fullPath) {
			logger.Debugf("Used endpoint: %s is on list to omit, no rules applied", fullPath)
			return next(c)
		}

		authValue := c.Request().Header.Get(AUTHORIZATION)
		token, err := apiUtils.GetTokenFromAuthHeader(authValue)
		if err != nil {
			e.HandleError(e.ApiErrNoAuthorizationHeader, c)
			return nil
		}

		providerId, err := auth.GetVerifiedProviderID(token)
		if err != nil {
			e.HandleError(e.Wrap(err, http.StatusUnauthorized), c)
			return nil
		}

		if !matchProviderID(c.Param("providerId"), providerId) {
			e.HandleError(e.ApiErrWrongProvider, c)
			return nil
		}

		role, err := auth.GetVerifiedRole(token)
		if err != nil {
			e.HandleError(e.ApiErrRoleTooLow, c)
			return nil
		}

		if !rbacConf.IsAllowed(role, fullPath, c.Request().Method) {
			e.HandleError(e.ApiErrEndpointForbidden, c)
			return nil
		}
		return next(c)
	}
}

func matchProviderID(fromPath string, fromToken int) bool {
	if fromPath == "" || fromPath == "0" {
		return true
	}

	if fromToken == 0 {
		return true
	}

	fromPathInt, err := strconv.Atoi(fromPath)
	if err != nil {
		return false
	}

	return fromPathInt == fromToken
}
