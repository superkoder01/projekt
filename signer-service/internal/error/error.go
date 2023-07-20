package error

import (
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"net/http"
)

var (
	logger = logging.MustGetLogger("error")
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

func New(msg string, code int) error {
	return &Error{
		Message: msg,
		Code:    code,
	}
}

func Wrap(e error, code int) error {
	return &Error{
		Cause: e,
		Code:  code,
	}
}

func HandleError(e error, ctx *gin.Context) {
	err, ok := e.(*Error)
	if ok {
		ctx.String(err.Code, err.Error())
	} else {
		ctx.String(http.StatusInternalServerError, e.Error())
	}
	logger.Errorf("Handle Error: %s", err)
	ctx.Abort()
}
