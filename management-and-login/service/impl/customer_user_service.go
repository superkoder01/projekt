package impl

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/customer_user"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	apiUtils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/auth"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/entity"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/mysql"
)

type customerUserService struct {
	Auth     auth.Authenticator
	User     dao.Dao
	Customer dao.Dao
}

func NewCustomerUserService(user dao.Dao, customer dao.Dao, auth auth.Authenticator) *customerUserService {
	return &customerUserService{
		User:     user,
		Customer: customer,
		Auth:     auth,
	}
}

func (w *customerUserService) Create(model model.Model) (model.Model, error) {
	logger.Debug("Customer and user create")
	customerUserModel, ok := model.(*customer_user.CustomerUser)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	// map model to customer entity
	customerEntity, err := w.CustomerModelToEntity(customerUserModel)
	if err != nil {
		return nil, err
	}
	customerEntity.(*entity.CustomerAccount).SetRegistrationNumber(generateRegistrationCode())

	// TRANSACTION BEGIN
	tx := w.User.BeginTransaction()
	defer tx.Rollback()

	if err = tx.Create(customerEntity).Error(); err != nil {
		return nil, err
	}

	en, ok := customerEntity.(*entity.CustomerAccount)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	// hash password
	var hash string
	if hash, err = w.Auth.Encode(*customerUserModel.Password); err != nil {
		return nil, e.ErrInternalServerError
	}
	customerUserModel.Password = &hash

	// map model to user entity
	userEntity, err := w.UserModelToEntity(en.ID, customerUserModel)
	if err != nil {
		return nil, err
	}

	// create user entity
	if err = createUserEntity(tx, userEntity, w.Auth.GenerateRandomCode()); err != nil {
		return nil, err
	}

	if err = tx.Commit().Error(); err != nil {
		return nil, e.ErrInternalServerError
	}
	// TRANSACTION COMMITED

	return w.EntitiesToModel(customerEntity, userEntity)
}

func (w *customerUserService) Check(query *apiUtils.Query) (int, error) {
	logger.Debug("Check users and customer accounts")

	var (
		count int
		err   error
		sql   *bd.Query
	)

	if isOnList(query.FilterFields[0], userFields) {
		query.TableName = "USER"
		sql, err = query.BuildSQLFromApiCheck()
		if err != nil {
			return 0, e.ApiErrImproperQuery
		}
		count, err = w.User.Check(sql)
	} else {
		sql, err = query.BuildSQLFromApiCheck()
		if err != nil {
			return 0, e.ApiErrImproperQuery
		}
		count, err = w.Customer.Check(sql)
	}

	if err != nil {
		return 0, e.ErrInternalServerError
	}

	return count, nil
}

// NOT USED
func (w *customerUserService) Query(v interface{}, query *apiUtils.Query) (int, []model.Model, error) {
	return 0, nil, nil
}

// NOT USED
func (w *customerUserService) DeleteByID(i int) error {
	return nil
}

// NOT USED
func (w *customerUserService) GetByID(i int) (model.Model, error) {
	return nil, nil
}

// NOT USED
func (w *customerUserService) List() ([]model.Model, error) {
	return nil, nil
}

// NOT USED
func (w *customerUserService) GetWithFilter(query interface{}, args ...interface{}) ([]model.Model, error) {
	return nil, nil
}

// NOT USED
func (w *customerUserService) UpdateByID(i int, model model.Model) (model.Model, error) {
	return nil, nil
}

// NOT USED
func (w *customerUserService) ModelToEntity(model model.Model) (entity.Entity, error) {
	return nil, nil
}

// NOT USED
func (w *customerUserService) EntityToModel(entity entity.Entity) (model.Model, error) {
	return nil, nil
}

func (w *customerUserService) EntitiesToModel(customer entity.Entity, user entity.Entity) (model.Model, error) {
	enCustomer, ok := customer.(*entity.CustomerAccount)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	enUser, ok := user.(*entity.User)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	return &customer_user.CustomerUser{
		UserID:             enUser.ID,
		Login:              &enUser.Login,
		Password:           &enUser.Password,
		Active:             enUser.Active,
		MustChangePassword: enUser.MustChangePassword,
		RoleID:             enUser.RoleID,
		ProviderID:         enCustomer.ProviderID,
		CustomerAccountID:  enCustomer.ID,
		CustomerTypeName:   enCustomer.CustomerTypeName,
		FirstName:          enCustomer.FirstName,
		LastName:           enCustomer.LastName,
		Status:             enCustomer.Status,
		Street:             enCustomer.Street,
		BuildingNumber:     enCustomer.BuildingNumber,
		ApartmentNumber:    enCustomer.ApartmentNumber,
		City:               enCustomer.City,
		PostalCode:         enCustomer.PostalCode,
		Province:           enCustomer.Province,
		Country:            enCustomer.Country,
		Phone:              enCustomer.Phone,
		Email:              enCustomer.Email,
		NIP:                enCustomer.NIP,
		REGON:              enCustomer.REGON,
		PESEL:              enCustomer.PESEL,
		BankAccNumber:      enCustomer.BankAccNumber,
		WorkerID:           enCustomer.WorkerID,
		RegistrationNumber: enCustomer.RegistrationNumber,
		BusinessType:       enCustomer.BusinessType,
		KRS:                enCustomer.KRS,
		LineOfBusiness:     enCustomer.LineOfBusiness,
	}, nil
}

func (w *customerUserService) CustomerModelToEntity(model model.Model) (entity.Entity, error) {
	mdCustomer, ok := model.(*customer_user.CustomerUser)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	return &entity.CustomerAccount{
		ProviderID:         mdCustomer.ProviderID,
		CustomerTypeName:   mdCustomer.CustomerTypeName,
		FirstName:          mdCustomer.FirstName,
		LastName:           mdCustomer.LastName,
		Status:             mdCustomer.Status,
		NIP:                mdCustomer.NIP,
		PESEL:              mdCustomer.PESEL,
		REGON:              mdCustomer.REGON,
		Email:              mdCustomer.Email,
		Phone:              mdCustomer.Phone,
		Street:             mdCustomer.Street,
		BuildingNumber:     mdCustomer.BuildingNumber,
		ApartmentNumber:    mdCustomer.ApartmentNumber,
		PostalCode:         mdCustomer.PostalCode,
		Province:           mdCustomer.Province,
		City:               mdCustomer.City,
		Country:            mdCustomer.Country,
		BankAccNumber:      mdCustomer.BankAccNumber,
		WorkerID:           mdCustomer.WorkerID,
		RegistrationNumber: mdCustomer.RegistrationNumber,
		BusinessType:       mdCustomer.BusinessType,
		KRS:                mdCustomer.KRS,
		LineOfBusiness:     mdCustomer.LineOfBusiness,
	}, nil
}

func (w *customerUserService) UserModelToEntity(customerId int, model model.Model) (entity.Entity, error) {
	mdUser, ok := model.(*customer_user.CustomerUser)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	var login, password string
	if mdUser.Login != nil {
		login = *mdUser.Login
	}
	if mdUser.Password != nil {
		password = *mdUser.Password
	}

	return &entity.User{
		Login:              login,
		Password:           password,
		Email:              mdUser.Email,
		ProviderID:         mdUser.ProviderID,
		Active:             mdUser.Active,
		MustChangePassword: mdUser.MustChangePassword,
		CustomerAccountID:  customerId,
		RoleID:             mdUser.RoleID,
	}, nil
}
