package impl

import (
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model"
	mp "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/provider"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service/impl"
	"net/http"
	"strconv"
)

const (
	PROVIDER = "PROVIDER"
)

type providerHandler struct {
	svc service.Service
}

func NewProviderHandler(svc service.Service) *providerHandler {
	return &providerHandler{svc: svc}
}

func (p *providerHandler) TableName() string {
	return PROVIDER
}

func (p *providerHandler) CheckIfExistWithFilter(ctx *gin.Context) {
	logger.Debug("Check if provider with given parameters already exist")

	query := api_utils.ParseQuery(p.TableName(), ctx.Request.RequestURI)
	count, err := p.svc.Check(query)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, checkResult(count))
}

func (p *providerHandler) GetDetails(ctx *gin.Context) {
	providerId, err := api_utils.GetTokenProviderID(ctx)
	if err != nil {
		HandleError(err, ctx)
		return
	}
	logger.Debugf("Get provider: %d details", providerId)

	model, err := p.svc.GetByID(providerId)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, model)
}

func (p *providerHandler) GetByID(ctx *gin.Context) {
	sId := ctx.Param("providerId")
	logger.Debugf("Get provider by ID: %s", sId)

	id, err := strconv.Atoi(sId)
	if err != nil {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	model, err := p.svc.GetByID(id)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, model)
}

func (p *providerHandler) List(ctx *gin.Context) {
	logger.Debug("Listing providers")
	models, err := p.svc.List()
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, models)
}

func (p *providerHandler) Query(ctx *gin.Context) {
	logger.Debug("Querying providers")

	query := api_utils.ParseQuery(p.TableName(), ctx.Request.RequestURI)
	count, models, err := p.svc.Query(mp.Provider{}, query)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, wrapQueryResult(count, models))
}

func (p *providerHandler) Create(ctx *gin.Context) {
	logger.Debug("Creating provider")
	var newProvider mp.Provider
	err := ctx.BindJSON(&newProvider)

	if err != nil {
		HandleError(err, ctx)
		return
	}

	var createCustomerAccount model.Model
	createCustomerAccount, err = p.svc.Create(&newProvider)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Infof("Provider: %s, has been created", newProvider.Name)
	ctx.JSON(200, createCustomerAccount)
}

func (p *providerHandler) UpdateByID(ctx *gin.Context) {
	sId := ctx.Param("providerId")
	logger.Debugf("Update provider by ID: %s", sId)

	id, err := strconv.Atoi(sId)
	if err != nil {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	var updateProvider mp.Provider
	err = ctx.BindJSON(&updateProvider)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	model, err := p.svc.UpdateByID(id, &updateProvider)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Infof("Provider of id %d, has been updated", sId)
	ctx.JSON(200, model)
}

func (p *providerHandler) DeleteByID(ctx *gin.Context) {
	sId := ctx.Param("providerId")
	logger.Debugf("Delete provider by ID: %s", sId)

	id, err := strconv.Atoi(sId)
	if err != nil {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	providerService, ok := p.svc.(impl.ProviderService)
	if !ok {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	err = providerService.DeleteWithAdmins(id)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Infof("Provider of id %d, has been deleted", sId)
	ctx.JSON(200, "OK")
}
