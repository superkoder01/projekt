package api_utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	er "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	a "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/auth"
	"strings"
)

const (
	authorization = "Authorization"
	bearer        = "Bearer"

	space        = " "
	or           = " or "
	and          = " and "
	leftBracket  = "("
	rightBracket = ")"
	wildcard     = "%"
	equal        = "="
	like         = "like"
)

func GetTokenUserID(ctx *gin.Context) (int, error) {
	auth := a.NewAuth(nil)

	authValue := ctx.GetHeader(authorization)
	token, err := GetTokenFromAuthHeader(authValue)
	if err != nil {
		return 0, err
	}

	userId, err := auth.GetVerifiedUserID(token)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func GetTokenProviderID(ctx *gin.Context) (int, error) {
	auth := a.NewAuth(nil)

	authValue := ctx.GetHeader(authorization)
	token, err := GetTokenFromAuthHeader(authValue)
	if err != nil {
		return 0, err
	}

	providerId, err := auth.GetVerifiedProviderID(token)
	if err != nil {
		return 0, err
	}

	return providerId, nil
}

func GetTokenRole(ctx *gin.Context) (string, error) {
	auth := a.NewAuth(nil)

	authValue := ctx.GetHeader(authorization)
	token, err := GetTokenFromAuthHeader(authValue)
	if err != nil {
		return "", err
	}

	role, err := auth.GetVerifiedRole(token)
	if err != nil {
		return "", err
	}

	return role, nil
}

func GetTokenCustomerAccountID(ctx *gin.Context) (int, error) {
	auth := a.NewAuth(nil)

	authValue := ctx.GetHeader(authorization)
	token, err := GetTokenFromAuthHeader(authValue)
	if err != nil {
		return 0, err
	}

	customerAccountId, err := auth.GetVerifiedCustomerAccountID(token)
	if err != nil {
		return 0, err
	}

	return customerAccountId, nil
}

func GetTokenWorkerID(ctx *gin.Context) (int, error) {
	auth := a.NewAuth(nil)

	authValue := ctx.GetHeader(authorization)
	token, err := GetTokenFromAuthHeader(authValue)
	if err != nil {
		return 0, err
	}

	workerId, err := auth.GetVerifiedWorkerID(token)
	if err != nil {
		return 0, err
	}

	return workerId, nil
}

func GetTokenUUID(ctx *gin.Context) (*uuid.UUID, error) {
	auth := a.NewAuth(nil)

	authValue := ctx.GetHeader(authorization)
	token, err := GetTokenFromAuthHeader(authValue)
	if err != nil {
		return nil, err
	}

	uuid, err := auth.GetVerifiedUUID(token)
	if err != nil {
		return nil, err
	}

	return uuid, nil
}

func GetToken(ctx *gin.Context) (string, error) {
	authValue := ctx.GetHeader(authorization)
	token, err := GetTokenFromAuthHeader(authValue)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetTokenFromAuthHeader(header string) (string, error) {
	header = strings.TrimSpace(header)
	if header == "" {
		return "", er.ApiErrNoAuthorizationHeader
	}

	headerContent := strings.Split(header, space)
	if len(headerContent) != 2 || headerContent[0] != bearer {
		return "", er.ApiErrAuthorizationHeaderInvalid
	}

	return headerContent[1], nil
}
