/*
BMSFES. 
Copyright (C) 2022-2035 

This file is part of the BMSFES solution. 
BMSFES is free software: you can redistribute it and/or modify 
it under the terms of the GNU Affero General Public License as 
published by the Free Software Foundation, either version 3 of the 
License, or (at your option) any later version.
 
BMSFES solution is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the 
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License 
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
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
