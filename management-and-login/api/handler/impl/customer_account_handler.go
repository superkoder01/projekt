package impl

import (
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model"
	mc "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/customer_account"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
	"net/http"
	"strconv"
)

const (
	CUSTOMER_ACCOUNT = "CUSTOMER_ACCOUNT"
)

type customerAccountHandler struct {
	svc service.Service
}

func NewCustomerAccountHandler(svc service.Service) *customerAccountHandler {
	return &customerAccountHandler{svc: svc}
}

func (c *customerAccountHandler) TableName() string {
	return CUSTOMER_ACCOUNT
}

func (c *customerAccountHandler) GetDetails(ctx *gin.Context) {
	customerId, err := api_utils.GetTokenCustomerAccountID(ctx)
	if err != nil {
		HandleError(err, ctx)
		return
	}
	logger.Debugf("Get customer account: %d details", customerId)

	model, err := c.svc.GetByID(customerId)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, model)
}

func (c *customerAccountHandler) List(ctx *gin.Context) {
	logger.Debug("Listing customer accounts")

	pId, err := api_utils.GetTokenProviderID(ctx)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	models, err := c.svc.GetWithFilter(mc.CustomerAccount{ProviderID: pId})
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, models)
}

func (c *customerAccountHandler) Query(ctx *gin.Context) {
	logger.Debug("Querying customer accounts")

	pId, err := api_utils.GetTokenProviderID(ctx)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	query := api_utils.ParseQuery(c.TableName(), ctx.Request.RequestURI)
	count, models, err := c.svc.Query(mc.CustomerAccount{ProviderID: pId}, query)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, wrapQueryResult(count, models))
}

func (c *customerAccountHandler) Create(ctx *gin.Context) {
	logger.Debug("Creating customer account")
	var newCustomerAccount mc.CustomerAccount
	err := ctx.BindJSON(&newCustomerAccount)

	if err != nil {
		HandleError(err, ctx)
		return
	}

	var createCustomerAccount model.Model
	createCustomerAccount, err = c.svc.Create(&newCustomerAccount)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Infof("CustomerAccount: %s, has been created",
		newCustomerAccount.Email)
	ctx.JSON(200, createCustomerAccount)
}

func (c *customerAccountHandler) CheckIfExistWithFilter(ctx *gin.Context) {
	logger.Debug("Check if customer account with given parameters already exist")

	query := api_utils.ParseQuery(c.TableName(), ctx.Request.RequestURI)
	count, err := c.svc.Check(query)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, checkResult(count))
}

func (c *customerAccountHandler) GetByID(ctx *gin.Context) {
	sId := ctx.Param("customerAccountId")
	logger.Debugf("Get customer account by ID: %s", sId)

	id, err := strconv.Atoi(sId)
	if err != nil {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	model, err := c.svc.GetByID(id)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, model)
}

func (c *customerAccountHandler) UpdateByID(ctx *gin.Context) {
	sId := ctx.Param("customerAccountId")
	logger.Debugf("Update customer account by ID: %s", sId)

	id, err := strconv.Atoi(sId)
	if err != nil {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	var updateCustomerAccount mc.CustomerAccount
	err = ctx.BindJSON(&updateCustomerAccount)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	model, err := c.svc.UpdateByID(id, &updateCustomerAccount)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Infof("CustomerAccount of id %d, has been updated", sId)
	ctx.JSON(200, model)
}

func (c *customerAccountHandler) DeleteByID(ctx *gin.Context) {
	sId := ctx.Param("customerAccountId")
	logger.Debugf("Delete customer account by ID: %s", sId)

	id, err := strconv.Atoi(sId)
	if err != nil {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	err = c.svc.DeleteByID(id)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Infof("CustomerAccount of id %d, has been deleted", sId)
	ctx.JSON(200, "OK")
}
