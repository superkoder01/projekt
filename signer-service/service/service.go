package service

import (
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/billing_management/contract"
)

type Service interface {
	InitSign(*gin.Context, contract.Contract) (string, error)
	SigningCompletedNotification(*gin.Context) error
}
