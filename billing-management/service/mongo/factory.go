package mongo

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mongo"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/service/mongo/impl"
)

type ServiceFactory interface {
	New(string) Service
}

type serviceFactory struct {
	cf mongo.CollectionFactory
}

func NewServiceFactory(s mongo.CollectionFactory) *serviceFactory {
	return &serviceFactory{cf: s}
}

const (
	CONTRACT           = "CONTRACT"
	OFFER              = "OFFER"
	INVOICE            = "INVOICE"
	PRICING            = "PRICING"
	OFFER_DRAFTS       = "OFFER_DRAFTS"
	TARIFF_GROUP_LABEL = "TARIFF_GROUP_LABEL"
)

func (sf *serviceFactory) New(name string) Service {
	switch name {
	case CONTRACT:
		return impl.NewContractService(sf.cf.New(CONTRACT))
	case OFFER:
		return impl.NewOfferService(sf.cf.New(OFFER))
	case INVOICE:
		return impl.NewInvoiceService(sf.cf.New(INVOICE))
	case PRICING:
		return impl.NewPricingService(sf.cf.New(PRICING))
	case OFFER_DRAFTS:
		return impl.NewOfferDraftService(sf.cf.New(OFFER_DRAFTS))
	case TARIFF_GROUP_LABEL:
		return impl.NewTariffGroupLabelService(sf.cf.New(TARIFF_GROUP_LABEL))
	default:
		return nil

	}

}
