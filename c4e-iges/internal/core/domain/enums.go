package domain

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing.git/pkg/invoice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
)

//InvoiceType type of invoice
type InvoiceType int

const (
	InvoiceTypeSell InvoiceType = iota
	InvoiceTypeRepurchase
)

func (t InvoiceType) Get() string {
	switch t {
	case InvoiceTypeSell:
		return "SP"
	case InvoiceTypeRepurchase:
		return "OD"
	default:
		return "XX"
	}
}

//SettlementType type of settlement
type SettlementType int

func (s SettlementType) String() string {
	switch s {
	case SettlementTypeSell:
		return "sell"
	case SettlementTypeSellRdn:
		return "sell-rdn"
	case SettlementTypeTrade:
		return "trade"
	case SettlementTypeDistribution:
		return "distribution"
	case SettlementTypeRepurchase:
		return "repurchase"
	case SettlementTypeRepurchaseRdn:
		return "repurchase-rdn"
	default:
		return "unknown"
	}
}

const (
	SettlementTypeSell SettlementType = iota
	SettlementTypeSellRdn
	SettlementTypeTrade
	SettlementTypeDistribution
	SettlementTypeRepurchase
	SettlementTypeRepurchaseRdn
)

//InvoiceLineItemCode invoice line item codes
type InvoiceLineItemCode string

func (i InvoiceLineItemCode) ToName(cfg *config.AppConfig) string {
	if value, ok := cfg.InvoiceContent.InvoiceLineItems[string(i)]; ok {
		return value
	}

	return "undefined"
}

func (i InvoiceLineItemCode) ToCode() string {
	return string(i)
}

const (
	InvoiceLineItemSellEnergy                 InvoiceLineItemCode = "ITEM_00"
	InvoiceLineItemTradeFee                   InvoiceLineItemCode = "ITEM_01"
	InvoiceLineItemVarDistributionFee         InvoiceLineItemCode = "ITEM_02"
	InvoiceLineItemConstDistributionFee       InvoiceLineItemCode = "ITEM_03"
	InvoiceLineItemSubscriptionFee            InvoiceLineItemCode = "ITEM_04"
	InvoiceLineItemRenewableEnergyFee         InvoiceLineItemCode = "ITEM_05"
	InvoiceLineItemCogenerationFee            InvoiceLineItemCode = "ITEM_06"
	InvoiceLineItemTransitionalFee            InvoiceLineItemCode = "ITEM_07"
	InvoiceLineItemQualityFee                 InvoiceLineItemCode = "ITEM_08"
	InvoiceLineItemPowerFee                   InvoiceLineItemCode = "ITEM_09"
	InvoiceLineItemRepurchaseEnergy           InvoiceLineItemCode = "ITEM_10"
	InvoiceLineItemDistributionEnergy         InvoiceLineItemCode = "ITEM_11"
	InvoiceLineItemSellEnergyValue            InvoiceLineItemCode = "ITEM_12"
	InvoiceLineItemDepositUsage               InvoiceLineItemCode = "ITEM_13"
	InvoiceLineItemExcessIncludedInSettlement InvoiceLineItemCode = "ITEM_14"
	InvoiceLineItemHistoryIncomeValue         InvoiceLineItemCode = "ITEM_15"
	InvoiceLineItemHistoryOutcomeValue        InvoiceLineItemCode = "ITEM_16"
	InvoiceLineItemHistoryDepositValue        InvoiceLineItemCode = "ITEM_17"
	InvoiceLineItemHistoryUsedValue           InvoiceLineItemCode = "ITEM_18"
	InvoiceLineItemHistoryResidualValue       InvoiceLineItemCode = "ITEM_19"
	InvoiceLineItemHistoryIncomeAmount        InvoiceLineItemCode = "ITEM_20"
	InvoiceLineItemHistoryOutcomeAmount       InvoiceLineItemCode = "ITEM_21"
	InvoiceLineItemHistoryDepositAmount       InvoiceLineItemCode = "ITEM_22"
	InvoiceLineItemHistoryUsedAmount          InvoiceLineItemCode = "ITEM_23"
	InvoiceLineItemHistoryResidualAmount      InvoiceLineItemCode = "ITEM_24"
)

//ItemCodeConverter item code converter
type ItemCodeConverter invoice.ItemCode

func (icc ItemCodeConverter) Convert() InvoiceLineItemCode {
	switch invoice.ItemCode(icc) {
	case invoice.SaleFeeCode:
		return InvoiceLineItemSellEnergy
	case invoice.TradeFeeCode:
		return InvoiceLineItemTradeFee
	case invoice.VarDistributionFeeCode:
		return InvoiceLineItemVarDistributionFee
	case invoice.ConstDistributionFeeCode:
		return InvoiceLineItemConstDistributionFee
	case invoice.SubscriptionFeeCode:
		return InvoiceLineItemSubscriptionFee
	case invoice.RenewableEnergyFeeCode:
		return InvoiceLineItemRenewableEnergyFee
	case invoice.CogenerationFeeCode:
		return InvoiceLineItemCogenerationFee
	case invoice.TransitionalFeeCode:
		return InvoiceLineItemTransitionalFee
	case invoice.QualityFeeCode:
		return InvoiceLineItemQualityFee
	case invoice.PowerFeeCode:
		return InvoiceLineItemPowerFee
	case invoice.RepurchaseCode:
		return InvoiceLineItemRepurchaseEnergy
	default:
		return "undefined"
	}
}
