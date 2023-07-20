package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/enum"
	er "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"net/http"
)

var (
	logger = logging.MustGetLogger("handler")
)

type (
	queryResult struct {
		Amount   int         `json:"amount,omitempty"`
		Elements interface{} `json:"elements,omitempty"`
	}

	result struct {
		Result int `json:"checkResult"`
	}
)

const (
	negative = 0
	positive = 1
)

func checkResult(count int) *result {
	var res int
	if count > 0 {
		res = positive
	}

	return &result{
		Result: res,
	}
}

func wrapQueryResult(count int, r interface{}) *queryResult {
	return &queryResult{
		Amount:   count,
		Elements: r,
	}
}

func HandleError(e error, ctx *gin.Context) {
	err, ok := e.(*er.Error)
	if ok {
		logger.Errorf("%s error, code: %s sent due to: %s", err.Type, err.Code, err.Error())
		ctx.String(err.Code, err.Error())
	} else {
		logger.Errorf("%s error, code: %s sent due to: %s", er.UNKNOWN, http.StatusInternalServerError, e.Error())
		ctx.String(http.StatusInternalServerError, e.Error())
	}
	ctx.Abort()
}

func userOperationPermission(ctx *gin.Context, role string) bool {
	tokenRole, err := api_utils.GetTokenRole(ctx)
	if err != nil {
		return false
	}
	logger.Debug("User with role: %s is trying to execute operation with at least %s role needed", tokenRole, role)

	return enum.RoleName(tokenRole) <= enum.RoleName(role)
}
