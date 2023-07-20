package mongo

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/mongodb"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mongo/collections"
)

type CollectionFactory interface {
	New(string) collections.Collection
}

type collectionFactory struct {
	session mongodb.MongoSession
}

func NewCollectionFactory(s mongodb.MongoSession) *collectionFactory {
	return &collectionFactory{session: s}
}

const (
	CONTRACT           = "CONTRACT"
	OFFER              = "OFFER"
	INVOICE            = "INVOICE"
	PRICING            = "PRICING"
	OFFER_DRAFTS       = "OFFER_DRAFTS"
	TARIFF_GROUP_LABEL = "TARIFF_GROUP_LABEL"
)

func (cf *collectionFactory) New(name string) collections.Collection {
	switch name {
	case CONTRACT:
		return collections.NewContractCollection(cf.session)
	case OFFER:
		return collections.NewOfferCollection(cf.session)
	case INVOICE:
		return collections.NewInvoiceCollection(cf.session)
	case PRICING:
		return collections.NewPricingCollection(cf.session)
	case OFFER_DRAFTS:
		return collections.NewOfferDraftCollection(cf.session)
	case TARIFF_GROUP_LABEL:
		return collections.NewTariffGroupLabelCollection(cf.session)
	default:
		return nil

	}

}
