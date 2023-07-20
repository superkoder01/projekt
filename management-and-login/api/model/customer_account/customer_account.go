package customer_account

import "fmt"

type CustomerAccount struct {

	// apartment number
	// Example: 21
	ApartmentNumber string `json:"apartmentNumber,omitempty"`

	// blockchain account address
	BlockchainAccountAddress string `json:"blockchainAccountAddress,omitempty"`

	// building number
	// Example: 33
	BuildingNumber string `json:"buildingNumber,omitempty"`

	// city
	// Example: Krakow
	City string `json:"city,omitempty"`

	// provider Id
	// Required: true
	ProviderID int `json:"providerId"`

	// country
	// Example: Poland
	Country string `json:"country,omitempty"`

	// customer type id
	CustomerTypeName string `json:"customerTypeName,omitempty"`

	// email
	// Example: adam@ovoo.pl
	Email string `json:"email,omitempty"`

	// first name
	// Example: Adam
	FirstName string `json:"firstName,omitempty"`

	// id
	// Read Only: true
	ID int `json:"id,omitempty"`

	// last name
	// Example: Adam
	LastName string `json:"lastName,omitempty"`

	// phone number
	// Example: 600700800
	Phone string `json:"phone,omitempty"`

	// postal code
	// Example: 31-519
	PostalCode string `json:"postalCode,omitempty"`

	// province
	// Example: Malopolskie
	Province string `json:"province,omitempty"`

	// street
	// Example: Kalwaryjska
	Street string `json:"street,omitempty"`

	// NIP number
	NIP string `json:"nip,omitempty"`

	// REGON number
	REGON string `json:"regon,omitempty"`

	Status bool `json:"status"`

	PESEL string `json:"pesel,omitempty"`

	BankAccNumber string `json:"bankAccNumber,omitempty"`

	WorkerID int `json:"workerId,omitempty"`

	RegistrationNumber int64 `json:"registrationNumber,omitempty"`

	BusinessType string `json:"businessType,omitempty"`

	KRS string `json:"krs,omitempty"`

	LineOfBusiness string `json:"lineOfBusiness,omitempty"`
}

func NewCustomerAccount() *CustomerAccount {
	return &CustomerAccount{}
}

type CustomerAccountOption func(c *CustomerAccount)

func (c *CustomerAccount) String() string {
	return fmt.Sprintf("%s", *c)
}

// SETTERS

func (c *CustomerAccount) SetProviderID(i int) {
	c.ProviderID = i
}

func (c *CustomerAccount) SetCustomerTypeID(s string) {
	c.CustomerTypeName = s
}

func (c *CustomerAccount) SetFirstName(s string) {
	c.FirstName = s
}

func (c *CustomerAccount) SetLastName(s string) {
	c.LastName = s
}

func (c *CustomerAccount) SetPESEL(s string) {
	c.PESEL = s
}

func (c *CustomerAccount) SetNIP(s string) {
	c.NIP = s
}

func (c *CustomerAccount) SetREGON(s string) {
	c.REGON = s
}

func (c *CustomerAccount) SetEmail(s string) {
	c.Email = s
}

func (c *CustomerAccount) SetPhone(s string) {
	c.Phone = s
}

func (c *CustomerAccount) SetStreet(s string) {
	c.Street = s
}

func (c *CustomerAccount) SetBuildingNumber(s string) {
	c.BuildingNumber = s
}

func (c *CustomerAccount) SetApartmentNumber(s string) {
	c.ApartmentNumber = s
}

func (c *CustomerAccount) SetPostalCode(s string) {
	c.PostalCode = s
}

func (c *CustomerAccount) SetProvince(s string) {
	c.Province = s
}

func (c *CustomerAccount) SetCity(s string) {
	c.City = s
}

func (c *CustomerAccount) SetCountry(s string) {
	c.Country = s
}

func (c *CustomerAccount) SetStatus(b bool) {
	c.Status = b
}

func (c *CustomerAccount) SetBankAccNumber(s string) {
	c.BankAccNumber = s
}

func (c *CustomerAccount) SetRegistrationNumber(i int64) {
	c.RegistrationNumber = i
}

func (c *CustomerAccount) SetBusinessType(s string) {
	c.BusinessType = s
}

func (c *CustomerAccount) SetKRS(s string) {
	c.KRS = s
}

func (c *CustomerAccount) SetLineOfBusiness(s string) {
	c.LineOfBusiness = s
}
