package generators

import (
	"context"
	"fmt"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing.git/pkg/invoice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/coreutils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/domain"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
	flowcontrol_util "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/pkg/utils/flowcontrol"
	invoiceutil "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/pkg/utils/invoice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
	"strconv"
	"time"
)

type invoiceGeneratorB2C struct {
	log            logger.Logger
	cfg            *config.AppConfig
	contract       *billing.Contract
	event          *invoice.InvoiceEvent
	invoiceNumber  string
	paymentSummary *domain.PaymentSummary
}

func NewInvoiceGeneratorB2C(contract *billing.Contract, event *invoice.InvoiceEvent, log logger.Logger, cfg *config.AppConfig) ports.InvoiceGenerator {
	return &invoiceGeneratorB2C{
		log:      log,
		cfg:      cfg,
		contract: contract,
		event:    event,
		paymentSummary: &domain.PaymentSummary{
			EnergySell:         make(map[int]*domain.Summary),
			EnergyRepurchase:   make(map[int]*domain.Summary),
			TradeFee:           make(map[int]*domain.Summary),
			EnergyDistribution: make(map[int]*domain.Summary),
			DepositIncluded:    make(map[int]*domain.Summary),
		},
	}
}

func (service *invoiceGeneratorB2C) GenerateInvoice(ctx context.Context, number string) (*billing.InvoiceProsument, interface{}, error) {
	service.invoiceNumber = number
	service.log.Infof("start generating invoice, contract: %s, customerId: %s, invoiceNumber: %s",
		service.contract.Payload.ContractDetails.Number,
		service.contract.Payload.CustomerDetails.CustomerId,
		service.invoiceNumber)

	var _invoice localInvoice
	_invoice.create(service)

	ret := billing.InvoiceProsument(_invoice)
	return &ret, nil, nil
}

