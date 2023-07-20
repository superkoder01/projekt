package impl

import (
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/enum"
	mw "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/worker"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service/impl"
	"net/http"
	"strconv"
)

const (
	WORKER = "WORKER"
)

type WorkerHandler interface {
	ListWorkers(ctx *gin.Context)
}

type workerHandler struct {
	svc service.Service
}

func (w *workerHandler) TableName() string {
	return WORKER
}

func NewWorkerHandler(svc service.Service) *workerHandler {
	return &workerHandler{svc: svc}
}

func (w *workerHandler) CheckIfExistWithFilter(ctx *gin.Context) {
	logger.Debug("Check if service access point with given parameters already exist")

	query := api_utils.ParseQuery(w.TableName(), ctx.Request.RequestURI)
	count, err := w.svc.Check(query)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, checkResult(count))
}

func (w *workerHandler) GetByID(ctx *gin.Context) {
	sId := ctx.Param("workerId")
	logger.Debugf("Get worker by ID: %s", sId)

	id, err := strconv.Atoi(sId)
	if err != nil {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	model, err := w.svc.GetByID(id)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, model)
}

func (w *workerHandler) GetDetails(ctx *gin.Context) {
	workerId, err := api_utils.GetTokenUserID(ctx)
	if err != nil {
		HandleError(err, ctx)
		return
	}
	logger.Debugf("Get worker: %d details", workerId)

	model, err := w.svc.GetByID(workerId)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, model)
}

func (w *workerHandler) List(ctx *gin.Context) {
	logger.Debug("Listing workers")

	pId, err := api_utils.GetTokenProviderID(ctx)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	models, err := w.svc.GetWithFilter(mw.Worker{ProviderID: pId})
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, models)
}

func (w *workerHandler) ListWorkers(ctx *gin.Context) {
	logger.Debug("List workers")

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

	workerService, ok := w.svc.(impl.WorkerService)
	if !ok {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	count, models, err := workerService.ListWorkers(role, providerId, workerId, query)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, wrapQueryResult(count, models))
}

func (w *workerHandler) Query(ctx *gin.Context) {
	logger.Debug("Querying workers")

	pId, err := api_utils.GetTokenProviderID(ctx)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	query := api_utils.ParseQuery(w.TableName(), ctx.Request.RequestURI)
	count, models, err := w.svc.Query(mw.Worker{ProviderID: pId}, query)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, wrapQueryResult(count, models))
}

func (w *workerHandler) Create(ctx *gin.Context) {
	logger.Debug("Creating worker")
	var newWorker mw.Worker
	err := ctx.BindJSON(&newWorker)

	if err != nil {
		HandleError(err, ctx)
		return
	}

	var createdWorker model.Model
	createdWorker, err = w.svc.Create(&newWorker)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Infof("Worker: %s %s, has been created", newWorker.FirstName, newWorker.LastName)
	ctx.JSON(200, createdWorker)
}

func (w *workerHandler) UpdateByID(ctx *gin.Context) {
	sId := ctx.Param("workerId")
	logger.Debugf("Update worker by ID: %s", sId)

	id, err := strconv.Atoi(sId)
	if err != nil {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	var updateWorker mw.Worker
	err = ctx.BindJSON(&updateWorker)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	model, err := w.svc.UpdateByID(id, &updateWorker)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Infof("Worker of id %d, has been updated", sId)
	ctx.JSON(200, model)
}

func (w *workerHandler) DeleteByID(ctx *gin.Context) {
	sId := ctx.Param("workerId")
	logger.Debugf("Delete worker by ID: %s", sId)

	id, err := strconv.Atoi(sId)
	if err != nil {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	workerService, ok := w.svc.(impl.WorkerService)
	if !ok {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	err = workerService.DeleteWithUser(id)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Infof("Worker of id %d, has been deleted", sId)
	ctx.JSON(200, "OK")
}
