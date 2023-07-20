package auth

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

type Authenticator interface {
	Encode(password string) (string, error)
	ComparePassword(hash string, pass string) error
	GenerateRandomCode() string
	CreateToken(role string, providerId int, userId int, customerAccountId int, workerId int, login string) (*TokensPair, error)
	VerifyToken(token string, kind string) (*tokenInput, error)
	GetVerifiedRole(token string) (string, error)
	GetVerifiedProviderID(token string) (int, error)
	GetVerifiedUserID(token string) (int, error)
	GetVerifiedCustomerAccountID(token string) (int, error)
	GetVerifiedWorkerID(token string) (int, error)
	GetVerifiedUUID(token string) (*uuid.UUID, error)
	RemoveToken(tokenID string) error
}

type auth struct {
	Encoder                    Encoder
	KeyLoader                  KeyLoader
	Redis                      *redis.Client
	AccessTokenExpirationTime  time.Duration
	RefreshTokenExpirationTime time.Duration
}

type AuthConfig struct {
	KeyFilePath                string
	RedisClient                *redis.Client
	AccessTokenExpirationTime  time.Duration
	RefreshTokenExpirationTime time.Duration
}

// NewAuth - new 'Authenticator' instance
func NewAuth(a *AuthConfig) *auth {
	if a != nil {
		return &auth{
			Encoder:                    newEncoder(),
			KeyLoader:                  newKeyLoader(a.KeyFilePath),
			Redis:                      a.RedisClient,
			AccessTokenExpirationTime:  a.AccessTokenExpirationTime,
			RefreshTokenExpirationTime: a.RefreshTokenExpirationTime,
		}
	}
	return nil
}

// Encode - returns hashed password
func (a *auth) Encode(pass string) (string, error) {
	return a.Encoder.Encode(pass)
}

// ComparePassword - validates given hash against password
func (a *auth) ComparePassword(hash string, pass string) error {
	return a.Encoder.Compare(hash, pass)
}

// CreateToken - returns signed JWT token with standard claims + 'role'
func (a *auth) CreateToken(role string, providerId int, userId int, customerAccountId int, workerId int, login string) (*TokensPair, error) {
	priv, err := a.KeyLoader.GetPrivateKey()
	if err != nil {
		return nil, err
	}
	token := a.newTokenInput(role, providerId, userId, customerAccountId, workerId, login)

	pair, err := token.NewPair(priv, a.RefreshTokenExpirationTime)
	if err != nil {
		return nil, err
	}

	err = a.saveAccessToken(pair.AccessToken.ID.String())
	if err != nil {
		return nil, err
	}

	err = a.saveRefreshToken(pair.RefreshToken.ID.String())
	if err != nil {
		return nil, err
	}

	return pair, nil
}

// VerifyToken - verifies token signature
func (a *auth) VerifyToken(token string, kind string) (*tokenInput, error) {
	pub, err := a.KeyLoader.GetPublicKey()
	if err != nil {
		return nil, err
	}

	ti, err := verifyToken(token, pub)
	if err != nil {
		return nil, err
	}

	return ti, a.getToken(ti.UUID.String(), kind)
}

// GetVerifiedRole - returns verified role (string)
func (a *auth) GetVerifiedRole(token string) (string, error) {
	ti, err := getClaims(token)
	if err != nil {
		return "", err
	}

	return ti.Role, nil
}

// GetVerifiedProviderID - returns verified provider id (int)
func (a *auth) GetVerifiedProviderID(token string) (int, error) {
	ti, err := getClaims(token)
	if err != nil {
		return 0, err
	}

	return ti.ProviderID, nil
}

// GetVerifiedUserID - returns verified user id (int)
func (a *auth) GetVerifiedUserID(token string) (int, error) {
	ti, err := getClaims(token)
	if err != nil {
		return 0, err
	}

	return ti.UserID, nil
}

// GetVerifiedCustomerAccountID - returns verified customer account id (int)
func (a *auth) GetVerifiedCustomerAccountID(token string) (int, error) {
	ti, err := getClaims(token)
	if err != nil {
		return 0, err
	}

	return ti.CustomerAccountID, nil
}

// GetVerifiedWorkerID - returns verified worker id (int)
func (a *auth) GetVerifiedWorkerID(token string) (int, error) {
	ti, err := getClaims(token)
	if err != nil {
		return 0, err
	}

	return ti.WorkerID, nil
}

// GetVerifiedUUID - returns verified token uuid (int)
func (a *auth) GetVerifiedUUID(token string) (*uuid.UUID, error) {
	ti, err := getClaims(token)
	if err != nil {
		return nil, err
	}

	return &ti.UUID, nil
}

// RemoveToken - deletes key in redis database
func (a *auth) RemoveToken(tokenID string) error {
	return a.removeToken(tokenID)
}

// GenerateRandomCode - returns random 32 char string built out of letters and digits
func (a *auth) GenerateRandomCode() string {
	return randStringRunes(32)
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
