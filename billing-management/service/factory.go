package service

import (
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/service/impl"
)

type ServiceFactory interface {
	New(string) Service
}

type serviceFactory struct {
	df bd.DaoFactory
}

func NewServiceFactory(df bd.DaoFactory) *serviceFactory {
	return &serviceFactory{df: df}
}

const (
	DISTRIBUTION_NETWORK_OPERATOR = "DISTRIBUTION_NETWORK_OPERATOR"
	PARAMETER_NAME                = "PARAMETER_NAME"
	TARIFF_GROUP                  = "TARIFF_GROUP"
	TARIFF_GROUP_PARAMETER        = "TARIFF_GROUP_PARAMETER"
)

func (sf *serviceFactory) New(name string) Service {
	switch name {
	case DISTRIBUTION_NETWORK_OPERATOR:
		return impl.NewDistributionNetworkOperatorService(sf.df.New(bd.DISTRIBUTION_NETWORK_OPERATOR))
	case PARAMETER_NAME:
		return impl.NewParameterNameService(sf.df.New(bd.PARAMETER_NAME))
	case TARIFF_GROUP:
		return impl.NewTariffGroupService(sf.df.New(bd.TARIFF_GROUP))
	default:
		return nil
	}
}
