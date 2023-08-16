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
package billing_dao

import (
	d "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/dao"
	db "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mysql"
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
	PROVIDER                      = "PROVIDER"
	WORKER                        = "WORKER"
	AUTH                          = "AUTH"
	USER                          = "USER"
	WORKER_USER                   = "WORKER_USER"
	CUSTOMER_USER                 = "CUSTOMER_USER"
	CUSTOMER_ACCOUNT              = "CUSTOMER_ACCOUNT"
	ACCOUNT_BALANCE               = "ACCOUNT_BALANCE"
	ACCOUNT_SUBSCRIPTION_CONTRACT = "ACCOUNT_SUBSCRIPTION_CONTRACT"
	ACCOUNT_SUBSCRIPTION          = "ACCOUNT_SUBSCRIPTION"
	BALANCE_TYPE                  = "BALANCE_TYPE"
	CONTRACT                      = "CONTRACT"
	CONTRACT_STATUSES             = "CONTRACT_STATUSES"
	DISTRIBUTION_NETWORK_OPERATOR = "DISTRIBUTION_NETWORK_OPERATOR"
	ENERGY_EXCESS                 = "ENERGY_EXCESS"
	ENERGY_PRODUCTION             = "ENERGY_PRODUCTION"
	INVOICE                       = "INVOICE"
	OFFER_STATUSES                = "OFFER_STATUSES"
	PARAMETER_NAME                = "PARAMETER_NAME"
	RATING_PLAN                   = "RATING_PLAN"
	TARIFF_GROUP_PARAMETER        = "TARIFF_GROUP_PARAMETER"
	RATING_PLAN_TYPE              = "RATING_PLAN_TYPE"
	REGION                        = "REGION"
	SERVICE                       = "SERVICE"
	SERVICE_ACCESS_POINT          = "SERVICE_ACCESS_POINT"
	SERVICE_ACCESS_POINT_CONTRACT = "SERVICE_ACCESS_POINT_CONTRACT"
	SERVICE_OFFER                 = "SERVICE_OFFER"
	SERVICE_OFFER_GROUP           = "SERVICE_OFFER_GROUP"
	SETTINGS                      = "SETTINGS"
	TARIFF_GROUP                  = "TARIFF_GROUP"
	TARIFF_GROUP_LABEL            = "TARIFF_GROUP_LABEL"
	UNIT_TYPE                     = "UNIT_TYPE"
	VAT_RATES                     = "VAT_RATES"
	ZONE                          = "ZONE"
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
	case ACCOUNT_BALANCE:
		return d.NewAccountBalanceDao(df.session)
	case ACCOUNT_SUBSCRIPTION_CONTRACT:
		return d.NewAccountSubscriptionContractDao(df.session)
	case ACCOUNT_SUBSCRIPTION:
		return d.NewAccountSubscriptionDao(df.session)
	case BALANCE_TYPE:
		return d.NewBalanceTypeDao(df.session)
	case CONTRACT:
		return d.NewContractDao(df.session)
	case CONTRACT_STATUSES:
		return d.NewContractStatusesDao(df.session)
	case DISTRIBUTION_NETWORK_OPERATOR:
		return d.NewDistributionNetworkOperatorDao(df.session)
	case ENERGY_EXCESS:
		return d.NewEnergyExcessDao(df.session)
	case ENERGY_PRODUCTION:
		return d.NewEnergyProductionDao(df.session)
	case INVOICE:
		return d.NewInvoiceDao(df.session)
	case OFFER_STATUSES:
		return d.NewOfferStatusesDao(df.session)
	case PARAMETER_NAME:
		return d.NewParameterNameDao(df.session)
	case RATING_PLAN:
		return d.NewRatingPlanDao(df.session)
	case TARIFF_GROUP_PARAMETER:
		return d.NewTariffGroupParameterDao(df.session)
	case RATING_PLAN_TYPE:
		return d.NewRatingPlanTypeDao(df.session)
	case REGION:
		return d.NewRegionDao(df.session)
	case SERVICE:
		return d.NewServiceDao(df.session)
	case SERVICE_ACCESS_POINT:
		return d.NewServiceAccessPointDao(df.session)
	case SERVICE_ACCESS_POINT_CONTRACT:
		return d.NewServiceAccessPointContractDao(df.session)
	case SERVICE_OFFER:
		return d.NewServiceOfferDao(df.session)
	case SERVICE_OFFER_GROUP:
		return d.NewServiceOfferGroupDao(df.session)
	case SETTINGS:
		return d.NewSettingsDao(df.session)
	case TARIFF_GROUP:
		return d.NewTariffGroupDao(df.session)
	case TARIFF_GROUP_LABEL:
		return d.NewTariffGroupLabelDao(df.session)
	case UNIT_TYPE:
		return d.NewUnitTypeDao(df.session)
	case VAT_RATES:
		return d.NewVatRateDao(df.session)
	case ZONE:
		return d.NewZoneDao(df.session)
	default:
		return nil
	}
}
