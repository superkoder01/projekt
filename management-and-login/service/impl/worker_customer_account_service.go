package impl

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/customer_account"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/enum"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/service_access_point"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	apiUtils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/entity"
)

type workerCustomerAccountService struct {
	customer dao.Dao
	worker   dao.Dao
	sap      dao.Dao
}

type WorkerCustomerAccountService interface {
	ListWorkerCustomerAccounts(role int, providerId int, workerId int, query *apiUtils.Query) (int, []model.Model, error)
	ListWorkerServiceAccessPoints(role int, providerId int, workerId int, query *apiUtils.Query) (int, []model.Model, error)
}

func NewWorkerCustomerAccountService(customer dao.Dao, worker dao.Dao, sap dao.Dao) *workerCustomerAccountService {
	return &workerCustomerAccountService{customer: customer, worker: worker, sap: sap}
}

func (w *workerCustomerAccountService) ListWorkerServiceAccessPoints(role int, providerId int, workerId int, query *apiUtils.Query) (int, []model.Model, error) {
	logger.Debugf("Get worker's customer's service access points - providerId: %d, workerId: %d", providerId, workerId)

	var (
		// customer filter
		m = make(map[string]interface{})
		// service access points filter
		sapM        = make(map[string]interface{})
		customerIds []int
		// entities count
		count int
		// entitites
		ens []entity.Entity
		// models to return
		mdls []model.Model
		err  error
	)

	if role > int(enum.SUPER_ADMIN) && role < int(enum.TRADER) {
		m["`CUSTOMER_ACCOUNT`.`PROVIDER_ID`"] = providerId
	} else if role > int(enum.ADMINISTRATOR_BASIC) {
		subordinates, err := w.worker.GetByFilter(entity.Worker{ProviderID: providerId, Supervisor: workerId})
		if err != nil {
			return 0, nil, e.DbErrEntityNotFound
		}

		ids := []int{workerId}
		for _, sub := range subordinates {
			ids = append(ids, sub.(*entity.WorkerJoinUserRole).ID)
		}

		m = make(map[string]interface{})
		m["`CUSTOMER_ACCOUNT`.`PROVIDER_ID`"] = providerId
		m["`CUSTOMER_ACCOUNT`.`WORKER_ID`"] = ids
	}

	sql, err := query.BuildSQLFromApiQuery()
	if err != nil {
		return 0, nil, e.ApiErrImproperQuery
	}

	if role > int(enum.SUPER_ADMIN) {
		customers, err := w.customer.GetByFilter(m)
		if err != nil {
			return 0, nil, e.DbErrEntityNotFound
		}

		for _, customer := range customers {
			customerIds = append(customerIds, customer.(*entity.CustomerAccount).ID)
		}
		sapM["`SERVICE_ACCESS_POINT`.`ACCOUNT_ID`"] = customerIds

		count, ens, err = w.sap.Query(sapM, sql)
	} else {
		count, ens, err = w.sap.Query(entity.ServiceAccessPoint{}, sql)
	}

	for _, e := range ens {
		var serviceAccessPointModel model.Model
		var err error

		serviceAccessPointModel, err = w.EntityToSapModel(e)
		if err != nil {
			return 0, nil, err
		}
		mdls = append(mdls, serviceAccessPointModel)
	}

	return count, mdls, nil
}

