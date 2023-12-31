package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Error struct {
	Code    int
	Message string
	Cause   error
}

func (e *Error) Error() string {
	if e.Cause != nil {
		return e.Cause.Error()
	}
	return e.Message
}

func NewError(msg string, code int) error {
	return &Error{
		Message: msg,
		Code:    code,
	}
}

func HandleError(e error, ctx *gin.Context) {
	err, ok := e.(*Error)
	if ok {
		ctx.String(err.Code, err.Error())
	} else {
		ctx.String(http.StatusInternalServerError, e.Error())
	}
	ctx.Abort()
}
