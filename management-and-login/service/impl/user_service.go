package impl

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/administrator"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/customer_account"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/user"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/user_activate"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	apiUtils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/auth"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/entity"
	"gorm.io/gorm/clause"
)

type UserService interface {
	Activate(string, model.Model) error
	ResetPassword(string, model.Model) error
	ForgotPassword(string) error
	SendActivationLink(int) error
	ListSuperAdministrators(interface{}, *apiUtils.Query) (int, []model.Model, error)
	ListAdministrators(interface{}, *apiUtils.Query) (int, []model.Model, error)
}

type userService struct {
	Auth auth.Authenticator
	Dao  dao.Dao
}

func NewUserService(dao dao.Dao, auth auth.Authenticator) *userService {
	return &userService{
		Dao:  dao,
		Auth: auth,
	}
}

func (u *userService) Query(v interface{}, query *apiUtils.Query) (int, []model.Model, error) {
	logger.Debug("Query users")

	sql, err := query.BuildSQLFromApiQuery()
	if err != nil {
		return 0, nil, e.ApiErrImproperQuery
	}

	count, ens, err := u.Dao.Query(v, sql)
	if err != nil {
		return 0, nil, e.ErrInternalServerError
	}
	mdls := []model.Model{}

	for _, e := range ens {
		var customerModel model.Model
		var err error

		customerModel, err = u.EntityToModel(e)
		if err != nil {
			return 0, nil, err
		}
		mdls = append(mdls, customerModel)
	}

	return count, mdls, nil
}

func (u *userService) Check(query *apiUtils.Query) (int, error) {
	logger.Debug("Check users")

	sql, err := query.BuildSQLFromApiCheck()
	if err != nil {
		return 0, e.ApiErrImproperQuery
	}

	count, err := u.Dao.Check(sql)
	if err != nil {
		return 0, e.ErrInternalServerError
	}

	return count, nil
}

func (u *userService) Activate(code string, model model.Model) error {
	logger.Debug("User activation")
	userActivateModel, ok := model.(*user_activate.UserActivate)
	if !ok {
		return e.ApiErrInvalidDataModel
	}

	ens, err := u.Dao.GetByFilter(entity.User{ActivationCode: code})
	if err != nil {
		return e.DbErrEntityNotFound
	}

	if len(ens) != 1 {
		return e.DbErrEntityNotFound
	}

	user, ok := ens[0].(*entity.User)
	if !ok {
		return e.ErrInternalServerError
	}

	newPass, err := u.Auth.Encode(userActivateModel.NewPassword)
	if err != nil {
		return e.ErrInternalServerError
	}

	tx := u.Dao.BeginTransaction()
	defer tx.Rollback()

	user.SetPassword(newPass)
	user.SetActive(true)
	user.SetMustChangePassword(false) // TODO zero values are omitted on update. Should be fixed somehow in DAO or should provide another function to handle situations of this kind
	user.SetActivationCode(" ")

	err = tx.Omit(clause.Associations).Where(&entity.User{ID: user.ID}).Updates(user).Error()
	if err != nil {
		return e.DbErrEntityNotFound
	}

	if user.CustomerAccountID != 0 {
		err = tx.Where(&entity.CustomerAccount{ID: user.CustomerAccountID}).Updates(&entity.CustomerAccount{Status: true}).Error()
	} else if user.WorkerID != 0 {
		err = tx.Where(&entity.Worker{ID: user.WorkerID}).Updates(&entity.Worker{Status: true}).Error()
	}
	if err != nil {
		return e.DbErrEntityNotFound
	}

	return tx.Commit().Error()
}

func (u *userService) ResetPassword(code string, model model.Model) error {
	logger.Debug("User password reset")
	userActivateModel, ok := model.(*user_activate.UserActivate)
	if !ok {
		return e.ApiErrInvalidDataModel
	}

	ens, err := u.Dao.GetByFilter(entity.User{ActivationCode: code})
	if err != nil {
		return e.DbErrEntityNotFound
	}

	if len(ens) != 1 {
		return e.DbErrEntityNotFound
	}

	user, ok := ens[0].(*entity.User)
	if !ok {
		return e.ErrInternalServerError
	}

	newPass, err := u.Auth.Encode(userActivateModel.NewPassword)
	if err != nil {
		return e.ErrInternalServerError
	}

	user.SetPassword(newPass)
	user.SetActivationCode(" ")
	_, err = u.Dao.UpdateByID(user.ID, user)
	if err != nil {
		return e.DbErrEntityNotFound
	}

	return nil
}

