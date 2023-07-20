package controller

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/controller/impl"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/service/mongo"
)

type ControllerFactory interface {
	New(string) Controller
}

type controllerFactory struct {
	sf mongo.ServiceFactory
}

func NewControllerFactory(s mongo.ServiceFactory) *controllerFactory {
	return &controllerFactory{sf: s}
}

const (
	CONTRACT           = "CONTRACT"
	OFFER              = "OFFER"
	INVOICE            = "INVOICE"
	PRICING            = "PRICING"
	OFFER_DRAFTS       = "OFFER_DRAFTS"
	TARIFF_GROUP_LABEL = "TARIFF_GROUP_LABEL"
)

func (hf *controllerFactory) New(name string) Controller {
	switch name {
	case CONTRACT:
		return impl.NewContractController(hf.sf.New(CONTRACT))
	case OFFER:
		return impl.NewOfferController(hf.sf.New(OFFER))
	case INVOICE:
		return impl.NewInvoiceController(hf.sf.New(INVOICE))
	case PRICING:
		return impl.NewPricingController(hf.sf.New(PRICING))
	case OFFER_DRAFTS:
		return impl.NewOfferDraftController(hf.sf.New(OFFER_DRAFTS))
	case TARIFF_GROUP_LABEL:
		return impl.NewTariffGroupLabelController(hf.sf.New(TARIFF_GROUP_LABEL))
	default:
		return nil
	}
}
