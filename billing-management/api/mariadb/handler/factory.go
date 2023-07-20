package handler

import (
	impl "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/handler/impl"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/service"
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
	DISTRIBUTION_NETWORK_OPERATOR = "DISTRIBUTION_NETWORK_OPERATOR"
	PARAMETER_NAME                = "PARAMETER_NAME"
	TARIFF_GROUP_OSD              = "TARIFF_GROUP_OSD"
)

func (hf *handlerFactory) New(name string) Handler {
	switch name {
	case DISTRIBUTION_NETWORK_OPERATOR:
		return impl.NewDistributionNetworkOperatorHandler(hf.svc.New(service.DISTRIBUTION_NETWORK_OPERATOR))
	case PARAMETER_NAME:
		return impl.NewParameterNameHandler(hf.svc.New(service.PARAMETER_NAME))
	case TARIFF_GROUP_OSD:
		return impl.NewTariffGroupOsdHandler(hf.svc.New(service.TARIFF_GROUP), hf.svc.New(service.TARIFF_GROUP_PARAMETER))
	default:
		return nil
	}
}
