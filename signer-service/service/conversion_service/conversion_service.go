package conversion_service

import (
	"github.com/gin-gonic/gin"
	b "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/configuration"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/error"
	http_utils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/utils/http"
	"io/ioutil"
)

const ()

func GetContractPdf(ctx *gin.Context, body *b.Model) {
	config := configuration.GetConversionServiceConfigConfig()
	url := config.Prefix

	resp := http_utils.Post(ctx, config.Host, config.Port, url, body)
	defer resp.Body.Close()

	_, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		e.HandleError(e.ErrInternalServerError, ctx)
	}

}