func (service *invoiceGeneratorB2C) aggregateSellSummary(record invoice.ConsumptionRecord, category domain.SettlementType) {
	localRecord := record
	if category == domain.SettlementTypeSell || category == domain.SettlementTypeTrade {
		if localRecord.ItemCode == invoice.TradeFeeCode {
			if _, ok := service.paymentSummary.TradeFee[localRecord.VatRate]; !ok {
				service.paymentSummary.TradeFee[localRecord.VatRate] = new(domain.Summary)
			}

			service.paymentSummary.TradeFee[localRecord.VatRate].VatRate = localRecord.VatRate
			service.paymentSummary.TradeFee[localRecord.VatRate].NetValue += localRecord.Net
			service.paymentSummary.TradeFee[localRecord.VatRate].VatValue += localRecord.Vat
			service.paymentSummary.TradeFee[localRecord.VatRate].GrossValue += localRecord.Gross
		} else {
			if _, ok := service.paymentSummary.EnergySell[localRecord.VatRate]; !ok {
				service.paymentSummary.EnergySell[localRecord.VatRate] = new(domain.Summary)
			}

			service.paymentSummary.EnergySell[localRecord.VatRate].VatRate = localRecord.VatRate
			service.paymentSummary.EnergySell[localRecord.VatRate].NetValue += localRecord.Net
			service.paymentSummary.EnergySell[localRecord.VatRate].VatValue += localRecord.Vat
			service.paymentSummary.EnergySell[localRecord.VatRate].GrossValue += localRecord.Gross
		}
	}

	if category == domain.SettlementTypeDistribution {
		if _, ok := service.paymentSummary.EnergyDistribution[localRecord.VatRate]; !ok {
			service.paymentSummary.EnergyDistribution[localRecord.VatRate] = new(domain.Summary)
		}

		service.paymentSummary.EnergyDistribution[localRecord.VatRate].VatRate = localRecord.VatRate
		service.paymentSummary.EnergyDistribution[localRecord.VatRate].NetValue += localRecord.Net
		service.paymentSummary.EnergyDistribution[localRecord.VatRate].VatValue += localRecord.Vat
		service.paymentSummary.EnergyDistribution[localRecord.VatRate].GrossValue += localRecord.Gross
	}
}
func (service *invoiceGeneratorB2C) aggregateRepurchaseSummary(record invoice.ProductionRecord, category domain.SettlementType) {
	localRecord := record
	if category == domain.SettlementTypeRepurchase {
		if _, ok := service.paymentSummary.EnergyRepurchase[localRecord.VatRate]; !ok {
			service.paymentSummary.EnergyRepurchase[localRecord.VatRate] = new(domain.Summary)
		}

		service.paymentSummary.EnergyRepurchase[localRecord.VatRate].VatRate = localRecord.VatRate
		service.paymentSummary.EnergyRepurchase[localRecord.VatRate].NetValue += localRecord.Net
		service.paymentSummary.EnergyRepurchase[localRecord.VatRate].VatValue += localRecord.Vat
		service.paymentSummary.EnergyRepurchase[localRecord.VatRate].GrossValue += localRecord.Gross
	}
}
func (service *invoiceGeneratorB2C) aggregateExciseSummary(excise float64) {
	service.paymentSummary.Excise += excise
}
func (service *invoiceGeneratorB2C) aggregateDepositSummary(deposit invoice.DepositRecord) {
	if _, ok := service.paymentSummary.DepositIncluded[0]; !ok {
		service.paymentSummary.DepositIncluded[0] = new(domain.Summary)
	}

	// TODO deposit should keep vat history
	service.paymentSummary.DepositIncluded[0].VatRate = 0
	service.paymentSummary.DepositIncluded[0].NetValue += deposit.Value.UsedValue
	service.paymentSummary.DepositIncluded[0].VatValue += 0
	service.paymentSummary.DepositIncluded[0].GrossValue += deposit.Value.UsedValue
}
func (service *invoiceGeneratorB2C) invoiceLines(vatMap map[int]*domain.Summary, invoiceLineItemCode domain.InvoiceLineItemCode) []billing.SellSummaryItem {
	items := make([]billing.SellSummaryItem, 0)

	for vat, invoiceLine := range vatMap {
		items = append(items, billing.SellSummaryItem{
			ItemName: invoiceLineItemCode.ToName(service.cfg),
			ItemCode: invoiceLineItemCode.ToCode(),
			VatRate:  vat,
			NetVal:   invoiceutil.RoundFloat(invoiceLine.NetValue),
			TaxVal:   invoiceutil.RoundFloat(invoiceLine.VatValue),
			GrossVal: invoiceutil.RoundFloat(invoiceLine.GrossValue),
		})
	}

	return items
}

type localInvoice billing.InvoiceProsument
type localInvoiceHeader billing.Header
type localInvoicePayload billing.ProsumentPayload

