package impl

import (
	"github.com/gin-gonic/gin"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/service"
)

type parameterNameHandler struct {
	svc service.Service
}

func NewParameterNameHandler(svc service.Service) *parameterNameHandler {
	return &parameterNameHandler{svc: svc}
}

func (c *parameterNameHandler) List(ctx *gin.Context) {
	models, err := c.svc.List()
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, models)
}

func (c *parameterNameHandler) Create(ctx *gin.Context) {
}

func (c *parameterNameHandler) GetWithFilter(ctx *gin.Context) {
}

func (c *parameterNameHandler) GetByID(ctx *gin.Context) {
}

func (c *parameterNameHandler) UpdateByID(ctx *gin.Context) {
}

func (c *parameterNameHandler) DeleteByID(ctx *gin.Context) {
}

func (c *parameterNameHandler) Get(ctx *gin.Context) {
}

func (c *parameterNameHandler) Update(ctx *gin.Context) {
}

func (c *parameterNameHandler) Delete(ctx *gin.Context) {
}
