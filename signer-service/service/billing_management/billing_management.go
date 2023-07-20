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
