package handler

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/handler/impl"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/service"
)

type HandlerFactory interface {
	New(string) Handler
}

type handlerFactory struct {
	svc service.ServiceFactory
}

func NewHandlerFactory(svc service.ServiceFactory) *handlerFactory {
	return &handlerFactory{svc: svc}
}

const (
	AUTH                    = "AUTH"
	USER                    = "USER"
	CUSTOMER_ACCOUNT        = "CUSTOMER_ACCOUNT"
	WORKER                  = "WORKER"
	PROVIDER                = "PROVIDER"
	WORKER_USER             = "WORKER_USER"
	CUSTOMER_USER           = "CUSTOMER_USER"
	SERVICE_ACCESS_POINT    = "SERVICE_ACCESS_POINT"
	WORKER_CUSTOMER_ACCOUNT = "WORKER_CUSTOMER_ACCOUNT"
)

func (hf *handlerFactory) New(name string) Handler {
	// Read config
	//ac := conf.GetAuthConfig()

	switch name {
	case AUTH:
		return impl.NewAuthHandler(hf.svc.New(service.AUTH))
	case USER:
		return impl.NewUserHandler(hf.svc.New(service.USER))
	case CUSTOMER_ACCOUNT:
		return impl.NewCustomerAccountHandler(hf.svc.New(service.CUSTOMER_ACCOUNT))
	case PROVIDER:
		return impl.NewProviderHandler(hf.svc.New(service.PROVIDER))
	case WORKER:
		return impl.NewWorkerHandler(hf.svc.New(service.WORKER))
	case WORKER_USER:
		return impl.NewWorkerUserHandler(hf.svc.New(service.WORKER_USER))
	case CUSTOMER_USER:
		return impl.NewCustomerUserHandler(hf.svc.New(service.CUSTOMER_USER))
	case SERVICE_ACCESS_POINT:
		return impl.NewServiceAccessPointHandler(hf.svc.New(service.SERVICE_ACCESS_POINT))
	case WORKER_CUSTOMER_ACCOUNT:
		return impl.NewWorkerCustomerAccountHandler(hf.svc.New(service.WORKER_CUSTOMER_ACCOUNT))
	default:
		return nil
	}
}
