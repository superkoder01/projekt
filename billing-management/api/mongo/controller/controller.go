package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	List(ctx *gin.Context)
	FindOne(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
}
