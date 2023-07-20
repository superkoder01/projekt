package billing

type OfferPayload struct {
	OfferDetails    OfferDetails       `json:"offerDetails" bson:"offerDetails,omitempty"`
	SellerDetails   OfferSellerDetails `json:"sellerDtls" bson:"sellerDtls,omitempty"`
	CustomerDetails CustomerDetails    `json:"customerDtls" bson:"customerDtls,omitempty"`
	OfferConditions OfferConditions    `json:"conditions" bson:"conditions"`
	PriceList       []OfferPriceList   `json:"priceList" bson:"priceList"`
	Repurchase      Repurchase         `json:"repurchase" bson:"repurchase"`
}

type OfferDetails struct {
	Title         string `json:"title" bson:"title,omitempty"`
	TypeName      string `json:"type" bson:"type,omitempty"`
	OfferDraftId  string `json:"offerDraftId,omitempty" bson:"offerDraftId,omitempty"`
	Number        string `json:"number" bson:"number,omitempty"`
	CreationDate  string `json:"creationDate" bson:"creationDate,omitempty"`
	Status        string `json:"status" bson:"status,omitempty"`
	CustomerId    string `json:"customerId" bson:"customerId,omitempty"`
	TariffGroup   string `json:"tariffGroup" bson:"tariffGroup"`
	AgreementType string `json:"agreementType" bson:"agreementType"`
}

type OfferSellerDetails struct {
	LegalName   string  `json:"legalName" bson:"legalName,omitempty"`
	DisplayName string  `json:"displayName" bson:"displayName,omitempty"`
	KRS         string  `json:"krs" bson:"krs,omitempty"`
	NIP         string  `json:"nip" bson:"nip,omitempty"`
	Address     Address `json:"address" bson:"address,omitempty"`
	Contact     Contact `json:"contact" bson:"contact,omitempty"`
}

type OfferConditions struct {
	Duration                              Duration          `json:"duration" bson:"duration"`
	BillingPeriod                         Duration          `json:"billingPeriod" bson:"billingPeriod"`
	InvoiceDueDate                        string            `json:"invoiceDueDate" bson:"invoiceDueDate"`
	NumberOfPPE                           string            `json:"numberOfPPE" bson:"numberOfPPE"`
	OfferActivePeriod                     OfferActivePeriod `json:"offerActivePeriod" bson:"offerActivePeriod"`
	EstimatedAnnualElectricityConsumption Energy            `json:"estimatedAnnualElectricityConsumption,omitempty" bson:"estimatedAnnualElectricityConsumption,omitempty"`
	EstimatedAnnualElectricityProduction  Energy            `json:"estimatedAnnualElectricityProduction,omitempty" bson:"estimatedAnnualElectricityProduction,omitempty"`
}

type OfferActivePeriod struct {
	StartDate string `json:"startDate" bson:"startDate"`
	EndDate   string `json:"endDate" bson:"endDate"`
}

type OfferPriceList struct {
	Name          string               `json:"name" bson:"name"`
	Id            string               `json:"objectName" bson:"id"`
	Type          string               `json:"type" bson:"type,omitempty"`
	StartDate     string               `json:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate       string               `json:"endDate,omitempty" bson:"endDate,omitempty"`
	Osd           string               `json:"osd,omitempty" bson:"osd,omitempty"`
	TariffGroup   string               `json:"tariffGroup" bson:"tariffGroup,omitempty"`
	Zones         []Zone               `json:"zones,omitempty" bson:"zones,omitempty"`
	CommercialFee []OfferCommercialFee `json:"commercialFee,omitempty" bson:"commercialFee,omitempty"`
}

type OfferCommercialFee struct {
	From  string        `json:"from,omitempty" bson:"from,omitempty"`
	To    string        `json:"to,omitempty" bson:"to,omitempty"`
	Unit  string        `json:"unit,omitempty" bson:"unit,omitempty"`
	Price CommercialFee `json:"price,omitempty" bson:"price,omitempty"`
}
