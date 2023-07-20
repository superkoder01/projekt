package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetWithFilter(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Get(ctx *gin.Context)

	List(ctx *gin.Context)

	Create(ctx *gin.Context)

	UpdateByID(ctx *gin.Context)
	Update(ctx *gin.Context)

	DeleteByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
