package server

import (
	"github.com/gin-gonic/gin"
	hutil "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/handler/impl"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/configuration"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	a "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/auth"
	"net/http"
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
	logger.Debugf("Incoming request: %s %s\n %s", c.Request.Method,
		c.Request.RequestURI,
		c.Request)
}

// Verify RBAC - match configured rules against token content.
// At this point we dont validate token itself.
func verifyRBAC(c *gin.Context) {
	logger.Debugf("Verifying RBAC rules")
	fullPath := c.Request.RequestURI

	rbacConf := conf.GetRBACConfig()
	auth := a.NewAuth(nil)

	if rbacConf.Omit(fullPath) {
		logger.Debugf("Used endpoint: %s is on list to omit, no rules applied", fullPath)
		return
	}

	authValue := c.GetHeader(AUTHORIZATION)
	token, err := api_utils.GetTokenFromAuthHeader(authValue)
	if err != nil {
		hutil.HandleError(e.Wrap(err, http.StatusUnauthorized), c)
		return
	}

	role, err := auth.GetVerifiedRole(token)
	if err != nil {
		hutil.HandleError(e.ApiErrRoleTooLow, c)
		return
	}

	if !rbacConf.IsAllowed(role, fullPath, c.Request.Method) {
		hutil.HandleError(e.ApiErrEndpointForbidden, c)
		return
	}

	logger.Debugf("Request meets the RBAC rules")
}
