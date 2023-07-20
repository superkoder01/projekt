package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	InitSign(ctx *gin.Context)
	SigningCompletedNotification(ctx *gin.Context)
}
