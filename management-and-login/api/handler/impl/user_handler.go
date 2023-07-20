package impl

import (
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/enum"
	mu "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/user"
	mua "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/user_activate"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service/impl"
	"net/http"
	"strconv"
)

type userHandler struct {
	svc service.Service
}

type UserHandler interface {
	Activate(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
	ForgotPassword(ctx *gin.Context)
	SendActivationLink(ctx *gin.Context)
	ListSuperAdmins(ctx *gin.Context)
	ListAdministrators(ctx *gin.Context)
}

const (
	USER = "USER"
)

func NewUserHandler(svc service.Service) *userHandler {
	return &userHandler{svc: svc}
}

func (u *userHandler) TableName() string {
	return USER
}

func (u *userHandler) Activate(ctx *gin.Context) {
	logger.Debug("Activating user")
	code := ctx.Param("activationCode")
	if len(code) != 32 {
		HandleError(e.ApiErrInvalidActivationCode, ctx)
		return
	}

	var userActivate mua.UserActivate
	err := ctx.BindJSON(&userActivate)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	us, ok := u.svc.(impl.UserService)
	if !ok {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	err = us.Activate(code, &userActivate)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Infof("User account with activation code %d, has been activated", code)
	ctx.JSON(200, "OK")
}

func (u *userHandler) ResetPassword(ctx *gin.Context) {
	logger.Debug("Reset user password")
	code := ctx.Param("resetCode")
	if len(code) != 32 {
		HandleError(e.ApiErrInvalidActivationCode, ctx)
		return
	}

	var userActivate mua.UserActivate
	err := ctx.BindJSON(&userActivate)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	us, ok := u.svc.(impl.UserService)
	if !ok {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	err = us.ResetPassword(code, &userActivate)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Infof("User with activation code %d, has password reset", code)
	ctx.JSON(200, "OK")
}

func (u *userHandler) ForgotPassword(ctx *gin.Context) {
	email := ctx.Param("email")
	logger.Debugf("Forgotten user password for email: %s", email)

	us, ok := u.svc.(impl.UserService)
	if !ok {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	err := us.ForgotPassword(email)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Infof("Link for password reset has been sent to: %s", email)
	ctx.JSON(200, "OK")
}

func (u *userHandler) SendActivationLink(ctx *gin.Context) {
	cId := ctx.Param("customerAccountId")
	logger.Debugf("Send activation link to customer with ID: %s", cId)

	id, err := strconv.Atoi(cId)
	if err != nil {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	us, ok := u.svc.(impl.UserService)
	if !ok {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	err = us.SendActivationLink(id)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Infof("Activation link sent to customer with ID: %s", cId)
	ctx.JSON(200, "OK")
}

func (u *userHandler) List(ctx *gin.Context) {
	logger.Debug("Listing users")
	models, err := u.svc.List()
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, models)
}

func (u *userHandler) Query(ctx *gin.Context) {
	logger.Debug("Querying users")

	query := api_utils.ParseQuery(u.TableName(), ctx.Request.RequestURI)
	count, models, err := u.svc.Query(mu.User{}, query)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, wrapQueryResult(count, models))
}

func (u *userHandler) ListSuperAdmins(ctx *gin.Context) {
	logger.Debug("Listing super admins")
	query := api_utils.ParseQuery(u.TableName(), ctx.Request.RequestURI)

	us, ok := u.svc.(impl.UserService)
	if !ok {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	count, models, err := us.ListSuperAdministrators(mu.User{RoleID: int(enum.SUPER_ADMIN)}, query)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, wrapQueryResult(count, models))
}

func (u *userHandler) ListAdministrators(ctx *gin.Context) {
	sId := ctx.Param("providerId")
	logger.Debugf("Listing providers: %s administrators", sId)

	pId, err := strconv.Atoi(sId)
	if err != nil {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	us, ok := u.svc.(impl.UserService)
	if !ok {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	query := api_utils.ParseQuery(u.TableName(), ctx.Request.RequestURI)
	count, models, err := us.ListAdministrators(mu.User{ProviderID: pId, RoleID: int(enum.ADMINISTRATOR_FULL)}, query)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, wrapQueryResult(count, models))
}

func (u *userHandler) Create(ctx *gin.Context) {
	logger.Debug("Creating user")
	var newUser mu.User
	err := ctx.BindJSON(&newUser)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	if !userOperationPermission(ctx, enum.Role(newUser.RoleID).Name()) {
		HandleError(e.ApiErrRoleTooLow, ctx)
		return
	}

	var createdUser model.Model
	createdUser, err = u.svc.Create(&newUser)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Infof("User: %s, has been created", *newUser.Login)
	ctx.JSON(200, createdUser)
}

func (u *userHandler) CheckIfExistWithFilter(ctx *gin.Context) {
	logger.Debug("Check if service access point with given parameters already exist")

	query := api_utils.ParseQuery(u.TableName(), ctx.Request.RequestURI)
	count, err := u.svc.Check(query)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, checkResult(count))
}

func (u *userHandler) GetByID(ctx *gin.Context) {
	sId := ctx.Param("userId")
	logger.Debugf("Get user by ID: %s", sId)

	id, err := strconv.Atoi(sId)
	if err != nil {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	model, err := u.svc.GetByID(id)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, model)
}

func (u *userHandler) GetDetails(ctx *gin.Context) {
	userId, err := api_utils.GetTokenUserID(ctx)
	if err != nil {
		HandleError(err, ctx)
		return
	}
	logger.Debugf("Get user: %d details", userId)

	model, err := u.svc.GetByID(userId)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, model)
}

func (u *userHandler) UpdateByID(ctx *gin.Context) {
	sId := ctx.Param("userId")
	logger.Debugf("Update user by ID: %s", sId)

	id, err := strconv.Atoi(sId)
	if err != nil {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	var updateUser mu.User
	err = ctx.BindJSON(&updateUser)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	model, err := u.svc.UpdateByID(id, &updateUser)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Infof("User of id %d, has been updated", sId)
	ctx.JSON(200, model)
}

func (u *userHandler) DeleteByID(ctx *gin.Context) {
	sId := ctx.Param("userId")
	logger.Debugf("Delete user by ID: %s", sId)

	id, err := strconv.Atoi(sId)
	if err != nil {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	err = u.svc.DeleteByID(id)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Infof("User of id %d, has been deleted", sId)
	ctx.JSON(200, "OK")
}