func (u *userService) SendActivationLink(id int) error {
	logger.Debug("Send activation link")
	ens, err := u.Dao.GetByFilter(entity.User{CustomerAccountID: id})
	if err != nil {
		return e.DbErrEntityNotFound
	}

	if len(ens) != 1 {
		return e.DbErrEntityNotFound
	}

	err = sendEmailNotificationOnUserCreation(ens[0])
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) ForgotPassword(email string) error {
	logger.Debugf("Forgot password link for email: %s", email)
	ens, err := u.Dao.GetByFilter(entity.User{Email: email})
	if err != nil {
		return e.DbErrEntityNotFound
	}

	if len(ens) != 1 {
		return e.DbErrEntityNotFound
	}

	user := ens[0].(*entity.User)
	user.SetActivationCode(u.Auth.GenerateRandomCode())
	if _, err := u.Dao.UpdateByID(user.ID, user); err != nil {
		return e.DbErrEntityNotFound
	}

	err = sendEmailNotificationOnPasswordForgot(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) Create(model model.Model) (model.Model, error) {
	logger.Debug("User create")
	userModel, ok := model.(*user.User)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	hash, err := u.Auth.Encode(*userModel.Password)
	if err != nil {
		return nil, e.ErrInternalServerError
	}
	userModel.Password = &hash

	userEntity, err := u.ModelToEntity(userModel)
	if err != nil {
		return nil, err
	}

	tx := u.Dao.BeginTransaction()
	defer tx.Rollback()

	// create user entity
	if err = createUserEntity(tx, userEntity, u.Auth.GenerateRandomCode()); err != nil {
		return nil, err
	}

	if err = tx.Commit().Error(); err != nil {
		return nil, e.ErrInternalServerError
	}

	return u.EntityToModel(userEntity)
}

func (u *userService) DeleteByID(id int) error {
	logger.Debugf("User delete by ID: %d", id)
	err := u.Dao.DeleteByID(id)
	if err != nil {
		return e.DbErrEntityNotFound
	}
	return nil
}

func (u *userService) UpdateByID(id int, model model.Model) (model.Model, error) {
	logger.Debugf("User update by ID: %d", id)
	userModel, ok := model.(*user.User)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	en, err := u.ModelToEntity(userModel)
	if err != nil {
		return nil, err
	}

	en, err = u.Dao.UpdateByID(id, en)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}

	return u.EntityToModel(en)
}

func (u *userService) GetByID(id int) (model.Model, error) {
	logger.Debugf("User get by ID: %d", id)
	en, err := u.Dao.GetByID(id)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}

	return u.EntityToModel(en)
}

func (u *userService) List() ([]model.Model, error) {
	logger.Debug("List users")
	ens, err := u.Dao.List()
	if err != nil {
		return nil, e.ErrInternalServerError
	}
	mdls := []model.Model{}

	for _, e := range ens {
		var userModel model.Model
		var err error

		userModel, err = u.EntityToModel(e)
		if err != nil {
			return nil, err
		}
		mdls = append(mdls, userModel)
	}

	return mdls, nil
}

func (u *userService) ListSuperAdministrators(v interface{}, query *apiUtils.Query) (int, []model.Model, error) {
	logger.Debug("Query super admins")

	sql, err := query.BuildSQLFromApiQuery()
	if err != nil {
		return 0, nil, e.ApiErrImproperQuery
	}

	count, ens, err := u.Dao.Query(v, sql)
	if err != nil {
		return 0, nil, e.ErrInternalServerError
	}
	mdls := []model.Model{}

	for _, e := range ens {
		var customerModel model.Model
		var err error

		customerModel, err = u.EntityToSuperAdministratorModel(e)
		if err != nil {
			return 0, nil, err
		}
		mdls = append(mdls, customerModel)
	}

	return count, mdls, nil
}

func (u *userService) ListAdministrators(v interface{}, query *apiUtils.Query) (int, []model.Model, error) {
	logger.Debug("Query administrators")

	sql, err := query.BuildSQLFromApiQuery()
	if err != nil {
		return 0, nil, e.ApiErrImproperQuery
	}

	count, ens, err := u.Dao.Query(v, sql)
	if err != nil {
		return 0, nil, e.ErrInternalServerError
	}
	mdls := []model.Model{}

	for _, e := range ens {
		var customerModel model.Model
		var err error

		customerModel, err = u.EntityToAdministratorModel(e)
		if err != nil {
			return 0, nil, err
		}
		mdls = append(mdls, customerModel)
	}

	return count, mdls, nil
}