func (i *localInvoice) create(service *invoiceGeneratorB2C) {
	var h localInvoiceHeader
	h.createHeader(service)
	i.Header = billing.Header(h)

	var p localInvoicePayload
	p.createInvoiceDetails(service)
	p.createSellerDetails(service)
	p.createCustomerDetails(service)
	p.createPaymentDetails(service)
	p.createPayerDetails(service)
	p.createSapDetails(service)
	p.createActiveEnergyConsumed(service)
	p.createActiveEnergyProduced(service)
	p.createDepositSummary(service)
	p.createExcessSalesBalance(service)
	p.createSellSummary(service)
	p.createPaymentSummary(service)
	p.createEnergyAnnualBalance(service)
	p.createSapSummary(service)
	i.Payload = billing.ProsumentPayload(p)
}
func (h *localInvoiceHeader) createHeader(service *invoiceGeneratorB2C) {
	h.Version = service.contract.Header.Version
	h.Provider = service.contract.Header.Provider
	h.Content = billing.Content{Type: service.cfg.InvoiceContent.HeaderContentType, Category: service.cfg.InvoiceContent.HeaderContentCategory}
}
func (p *localInvoicePayload) createInvoiceDetails(service *invoiceGeneratorB2C) {
	p.InvoiceDetails = billing.InvoiceDetails{
		Number:         service.invoiceNumber,
		IssueDt:        time.Now(),
		ServiceDt:      service.event.EndDate,
		Type:           service.cfg.InvoiceContent.PayloadInvoiceDetailsType,
		CustomerId:     service.contract.Payload.ContractDetails.CustomerId,
		BillingStartDt: service.event.StartDate,
		BillingEndDt:   service.event.EndDate,
		Catg:           service.cfg.InvoiceContent.HeaderContentCategory,
		Status:         "issued",
	}
}
func (p *localInvoicePayload) createSellerDetails(service *invoiceGeneratorB2C) {
	p.SellerDetails = billing.PartyDetails{
		LegalName:         service.contract.Payload.SellerDetails.LegalName,
		DisplayName:       service.contract.Payload.SellerDetails.DisplayName,
		Krs:               service.contract.Payload.SellerDetails.Krs,
		Nip:               service.contract.Payload.SellerDetails.Nip,
		Regon:             service.contract.Payload.SellerDetails.Regon,
		BankAccountNumber: service.contract.Payload.SellerDetails.BankAccountNumber,
		Address: billing.Address{
			Street:   service.contract.Payload.SellerDetails.Address.Street,
			PostCode: service.contract.Payload.SellerDetails.Address.PostCode,
			City:     service.contract.Payload.SellerDetails.Address.City,
		},
		Contact: billing.Contact{
			Address: billing.Address{
				Street:   service.contract.Payload.SellerDetails.Contact.Address.Street,
				PostCode: service.contract.Payload.SellerDetails.Contact.Address.PostCode,
				City:     service.contract.Payload.SellerDetails.Contact.Address.City,
			},
			PhoneNumbers: clonePhoneNumbers(service.contract.Payload.SellerDetails.Contact.PhoneNumbers),
			Email:        service.contract.Payload.SellerDetails.Contact.Email,
			WWW:          service.contract.Payload.SellerDetails.Contact.WWW,
		},
	}
}
func (p *localInvoicePayload) createCustomerDetails(service *invoiceGeneratorB2C) {
	p.CustomerDetails = billing.PartyDetails{
		CustomerId:  service.contract.Payload.CustomerDetails.CustomerId,
		FirstName:   service.contract.Payload.CustomerDetails.FirstName,
		LastName:    service.contract.Payload.CustomerDetails.LastName,
		DisplayName: service.contract.Payload.CustomerDetails.DisplayName,
		Pesel:       service.contract.Payload.CustomerDetails.Pesel,
		Nip:         service.contract.Payload.CustomerDetails.Nip,
		Regon:       service.contract.Payload.CustomerDetails.Regon,
		Address: billing.Address{
			Street:   service.contract.Payload.CustomerDetails.Address.Street,
			PostCode: service.contract.Payload.CustomerDetails.Address.PostCode,
			City:     service.contract.Payload.CustomerDetails.Address.City,
		},
		Contact: billing.Contact{
			Address: billing.Address{
				Street:   service.contract.Payload.CustomerDetails.Contact.Address.Street,
				PostCode: service.contract.Payload.CustomerDetails.Contact.Address.PostCode,
				City:     service.contract.Payload.CustomerDetails.Contact.Address.City,
			},
			PhoneNumbers: clonePhoneNumbers(service.contract.Payload.CustomerDetails.Contact.PhoneNumbers),
			Email:        service.contract.Payload.CustomerDetails.Contact.Email,
			WWW:          service.contract.Payload.CustomerDetails.Contact.WWW,
		},
	}
}
func (p *localInvoicePayload) createPaymentDetails(service *invoiceGeneratorB2C) {
	p.PaymentDetails = billing.PaymentDetails{
		BankDetails: billing.BankDetails{
			Account: service.contract.Payload.SellerDetails.BankAccountNumber,
		},
		PaymentTitle: fmt.Sprintf(service.cfg.InvoiceContent.PayloadPaymentDetailsPaymentTitleTemplate, service.contract.Payload.CustomerDetails.CustomerId),
		PaymentDueDt: paymentDueDate(service.contract.Payload.ContractConditions.InvoiceDueDate),
	}
}
func (p *localInvoicePayload) createPayerDetails(service *invoiceGeneratorB2C) {
	p.PayerDetails = billing.PartyDetails{
		CustomerId:  service.contract.Payload.CustomerDetails.CustomerId,
		FirstName:   service.contract.Payload.CustomerDetails.FirstName,
		LastName:    service.contract.Payload.CustomerDetails.LastName,
		DisplayName: service.contract.Payload.CustomerDetails.DisplayName,
		Nip:         service.contract.Payload.CustomerDetails.Nip,
		Regon:       service.contract.Payload.CustomerDetails.Regon,
		Address: billing.Address{
			Street:   service.contract.Payload.CustomerDetails.Address.Street,
			PostCode: service.contract.Payload.CustomerDetails.Address.PostCode,
			City:     service.contract.Payload.CustomerDetails.Address.City,
		},
		Contact: billing.Contact{
			Address: billing.Address{
				Street:   service.contract.Payload.CustomerDetails.Contact.Address.Street,
				PostCode: service.contract.Payload.CustomerDetails.Contact.Address.PostCode,
				City:     service.contract.Payload.CustomerDetails.Contact.Address.City,
			},
			PhoneNumbers: clonePhoneNumbers(service.contract.Payload.CustomerDetails.Contact.PhoneNumbers),
			Email:        service.contract.Payload.CustomerDetails.Contact.Email,
			WWW:          service.contract.Payload.CustomerDetails.Contact.WWW,
		},
	}
}
func (p *localInvoicePayload) createSapDetails(service *invoiceGeneratorB2C) {
	saps := service.event.ServiceAccessPoints

	ppeItems := make([]billing.PpeItem, 0, len(saps))

	sapCodes := make([]string, 0, len(saps))
	for k, _ := range saps {
		sapCodes = append(sapCodes, k)
	}

	for _, sapCode := range sapCodes {
		sap := coreutils.FindSap(sapCode, service.contract.Payload.ServiceAccessPoints)
		if sap == nil {
			service.log.Errorf("seems like service access point %s does not exist in client contract %s - ignoring", sapCode, service.contract.Payload.ContractDetails.Number)
			continue
		}

		contractedPower, err := strconv.ParseFloat(sap.ContractedPower.Amount, 64)
		if err != nil {
			service.log.Warnf("sap code %s, error when parsing contracted power", sapCode)
			contractedPower = 0
		}

		ppeItem := billing.PpeItem{
			PpeCode:   sapCode,
			PpeName:   sap.ObjectName,
			PpeObName: sap.ObjectName,
			Address: billing.Address{
				// TODO sap address shall be taken from sap not from customerDetails -> fix contract
				Street:   service.contract.Payload.CustomerDetails.Address.Street,
				PostCode: service.contract.Payload.CustomerDetails.Address.PostCode,
				City:     service.contract.Payload.CustomerDetails.Address.City,
			},
			TariffGroup: sap.TariffGroup,
			ContractedPower: billing.ContractedPower{
				Value: contractedPower,
				Unit:  sap.ContractedPower.Unit,
			},
		}

		ppeItems = append(ppeItems, ppeItem)
	}

	p.PpeDetails = ppeItems
}
func (p *localInvoicePayload) createActiveEnergyConsumed(service *invoiceGeneratorB2C) {
	energySell := make([]billing.EnergyReading, len(service.event.ServiceAccessPoints))
	energyDistribution := make([]billing.EnergyReading, len(service.event.ServiceAccessPoints))
	excise := make([]float64, 0)

	for sap, sapBillingData := range service.event.ServiceAccessPoints {
		service.log.Infof("active energy consumption, processing sapCode: %s", sap)
		energySell = append(energySell, *processSapEnergySell(service, sap, sapBillingData.Sale))
		energyDistribution = append(energyDistribution, *processSapEnergyDistribution(service, sap, sapBillingData.Distribution))
		excise = append(excise, processSapExcise(service, sap, sapBillingData.Excise))
	}

	// -----------------------------------------------------------------------------------------------------------------
	p.ActiveEnergyConsumed = billing.ActiveEnergyConsumed{
		EnergySell:         *mergeMeters(energySell),
		EnergyDistribution: *mergeMeters(energyDistribution),
	}

	p.ActiveEnergyConsumed.EnergySell.ExciseTax = mergeExcise(excise)
}
func (p *localInvoicePayload) createActiveEnergyProduced(service *invoiceGeneratorB2C) {
	energyProduction := make([]billing.ActiveEnergyProduced, 0)

	for sap, sapBillingData := range service.event.ServiceAccessPoints {
		service.log.Infof("active energy production, processing sapCode: %s", sap)
		energyProduction = append(energyProduction, *processSapEnergyProduction(service, sap, sapBillingData.Repurchase))
	}

	// -----------------------------------------------------------------------------------------------------------------
	p.ActiveEnergyProduced = *mergeEnergyProduction(energyProduction)
}
func (p *localInvoicePayload) createDepositSummary(service *invoiceGeneratorB2C) {
	deposit := new(billing.PpeDeposit)

	depositItems := make([]billing.PpeDepositItem, 0)

	for sap, sapBillingData := range service.event.ServiceAccessPoints {
		service.log.Infof("deposit summary, processing sapCode: %s", sap)
		depositItems = append(depositItems, *processSapDepositItem(service, sap, sapBillingData.Deposit))
	}

	deposit.Deposit = depositItems
	deposit.PpeDepositTotal = *mergeDepositSummary(depositItems)
	// -----------------------------------------------------------------------------------------------------------------
	p.DepositSummary = *deposit
}
func (p *localInvoicePayload) createExcessSalesBalance(service *invoiceGeneratorB2C) {
	excessSaleBalance := new(billing.ExcessSalesBalance)

	excessSaleBalance.Items = append(excessSaleBalance.Items, billing.ExcessSalesBalanceItem{
		ItemName: domain.InvoiceLineItemSellEnergyValue.ToName(service.cfg),
		ItemCode: domain.InvoiceLineItemSellEnergyValue.ToCode(),
		GrossVal: domain.SummaryMap(service.paymentSummary.EnergySell).Sum().GrossValue,
	})

	excessSaleBalance.Items = append(excessSaleBalance.Items, billing.ExcessSalesBalanceItem{
		ItemName: domain.InvoiceLineItemDepositUsage.ToName(service.cfg),
		ItemCode: domain.InvoiceLineItemDepositUsage.ToCode(),
		GrossVal: domain.SummaryMap(service.paymentSummary.DepositIncluded).Sum().Negate().GrossValue,
	})

	for _, item := range excessSaleBalance.Items {
		excessSaleBalance.Summary.GrossVal += item.GrossVal
	}

	excessSaleBalance.Summary.GrossVal = invoiceutil.RoundFloat(excessSaleBalance.Summary.GrossVal)

	// -----------------------------------------------------------------------------------------------------------------
	p.ExcessSalesBalance = *excessSaleBalance
}
func (p *localInvoicePayload) createSellSummary(service *invoiceGeneratorB2C) {
	sellSummary := new(billing.SellSummary)

	sellSummary.Items = append(sellSummary.Items, service.invoiceLines(service.paymentSummary.EnergySell, domain.InvoiceLineItemSellEnergy)...)
	sellSummary.Items = append(sellSummary.Items, service.invoiceLines(service.paymentSummary.EnergyDistribution, domain.InvoiceLineItemDistributionEnergy)...)
	sellSummary.Items = append(sellSummary.Items, service.invoiceLines(service.paymentSummary.TradeFee, domain.InvoiceLineItemTradeFee)...)

	for _, item := range sellSummary.Items {
		sellSummary.Total.NetVal += item.NetVal
		sellSummary.Total.TaxVal += item.TaxVal
		sellSummary.Total.GrossVal += item.GrossVal
	}

	sellSummary.Total.NetVal = invoiceutil.RoundFloat(sellSummary.Total.NetVal)
	sellSummary.Total.TaxVal = invoiceutil.RoundFloat(sellSummary.Total.TaxVal)
	sellSummary.Total.GrossVal = invoiceutil.RoundFloat(sellSummary.Total.GrossVal)

	// -----------------------------------------------------------------------------------------------------------------
	p.SellSummary = *sellSummary
}
func (p *localInvoicePayload) createPaymentSummary(service *invoiceGeneratorB2C) {
	paymentSummary := new(billing.PaymentSummary)

	paymentSummary.Items = append(paymentSummary.Items, service.invoiceLines(service.paymentSummary.EnergySell, domain.InvoiceLineItemSellEnergy)...)
	paymentSummary.Items = append(paymentSummary.Items, service.invoiceLines(service.paymentSummary.EnergyDistribution, domain.InvoiceLineItemDistributionEnergy)...)
	paymentSummary.Items = append(paymentSummary.Items, service.invoiceLines(service.paymentSummary.TradeFee, domain.InvoiceLineItemTradeFee)...)

	paymentSummary.Items = append(paymentSummary.Items, service.invoiceLines(domain.SummaryMap(service.paymentSummary.DepositIncluded).Negate(), domain.InvoiceLineItemExcessIncludedInSettlement)...)

	for _, item := range paymentSummary.Items {
		paymentSummary.Total.NetVal += item.NetVal
		paymentSummary.Total.TaxVal += item.TaxVal
		paymentSummary.Total.GrossVal += item.GrossVal
	}

	paymentSummary.Total.NetVal = invoiceutil.RoundFloat(paymentSummary.Total.NetVal)
	paymentSummary.Total.TaxVal = invoiceutil.RoundFloat(paymentSummary.Total.TaxVal)
	paymentSummary.Total.GrossVal = invoiceutil.RoundFloat(paymentSummary.Total.GrossVal)

	// -----------------------------------------------------------------------------------------------------------------
	p.PaymentSummary = *paymentSummary
}
func (p *localInvoicePayload) createEnergyAnnualBalance(service *invoiceGeneratorB2C) {
	energyAnnualBalanceValue := new(billing.EnergyAnnualBalance)
	energyAnnualBalanceAmount := new(billing.EnergyAnnualBalance)

	for sap, sapBillingData := range service.event.ServiceAccessPoints {
		service.log.Infof("energy annual balance, processing sapCode: %s", sap)
		history := domain.EnergyDepositHistory{Sap: sap, Deposit: sapBillingData.Deposit, Cfg: service.cfg}
		value, amount := history.GetLast(13)
		energyAnnualBalanceValue.History = append(energyAnnualBalanceValue.History, *value)
		energyAnnualBalanceAmount.History = append(energyAnnualBalanceAmount.History, *amount)
	}

	// -----------------------------------------------------------------------------------------------------------------
	p.EnergyValueAnnualBalance = *energyAnnualBalanceValue
	p.EnergyAmountAnnualBalance = *energyAnnualBalanceAmount
}
func (p *localInvoicePayload) createSapSummary(service *invoiceGeneratorB2C) {
	sapSummary := new(billing.PpeSummary)

	for sap, _ := range service.event.ServiceAccessPoints {
		service.log.Infof("sap summary, processing sapCode: %s", sap)

		sapSummary.Items = append(sapSummary.Items, billing.PpeSummaryItem{
			PpeCode:        sap,
			Value:          invoiceutil.RoundFloat(p.PaymentSummary.Total.GrossVal),
			EnergyConsumed: p.ActiveEnergyConsumed.EnergySell.Subtotal.Amount,
			EnergyProduced: p.ActiveEnergyProduced.Summary.Production,
		})
	}

	for _, item := range sapSummary.Items {
		sapSummary.Total.Value += item.Value
		sapSummary.Total.EnergyProduced += item.EnergyProduced
		sapSummary.Total.EnergyConsumed += item.EnergyConsumed
	}

	// -----------------------------------------------------------------------------------------------------------------
	p.PpeSummary = *sapSummary

	service.log.Infof("invoice payment summary: %v", service.paymentSummary)
}

