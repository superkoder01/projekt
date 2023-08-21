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
package billing_management

import (
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/configuration"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/error"
	http_utils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/utils/http"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/utils/http/mSzafir"
	"io/ioutil"
)

func UpdateContractSigningStatus(ctx *gin.Context, signingId string, status *mSzafir.StatusResponse) {
	config := configuration.GetBillingManagementConfig()
	url := config.Prefix + "/contract/signingComplete/" + signingId

	resp := http_utils.Post(ctx, config.Host, config.Port, url, status)
	defer resp.Body.Close()

	_, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		e.HandleError(e.ErrInternalServerError, ctx)
	}

}
