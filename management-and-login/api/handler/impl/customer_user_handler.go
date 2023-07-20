package impl

import (
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model"
	mcu "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/customer_user"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
	"net/http"
)

type customerUserHandler struct {
	svc service.Service
}

func (w *customerUserHandler) TableName() string {
	return CUSTOMER_ACCOUNT
}

func NewCustomerUserHandler(svc service.Service) *customerUserHandler {
	return &customerUserHandler{svc: svc}
}

func (w *customerUserHandler) GetByID(ctx *gin.Context) {
}

func (w *customerUserHandler) GetDetails(ctx *gin.Context) {
}

func (w *customerUserHandler) List(ctx *gin.Context) {
}

func (w *customerUserHandler) Create(ctx *gin.Context) {
	logger.Debug("Create customer account with user")
	var newCustomerUser mcu.CustomerUser
	err := ctx.BindJSON(&newCustomerUser)

	if err != nil {
		HandleError(err, ctx)
		return
	}

	var createdCustomerUser model.Model
	createdCustomerUser, err = w.svc.Create(&newCustomerUser)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Infof("CustomerAccount: %s and User: %s, has been created",
		newCustomerUser.Email,
		*newCustomerUser.Login)
	ctx.JSON(200, createdCustomerUser)
}

func (w *customerUserHandler) CheckIfExistWithFilter(ctx *gin.Context) {
	logger.Debug("Check if user or customer account with given parameters already exist")

	query := api_utils.ParseQuery(w.TableName(), ctx.Request.RequestURI)
	count, err := w.svc.Check(query)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, checkResult(count))
}

func (w *customerUserHandler) UpdateByID(ctx *gin.Context) {
}

func (w *customerUserHandler) DeleteByID(ctx *gin.Context) {
}

func (w *customerUserHandler) Query(ctx *gin.Context) {
}
