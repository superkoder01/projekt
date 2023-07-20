package auth

import "errors"

var (
	ErrInvalidToken            = errors.New("token is invalid")
	ErrInvalidTokenFormat      = errors.New("token does not match expected format")
	ErrExpiredToken            = errors.New("token has expired")
	ErrBeforeTimeToken         = errors.New("too early to use the token")
	ErrInvalidTokenIssuer      = errors.New("token issuer is incorrect")
	ErrInvalidSigningMethod    = errors.New("token is not signed with Ed25519 key")
	ErrNotEd25519PublicKey     = errors.New("public key is not of type Ed25519")
	ErrRedisSet                = errors.New("redis set error")
	ErrRedisKeyNotFound        = errors.New("token has expired or is invalid")
	ErrRedisWrongTokenTypeUsed = errors.New("wrong token type used")
	ErrRedisDel                = errors.New("redis deletion error")
)
