package error

import "net/http"

var (
	// API errors
	ApiErrWrongProvider              = New("wrong provider", http.StatusUnauthorized)
	ApiErrRoleTooLow                 = New("too low level of permissions", http.StatusUnauthorized)
	ApiErrNoAuthorizationHeader      = New("'Authorization' header absence", http.StatusBadRequest)
	ApiErrAuthorizationHeaderInvalid = New("authorization header value is invalid", http.StatusBadRequest)
	ApiErrInvalidDataModel           = New("invalid entry data", http.StatusBadRequest)
	ApiErrInvalidActivationCode      = New("invalid activation code", http.StatusBadRequest)
	ApiErrEndpointForbidden          = New("resource fetch forbidden", http.StatusForbidden)
	ApiRequestParamsMismatch         = New("id param is different than id of body request", http.StatusBadRequest)

	// COMMON errors
	ErrInternalServerError = New("internal server error", http.StatusInternalServerError)

	// DB errors
	DbErrEntityNotFound   = New("entity not found", http.StatusNotFound)
	DbErrPasswordMismatch = New("incorrect password", http.StatusUnauthorized)

	//Mongo errors
	FilteredFieldsAndFilteredValuesLengthsDoNotMatch = New("entity not found", http.StatusBadRequest)

	// Domain Errors
	CannotUpdateWhenStateIsAccepted = New("Cannot update when state is accepted", http.StatusBadRequest)
)
