package impl

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/enum"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/worker"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	apiUtils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/entity"
)

type WorkerService interface {
	DeleteWithUser(int) error
	ListWorkers(role int, providerId int, workerId int, query *apiUtils.Query) (int, []model.Model, error)
}

type workerService struct {
	Dao dao.Dao
}

func NewWorkerService(dao dao.Dao) *workerService {
	return &workerService{Dao: dao}
}

func (w *workerService) Query(v interface{}, query *apiUtils.Query) (int, []model.Model, error) {
	logger.Debug("Query workers")

	sql, err := query.BuildSQLFromApiQuery()
	if err != nil {
		return 0, nil, e.ApiErrImproperQuery
	}

	count, ens, err := w.Dao.Query(v, sql)
	if err != nil {
		return 0, nil, e.ErrInternalServerError
	}
	mdls := []model.Model{}

	for _, e := range ens {
		var workerModel model.Model
		var err error

		workerModel, err = w.EntityToWorkerWithRole(e)
		if err != nil {
			return 0, nil, err
		}
		mdls = append(mdls, workerModel)
	}

	return count, mdls, nil
}

func (w *workerService) Check(query *apiUtils.Query) (int, error) {
	logger.Debug("Check workers")

	sql, err := query.BuildSQLFromApiCheck()
	if err != nil {
		return 0, e.ApiErrImproperQuery
	}

	count, err := w.Dao.Check(sql)
	if err != nil {
		return 0, e.ErrInternalServerError
	}

	return count, nil
}

func (w *workerService) ListWorkers(role int, providerId int, workerId int, query *apiUtils.Query) (int, []model.Model, error) {
	logger.Debugf("Get worker's workers - providerId: %d, workerId: %d", providerId, workerId)

	m := make(map[string]interface{})
	if role > int(enum.SUPER_ADMIN) && role < int(enum.TRADER) {
		m["`WORKER`.`PROVIDER_ID`"] = providerId
	} else if role > int(enum.ADMINISTRATOR_BASIC) {
		subordinates, err := w.Dao.GetByFilter(entity.Worker{ProviderID: providerId, Supervisor: workerId})
		if err != nil {
			return 0, nil, e.DbErrEntityNotFound
		}

		ids := []int{workerId}
		for _, sub := range subordinates {
			ids = append(ids, sub.(*entity.WorkerJoinUserRole).ID)
		}

		m["`WORKER`.`PROVIDER_ID`"] = providerId
		m["`WORKER`.`ID`"] = ids
	}

	sql, err := query.BuildSQLFromApiQuery()
	if err != nil {
		return 0, nil, e.ApiErrImproperQuery
	}

	count, ens, err := w.Dao.Query(m, sql)

	mdls := []model.Model{}
	for _, e := range ens {
		var workerModel model.Model
		var err error

		workerModel, err = w.EntityToWorkerWithRole(e)
		if err != nil {
			return 0, nil, err
		}
		mdls = append(mdls, workerModel)
	}

	return count, mdls, nil
}

func (w *workerService) Create(model model.Model) (model.Model, error) {
	logger.Debug("Worker create")
	workerModel, ok := model.(*worker.Worker)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	workerEntity, err := w.ModelToEntity(workerModel)
	if err != nil {
		return nil, err
	}

	en, err := w.Dao.Create(workerEntity)
	if err != nil {
		return nil, err
	}

	return w.EntityToModel(en)
}

func (w *workerService) DeleteByID(id int) error {
	logger.Debugf("Worker delete by ID: %d", id)
	err := w.Dao.DeleteByID(id)
	if err != nil {
		return e.DbErrEntityNotFound
	}
	return nil
}

func (w *workerService) DeleteWithUser(id int) error {
	logger.Debugf("Worker with user delete by ID: %d", id)
	var err error

	var ens []entity.CustomerAccount

	tx := w.Dao.BeginTransaction()
	defer tx.Rollback()

	if err = tx.Where(&entity.CustomerAccount{WorkerID: id}).Find(&ens).Error(); err != nil {
		return err
	}

	if len(ens) > 0 {
		return e.DbErrEntityDeleteHasChild
	}

	logger.Debugf("Deleting users with worker id: %d", id)
	if err = tx.Where(&entity.User{WorkerID: id}).Delete(&entity.User{}).Error(); err != nil {
		return err
	}

	logger.Debugf("Deleting workers with id: %d", id)
	if err = tx.Where(&entity.Worker{ID: id}).Delete(&entity.Worker{}).Error(); err != nil {
		return err
	}

	return tx.Commit().Error()
}

func (w *workerService) GetByID(id int) (model.Model, error) {
	logger.Debugf("Worker get by ID: %d", id)
	en, err := w.Dao.GetByID(id)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}

	return w.EntityToWorkerWithRole(en)
}

func (w *workerService) List() ([]model.Model, error) {
	logger.Debug("List workers")
	ens, err := w.Dao.List()
	if err != nil {
		return nil, e.ErrInternalServerError
	}
	mdls := []model.Model{}

	for _, e := range ens {
		var workerModel model.Model
		var err error

		workerModel, err = w.EntityToWorkerWithRole(e)
		if err != nil {
			return nil, err
		}
		mdls = append(mdls, workerModel)
	}

	return mdls, nil
}

