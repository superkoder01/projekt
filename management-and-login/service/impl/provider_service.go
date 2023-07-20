package impl

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/provider"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	apiUtils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/entity"
)

type ProviderService interface {
	DeleteWithAdmins(int) error
}

type providerService struct {
	Dao dao.Dao
}

func NewProviderService(dao dao.Dao) *providerService {
	return &providerService{Dao: dao}
}

func (p *providerService) Query(v interface{}, query *apiUtils.Query) (int, []model.Model, error) {
	logger.Debug("Query providers")

	sql, err := query.BuildSQLFromApiQuery()
	if err != nil {
		return 0, nil, e.ApiErrImproperQuery
	}

	count, ens, err := p.Dao.Query(v, sql)
	if err != nil {
		return 0, nil, e.ErrInternalServerError
	}
	mdls := []model.Model{}

	for _, e := range ens {
		var customerModel model.Model
		var err error

		customerModel, err = p.EntityToModel(e)
		if err != nil {
			return 0, nil, err
		}
		mdls = append(mdls, customerModel)
	}

	return count, mdls, nil
}

func (p *providerService) Check(query *apiUtils.Query) (int, error) {
	logger.Debug("Check providers")

	sql, err := query.BuildSQLFromApiCheck()
	if err != nil {
		return 0, e.ApiErrImproperQuery
	}

	count, err := p.Dao.Check(sql)
	if err != nil {
		return 0, e.ErrInternalServerError
	}

	return count, nil
}

func (p *providerService) Create(model model.Model) (model.Model, error) {
	logger.Debug("Provider create")
	providerModel, ok := model.(*provider.Provider)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	providerEntity, err := p.ModelToEntity(providerModel)
	if err != nil {
		return nil, err
	}

	en, err := p.Dao.Create(providerEntity)
	if err != nil {
		return nil, err
	}

	return p.EntityToModel(en)
}

func (p *providerService) DeleteByID(id int) error {
	logger.Debugf("Provider delete by ID: %d", id)
	err := p.Dao.DeleteByID(id)
	if err != nil {
		return e.DbErrEntityNotFound
	}
	return nil
}

func (p *providerService) DeleteWithAdmins(id int) error {
	logger.Debugf("Provider with administrators delete by ID: %d", id)
	var err error

	var ens []entity.CustomerAccount

	tx := p.Dao.BeginTransaction()
	defer tx.Rollback()

	if err = tx.Where(&entity.CustomerAccount{ProviderID: id}).Find(&ens).Error(); err != nil {
		return err
	}

	if len(ens) > 0 {
		return e.DbErrEntityDeleteHasChild
	}

	logger.Debugf("Deleting users with provider id: %d", id)
	if err = tx.Where(&entity.User{ProviderID: id}).Delete(&entity.User{}).Error(); err != nil {
		return err
	}

	logger.Debugf("Deleting workers with provider id: %d", id)
	if err = tx.Where(&entity.Worker{ProviderID: id}).Delete(&entity.Worker{}).Error(); err != nil {
		return err
	}

	logger.Debugf("Deleting provider with id: %d", id)
	if err = tx.Where(&entity.Provider{ID: id}).Delete(&entity.Provider{}).Error(); err != nil {
		return err
	}

	return tx.Commit().Error()
}

func (p *providerService) GetByID(id int) (model.Model, error) {
	logger.Debugf("Provider get by ID: %d", id)
	en, err := p.Dao.GetByID(id)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}

	return p.EntityToModel(en)
}

func (p *providerService) List() ([]model.Model, error) {
	logger.Debug("List providers")
	ens, err := p.Dao.List()
	if err != nil {
		return nil, e.ErrInternalServerError
	}
	mdls := []model.Model{}

	for _, e := range ens {
		var providerModel model.Model
		var err error

		providerModel, err = p.EntityToModel(e)
		if err != nil {
			return nil, err
		}
		mdls = append(mdls, providerModel)
	}

	return mdls, nil
}

func (p *providerService) GetWithFilter(query interface{}, args ...interface{}) ([]model.Model, error) {
	logger.Debugf("Get providers with filter: %v", query)
	ens, err := p.Dao.GetByFilter(query, args)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}
	mdls := []model.Model{}

	for _, e := range ens {
		var providerModel model.Model
		var err error

		providerModel, err = p.EntityToModel(e)
		if err != nil {
			return nil, err
		}
		mdls = append(mdls, providerModel)
	}

	return mdls, nil
}

func (p *providerService) UpdateByID(id int, model model.Model) (model.Model, error) {
	logger.Debugf("Provider update by ID: %d", id)
	providerModel, ok := model.(*provider.Provider)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	en, err := p.ModelToEntity(providerModel)
	if err != nil {
		return nil, err
	}

	en, err = p.Dao.UpdateByID(id, en)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}

	return p.EntityToModel(en)
}

func (p *providerService) ModelToEntity(model model.Model) (entity.Entity, error) {
	mdProvider, ok := model.(*provider.Provider)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	return &entity.Provider{
		Name:                  mdProvider.Name,
		Type:                  mdProvider.Type,
		Status:                mdProvider.Status,
		NIP:                   mdProvider.NIP,
		REGON:                 mdProvider.REGON,
		KRS:                   mdProvider.KRS,
		Email:                 mdProvider.Email,
		PhoneNumber:           mdProvider.PhoneNumber,
		Street:                mdProvider.Street,
		BuildingNumber:        mdProvider.BuildingNumber,
		ApartmentNumber:       mdProvider.ApartmentNumber,
		PostalCode:            mdProvider.PostalCode,
		Province:              mdProvider.Province,
		City:                  mdProvider.City,
		Country:               mdProvider.Country,
		LicenseID:             mdProvider.LicenseID,
		LicenseExpirationDate: mdProvider.LicenseExpirationDate,
		LicenseArea:           mdProvider.LicenseArea,
		WWW:                   mdProvider.WWW,
	}, nil
}

func (p *providerService) EntityToModel(en entity.Entity) (model.Model, error) {
	enProvider, ok := en.(*entity.Provider)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	return &provider.Provider{
		ID:                    enProvider.ID,
		Name:                  enProvider.Name,
		Type:                  enProvider.Type,
		Status:                enProvider.Status,
		NIP:                   enProvider.NIP,
		REGON:                 enProvider.REGON,
		KRS:                   enProvider.KRS,
		Email:                 enProvider.Email,
		PhoneNumber:           enProvider.PhoneNumber,
		Street:                enProvider.Street,
		BuildingNumber:        enProvider.BuildingNumber,
		ApartmentNumber:       enProvider.ApartmentNumber,
		PostalCode:            enProvider.PostalCode,
		Province:              enProvider.Province,
		City:                  enProvider.City,
		Country:               enProvider.Country,
		LicenseID:             enProvider.LicenseID,
		LicenseExpirationDate: enProvider.LicenseExpirationDate,
		LicenseArea:           enProvider.LicenseArea,
		WWW:                   enProvider.WWW,
	}, nil
}
