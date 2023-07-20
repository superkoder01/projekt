package impl

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/worker_user"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	apiUtils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/auth"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/entity"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/mysql"
)

type workerUserService struct {
	Auth   auth.Authenticator
	User   dao.Dao
	Worker dao.Dao
}

func NewWorkerUserService(user dao.Dao, worker dao.Dao, auth auth.Authenticator) *workerUserService {
	return &workerUserService{
		User:   user,
		Worker: worker,
		Auth:   auth,
	}
}

func (w *workerUserService) Create(model model.Model) (model.Model, error) {
	logger.Debug("Worker and user create")
	workerUserModel, ok := model.(*worker_user.WorkerUser)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	// map model to worker entity
	workerEntity, err := w.WorkerModelToEntity(workerUserModel)
	if err != nil {
		return nil, err
	}

	// TRANSACTION BEGIN
	tx := w.User.BeginTransaction()
	defer tx.Rollback()

	if err = tx.Create(workerEntity).Error(); err != nil {
		return nil, err
	}

	en, ok := workerEntity.(*entity.Worker)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	// hash password
	var hash string
	if hash, err = w.Auth.Encode(*workerUserModel.Password); err != nil {
		return nil, e.ErrInternalServerError
	}
	workerUserModel.Password = &hash

	// map model to user entity
	userEntity, err := w.UserModelToEntity(en.ID, workerUserModel)
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

	return w.EntitiesToModel(workerEntity, userEntity)
}

func (w *workerUserService) Check(query *apiUtils.Query) (int, error) {
	logger.Debug("Check users and workers")

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
		count, err = w.Worker.Check(sql)
	}

	if err != nil {
		return 0, e.ErrInternalServerError
	}

	return count, nil
}

// NOT USED
func (w *workerUserService) Query(v interface{}, query *apiUtils.Query) (int, []model.Model, error) {
	return 0, nil, nil
}

// NOT USED
func (w *workerUserService) DeleteByID(i int) error {
	return nil
}

// NOT USED
func (w *workerUserService) GetByID(i int) (model.Model, error) {
	return nil, nil
}

// NOT USED
func (w *workerUserService) List() ([]model.Model, error) {
	return nil, nil
}

// NOT USED
func (w *workerUserService) GetWithFilter(query interface{}, args ...interface{}) ([]model.Model, error) {
	return nil, nil
}

// NOT USED
func (w *workerUserService) UpdateByID(i int, model model.Model) (model.Model, error) {
	return nil, nil
}

// NOT USED
func (w *workerUserService) ModelToEntity(model model.Model) (entity.Entity, error) {
	return nil, nil
}

// NOT USED
func (w *workerUserService) EntityToModel(entity entity.Entity) (model.Model, error) {
	return nil, nil
}

func (w *workerUserService) EntitiesToModel(worker entity.Entity, user entity.Entity) (model.Model, error) {
	enWorker, ok := worker.(*entity.Worker)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	enUser, ok := user.(*entity.User)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	return &worker_user.WorkerUser{
		WorkerID:             enWorker.ID,
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
		NIP:                  enWorker.NIP,
		REGON:                enWorker.REGON,
		PESEL:                enWorker.PESEL,
		KRS:                  enWorker.KRS,
		ExtraInfo:            enWorker.ExtraInfo,
		Status:               enWorker.Status,
		UserID:               enUser.ID,
		Login:                &enUser.Login,
		Password:             &enUser.Password,
		RoleID:               enUser.RoleID,
		Active:               enUser.Active,
		MustChangePassword:   enUser.MustChangePassword,
	}, nil
}

func (w *workerUserService) WorkerModelToEntity(model model.Model) (entity.Entity, error) {
	mdWorker, ok := model.(*worker_user.WorkerUser)
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

func (w *workerUserService) UserModelToEntity(workerId int, model model.Model) (entity.Entity, error) {
	mdUser, ok := model.(*worker_user.WorkerUser)
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
		RoleID:             mdUser.RoleID,
		Active:             mdUser.Active,
		MustChangePassword: mdUser.MustChangePassword,
		WorkerID:           workerId,
	}, nil
}
