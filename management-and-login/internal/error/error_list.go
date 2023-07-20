package error

import "net/http"

var (
	// API errors
	ApiErrWrongProvider              = New("wrong provider", http.StatusUnauthorized, API)
	ApiErrRoleTooLow                 = New("too low level of permissions", http.StatusUnauthorized, API)
	ApiErrNoAuthorizationHeader      = New("'Authorization' header absence", http.StatusBadRequest, API)
	ApiErrAuthorizationHeaderInvalid = New("authorization header value is invalid", http.StatusBadRequest, API)
	ApiErrInvalidDataModel           = New("invalid entry data", http.StatusBadRequest, API)
	ApiErrInvalidActivationCode      = New("invalid activation code", http.StatusBadRequest, API)
	ApiErrEndpointForbidden          = New("resource fetch forbidden", http.StatusForbidden, API)
	ApiErrImproperQuery              = New("improper api query", http.StatusBadRequest, API)
	ApiErrInvalidEmail               = New("invalid email address", http.StatusBadRequest, API)
	ApiErrNotEnoughData              = New("not enough data provided in request", http.StatusBadRequest, API)

	// COMMON errors
	ErrInternalServerError = New("internal server error", http.StatusInternalServerError, UNKNOWN)

	// Blockchain errors
	BlockchainErrAccountCreation = New("blockchain account creation failed", http.StatusInternalServerError, BLOCKCHAIN)
	BlockchainErrAccountInvalid  = New("blockchain account creation failed, invalid data received", http.StatusInternalServerError, BLOCKCHAIN)

	// DB errors
	DbErrEntityDeleteHasChild = New("entity has child relation that cannot be deleted", http.StatusForbidden, DATABASE)
	DbErrEntityNotFound       = New("entity not found", http.StatusNotFound, DATABASE)
	DbErrPasswordMismatch     = New("incorrect password", http.StatusUnauthorized, DATABASE)
	DbErrUserNotActivated     = New("user account not activated", http.StatusUnauthorized, DATABASE)
)
