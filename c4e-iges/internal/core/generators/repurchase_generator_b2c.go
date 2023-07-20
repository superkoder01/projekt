package generators

import (
	"context"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing.git/pkg/invoice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/coreutils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/domain"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
	invoiceutil "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/pkg/utils/invoice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
	"strconv"
	"time"
)

type repurchaseInvoiceGeneratorB2C struct {
	log            logger.Logger
	cfg            *config.AppConfig
	contract       *billing.Contract
	event          *invoice.InvoiceEvent
	invoiceNumber  string
	paymentSummary *domain.PaymentSummary
}

func NewRepurchaseInvoiceGeneratorB2C(contract *billing.Contract, event *invoice.InvoiceEvent, log logger.Logger, cfg *config.AppConfig) ports.RepurchaseInvoiceGenerator {
	return &repurchaseInvoiceGeneratorB2C{
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

func (service *repurchaseInvoiceGeneratorB2C) GenerateRepurchaseInvoice(ctx context.Context, number string) (*billing.InvoiceProsumentRepurchase, *billing.InvoiceProsumentRepurchaseDetails, error) {
	service.invoiceNumber = number
	service.log.Infof("start generating repurchase invoice, contract: %s, customerId: %s, repurchaseInvoiceNumber: %s",
		service.contract.Payload.ContractDetails.Number,
		service.contract.Payload.CustomerDetails.CustomerId,
		service.invoiceNumber)

	var _invoice localRepurchaseInvoice
	_invoice.create(service)
	invoice := billing.InvoiceProsumentRepurchase(_invoice)

	var _invoiceDetails localRepurchaseInvoiceDetails
	_invoiceDetails.create(service)
	invoiceDetails := billing.InvoiceProsumentRepurchaseDetails(_invoiceDetails)

	return &invoice, &invoiceDetails, nil

}

// Repurchase Invoice Details
type localRepurchaseInvoiceDetails billing.InvoiceProsumentRepurchaseDetails
type localRepurchaseInvoiceDetailsHeader billing.Header
type localRepurchaseInvoiceDetailsPayload billing.ProsumentRepurchaseDetailsPayload

func (i *localRepurchaseInvoiceDetails) create(service *repurchaseInvoiceGeneratorB2C) {
	var h localRepurchaseInvoiceDetailsHeader
	h.createHeader(service)
	i.Header = billing.Header(h)

	var p localRepurchaseInvoiceDetailsPayload
	p.createInvoiceDetails(service)
	p.createRepurchaseRdnDetails(service)
	i.Payload = billing.ProsumentRepurchaseDetailsPayload(p)
}
func (h *localRepurchaseInvoiceDetailsHeader) createHeader(service *repurchaseInvoiceGeneratorB2C) {
	h.Version = service.contract.Header.Version
	h.Provider = service.contract.Header.Provider
	h.Content = billing.Content{Type: service.cfg.RepurchaseDetailsContent.HeaderContentType, Category: service.cfg.RepurchaseDetailsContent.HeaderContentCategory}
}
func (p *localRepurchaseInvoiceDetailsPayload) createInvoiceDetails(service *repurchaseInvoiceGeneratorB2C) {
	p.InvoiceDetails = billing.InvoiceDetails{
		Number:         service.invoiceNumber,
		IssueDt:        time.Now(),
		ServiceDt:      service.event.EndDate,
		Type:           service.cfg.RepurchaseDetailsContent.PayloadInvoiceDetailsType,
		CustomerId:     service.contract.Payload.ContractDetails.CustomerId,
		BillingStartDt: service.event.StartDate,
		BillingEndDt:   service.event.EndDate,
		Catg:           service.cfg.RepurchaseDetailsContent.HeaderContentCategory,
		Status:         "issued",
	}
}
func (p *localRepurchaseInvoiceDetailsPayload) createRepurchaseRdnDetails(service *repurchaseInvoiceGeneratorB2C) {
	repurchaseRdnDetails := make([]billing.RdnMeterRecord, 0, len(service.event.ServiceAccessPoints))

	for sap, sapBillingData := range service.event.ServiceAccessPoints {
		service.log.Infof("repurchase rdn details, processing sapCode: %s", sap)
		repurchaseRdnDetails = append(repurchaseRdnDetails, *processSapRepurchaseRdnDetails(service, sap, sapBillingData.RepurchaseDetails))
	}

	// -----------------------------------------------------------------------------------------------------------------
	p.RepurchaseDetails = repurchaseRdnDetails
}

func processSapRepurchaseRdnDetails(service *repurchaseInvoiceGeneratorB2C, sapCode string, rdnDetails []invoice.RepurchaseDetailRecord) *billing.RdnMeterRecord {
	service.log.Infof("processing %s %s details", sapCode, domain.SettlementTypeRepurchaseRdn)

	rdnItems := make([]billing.RdnItem, 0, len(rdnDetails))

	sap := coreutils.FindSap(sapCode, service.contract.Payload.ServiceAccessPoints)
	if sap == nil {
		service.log.Errorf("seems like service access point %s does not exist in client contract %s - exiting", sapCode, service.contract.Payload.ContractDetails.Number)
		return new(billing.RdnMeterRecord)
	}

	for _, rdn := range rdnDetails {
		rdnItem := billing.RdnItem{
			Date:         rdn.Date,
			Hour:         string(rdn.HourPeriod),
			EnergyAmount: rdn.Units,
			NetPrice:     rdn.UnitPrice,
			VatRate:      rdn.VatRate,
			NetVal:       rdn.Net,
			TaxVal:       rdn.Vat,
			GrossVal:     rdn.Gross,
		}

		rdnItems = append(rdnItems, rdnItem)
	}

	return &billing.RdnMeterRecord{
		MeterNumber: sap.MeterNumber,
		RdnItems:    rdnItems,
	}
}

// Repurchase Invoice
func (service *repurchaseInvoiceGeneratorB2C) aggregateRepurchaseSummary(record invoice.ProductionRecord, category domain.SettlementType) {
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
func (service *repurchaseInvoiceGeneratorB2C) invoiceLines(vatMap map[int]*domain.Summary, invoiceLineItemCode domain.InvoiceLineItemCode) []billing.SellSummaryItem {
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

type localRepurchaseInvoice billing.InvoiceProsumentRepurchase
type localRepurchaseInvoiceHeader billing.Header
type localRepurchaseInvoicePayload billing.ProsumentRepurchasePayload

func (i *localRepurchaseInvoice) create(service *repurchaseInvoiceGeneratorB2C) {
	var h localRepurchaseInvoiceHeader
	h.createHeader(service)
	i.Header = billing.Header(h)

	var p localRepurchaseInvoicePayload
	p.createInvoiceDetails(service)
	p.createSellerDetails(service)
	p.createCustomerDetails(service)
	p.createSapDetails(service)
	p.createActiveEnergyConsumed(service)
	p.createSellSummary(service)
	i.Payload = billing.ProsumentRepurchasePayload(p)
}
func (h *localRepurchaseInvoiceHeader) createHeader(service *repurchaseInvoiceGeneratorB2C) {
	h.Version = service.contract.Header.Version
	h.Provider = service.contract.Header.Provider
	h.Content = billing.Content{Type: service.cfg.RepurchaseContent.HeaderContentType, Category: service.cfg.RepurchaseContent.HeaderContentCategory}
}
func (p *localRepurchaseInvoicePayload) createInvoiceDetails(service *repurchaseInvoiceGeneratorB2C) {
	p.InvoiceDetails = billing.InvoiceDetails{
		Number:         service.invoiceNumber,
		IssueDt:        time.Now(),
		ServiceDt:      service.event.EndDate,
		Type:           service.cfg.RepurchaseContent.PayloadInvoiceDetailsType,
		CustomerId:     service.contract.Payload.ContractDetails.CustomerId,
		BillingStartDt: service.event.StartDate,
		BillingEndDt:   service.event.EndDate,
		Catg:           service.cfg.RepurchaseContent.HeaderContentCategory,
		Status:         "issued",
	}
}
func (p *localRepurchaseInvoicePayload) createSellerDetails(service *repurchaseInvoiceGeneratorB2C) {
	p.SellerDetails = billing.PartyDetails{
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
func (p *localRepurchaseInvoicePayload) createCustomerDetails(service *repurchaseInvoiceGeneratorB2C) {
	p.CustomerDetails = billing.PartyDetails{
		LegalName:   service.contract.Payload.SellerDetails.LegalName,
		DisplayName: service.contract.Payload.SellerDetails.DisplayName,
		Nip:         service.contract.Payload.SellerDetails.Nip,
		Regon:       service.contract.Payload.SellerDetails.Regon,
		Krs:         service.contract.Payload.SellerDetails.Krs,
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
func (p *localRepurchaseInvoicePayload) createSapDetails(service *repurchaseInvoiceGeneratorB2C) {
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
			PpeCode:   sap.SapCode,
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
func (p *localRepurchaseInvoicePayload) createActiveEnergyConsumed(service *repurchaseInvoiceGeneratorB2C) {
	energyRepurchase := make([]billing.EnergyReading, len(service.event.ServiceAccessPoints))

	for sap, sapBillingData := range service.event.ServiceAccessPoints {
		service.log.Infof("active energy repurchase, processing sapCode: %s", sap)
		energyRepurchase = append(energyRepurchase, *processSapEnergyRepurchase(service, sap, sapBillingData.Repurchase))
	}

	// -----------------------------------------------------------------------------------------------------------------
	p.ActiveEnergyConsumed = billing.ActiveEnergyConsumed{
		EnergySell: *mergeMeters(energyRepurchase),
	}
}
func (p *localRepurchaseInvoicePayload) createSellSummary(service *repurchaseInvoiceGeneratorB2C) {
	sellSummary := new(billing.SellSummary)

	sellSummary.Items = append(sellSummary.Items, service.invoiceLines(service.paymentSummary.EnergyRepurchase, domain.InvoiceLineItemSellEnergy)...)

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

	service.log.Infof("repurchase invoice sell summary: %v", service.paymentSummary)
}

func processSapEnergyRepurchase(service *repurchaseInvoiceGeneratorB2C, sap string, records []invoice.ProductionRecord) *billing.EnergyReading {
	return processSapEnergyRepurchaseRecord(service, sap, records, domain.SettlementTypeRepurchase)
}
func processSapEnergyRepurchaseRecord(service *repurchaseInvoiceGeneratorB2C, sap string, records []invoice.ProductionRecord, feeCategory domain.SettlementType) *billing.EnergyReading {
	service.log.Infof("processing %s %s records", sap, feeCategory)

	fees := new(billing.EnergyReading)
	var totalRepurchase float64
	var totalNetVal float64

	meters := make(map[string][]billing.Item)

	for _, fee := range records {
		service.aggregateRepurchaseSummary(fee, feeCategory)

		item := billing.Item{
			ItemName: domain.ItemCodeConverter(invoice.SaleFeeCode).Convert().ToName(service.cfg),
			ItemCode: domain.ItemCodeConverter(invoice.SaleFeeCode).Convert().ToCode(),
			PrevMeterRead: billing.MeterReading{
				Date: fee.From,
			},
			CurrMeterRead: billing.MeterReading{
				Date: fee.To,
			},
			Factor:       1, // always 1 in case repurchase
			Consumption:  fee.Excess,
			NetUnitPrice: fee.UnitPrice,
			NetVal:       fee.Net,
			VatRate:      fee.VatRate,
		}

		totalRepurchase += fee.Excess
		totalNetVal += fee.Net

		meters[fee.MeterNumber] = append(meters[fee.MeterNumber], item)
	}

	for meterNumber, meterItems := range meters {
		fees.Meters = append(fees.Meters, billing.Meter{
			MeterNumber: meterNumber,
			Items:       meterItems,
		})
	}

	fees.Subtotal.Amount = totalRepurchase
	fees.Subtotal.NetValue = invoiceutil.RoundFloat(totalNetVal)

	return fees
}
