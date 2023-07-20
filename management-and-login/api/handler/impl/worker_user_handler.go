package impl

import (
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/enum"
	mwu "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/worker_user"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
	"net/http"
)

type workerUserHandler struct {
	svc service.Service
}

func (w *workerUserHandler) TableName() string {
	return WORKER
}

func NewWorkerUserHandler(svc service.Service) *workerUserHandler {
	return &workerUserHandler{svc: svc}
}

func (w *workerUserHandler) GetByID(ctx *gin.Context) {
}

func (w *workerUserHandler) GetDetails(ctx *gin.Context) {
}

func (w *workerUserHandler) List(ctx *gin.Context) {
}

func (w *workerUserHandler) Query(ctx *gin.Context) {
}

func (w *workerUserHandler) Create(ctx *gin.Context) {
	logger.Debug("Create worker with user")
	var newWorkerUser mwu.WorkerUser
	err := ctx.BindJSON(&newWorkerUser)

	if err != nil {
		HandleError(err, ctx)
		return
	}

	if !userOperationPermission(ctx, enum.Role(newWorkerUser.RoleID).Name()) {
		HandleError(e.ApiErrRoleTooLow, ctx)
		return
	}

	var createdWorkerUser model.Model
	createdWorkerUser, err = w.svc.Create(&newWorkerUser)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Infof("Worker: %s %s and User: %s, has been created",
		newWorkerUser.FirstName,
		newWorkerUser.LastName,
		*newWorkerUser.Login)
	ctx.JSON(200, createdWorkerUser)
}

func (w *workerUserHandler) CheckIfExistWithFilter(ctx *gin.Context) {
	logger.Debug("Check if user or worker with given parameters already exist")

	query := api_utils.ParseQuery(w.TableName(), ctx.Request.RequestURI)
	count, err := w.svc.Check(query)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, checkResult(count))
}

func (w *workerUserHandler) UpdateByID(ctx *gin.Context) {
}

func (w *workerUserHandler) DeleteByID(ctx *gin.Context) {
}
