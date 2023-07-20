package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/billing_management/contract"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/service"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	logger = logging.MustGetLogger("controller")
)

type signerController struct {
	service service.Service
}

func NewReportController(s service.Service) *signerController {
	return &signerController{service: s}
}

func (c *signerController) InitSign(ctx *gin.Context) {
	//// TODO: Handle new parameter - jsonType
	//// TODO: var newAuthorisation - authorisation.Authorisation

	var newContract contract.Contract
	err := ctx.BindJSON(&newContract)

	if err != nil {
		logger.Errorf("Cannot read contract")
		e.HandleError(err, ctx)
		return
	}
	loginHauth, err := c.service.InitSign(ctx, newContract)
	//// TODO: c.service.InitSign(ctx, newAuthorisation)

	if err != nil {
		e.HandleError(err, ctx)
		return
	}
	ctx.JSON(200, bson.M{"hauth": loginHauth})
}

func (c *signerController) SigningCompletedNotification(ctx *gin.Context) {
	err := c.service.SigningCompletedNotification(ctx)
	if err != nil {
		e.HandleError(err, ctx)
		return
	}
	ctx.JSON(200, nil)
}
