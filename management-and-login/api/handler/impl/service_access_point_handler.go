package impl

import (
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model"
	ms "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/service_access_point"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
	"net/http"
	"strconv"
)

const (
	SERVICE_ACCESS_POINT = "SERVICE_ACCESS_POINT"
)

type ServiceAccessPointHandler interface {
	ListForCustomer(ctx *gin.Context)
}

type serviceAccessPointHandler struct {
	svc service.Service
}

func NewServiceAccessPointHandler(svc service.Service) *serviceAccessPointHandler {
	return &serviceAccessPointHandler{svc: svc}
}

func (c *serviceAccessPointHandler) TableName() string {
	return SERVICE_ACCESS_POINT
}

func (c *serviceAccessPointHandler) List(ctx *gin.Context) {
	logger.Debug("Listing service access points")

	pId, err := api_utils.GetTokenProviderID(ctx)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	models, err := c.svc.GetWithFilter(ms.ServiceAccessPoint{ProviderID: pId})
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, models)
}

func (c *serviceAccessPointHandler) ListForCustomer(ctx *gin.Context) {
	logger.Debug("Listing customer's service access points")

	pId, err := api_utils.GetTokenProviderID(ctx)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	sId := ctx.Param("customerAccountId")
	cId, err := strconv.Atoi(sId)
	if err != nil {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	query := api_utils.ParseQuery(c.TableName(), ctx.Request.RequestURI)
	count, models, err := c.svc.Query(ms.ServiceAccessPoint{ProviderID: pId, AccountID: cId}, query)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, wrapQueryResult(count, models))
}

func (c *serviceAccessPointHandler) Query(ctx *gin.Context) {
	logger.Debug("Querying service access points")

	query := api_utils.ParseQuery(c.TableName(), ctx.Request.RequestURI)
	count, models, err := c.svc.Query(ms.ServiceAccessPoint{}, query)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, wrapQueryResult(count, models))
}

func (c *serviceAccessPointHandler) Create(ctx *gin.Context) {
	logger.Debug("Creating service access point")
	var newServiceAccessPoint ms.ServiceAccessPoint
	err := ctx.BindJSON(&newServiceAccessPoint)

	if err != nil {
		HandleError(err, ctx)
		return
	}

	var createServiceAccessPoint model.Model
	createServiceAccessPoint, err = c.svc.Create(&newServiceAccessPoint)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Infof("ServiceAccessPoint: accountId: %d, meterNumber: %s, has been created",
		newServiceAccessPoint.AccountID,
		newServiceAccessPoint.MeterNumber)
	ctx.JSON(200, createServiceAccessPoint)
}

func (c *serviceAccessPointHandler) GetDetails(ctx *gin.Context) {
}

func (c *serviceAccessPointHandler) CheckIfExistWithFilter(ctx *gin.Context) {
	logger.Debug("Check if service access point with given parameters already exist")

	query := api_utils.ParseQuery(c.TableName(), ctx.Request.RequestURI)
	count, err := c.svc.Check(query)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, checkResult(count))
}

func (c *serviceAccessPointHandler) GetByID(ctx *gin.Context) {
	sId := ctx.Param("serviceAccessPointId")
	logger.Debugf("Get service access point by ID: %s", sId)

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

func (c *serviceAccessPointHandler) UpdateByID(ctx *gin.Context) {
	sId := ctx.Param("serviceAccessPointId")
	logger.Debugf("Update service access point by ID: %s", sId)

	id, err := strconv.Atoi(sId)
	if err != nil {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	var updateServiceAccessPoint ms.ServiceAccessPoint
	err = ctx.BindJSON(&updateServiceAccessPoint)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	model, err := c.svc.UpdateByID(id, &updateServiceAccessPoint)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Infof("ServiceAccessPoint of id %d, has been updated", sId)
	ctx.JSON(200, model)
}

func (c *serviceAccessPointHandler) DeleteByID(ctx *gin.Context) {
	sId := ctx.Param("serviceAccessPointId")
	logger.Debugf("Delete service access point by ID: %s", sId)

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

	logger.Infof("ServiceAccessPoint of id %d, has been deleted", sId)
	ctx.JSON(200, "OK")
}
