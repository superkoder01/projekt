package auth

import (
	"crypto/ed25519"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"strings"
)

func verifyToken(token string, key ed25519.PublicKey) (*tokenInput, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodEd25519)
		if !ok {
			return nil, ErrInvalidSigningMethod
		}
		return key, nil
	}

	parsed, err := jwt.ParseWithClaims(token, &tokenInput{}, keyFunc)
	if err != nil {
		return nil, ErrInvalidToken
	}

	ti, ok := parsed.Claims.(*tokenInput)
	if !ok {
		return nil, ErrInvalidTokenFormat
	}
	return ti, nil
}

func getClaims(token string) (*tokenInput, error) {
	segments := strings.Split(token, ".")
	if len(segments) != 3 {
		return nil, ErrInvalidToken
	}

	decoded, err := jwt.DecodeSegment(segments[1])
	if err != nil {
		return nil, ErrInvalidToken
	}

	var ti tokenInput
	err = json.Unmarshal(decoded, &ti)
	if err != nil {
		return nil, ErrInvalidToken
	}

	return &ti, nil
}
