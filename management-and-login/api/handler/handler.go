package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	CheckIfExistWithFilter(ctx *gin.Context)
	GetByID(ctx *gin.Context)

	List(ctx *gin.Context)

	Create(ctx *gin.Context)

	//Update(ctx *gin.Context)
	UpdateByID(ctx *gin.Context)

	//Delete(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)

	GetDetails(ctx *gin.Context)

	Query(ctx *gin.Context)

	TableName() string
}
