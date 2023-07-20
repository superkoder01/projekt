package error

import (
	"net/http"
	"github.com/op/go-logging"

	"github.com/labstack/echo/v4"
)

var (
	// API errors
	ApiErrWrongProvider              = New("wrong provider", http.StatusUnauthorized)
	ApiErrRoleTooLow                 = New("too low level of permissions", http.StatusUnauthorized)
	ApiErrNoAuthorizationHeader      = New("'Authorization' header absence", http.StatusBadRequest)
	ApiErrAuthorizationHeaderInvalid = New("authorization header value is invalid", http.StatusBadRequest)
	ApiErrInvalidDataModel           = New("invalid entry data", http.StatusBadRequest)
	ApiErrInvalidActivationCode      = New("invalid activation code", http.StatusBadRequest)
	ApiErrEndpointForbidden          = New("resource fetch forbidden", http.StatusForbidden)

	// COMMON errors
	ErrInternalServerError = New("internal server error", http.StatusInternalServerError)

	// DB errors
	DbErrEntityNotFound   = New("entity not found", http.StatusNotFound)
	DbErrPasswordMismatch = New("incorrect password", http.StatusUnauthorized)

	//Mongo errors
	FilteredFieldsAndFilteredValuesLengthsDoNotMatch = New("entity not found", http.StatusBadRequest)
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

func HandleError(e error, ctx echo.Context) {
	err, ok := e.(*Error)
	logger.Error("handle error: ", e)
	if ok {
		ctx.Error(echo.NewHTTPError(err.Code, err.Message))
	} else {
		ctx.Error(echo.NewHTTPError(http.StatusInternalServerError, err.Message))
	}
}
