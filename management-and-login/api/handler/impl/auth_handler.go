package impl

import (
	"github.com/gin-gonic/gin"
	ma "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/auth"
	ml "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/logout"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service/impl"
)

const authorization = "Authorization"

type AuthHandler interface {
	Authenticate(ctx *gin.Context)
	Logout(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
}

type authHandler struct {
	svc service.Service
}

func (l *authHandler) TableName() string {
	return ""
}

func NewAuthHandler(svc service.Service) *authHandler {
	return &authHandler{svc: svc}
}

func (l *authHandler) Authenticate(ctx *gin.Context) {
	var authUser ma.Auth
	err := ctx.BindJSON(&authUser)

	if err != nil {
		HandleError(err, ctx)
		return
	}
	logger.Debugf("User: %s is trying to authenticate", authUser.Login)

	as, ok := l.svc.(impl.AuthService)
	if !ok {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	token, err := as.Authenticate(&authUser)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	logger.Debugf("User: %s authenticated", authUser.Login)
	ctx.JSON(200, token)
}

func (l *authHandler) Logout(ctx *gin.Context) {
	logger.Debug("Logout attempt")
	var logout ml.Logout
	err := ctx.BindJSON(&logout)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	as, ok := l.svc.(impl.AuthService)
	if !ok {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	err = as.Logout(&logout)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.String(200, "OK")
}

func (l *authHandler) RefreshToken(ctx *gin.Context) {
	logger.Debug("Refresh token attempt")

	token, err := api_utils.GetToken(ctx)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	as, ok := l.svc.(impl.AuthService)
	if !ok {
		HandleError(e.ErrInternalServerError, ctx)
		return
	}

	tokenPair, err := as.RefreshToken(token)
	if err != nil {
		HandleError(err, ctx)
		return
	}

	ctx.JSON(200, tokenPair)
}

// NOT USED
func (l *authHandler) Query(ctx *gin.Context) {
}

// NOT USED
func (l *authHandler) GetDetails(ctx *gin.Context) {
}

// NOT USED
func (l *authHandler) CheckIfExistWithFilter(ctx *gin.Context) {
}

// NOT USED
func (l *authHandler) GetByID(ctx *gin.Context) {
}

// NOT USED
func (l *authHandler) List(ctx *gin.Context) {
}

// NOT USED
func (l *authHandler) Create(ctx *gin.Context) {
}

// NOT USED
func (l *authHandler) UpdateByID(ctx *gin.Context) {
}

// NOT USED
func (l *authHandler) DeleteByID(ctx *gin.Context) {
}
