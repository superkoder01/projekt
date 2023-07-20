package impl

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model"
	a "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/auth"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/enum"
	ml "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/logout"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	apiUtils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/auth"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/entity"
	"net/http"
)

type AuthService interface {
	Authenticate(model model.Model) (*auth.TokensPair, error)
	RefreshToken(string) (*auth.TokensPair, error)
	Logout(model model.Model) error
}

type loginService struct {
	Auth auth.Authenticator
	Dao  dao.Dao
}

func NewAuthService(dao dao.Dao, auth auth.Authenticator) *loginService {
	return &loginService{
		Dao:  dao,
		Auth: auth,
	}
}

func (l *loginService) Authenticate(model model.Model) (*auth.TokensPair, error) {
	logger.Debug("User authentication")
	authModel, ok := model.(*a.Auth)
	if !ok {
		return nil, e.ApiErrInvalidDataModel
	}

	en, err := l.Dao.GetByFilter(entity.User{Login: authModel.Login})
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}

	if len(en) != 1 {
		return nil, e.DbErrEntityNotFound
	}

	user, ok := en[0].(*entity.User)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	err = l.Auth.ComparePassword(user.Password, authModel.Password)
	if err != nil {
		return nil, e.DbErrPasswordMismatch
	}

	if !user.Active {
		return nil, e.DbErrUserNotActivated
	}

	tokenPair, err := l.Auth.CreateToken(enum.Role(user.RoleID).Name(),
		user.ProviderID,
		user.ID,
		user.CustomerAccountID,
		user.WorkerID,
		user.Login)
	if err != nil {
		return nil, e.ErrInternalServerError
	}

	return tokenPair, nil
}

func (l *loginService) RefreshToken(token string) (*auth.TokensPair, error) {
	ti, err := l.Auth.VerifyToken(token, auth.REFRESH)
	if err != nil {
		return nil, e.Wrap(err, http.StatusUnauthorized)
	}

	logger.Debugf("Token refreshing for user with ID: %d", ti.UserID)
	en, err := l.Dao.GetByID(ti.UserID)
	if err != nil {
		return nil, e.DbErrEntityNotFound
	}

	user, ok := en.(*entity.User)
	if !ok {
		return nil, e.ErrInternalServerError
	}

	if !user.Active {
		return nil, e.DbErrUserNotActivated
	}

	tokenPair, err := l.Auth.CreateToken(enum.Role(user.RoleID).Name(),
		user.ProviderID,
		user.ID,
		user.CustomerAccountID,
		user.WorkerID,
		user.Login)
	if err != nil {
		return nil, e.ErrInternalServerError
	}

	uid, err := l.Auth.GetVerifiedUUID(token)
	if err != nil {
		logger.Warningf("error on reading JWT token: %v", err)
	}
	if err := l.Auth.RemoveToken(uid.String()); err != nil {
		logger.Warningf("error on removing token: %v", err)
	}

	return tokenPair, nil
}

func (l *loginService) Logout(model model.Model) error {
	logger.Debug("User logout")
	logoutModel, ok := model.(*ml.Logout)
	if !ok {
		return e.ApiErrInvalidDataModel
	}

	if err := l.Auth.RemoveToken(logoutModel.AccessID); err != nil {
		logger.Errorf("error on token id deletion: %v", err)
		return err
	}
	if err := l.Auth.RemoveToken(logoutModel.RefreshID); err != nil {
		logger.Errorf("error on token id deletion: %v", err)
		return err
	}

	return nil
}

// NOT USED
func (l *loginService) Check(query *apiUtils.Query) (int, error) {
	return 0, nil
}

// NOT USED
func (l *loginService) Query(v interface{}, query *apiUtils.Query) (int, []model.Model, error) {
	return 0, nil, nil
}

// NOT USED
func (l *loginService) Create(model model.Model) (model.Model, error) {
	return nil, nil
}

// NOT USED
func (l *loginService) DeleteByID(i int) error {
	return nil
}

// NOT USED
func (l *loginService) UpdateByID(i int, model model.Model) (model.Model, error) {
	return nil, nil
}

// NOT USED
func (l *loginService) GetByID(i int) (model.Model, error) {
	return nil, nil
}

// NOT USED
func (l *loginService) GetWithFilter(query interface{}, args ...interface{}) ([]model.Model, error) {
	return nil, nil
}

// NOT USED
func (l *loginService) List() ([]model.Model, error) {
	return nil, nil
}

// NOT USED
func (l *loginService) EntityToModel(e entity.Entity) (model.Model, error) {
	return nil, nil
}

// NOT USED
func (l *loginService) ModelToEntity(model model.Model) (entity.Entity, error) {
	return nil, nil
}
