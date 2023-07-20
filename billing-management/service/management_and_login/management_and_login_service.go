package management_and_login

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/utils/http"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/configuration"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	http_utils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/utils/http"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	"io/ioutil"
)

const (
	GET_PROVIDER_INFO = "/providers/details"
	//GET_WORKERS_CUSTOMER_ACCOUNTS = "/workers/customerAccounts?limit=10000000000"
	GET_WORKERS_CUSTOMER_ACCOUNTS = "/workers/customerAccounts"
	GET_CUSTOMER_ACCOUNT_INFO     = "/customerAccounts"
)

func GetProvider(ctx *gin.Context) *entity.Provider {
	config := configuration.GetManagementAndLoginConfig()
	url := config.Prefix + GET_PROVIDER_INFO

	resp := http_utils.Get(ctx, config.Host, config.Port, url)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		e.HandleError(e.ErrInternalServerError, ctx)
	}
	var provider entity.Provider
	if err = json.Unmarshal(body, &provider); err != nil {
		e.HandleError(e.ErrInternalServerError, ctx)
	}
	return &provider
}

func GetWorkersCustomerAccounts(ctx *gin.Context) *http.ResponseBody[entity.CustomerAccount] {
	config := configuration.GetManagementAndLoginConfig()
	url := config.Prefix + GET_WORKERS_CUSTOMER_ACCOUNTS

	resp := http_utils.Get(ctx, config.Host, config.Port, url)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		e.HandleError(e.ErrInternalServerError, ctx)
	}
	var accounts http.ResponseBody[entity.CustomerAccount]
	if err = json.Unmarshal(body, &accounts); err != nil {
		e.HandleError(e.ErrInternalServerError, ctx)
	}
	return &accounts
}

func GetCustomerAccountInfo(ctx *gin.Context, customerAccountId string) *entity.CustomerAccount {
	config := configuration.GetManagementAndLoginConfig()
	url := config.Prefix + GET_CUSTOMER_ACCOUNT_INFO + "/" + customerAccountId

	resp := http_utils.Get(ctx, config.Host, config.Port, url)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		e.HandleError(e.ErrInternalServerError, ctx)
	}
	var customer entity.CustomerAccount
	if err = json.Unmarshal(body, &customer); err != nil {
		e.HandleError(e.ErrInternalServerError, ctx)
	}
	return &customer

}
