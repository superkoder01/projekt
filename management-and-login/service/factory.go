package service

import (
	"github.com/go-redis/redis/v8"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/configuration"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/auth"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service/impl"
	"time"
)

type ServiceFactory interface {
	New(string) Service
}

type serviceFactory struct {
	df   bd.DaoFactory
	rCli *redis.Client
}

func NewServiceFactory(df bd.DaoFactory, rCli *redis.Client) *serviceFactory {
	return &serviceFactory{
		df:   df,
		rCli: rCli,
	}
}

const (
	AUTH                    = "AUTH"
	USER                    = "USER"
	CUSTOMER_ACCOUNT        = "CUSTOMER_ACCOUNT"
	PROVIDER                = "PROVIDER"
	WORKER                  = "WORKER"
	WORKER_USER             = "WORKER_USER"
	CUSTOMER_USER           = "CUSTOMER_USER"
	SERVICE_ACCESS_POINT    = "SERVICE_ACCESS_POINT"
	WORKER_CUSTOMER_ACCOUNT = "WORKER_CUSTOMER_ACCOUNT"
)

func (cf *serviceFactory) New(name string) Service {
	// Read config
	ac := conf.GetAuthConfig()
	authConf := auth.AuthConfig{
		KeyFilePath:                ac.KeyPath,
		RedisClient:                cf.rCli,
		AccessTokenExpirationTime:  time.Second * time.Duration(ac.AccessExpirationTime),
		RefreshTokenExpirationTime: time.Second * time.Duration(ac.RefreshExpirationTime),
	}

	switch name {
	case AUTH:
		return impl.NewAuthService(cf.df.New(bd.AUTH), auth.NewAuth(&authConf))
	case USER:
		return impl.NewUserService(cf.df.New(bd.USER), auth.NewAuth(&authConf))
	case CUSTOMER_ACCOUNT:
		return impl.NewCustomerAccountService(cf.df.New(bd.CUSTOMER_ACCOUNT))
	case PROVIDER:
		return impl.NewProviderService(cf.df.New(bd.PROVIDER))
	case WORKER:
		return impl.NewWorkerService(cf.df.New(bd.WORKER))
	case WORKER_USER:
		return impl.NewWorkerUserService(cf.df.New(bd.USER), cf.df.New(bd.WORKER), auth.NewAuth(&authConf))
	case CUSTOMER_USER:
		return impl.NewCustomerUserService(cf.df.New(bd.USER), cf.df.New(bd.CUSTOMER_ACCOUNT), auth.NewAuth(&authConf))
	case SERVICE_ACCESS_POINT:
		return impl.NewServiceAccessPointService(cf.df.New(bd.SERVICE_ACCESS_POINT))
	case WORKER_CUSTOMER_ACCOUNT:
		return impl.NewWorkerCustomerAccountService(cf.df.New(bd.CUSTOMER_ACCOUNT), cf.df.New(bd.WORKER), cf.df.New(bd.SERVICE_ACCESS_POINT))
	default:
		return nil
	}
}
