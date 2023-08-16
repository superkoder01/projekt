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