func (w *workerService) GetWithFilter(query interface{}, args ...interface{}) ([]model.Model, error) {
	logger.Debugf("Get workers with filter: %v", query)
	ens, err := w.Dao.GetByFilter(query, args)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}
	mdls := []model.Model{}

	for _, e := range ens {
		var workerModel model.Model
		var err error

		workerModel, err = w.EntityToWorkerWithRole(e)
		if err != nil {
			return nil, err
		}
		mdls = append(mdls, workerModel)
	}

	return mdls, nil
}

func (w *workerService) UpdateByID(id int, model model.Model) (model.Model, error) {
	logger.Debugf("Worker update by ID: %d", id)
	workerModel, ok := model.(*worker.Worker)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	en, err := w.ModelToEntity(workerModel)
	if err != nil {
		return nil, err
	}

	en, err = w.Dao.UpdateByID(id, en)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}

	return w.EntityToModel(en)
}

func (w *workerService) ModelToEntity(model model.Model) (entity.Entity, error) {
	mdWorker, ok := model.(*worker.Worker)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	return &entity.Worker{
		FirstName:            mdWorker.FirstName,
		LastName:             mdWorker.LastName,
		Email:                mdWorker.Email,
		Phone:                mdWorker.Phone,
		WorkStartDate:        mdWorker.WorkStartDate,
		WorkEndDate:          mdWorker.WorkEndDate,
		BlockchainAccAddress: mdWorker.BlockchainAccAddress,
		Street:               mdWorker.Street,
		City:                 mdWorker.City,
		PostalCode:           mdWorker.PostalCode,
		Province:             mdWorker.Province,
		BuildingNumber:       mdWorker.BuildingNumber,
		ApartmentNumber:      mdWorker.ApartmentNumber,
		Country:              mdWorker.Country,
		ProviderID:           mdWorker.ProviderID,
		Supervisor:           mdWorker.Supervisor,
		Status:               mdWorker.Status,
		NIP:                  mdWorker.NIP,
		REGON:                mdWorker.REGON,
		PESEL:                mdWorker.PESEL,
		KRS:                  mdWorker.KRS,
		ExtraInfo:            mdWorker.ExtraInfo,
	}, nil
}

func (w *workerService) EntityToModel(en entity.Entity) (model.Model, error) {
	enWorker, ok := en.(*entity.Worker)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	return &worker.Worker{
		ID:                   enWorker.ID,
		FirstName:            enWorker.FirstName,
		LastName:             enWorker.LastName,
		Email:                enWorker.Email,
		Phone:                enWorker.Phone,
		WorkStartDate:        enWorker.WorkStartDate,
		WorkEndDate:          enWorker.WorkEndDate,
		BlockchainAccAddress: enWorker.BlockchainAccAddress,
		Street:               enWorker.Street,
		City:                 enWorker.City,
		PostalCode:           enWorker.PostalCode,
		Province:             enWorker.Province,
		BuildingNumber:       enWorker.BuildingNumber,
		ApartmentNumber:      enWorker.ApartmentNumber,
		Country:              enWorker.Country,
		ProviderID:           enWorker.ProviderID,
		Supervisor:           enWorker.Supervisor,
		Status:               enWorker.Status,
		NIP:                  enWorker.NIP,
		REGON:                enWorker.REGON,
		PESEL:                enWorker.PESEL,
		KRS:                  enWorker.KRS,
		ExtraInfo:            enWorker.ExtraInfo,
	}, nil
}

func (w *workerService) EntityToWorkerWithRole(en entity.Entity) (model.Model, error) {
	enWorker, ok := en.(*entity.WorkerJoinUserRole)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	return &worker.Worker{
		ID:                   enWorker.ID,
		FirstName:            enWorker.FirstName,
		LastName:             enWorker.LastName,
		Email:                enWorker.Email,
		Phone:                enWorker.Phone,
		WorkStartDate:        enWorker.WorkStartDate,
		WorkEndDate:          enWorker.WorkEndDate,
		BlockchainAccAddress: enWorker.BlockchainAccAddress,
		Street:               enWorker.Street,
		City:                 enWorker.City,
		PostalCode:           enWorker.PostalCode,
		Province:             enWorker.Province,
		BuildingNumber:       enWorker.BuildingNumber,
		ApartmentNumber:      enWorker.ApartmentNumber,
		Country:              enWorker.Country,
		ProviderID:           enWorker.ProviderID,
		Supervisor:           enWorker.Supervisor,
		Status:               enWorker.Status,
		NIP:                  enWorker.NIP,
		REGON:                enWorker.REGON,
		PESEL:                enWorker.PESEL,
		KRS:                  enWorker.KRS,
		ExtraInfo:            enWorker.ExtraInfo,
		Role:                 enum.Role(enWorker.RoleID).Name(),
	}, nil
}
