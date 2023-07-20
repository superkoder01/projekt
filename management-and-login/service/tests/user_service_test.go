package tests

import (
	"github.com/stretchr/testify/assert"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/customer_account"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/enum"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/user"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service/impl"
	"testing"
)

func TestUserServiceCreate(t *testing.T) {
	mariaDBFlushData(s)

	login := "login101"
	password := "pass123!"

	userModel := &user.User{
		ProviderID: 1,
		Login:      &login,
		Password:   &password,
		Email:      "test@email.com",
	}

	userService := sf.New(service.USER)

	model, err := userService.Create(userModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	userModelResp, ok := model.(*user.User)
	assert.True(t, ok)
	assert.NotNil(t, userModelResp)
	assert.Equal(t, login, *userModelResp.Login)
	assert.NotEqual(t, password, *userModelResp.Password)
}

func TestUserServiceDeleteByID(t *testing.T) {
	mariaDBFlushData(s)

	login := "login101"
	password := "pass123!"

	userModel := &user.User{
		ProviderID: 1,
		Login:      &login,
		Password:   &password,
		Email:      "email1",
	}

	login2 := "login102"
	password2 := "pass124!"

	userModel2 := &user.User{
		ProviderID: 1,
		Login:      &login2,
		Password:   &password2,
		Email:      "email2",
	}

	userService := sf.New(service.USER)

	model, err := userService.Create(userModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := userService.Create(userModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	userModelResp, ok := model.(*user.User)
	assert.True(t, ok)
	assert.NotNil(t, userModelResp)
	assert.Equal(t, login, *userModelResp.Login)
	assert.NotEqual(t, password, *userModelResp.Password)

	userModelResp2, ok := model2.(*user.User)
	assert.True(t, ok)
	assert.NotNil(t, userModelResp2)
	assert.Equal(t, login2, *userModelResp2.Login)
	assert.NotEqual(t, password, *userModelResp2.Password)

	err = userService.DeleteByID(userModelResp.ID)
	assert.Nil(t, err)

	getModel2, err := userService.GetByID(userModelResp2.ID)
	assert.Nil(t, err)
	assert.NotNil(t, getModel2)
	getModel1, err := userService.GetByID(userModelResp.ID)
	assert.Error(t, err)
	assert.Nil(t, getModel1)
}

func TestUserServiceGetByID(t *testing.T) {
	mariaDBFlushData(s)

	login := "login101"
	password := "pass123!"

	userModel := &user.User{
		ProviderID: 1,
		Login:      &login,
		Password:   &password,
		Email:      "email1",
	}

	login2 := "login102"
	password2 := "pass124!"

	userModel2 := &user.User{
		ProviderID: 1,
		Login:      &login2,
		Password:   &password2,
		Email:      "email2",
	}

	userService := sf.New(service.USER)

	model, err := userService.Create(userModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := userService.Create(userModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	userModelResp, ok := model.(*user.User)
	assert.True(t, ok)
	assert.NotNil(t, userModelResp)
	assert.Equal(t, login, *userModelResp.Login)
	assert.NotEqual(t, password, *userModelResp.Password)

	userModelResp2, ok := model2.(*user.User)
	assert.True(t, ok)
	assert.NotNil(t, userModelResp2)
	assert.Equal(t, login2, *userModelResp2.Login)
	assert.NotEqual(t, password, *userModelResp2.Password)

	getModel1, err := userService.GetByID(userModelResp.ID)
	assert.Nil(t, err)
	assert.NotNil(t, getModel1)

	userGetModelResp, ok := getModel1.(*user.User)
	assert.True(t, ok)
	assert.Equal(t, login, *userGetModelResp.Login)
	assert.NotEqual(t, password, *userGetModelResp.Password)

	getModel2, err := userService.GetByID(userModelResp2.ID)
	assert.Nil(t, err)
	assert.NotNil(t, getModel2)

	userGetModelResp2, ok := getModel2.(*user.User)
	assert.True(t, ok)
	assert.Equal(t, login2, *userGetModelResp2.Login)
	assert.NotEqual(t, password, *userGetModelResp2.Password)
}

func TestUserServiceGetWithFilter(t *testing.T) {
	mariaDBFlushData(s)

	login := "login101"
	password := "pass123!"

	userModel := &user.User{
		ProviderID: 1,
		Login:      &login,
		Password:   &password,
		Email:      "email1",
	}

	login2 := "login102"
	password2 := "pass124!"

	userModel2 := &user.User{
		ProviderID: 1,
		Login:      &login2,
		Password:   &password2,
		Email:      "email2",
	}

	userService := sf.New(service.USER)

	model, err := userService.Create(userModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := userService.Create(userModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	ens, err := userService.GetWithFilter("LOGIN = ?", login2)
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 1, len(ens))

	userModelResp, ok := ens[0].(*user.User)
	assert.True(t, ok)
	assert.Equal(t, login2, *userModelResp.Login)
	assert.NotEqual(t, password, *userModelResp.Password)
}

func TestUserServiceList(t *testing.T) {
	mariaDBFlushData(s)

	login := "login101"
	password := "pass123!"

	userModel := &user.User{
		ProviderID: 1,
		Login:      &login,
		Password:   &password,
		Email:      "email1",
	}

	login2 := "login102"
	password2 := "pass124!"

	userModel2 := &user.User{
		ProviderID: 1,
		Login:      &login2,
		Password:   &password2,
		Email:      "email2",
	}

	userService := sf.New(service.USER)

	model, err := userService.Create(userModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := userService.Create(userModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	ens, err := userService.List()
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 2, len(ens))
}

func TestUserServiceUpdateByID(t *testing.T) {
	mariaDBFlushData(s)

	login := "login101"
	password := "pass123!"
	roleId := enum.TRADER

	userModel := &user.User{
		ProviderID: 1,
		Login:      &login,
		Password:   &password,
		Email:      "email1",
	}

	login2 := "login102"
	password2 := "pass124!"

	userModel2 := &user.User{
		ProviderID: 1,
		Login:      &login2,
		Password:   &password2,
		Email:      "email2",
	}

	userService := sf.New(service.USER)

	model, err := userService.Create(userModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := userService.Create(userModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	userModelResp, ok := model.(*user.User)
	assert.True(t, ok)
	assert.NotNil(t, userModelResp)
	assert.Equal(t, login, *userModelResp.Login)
	assert.NotEqual(t, password, *userModelResp.Password)

	userModelResp2, ok := model2.(*user.User)
	assert.True(t, ok)
	assert.NotNil(t, userModelResp2)
	assert.Equal(t, login2, *userModelResp2.Login)
	assert.NotEqual(t, password, *userModelResp2.Password)

	updateModel, err := userService.UpdateByID(userModelResp.ID, &user.User{
		Login:    &login,
		Password: &password2,
		RoleID:   int(roleId),
	})
	assert.Nil(t, err)
	assert.NotNil(t, updateModel)

	getModel, err := userService.GetByID(userModelResp.ID)
	assert.Nil(t, err)
	assert.NotNil(t, getModel)

	userGetModelResp, ok := getModel.(*user.User)
	assert.True(t, ok)
	assert.Equal(t, login, *userGetModelResp.Login)
	assert.NotEqual(t, password, *userGetModelResp.Password)
	assert.Equal(t, int(roleId), userGetModelResp.RoleID)
}

func TestUserServiceSendActivationLink(t *testing.T) {
	mariaDBFlushData(s)

	firstName := "Jan"
	lastName := "Kowalski"

	customerAccountModel := &customer_account.CustomerAccount{
		ProviderID:       1,
		CustomerTypeName: "PROSUMER",
		FirstName:        firstName,
		LastName:         lastName,
	}

	customerAccountService := sf.New(service.CUSTOMER_ACCOUNT)

	model, err := customerAccountService.Create(customerAccountModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	customerAccountModelResp, ok := model.(*customer_account.CustomerAccount)

	login := "login101"
	password := "pass123!"

	userModel := &user.User{
		ProviderID:        1,
		Login:             &login,
		Password:          &password,
		Email:             "test@email.com",
		CustomerAccountID: customerAccountModelResp.ID,
	}

	userService := sf.New(service.USER)

	modelU, err := userService.Create(userModel)
	assert.Nil(t, err)
	assert.NotNil(t, modelU)

	userModelResp, ok := modelU.(*user.User)
	assert.True(t, ok)
	assert.NotNil(t, userModelResp)
	assert.Equal(t, login, *userModelResp.Login)
	assert.NotEqual(t, password, *userModelResp.Password)

	us, ok := userService.(impl.UserService)
	assert.True(t, ok)
	assert.Nil(t, us.SendActivationLink(customerAccountModelResp.ID))
}

func TestUserServiceQuery(t *testing.T) {
	mariaDBFlushData(s)

	login := "login101"
	password := "pass123!"

	userModel := &user.User{
		ProviderID: 1,
		Login:      &login,
		Password:   &password,
		Email:      "email1",
	}

	login2 := "login102"
	password2 := "pass124!"

	userModel2 := &user.User{
		ProviderID: 1,
		Login:      &login2,
		Password:   &password2,
		Email:      "email2",
	}

	login3 := "login103"
	password3 := "pass125!"

	userModel3 := &user.User{
		ProviderID: 1,
		Login:      &login3,
		Password:   &password3,
		Email:      "email3",
	}

	userService := sf.New(service.USER)

	model, err := userService.Create(userModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	model2, err := userService.Create(userModel2)
	assert.Nil(t, err)
	assert.NotNil(t, model2)

	model3, err := userService.Create(userModel3)
	assert.Nil(t, err)
	assert.NotNil(t, model3)

	queryUrl := "https://example.com/api/management/users?filterFields=email&filterValues=email3"
	query := api_utils.ParseQuery("USER", queryUrl)
	count, ens, err := userService.Query(user.User{ProviderID: 1}, query)
	assert.Nil(t, err)
	assert.NotNil(t, ens)
	assert.Equal(t, 1, count)

	userModelResp, ok := ens[0].(*user.User)
	assert.True(t, ok)
	assert.Equal(t, login3, *userModelResp.Login)
	assert.NotEqual(t, password3, *userModelResp.Password)
}
