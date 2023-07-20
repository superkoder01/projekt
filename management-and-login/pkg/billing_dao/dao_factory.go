package billing_dao

import (
	d "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/dao"
	db "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/mysql"
)

type DaoFactory interface {
	New(string) d.Dao
}

type daoFactory struct {
	session db.Session
}

func NewDaoFactory(s db.Session) *daoFactory {
	return &daoFactory{session: s}
}

const (
	PROVIDER             = "PROVIDER"
	WORKER               = "WORKER"
	AUTH                 = "AUTH"
	USER                 = "USER"
	CUSTOMER_ACCOUNT     = "CUSTOMER_ACCOUNT"
	SERVICE_ACCESS_POINT = "SERVICE_ACCESS_POINT"
)

func (df *daoFactory) New(name string) d.Dao {
	switch name {
	case PROVIDER:
		return d.NewProviderDao(df.session)
	case AUTH, USER:
		return d.NewUserDao(df.session)
	case CUSTOMER_ACCOUNT:
		return d.NewCustomerAccountDao(df.session)
	case WORKER:
		return d.NewWorkerDao(df.session)
	case SERVICE_ACCESS_POINT:
		return d.NewServiceAccessPointDao(df.session)
	default:
		return nil
	}
}