func (u *userService) GetWithFilter(query interface{}, args ...interface{}) ([]model.Model, error) {
	logger.Debugf("Get users with filter: %v", query)
	ens, err := u.Dao.GetByFilter(query, args)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}
	mdls := []model.Model{}

	for _, e := range ens {
		var userModel model.Model
		var err error

		userModel, err = u.EntityToModel(e)
		if err != nil {
			return nil, err
		}
		mdls = append(mdls, userModel)
	}

	return mdls, nil
}

func (u *userService) ModelToEntity(model model.Model) (entity.Entity, error) {
	mdUser, ok := model.(*user.User)
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
		Email:              mdUser.Email,
		Login:              login,
		Password:           password,
		ProviderID:         mdUser.ProviderID,
		CustomerAccountID:  mdUser.CustomerAccountID,
		RoleID:             mdUser.RoleID,
		Active:             mdUser.Active,
		MustChangePassword: mdUser.MustChangePassword,
		WorkerID:           mdUser.WorkerID,
		AddedDate:          mdUser.AddedDate,
	}, nil
}

func (u *userService) EntityToModel(en entity.Entity) (model.Model, error) {
	enUser, ok := en.(*entity.User)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	return &user.User{
		Email:              enUser.Email,
		CustomerAccountID:  enUser.CustomerAccountID,
		ProviderID:         enUser.ProviderID,
		ID:                 enUser.ID,
		Active:             enUser.Active,
		Login:              &enUser.Login,
		MustChangePassword: enUser.MustChangePassword,
		Password:           &enUser.Password,
		RoleID:             enUser.RoleID,
		WorkerID:           enUser.WorkerID,
		AddedDate:          enUser.AddedDate,
	}, nil
}

func (u *userService) EntityToAdministratorModel(en entity.Entity) (model.Model, error) {
	enUser, ok := en.(*entity.User)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	return &administrator.Administrator{
		ProviderID: enUser.ProviderID,
		ID:         enUser.ID,
		Active:     enUser.Active,
		Login:      &enUser.Login,
		Email:      enUser.Email,
		RoleID:     enUser.RoleID,
		WorkerID:   enUser.WorkerID,
		FirstName:  enUser.Worker.FirstName,
		LastName:   enUser.Worker.LastName,
		Phone:      enUser.Worker.Phone,
		ExtraInfo:  enUser.Worker.ExtraInfo,
		AddedDate:  enUser.AddedDate,
	}, nil
}

func (u *userService) EntityToSuperAdministratorModel(en entity.Entity) (model.Model, error) {
	enUser, ok := en.(*entity.User)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	return &administrator.SuperAdministrator{
		ProviderID: enUser.ProviderID,
		ID:         enUser.ID,
		Active:     enUser.Active,
		Login:      &enUser.Login,
		Email:      enUser.Email,
		RoleID:     enUser.RoleID,
		WorkerID:   enUser.WorkerID,
		FirstName:  enUser.Worker.FirstName,
		LastName:   enUser.Worker.LastName,
		AddedDate:  enUser.AddedDate,
	}, nil
}

func (u *userService) CustomerEntityToModel(en entity.Entity) (model.Model, error) {
	enCustomer, ok := en.(*entity.CustomerAccount)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	return &customer_account.CustomerAccount{
		ApartmentNumber:    enCustomer.ApartmentNumber,
		BuildingNumber:     enCustomer.BuildingNumber,
		City:               enCustomer.City,
		ProviderID:         enCustomer.ProviderID,
		Country:            enCustomer.Country,
		CustomerTypeName:   enCustomer.CustomerTypeName,
		Email:              enCustomer.Email,
		FirstName:          enCustomer.FirstName,
		ID:                 enCustomer.ID,
		LastName:           enCustomer.LastName,
		Phone:              enCustomer.Phone,
		PostalCode:         enCustomer.PostalCode,
		Province:           enCustomer.Province,
		Street:             enCustomer.Street,
		NIP:                enCustomer.NIP,
		REGON:              enCustomer.REGON,
		Status:             enCustomer.Status,
		PESEL:              enCustomer.PESEL,
		BankAccNumber:      enCustomer.BankAccNumber,
		WorkerID:           enCustomer.WorkerID,
		RegistrationNumber: enCustomer.RegistrationNumber,
		BusinessType:       enCustomer.BusinessType,
		KRS:                enCustomer.KRS,
		LineOfBusiness:     enCustomer.LineOfBusiness,
	}, nil
}

func (u *userService) CustomerModelToEntity(model model.Model) (entity.Entity, error) {
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
