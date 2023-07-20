package tests

import (
	"github.com/stretchr/testify/assert"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/auth"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/enum"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/user"
	ath "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/auth"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service/impl"
	"testing"
)

func TestAuthenticateServiceOk(t *testing.T) {
	mariaDBFlushData(s)

	login := "jkowalski"
	password := "kowal1990!"

	userModel := &user.User{
		Login:      &login,
		Password:   &password,
		RoleID:     int(enum.ADMINISTRATOR_BASIC),
		ProviderID: 1,
		Active:     true,
		Email:      "email",
	}

	userService := sf.New(service.USER)
	userCreated, err := userService.Create(userModel)
	assert.Nil(t, err)
	assert.NotNil(t, userCreated)

	authModel := &auth.Auth{
		Login:    login,
		Password: password,
	}

	authService := sf.New(service.AUTH)
	authServiceImpl, ok := authService.(impl.AuthService)
	assert.True(t, ok)

	tokenPair, err := authServiceImpl.Authenticate(authModel)
	assert.Nil(t, err)
	assert.NotNil(t, tokenPair)

	authO := ath.NewAuth(authConfig)

	tokenRole, err := authO.GetVerifiedRole(tokenPair.AccessToken.Token)
	assert.Nil(t, err)
	assert.NotNil(t, tokenRole)
	assert.Equal(t, enum.ADMINISTRATOR_BASIC.Name(), tokenRole)

	tokenProvId, err := authO.GetVerifiedProviderID(tokenPair.AccessToken.Token)
	assert.Nil(t, err)
	assert.NotNil(t, 1, tokenProvId)
}

func TestAuthenticateServiceUserNotFound(t *testing.T) {
	mariaDBFlushData(s)

	login := "jkowalski"
	password := "kowal1990!"
	loginNotExist := "jkowaskli"

	userModel := &user.User{
		Login:      &login,
		Password:   &password,
		RoleID:     int(enum.ADMINISTRATOR_BASIC),
		ProviderID: 1,
		Email:      "test@email.com",
	}

	userService := sf.New(service.USER)
	userCreated, err := userService.Create(userModel)
	assert.Nil(t, err)
	assert.NotNil(t, userCreated)

	authModel := &auth.Auth{
		Login:    loginNotExist,
		Password: password,
	}

	authService := sf.New(service.AUTH)
	authServiceImpl, ok := authService.(impl.AuthService)
	assert.True(t, ok)

	signedToken, err := authServiceImpl.Authenticate(authModel)
	assert.Error(t, err)
	assert.Nil(t, signedToken)
}

func TestAuthenticateServiceIncorrectPassword(t *testing.T) {
	mariaDBFlushData(s)

	login := "jkowalski"
	password := "kowal1990!"
	passwordIncorrect := "kowal1991!@"

	userModel := &user.User{
		Login:      &login,
		Password:   &password,
		RoleID:     int(enum.ADMINISTRATOR_BASIC),
		ProviderID: 1,
		Email:      "test@email.com",
	}

	userService := sf.New(service.USER)
	userCreated, err := userService.Create(userModel)
	assert.Nil(t, err)
	assert.NotNil(t, userCreated)

	authModel := &auth.Auth{
		Login:    login,
		Password: passwordIncorrect,
	}

	authService := sf.New(service.AUTH)
	authServiceImpl, ok := authService.(impl.AuthService)
	assert.True(t, ok)

	signedToken, err := authServiceImpl.Authenticate(authModel)
	assert.Error(t, err)
	assert.Nil(t, signedToken)
}
