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
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	apiUtils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	a "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/auth"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/configuration"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/utils"
	"net/http"
	"net/url"
)

const (
	ACCESS_CONTROL_ALLOW_ORIGIN  = "Access-Control-Allow-Origin"
	ACCESS_CONTROL_ALLOW_HEADERS = "Access-Control-Allow-Headers"
	ACCESS_CONTROL_ALLOW_METHODS = "Access-Control-Allow-Methods"
	MATCH_ALL                    = "*"
	QUESTION_MARK                = "?"
	AUTHORIZATION                = "Authorization"
)

var (
	logger = logging.MustGetLogger("http_server")
)

// Verify RBAC - match configured rules against token content.
// At this point we don't validate token itself.
func handleOptions(c *gin.Context) {
	if c.Request.Method == http.MethodOptions {
		c.String(http.StatusOK, "Options handled")
		c.Abort()
	}
}

// CORS - Accept request from any address
func attachHeaders(c *gin.Context) {
	c.Header(ACCESS_CONTROL_ALLOW_ORIGIN, MATCH_ALL)
	c.Header(ACCESS_CONTROL_ALLOW_HEADERS, MATCH_ALL)
	c.Header(ACCESS_CONTROL_ALLOW_METHODS, MATCH_ALL)
}

func logRequest(c *gin.Context) {
	logger.Debugf("Incoming request: %s %s", c.Request.Method, c.Request.RequestURI)
}

// Verify RBAC - match configured rules against token content.
// At this point we don't validate token itself.
func verifyRBAC(c *gin.Context) {

	logger.Debugf("Verifying RBAC rules")
	urlObject, err := url.Parse(c.Request.RequestURI)
	if err != nil {
		e.HandleError(e.Wrap(err, http.StatusBadRequest), c)
		return
	}
	fullPath := urlObject.Path

	rbacConf := conf.GetRBACConfig()
	auth := a.NewAuth(nil)

	if rbacConf.Omit(fullPath) {
		logger.Debugf("Used endpoint: %s is on list to omit, no rules applied", fullPath)
		return
	}

	authValue := c.GetHeader(AUTHORIZATION)
	token, err := apiUtils.GetTokenFromAuthHeader(authValue)
	if err != nil {
		e.HandleError(e.Wrap(err, http.StatusUnauthorized), c)
		return
	}

	providerId, err := auth.GetVerifiedProviderID(token)
	if err != nil {
		e.HandleError(e.Wrap(err, http.StatusUnauthorized), c)
		return
	}

	if !utils.MatchProviderID(c.Param("providerId"), providerId) {
		e.HandleError(e.ApiErrWrongProvider, c)
		return
	}

	role, err := auth.GetVerifiedRole(token)
	if err != nil {
		e.HandleError(e.ApiErrRoleTooLow, c)
		return
	}

	if !rbacConf.IsAllowed(role, fullPath, c.Request.Method) {
		e.HandleError(e.ApiErrEndpointForbidden, c)
		return
	}
}