func mergeDepositSummary(items []billing.PpeDepositItem) *billing.PpeDepositSummary {
	result := new(billing.PpeDepositSummary)

	for _, item := range items {
		result.DepositCurrent += item.PpeDepositSummary.DepositCurrent
		result.DepositConsumed += item.PpeDepositSummary.DepositConsumed
		result.DepositNext += item.PpeDepositSummary.DepositNext
	}

	return result
}

func processSapDepositItem(service *invoiceGeneratorB2C, sap string, deposit invoice.Deposit) *billing.PpeDepositItem {
	service.log.Infof("processing %s deposit summary", sap)
	sapDepositItem := new(billing.PpeDepositItem)

	service.aggregateDepositSummary(deposit.Records[len(deposit.Records)-1])

	sapDepositItem.PpeCode = sap
	sapDepositItem.PpeDepositSummary.DepositCurrent = deposit.Records[len(deposit.Records)-1].Value.Deposit
	sapDepositItem.PpeDepositSummary.DepositConsumed = deposit.Records[len(deposit.Records)-1].Value.UsedValue
	sapDepositItem.PpeDepositSummary.DepositNext = deposit.Records[len(deposit.Records)-1].Value.ResidualValue

	return sapDepositItem
}

func mergeEnergyProduction(production []billing.ActiveEnergyProduced) *billing.ActiveEnergyProduced {
	result := new(billing.ActiveEnergyProduced)

	for _, p := range production {
		result.Meters = append(result.Meters, p.Meters...)
		result.Summary.Production += p.Summary.Production
		result.Summary.TaxVal += p.Summary.TaxVal
		result.Summary.NetVal += p.Summary.NetVal
		result.Summary.GrossVal += p.Summary.GrossVal
	}

	return result
}

