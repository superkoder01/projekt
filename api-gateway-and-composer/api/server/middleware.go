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
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/api-gateway-and-composer/internal/configuration"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/api-gateway-and-composer/internal/error"
	a "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/auth"
	"net/http"
	"regexp"
	"strings"
)

const (
	ACCESS_CONTROL_ALLOW_ORIGIN  = "Access-Control-Allow-Origin"
	ACCESS_CONTROL_ALLOW_HEADERS = "Access-Control-Allow-Headers"
	ACCESS_CONTROL_ALLOW_METHODS = "Access-Control-Allow-Methods"
	MATCH_ALL                    = "*"
	AUTHORIZATION                = "Authorization"
)

// Verify RBAC - match configured rules against token content.
// At this point we dont validate token itself.
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
	logger.Debugf("Incoming request: %s %s", c.Request.Method, c.FullPath())
}

func verifySignature(c *gin.Context) {
	requestUri := c.Request.RequestURI
	urls := conf.GetSkipJwtCheckConfig().Urls
	if contains(urls, requestUri) {
		logger.Debug("Skipping jwt check for uri: %s", requestUri)
		return
	}
	ac := conf.GetAuthConfig()
	auth := a.NewAuth(ac.KeyPath)
	authValue := c.GetHeader(AUTHORIZATION)
	token, err := getTokenFromAuthHeader(authValue)
	if err != nil {
		e.HandleError(e.Wrap(err, http.StatusUnauthorized), c)
		return
	}

	tokenInput, err := auth.VerifyToken(token)
	if tokenInput == nil {
		e.HandleError(e.ApiErrInvalidJWTSignature, c)
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		r := strings.Replace(v, "*any", ".*", -1)
		if matched, err := regexp.MatchString(r, str); err == nil {
			if matched {
				return true
			}
		} else {
			logger.Errorf("Error parsing JWT skip uri pattern %s with error %v", v, err)
		}
	}
	return false
}
