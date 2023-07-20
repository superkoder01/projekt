package tests

import (
	"github.com/stretchr/testify/assert"
	cu "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/customer_user"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/enum"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
	"testing"
)

func TestCreateCustomerUserService(t *testing.T) {
	mariaDBFlushData(s)

	login := "jkowalski"
	password := "kowal1990!"

	customerUserModel := &cu.CustomerUser{
		ProviderID:       1,
		Email:            "jan.kowalski@email.com",
		CustomerTypeName: "PROSUMER",
		Login:            &login,
		Password:         &password,
		RoleID:           int(enum.PROSUMER),
	}

	customerUserService := sf.New(service.CUSTOMER_USER)

	model, err := customerUserService.Create(customerUserModel)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	customerUserModelResp, ok := model.(*cu.CustomerUser)
	assert.True(t, ok)
	assert.NotNil(t, customerUserModelResp)
	assert.NotNil(t, customerUserModelResp.CustomerAccountID)
	assert.NotNil(t, customerUserModelResp.UserID)
	assert.Equal(t, customerUserModel.ProviderID, customerUserModelResp.ProviderID)
	assert.Equal(t, customerUserModel.Email, customerUserModelResp.Email)
	assert.Equal(t, customerUserModel.Login, customerUserModelResp.Login)
	assert.NotEqual(t, password, customerUserModelResp.Password)
}