func processSapEnergyProduction(service *invoiceGeneratorB2C, sap string, records []invoice.ProductionRecord) *billing.ActiveEnergyProduced {
	service.log.Infof("processing %s %s records", sap, domain.SettlementTypeRepurchase)

	energyProduced := new(billing.ActiveEnergyProduced)
	var totalProduction float64
	var totalNetVal float64
	var totalTaxVal float64
	var totalGrossVal float64

	meters := make(map[string][]billing.MeterProductionItem)

	for _, rep := range records {
		service.aggregateRepurchaseSummary(rep, domain.SettlementTypeRepurchase)

		item := billing.MeterProductionItem{
			ItemName:   domain.ItemCodeConverter(rep.ItemCode).Convert().ToName(service.cfg),
			ItemCode:   domain.ItemCodeConverter(rep.ItemCode).Convert().ToCode(),
			DateFrom:   rep.From,
			DateTo:     rep.To,
			Production: rep.Excess,
			VatRate:    rep.VatRate,
			NetVal:     rep.Net,
			TaxVal:     rep.Vat,
			GrossVal:   rep.Gross,
		}

		totalNetVal += rep.Net
		totalTaxVal += rep.Vat
		totalGrossVal += rep.Gross
		totalProduction += rep.Excess

		meters[rep.MeterNumber] = append(meters[rep.MeterNumber], item)
	}

	for meterNumber, meterProductionItems := range meters {
		energyProduced.Meters = append(energyProduced.Meters, billing.MeterProduction{
			MeterNumber: meterNumber,
			Items:       meterProductionItems,
		})
	}

	energyProduced.Summary.Production = totalProduction
	energyProduced.Summary.NetVal = totalNetVal
	energyProduced.Summary.TaxVal = totalTaxVal
	energyProduced.Summary.GrossVal = totalGrossVal

	return energyProduced
}

