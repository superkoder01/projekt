package auth

import (
	"crypto/ed25519"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

const (
	issuer = "c4e"
)

type Token interface {
	NewPair() (*TokensPair, error)
	Valid() error
}

type TokensPair struct {
	AccessToken  *AccessToken  `json:"access_token"`
	RefreshToken *RefreshToken `json:"refresh_token"`
}

type AccessToken struct {
	ID    uuid.UUID `json:"id,omitempty"`
	Token string    `json:"token,omitempty"`
}

type RefreshToken struct {
	ID    uuid.UUID `json:"id,omitempty"`
	Token string    `json:"token,omitempty"`
}

type tokenInput struct {
	Claims            jwt.StandardClaims
	Role              string    `json:"role,omitempty"`
	ProviderID        int       `json:"providerId,omitempty"`
	UserID            int       `json:"userId,omitempty"`
	CustomerAccountID int       `json:"customerAccountId,omitempty"`
	WorkerID          int       `json:"workerId,omitempty"`
	UUID              uuid.UUID `json:"uuid,omitempty"`
}

type refreshTokenInput struct {
	Claims jwt.StandardClaims
	UserID int       `json:"userId,omitempty"`
	UUID   uuid.UUID `json:"uuid,omitempty"`
}

func (a *auth) newTokenInput(role string, providerId int, userId int, customerAccountId int, workerId int, login string) *tokenInput {
	return &tokenInput{
		Claims: jwt.StandardClaims{
			ExpiresAt: time.Now().UnixMilli() + a.AccessTokenExpirationTime.Milliseconds(),
			Issuer:    issuer,
			NotBefore: time.Now().UnixMilli(),
			Subject:   login,
		},
		Role:              role,
		ProviderID:        providerId,
		UserID:            userId,
		CustomerAccountID: customerAccountId,
		WorkerID:          workerId,
	}
}

func (t *tokenInput) NewPair(key ed25519.PrivateKey, exp time.Duration) (*TokensPair, error) {
	// Access token
	accessId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	t.UUID = accessId
	aTk := jwt.NewWithClaims(jwt.SigningMethodEdDSA, t)
	accessToken, err := aTk.SignedString(key)
	if err != nil {
		return nil, err
	}

	// Refresh token
	refreshId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	rTk := jwt.NewWithClaims(jwt.SigningMethodEdDSA, &refreshTokenInput{
		Claims: jwt.StandardClaims{
			ExpiresAt: time.Now().UnixMilli() + exp.Milliseconds(),
			Issuer:    issuer,
			NotBefore: time.Now().UnixMilli(),
		},
		UserID: t.UserID,
		UUID:   refreshId,
	})
	refreshToken, err := rTk.SignedString(key)
	if err != nil {
		return nil, err
	}

	return &TokensPair{
		AccessToken: &AccessToken{
			ID:    accessId,
			Token: accessToken,
		},
		RefreshToken: &RefreshToken{
			ID:    refreshId,
			Token: refreshToken,
		},
	}, nil
}

func (t *tokenInput) Valid() error {
	expiration := time.UnixMilli(t.Claims.ExpiresAt)
	if time.Now().After(expiration) {
		return ErrExpiredToken
	}

	before := time.UnixMilli(t.Claims.NotBefore)
	if time.Now().Before(before) {
		return ErrBeforeTimeToken
	}

	if t.Claims.Issuer != issuer {
		return ErrInvalidTokenIssuer
	}
	return nil
}

func (t *refreshTokenInput) Valid() error {
	expiration := time.UnixMilli(t.Claims.ExpiresAt)
	if time.Now().After(expiration) {
		return ErrExpiredToken
	}

	before := time.UnixMilli(t.Claims.NotBefore)
	if time.Now().Before(before) {
		return ErrBeforeTimeToken
	}

	if t.Claims.Issuer != issuer {
		return ErrInvalidTokenIssuer
	}
	return nil
}
