package impl

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/service_access_point"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	apiUtils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/entity"
)

type serviceAccessPointService struct {
	Dao dao.Dao
}

func NewServiceAccessPointService(dao dao.Dao) *serviceAccessPointService {
	return &serviceAccessPointService{Dao: dao}
}

func (c *serviceAccessPointService) Query(v interface{}, query *apiUtils.Query) (int, []model.Model, error) {
	logger.Debug("Query service access points")

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

func (c *serviceAccessPointService) Check(query *apiUtils.Query) (int, error) {
	logger.Debug("Check service access points")

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

func (c *serviceAccessPointService) Create(model model.Model) (model.Model, error) {
	logger.Debug("Service access point create")
	customerModel, ok := model.(*service_access_point.ServiceAccessPoint)
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

func (c *serviceAccessPointService) DeleteByID(id int) error {
	logger.Debugf("Service access point delete by ID: %d", id)
	err := c.Dao.DeleteByID(id)
	if err != nil {
		return e.DbErrEntityNotFound
	}
	return nil
}

func (c *serviceAccessPointService) UpdateByID(id int, model model.Model) (model.Model, error) {
	logger.Debugf("Service access point update by ID: %d", id)
	customerModel, ok := model.(*service_access_point.ServiceAccessPoint)
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

func (c *serviceAccessPointService) GetByID(id int) (model.Model, error) {
	logger.Debugf("Service access point get by ID: %d", id)
	en, err := c.Dao.GetByID(id)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}

	return c.EntityToModel(en)
}

func (c *serviceAccessPointService) List() ([]model.Model, error) {
	logger.Debug("List service access points")
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

func (c *serviceAccessPointService) GetWithFilter(query interface{}, args ...interface{}) ([]model.Model, error) {
	logger.Debugf("Get service access points with filter: %v", query)
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

func (c *serviceAccessPointService) ModelToEntity(model model.Model) (entity.Entity, error) {
	mdServiceAccessPoint, ok := model.(*service_access_point.ServiceAccessPoint)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	return &entity.ServiceAccessPoint{
		AccountID:   mdServiceAccessPoint.AccountID,
		City:        mdServiceAccessPoint.City,
		Address:     mdServiceAccessPoint.Address,
		SapCode:     mdServiceAccessPoint.SapCode,
		MeterNumber: mdServiceAccessPoint.MeterNumber,
		ProviderID:  mdServiceAccessPoint.ProviderID,
		Name:        mdServiceAccessPoint.Name,
	}, nil
}

func (c *serviceAccessPointService) EntityToModel(en entity.Entity) (model.Model, error) {
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