func mergeExcise(excise []float64) float64 {
	var result float64

	for _, e := range excise {
		result += e
	}

	return result
}

func mergeMeters(meters []billing.EnergyReading) *billing.EnergyReading {
	result := new(billing.EnergyReading)

	for _, meter := range meters {
		result.Meters = append(result.Meters, meter.Meters...)
		result.ExciseTax += meter.ExciseTax
		result.Subtotal.NetValue += meter.Subtotal.NetValue
		result.Subtotal.Amount += meter.Subtotal.Amount
	}

	return result
}

func processSapExcise(service *invoiceGeneratorB2C, sap string, excise float64) float64 {
	service.log.Infof("processing %s excise", sap)

	service.aggregateExciseSummary(excise)

	return excise
}

func processSapEnergySell(service *invoiceGeneratorB2C, sap string, records []invoice.ConsumptionRecord) *billing.EnergyReading {
	return processSapEnergyConsumptionRecord(service, sap, records, domain.SettlementTypeSell)
}

func processSapEnergyDistribution(service *invoiceGeneratorB2C, sap string, records []invoice.ConsumptionRecord) *billing.EnergyReading {
	return processSapEnergyConsumptionRecord(service, sap, records, domain.SettlementTypeDistribution)
}

func processSapEnergyConsumptionRecord(service *invoiceGeneratorB2C, sap string, records []invoice.ConsumptionRecord, feeCategory domain.SettlementType) *billing.EnergyReading {
	service.log.Infof("processing %s %s records", sap, feeCategory)

	fees := new(billing.EnergyReading)
	var totalConsumption float64
	var totalNetVal float64
	var distTotalConsumption float64

	meters := make(map[string][]billing.Item)
	distConsumptionMap := make(map[int]float64)

	for _, fee := range records {
		service.aggregateSellSummary(fee, feeCategory)

		item := billing.Item{
			ItemName: domain.ItemCodeConverter(fee.ItemCode).Convert().ToName(service.cfg),
			ItemCode: domain.ItemCodeConverter(fee.ItemCode).Convert().ToCode(),
			PrevMeterRead: billing.MeterReading{
				Date:     fee.From,
				Value:    fee.Previous,
				ReadType: invoiceutil.FormatReadType(fee.ReadType),
			},
			CurrMeterRead: billing.MeterReading{
				Date:     fee.To,
				Value:    fee.Current,
				ReadType: invoiceutil.FormatReadType(fee.ReadType),
			},
			Factor:       fee.Factor.Value,
			Consumption:  fee.Consumed,
			NetUnitPrice: fee.UnitPrice,
			NetVal:       fee.Net,
			VatRate:      fee.VatRate,
		}

		if feeCategory == domain.SettlementTypeSell {
			totalConsumption += fee.Consumed
		}

		if feeCategory == domain.SettlementTypeDistribution && fee.Consumed > 0 {
			distConsumptionMap[int(fee.ItemCode)] += fee.Consumed
			distTotalConsumption += fee.Consumed
		}

		totalNetVal += fee.Net

		meters[fee.MeterNumber] = append(meters[fee.MeterNumber], item)
	}

	for meterNumber, meterItems := range meters {
		fees.Meters = append(fees.Meters, billing.Meter{
			MeterNumber: meterNumber,
			Items:       meterItems,
		})
	}

	fees.Subtotal.Amount = flowcontrol_util.Ternary(feeCategory == domain.SettlementTypeSell, totalConsumption, distTotalConsumption/float64(len(distConsumptionMap))).(float64)
	fees.Subtotal.NetValue = invoiceutil.RoundFloat(totalNetVal)

	return fees
}

func paymentDueDate(dueDate string) time.Time {
	d, err := strconv.Atoi(dueDate)
	if err != nil {
		d = 14
	}

	return time.Now().Add(time.Duration(d) * time.Hour * 24)
}

func clonePhoneNumbers(in []billing.PhoneNumber) []billing.PhoneNumber {
	var out []billing.PhoneNumber

	for _, p := range in {
		out = append(out, billing.PhoneNumber{
			Type:   p.Type,
			Number: p.Number,
		})
	}

	return out
}
