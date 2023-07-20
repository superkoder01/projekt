package impl

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/customer_account"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	apiUtils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/entity"
)

type customerAccountService struct {
	Dao dao.Dao
}

func NewCustomerAccountService(dao dao.Dao) *customerAccountService {
	return &customerAccountService{Dao: dao}
}

func (c *customerAccountService) Query(v interface{}, query *apiUtils.Query) (int, []model.Model, error) {
	logger.Debug("Query users")

	sql, err := query.BuildSQLFromApiQuery()
	if err != nil {
		return 0, nil, e.ApiErrImproperQuery
	}

	count, ens, err := c.Dao.Query(v, sql)
	if err != nil {
		return 0, nil, e.ErrInternalServerError
	}
	mdls := []model.Model{}

	for _, e := range ens {
		var customerModel model.Model
		var err error

		customerModel, err = c.EntityToModel(e)
		if err != nil {
			return 0, nil, err
		}
		mdls = append(mdls, customerModel)
	}

	return count, mdls, nil
}

func (c *customerAccountService) Check(query *apiUtils.Query) (int, error) {
	logger.Debug("Check customer accounts")

	sql, err := query.BuildSQLFromApiCheck()
	if err != nil {
		return 0, e.ApiErrImproperQuery
	}

	count, err := c.Dao.Check(sql)
	if err != nil {
		return 0, e.ErrInternalServerError
	}

	return count, nil
}

func (c *customerAccountService) Create(model model.Model) (model.Model, error) {
	logger.Debug("User create")
	customerModel, ok := model.(*customer_account.CustomerAccount)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	customerEntity, err := c.ModelToEntity(customerModel)
	if err != nil {
		return nil, err
	}

	en, err := c.Dao.Create(customerEntity)
	if err != nil {
		return nil, err
	}

	return c.EntityToModel(en)
}

func (c *customerAccountService) DeleteByID(id int) error {
	logger.Debugf("User delete by ID: %d", id)
	err := c.Dao.DeleteByID(id)
	if err != nil {
		return e.DbErrEntityNotFound
	}
	return nil
}

func (c *customerAccountService) UpdateByID(id int, model model.Model) (model.Model, error) {
	logger.Debugf("User update by ID: %d", id)
	customerModel, ok := model.(*customer_account.CustomerAccount)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	en, err := c.ModelToEntity(customerModel)
	if err != nil {
		return nil, err
	}

	en, err = c.Dao.UpdateByID(id, en)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}

	return c.EntityToModel(en)
}

func (c *customerAccountService) GetByID(id int) (model.Model, error) {
	logger.Debugf("User get by ID: %d", id)
	en, err := c.Dao.GetByID(id)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}

	return c.EntityToModel(en)
}

func (c *customerAccountService) List() ([]model.Model, error) {
	logger.Debug("List users")
	ens, err := c.Dao.List()
	if err != nil {
		return nil, e.ErrInternalServerError
	}
	mdls := []model.Model{}

	for _, e := range ens {
		var customerModel model.Model
		var err error

		customerModel, err = c.EntityToModel(e)
		if err != nil {
			return nil, err
		}
		mdls = append(mdls, customerModel)
	}

	return mdls, nil
}

func (c *customerAccountService) GetWithFilter(query interface{}, args ...interface{}) ([]model.Model, error) {
	logger.Debugf("Get users with filter: %v", query)
	ens, err := c.Dao.GetByFilter(query, args)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}
	mdls := []model.Model{}

	for _, e := range ens {
		var customerModel model.Model
		var err error

		customerModel, err = c.EntityToModel(e)
		if err != nil {
			return nil, err
		}
		mdls = append(mdls, customerModel)
	}

	return mdls, nil
}

func (c *customerAccountService) ModelToEntity(model model.Model) (entity.Entity, error) {
	mdCustomerAccount, ok := model.(*customer_account.CustomerAccount)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	return &entity.CustomerAccount{
		ProviderID:         mdCustomerAccount.ProviderID,
		CustomerTypeName:   mdCustomerAccount.CustomerTypeName,
		FirstName:          mdCustomerAccount.FirstName,
		LastName:           mdCustomerAccount.LastName,
		Status:             mdCustomerAccount.Status,
		NIP:                mdCustomerAccount.NIP,
		PESEL:              mdCustomerAccount.PESEL,
		REGON:              mdCustomerAccount.REGON,
		Email:              mdCustomerAccount.Email,
		Phone:              mdCustomerAccount.Phone,
		Street:             mdCustomerAccount.Street,
		BuildingNumber:     mdCustomerAccount.BuildingNumber,
		ApartmentNumber:    mdCustomerAccount.ApartmentNumber,
		PostalCode:         mdCustomerAccount.PostalCode,
		Province:           mdCustomerAccount.Province,
		City:               mdCustomerAccount.City,
		Country:            mdCustomerAccount.Country,
		BankAccNumber:      mdCustomerAccount.BankAccNumber,
		WorkerID:           mdCustomerAccount.WorkerID,
		RegistrationNumber: mdCustomerAccount.RegistrationNumber,
		BusinessType:       mdCustomerAccount.BusinessType,
		KRS:                mdCustomerAccount.KRS,
		LineOfBusiness:     mdCustomerAccount.LineOfBusiness,
	}, nil
}

func (c *customerAccountService) EntityToModel(en entity.Entity) (model.Model, error) {
	enCustomerAccount, ok := en.(*entity.CustomerAccount)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	return &customer_account.CustomerAccount{
		ApartmentNumber:    enCustomerAccount.ApartmentNumber,
		BuildingNumber:     enCustomerAccount.BuildingNumber,
		City:               enCustomerAccount.City,
		ProviderID:         enCustomerAccount.ProviderID,
		Country:            enCustomerAccount.Country,
		CustomerTypeName:   enCustomerAccount.CustomerTypeName,
		Email:              enCustomerAccount.Email,
		FirstName:          enCustomerAccount.FirstName,
		ID:                 enCustomerAccount.ID,
		LastName:           enCustomerAccount.LastName,
		Phone:              enCustomerAccount.Phone,
		PostalCode:         enCustomerAccount.PostalCode,
		Province:           enCustomerAccount.Province,
		Street:             enCustomerAccount.Street,
		NIP:                enCustomerAccount.NIP,
		REGON:              enCustomerAccount.REGON,
		Status:             enCustomerAccount.Status,
		PESEL:              enCustomerAccount.PESEL,
		BankAccNumber:      enCustomerAccount.BankAccNumber,
		WorkerID:           enCustomerAccount.WorkerID,
		RegistrationNumber: enCustomerAccount.RegistrationNumber,
		BusinessType:       enCustomerAccount.BusinessType,
		KRS:                enCustomerAccount.KRS,
		LineOfBusiness:     enCustomerAccount.LineOfBusiness,
	}, nil
}
