package impl

import (
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/enum"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service/impl"
)

type WorkerCustomerAccountHandler interface {
	ListWorkerCustomerAccounts(ctx *gin.Context)
	ListWorkerServiceAccessPoints(ctx *gin.Context)
}

type workerCustomerAccountHandler struct {
	svc service.Service
}

func NewWorkerCustomerAccountHandler(svc service.Service) *workerCustomerAccountHandler {
	return &workerCustomerAccountHandler{svc: svc}
}

func (c *workerCustomerAccountHandler) TableName() string {
	return CUSTOMER_ACCOUNT
}

func (w *workerCustomerAccountHandler) ListWorkerServiceAccessPoints(ctx *gin.Context) {
	logger.Debug("Listing worker service access points")

	roleName, err := api_utils.GetTokenRole(ctx)
	if err != nil {
		HandleError(err, ctx)
		return
	}
	role := int(enum.RoleName(roleName))

	providerId, err := api_utils.GetTokenProviderID(ctx)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	workerId, err := api_utils.GetTokenWorkerID(ctx)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	if workerId < 1 || providerId < 1 {
		HandleError(e.ApiErrNotEnoughData, ctx)
		return
	}

	query := api_utils.ParseQuery(SERVICE_ACCESS_POINT, ctx.Request.RequestURI)

	count, models, err := w.svc.(impl.WorkerCustomerAccountService).ListWorkerServiceAccessPoints(role, providerId, workerId, query)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, wrapQueryResult(count, models))
}

func (w *workerCustomerAccountHandler) ListWorkerCustomerAccounts(ctx *gin.Context) {
	logger.Debug("Listing worker customer accounts")

	roleName, err := api_utils.GetTokenRole(ctx)
	if err != nil {
		HandleError(err, ctx)
		return
	}
	role := int(enum.RoleName(roleName))

	providerId, err := api_utils.GetTokenProviderID(ctx)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	workerId, err := api_utils.GetTokenWorkerID(ctx)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	if workerId < 1 || providerId < 1 {
		HandleError(e.ApiErrNotEnoughData, ctx)
		return
	}

	query := api_utils.ParseQuery(w.TableName(), ctx.Request.RequestURI)

	count, models, err := w.svc.(impl.WorkerCustomerAccountService).ListWorkerCustomerAccounts(role, providerId, workerId, query)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, wrapQueryResult(count, models))
}

func (w *workerCustomerAccountHandler) Query(ctx *gin.Context) {
}

func (w workerCustomerAccountHandler) CheckIfExistWithFilter(ctx *gin.Context) {
}

func (w workerCustomerAccountHandler) GetByID(ctx *gin.Context) {
}

func (w workerCustomerAccountHandler) List(ctx *gin.Context) {
}

func (w workerCustomerAccountHandler) Create(ctx *gin.Context) {
}

func (w workerCustomerAccountHandler) UpdateByID(ctx *gin.Context) {
}

func (w workerCustomerAccountHandler) DeleteByID(ctx *gin.Context) {
}

func (w workerCustomerAccountHandler) GetDetails(ctx *gin.Context) {
}
