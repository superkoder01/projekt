package auth

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
	"time"
)

const (
	privTestKeyFilePath = "../../tests/data/key/ed25519key.pem"
)

var (
	mock       redismock.ClientMock
	rdb        *redis.Client
	authConfig *AuthConfig
)

func TestMain(m *testing.M) {
	rdb, mock = redismock.NewClientMock()

	authConfig = &AuthConfig{
		KeyFilePath:                privTestKeyFilePath,
		RedisClient:                rdb,
		AccessTokenExpirationTime:  time.Minute * time.Duration(1),
		RefreshTokenExpirationTime: time.Minute * time.Duration(30),
	}
	// Run test cases
	code := m.Run()

	os.Exit(code)
}

func TestCorrectPassword(t *testing.T) {
	auth := NewAuth(authConfig)
	defer clearMock()
	testPass := "ovoo"

	hash, err := auth.Encode(testPass)
	assert.Nil(t, err)

	err = auth.ComparePassword(hash, testPass)
	assert.Nil(t, err)
}

func TestIncorrectPassword(t *testing.T) {
	auth := NewAuth(authConfig)
	defer clearMock()
	testPass := "pass123!"
	testIncorrectPass := "pass321!"

	hash, err := auth.Encode(testPass)
	assert.Nil(t, err)

	err = auth.ComparePassword(hash, testIncorrectPass)
	assert.Error(t, err)
}

func TestToken(t *testing.T) {
	auth := NewAuth(authConfig)
	defer clearMock()
	userId := 1
	role := "SUPER_ADMIN"
	login := "login"

	mock.Regexp().ExpectSetEX(`.*`, ACCESS, time.Second*60).SetVal("")
	mock.Regexp().ExpectSetEX(`.*`, REFRESH, time.Second*1800).SetVal("")
	tokenPair, err := auth.CreateToken(role, 1, userId, 0, 0, login)
	assert.Nil(t, err)
	assert.NotNil(t, tokenPair)

	tokenId := tokenPair.AccessToken.ID
	tokenString := tokenPair.AccessToken.Token

	mock.ExpectGet(tokenId.String()).SetVal(ACCESS)
	tokenInput, err := auth.VerifyToken(tokenString, ACCESS)
	assert.Nil(t, err)
	assert.NotNil(t, tokenInput)

	assert.Equal(t, role, tokenInput.Role)
	assert.Equal(t, userId, tokenInput.UserID)
	assert.Equal(t, 1, tokenInput.ProviderID)
}

func TestMalformedToken(t *testing.T) {
	auth := NewAuth(authConfig)
	defer clearMock()
	userId := 1
	role := "AGENT"
	malformedRole := "SUPER_ADMIN"
	login := "login"

	mock.Regexp().ExpectSetEX(`.*`, ACCESS, time.Second*60).SetVal("")
	mock.Regexp().ExpectSetEX(`.*`, REFRESH, time.Second*1800).SetVal("")
	tokenPair, err := auth.CreateToken(role, 1, userId, 0, 0, login)
	assert.Nil(t, err)
	assert.NotNil(t, tokenPair)

	// create another token to get second segment containing claims
	mock.Regexp().ExpectSetEX(`.*`, ACCESS, time.Second*60).SetVal("")
	mock.Regexp().ExpectSetEX(`.*`, REFRESH, time.Second*1800).SetVal("")
	anotherTokenPair, err := auth.CreateToken(malformedRole, 1, userId, 0, 0, login)
	assert.Nil(t, err)
	assert.NotNil(t, anotherTokenPair)

	tokenString := tokenPair.AccessToken.Token
	anotherTokenString := anotherTokenPair.AccessToken.Token
	anotherTokenId := anotherTokenPair.AccessToken.ID

	mock.ExpectGet(anotherTokenId.String()).SetVal(ACCESS)
	anotherInput, err := auth.VerifyToken(anotherTokenString, ACCESS)
	assert.Nil(t, err)
	assert.NotNil(t, anotherInput)

	segments := strings.Split(tokenString, ".")
	anotherSegments := strings.Split(anotherTokenString, ".")

	// create token with changed role
	malformedTokenString := segments[0] + "." + anotherSegments[1] + "." + segments[2]

	mock.ExpectGet(anotherTokenId.String()).RedisNil()
	tokenInput, err := auth.VerifyToken(malformedTokenString, ACCESS)
	assert.Error(t, err)
	assert.Nil(t, tokenInput)
}

func TestRandomCode(t *testing.T) {
	auth := NewAuth(authConfig)
	defer clearMock()
	randomCode := auth.GenerateRandomCode()

	assert.True(t, len(randomCode) == 32)
}

func TestTokenCache(t *testing.T) {
	auth := NewAuth(authConfig)
	defer clearMock()
	userId := 1
	role := "SUPER_ADMIN"
	login := "login"

	mock.Regexp().ExpectSetEX(`.*`, ACCESS, time.Second*60).SetVal("")
	mock.Regexp().ExpectSetEX(`.*`, REFRESH, time.Second*1800).SetVal("")
	tokenPair, err := auth.CreateToken(role, 1, userId, 0, 0, login)
	assert.Nil(t, err)
	assert.NotNil(t, tokenPair)
	assert.NotNil(t, tokenPair.AccessToken)
	assert.NotNil(t, tokenPair.RefreshToken)

	mock.ExpectSetEX(tokenPair.AccessToken.ID.String(), ACCESS, time.Second*60).RedisNil()
	err = auth.saveAccessToken(tokenPair.AccessToken.ID.String())
	assert.Equal(t, err.Error(), "redis: nil")

	mock.ExpectSetEX(tokenPair.RefreshToken.ID.String(), REFRESH, time.Second*1800).RedisNil()
	err = auth.saveRefreshToken(tokenPair.RefreshToken.ID.String())
	assert.Equal(t, err.Error(), "redis: nil")
}

func clearMock() {
	mock.ClearExpect()
}