func (w *workerCustomerAccountService) ListWorkerCustomerAccounts(role int, providerId int, workerId int, query *apiUtils.Query) (int, []model.Model, error) {
	logger.Debugf("Get worker's customer accounts - providerId: %d, workerId: %d", providerId, workerId)

	m := make(map[string]interface{})
	if role > int(enum.SUPER_ADMIN) && role < int(enum.TRADER) {
		m["`CUSTOMER_ACCOUNT`.`PROVIDER_ID`"] = providerId
	} else if role > int(enum.ADMINISTRATOR_BASIC) {
		subordinates, err := w.worker.GetByFilter(entity.Worker{ProviderID: providerId, Supervisor: workerId})
		if err != nil {
			return 0, nil, e.DbErrEntityNotFound
		}

		ids := []int{workerId}
		for _, sub := range subordinates {
			ids = append(ids, sub.(*entity.WorkerJoinUserRole).ID)
		}

		m = make(map[string]interface{})
		m["`CUSTOMER_ACCOUNT`.`PROVIDER_ID`"] = providerId
		m["`CUSTOMER_ACCOUNT`.`WORKER_ID`"] = ids
	}

	sql, err := query.BuildSQLFromApiQuery()
	if err != nil {
		return 0, nil, e.ApiErrImproperQuery
	}

	count, ens, err := w.customer.Query(m, sql)

	mdls := []model.Model{}
	for _, e := range ens {
		var workerModel model.Model
		var err error

		workerModel, err = w.EntityToModel(e)
		if err != nil {
			return 0, nil, err
		}
		mdls = append(mdls, workerModel)
	}

	return count, mdls, nil
}

func (w *workerCustomerAccountService) Check(query *apiUtils.Query) (int, error) {
	return 0, nil
}

func (w *workerCustomerAccountService) Create(model model.Model) (model.Model, error) {
	return nil, nil
}

func (w *workerCustomerAccountService) DeleteByID(i int) error {
	return nil
}

func (w *workerCustomerAccountService) GetByID(i int) (model.Model, error) {
	return nil, nil
}

func (w *workerCustomerAccountService) GetWithFilter(query interface{}, args ...interface{}) ([]model.Model, error) {
	return nil, nil
}

func (w *workerCustomerAccountService) List() ([]model.Model, error) {
	return nil, nil
}

func (w *workerCustomerAccountService) Query(v interface{}, query *apiUtils.Query) (int, []model.Model, error) {
	return 0, nil, nil
}

func (w *workerCustomerAccountService) UpdateByID(i int, model model.Model) (model.Model, error) {
	return nil, nil
}

func (c *workerCustomerAccountService) ModelToEntity(model model.Model) (entity.Entity, error) {
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

func (c *workerCustomerAccountService) EntityToSapModel(en entity.Entity) (model.Model, error) {
	enServiceAccessPoint, ok := en.(*entity.ServiceAccessPoint)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	return &service_access_point.ServiceAccessPoint{
		ID:          enServiceAccessPoint.ID,
		AccountID:   enServiceAccessPoint.AccountID,
		City:        enServiceAccessPoint.City,
		Address:     enServiceAccessPoint.Address,
		SapCode:     enServiceAccessPoint.SapCode,
		MeterNumber: enServiceAccessPoint.MeterNumber,
		ProviderID:  enServiceAccessPoint.ProviderID,
		Name:        enServiceAccessPoint.Name,
	}, nil
}

func (c *workerCustomerAccountService) EntityToModel(en entity.Entity) (model.Model, error) {
	enCustomerAccount, ok := en.(*entity.CustomerAccount)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	return &customer_account.CustomerAccount{
		ID:                 enCustomerAccount.ID,
		ApartmentNumber:    enCustomerAccount.ApartmentNumber,
		BuildingNumber:     enCustomerAccount.BuildingNumber,
		City:               enCustomerAccount.City,
		ProviderID:         enCustomerAccount.ProviderID,
		Country:            enCustomerAccount.Country,
		Email:              enCustomerAccount.Email,
		NIP:                enCustomerAccount.NIP,
		REGON:              enCustomerAccount.REGON,
		FirstName:          enCustomerAccount.FirstName,
		LastName:           enCustomerAccount.LastName,
		Phone:              enCustomerAccount.Phone,
		PostalCode:         enCustomerAccount.PostalCode,
		Province:           enCustomerAccount.Province,
		Street:             enCustomerAccount.Street,
		CustomerTypeName:   enCustomerAccount.CustomerTypeName,
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
