package entity

const (
	CUSTOMER_ACCOUNT = "CUSTOMER_ACCOUNT"
)

type CustomerAccountEntity interface {
	SetProviderID(int)
	SetCustomerTypeName(string)
	SetFirstName(string)
	SetLastName(string)
	SetStatus(bool)
	SetPESEL(string)
	SetNIP(string)
	SetREGON(string)
	SetEmail(string)
	SetPhone(string)
	SetStreet(string)
	SetBuildingNumber(string)
	SetApartmentNumber(string)
	SetPostalCode(string)
	SetProvince(string)
	SetCity(string)
	SetCountry(string)
	SetBankAccNumber(string)
	SetWorkerID(int)
	SetRegistrationNumber(int64)
	SetBusinessType(string)
	SetKRS(string)
	SetLineOfBusiness(string)
}

type CustomerAccount struct {
	ID                 int     `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	ProviderID         int     `gorm:"column:PROVIDER_ID;size:11;not null"`
	CustomerTypeName   string  `gorm:"column:CUSTOMER_TYPE_NAME;size:45;not null"`
	FirstName          string  `gorm:"column:FIRST_NAME;size:45;default:null"`
	LastName           string  `gorm:"column:LAST_NAME;size:45;default:null"`
	Status             bool    `gorm:"column:STATUS;default:0"`
	NIP                string  `gorm:"column:NIP;size:45;default:null"`
	PESEL              string  `gorm:"column:PESEL;size:45;default:null"`
	REGON              string  `gorm:"column:REGON;size:45;default:null"`
	Email              string  `gorm:"column:EMAIL;size:45;default:null"`
	Phone              string  `gorm:"column:PHONE;size:45;default:null"`
	Street             string  `gorm:"column:STREET;size:45;default:null"`
	BuildingNumber     string  `gorm:"column:BUILDING_NUMBER;size:45;default:null"`
	ApartmentNumber    string  `gorm:"column:APARTMENT_NUMBER;size:45;default:null"`
	PostalCode         string  `gorm:"column:POSTAL_CODE;size:45;default:null"`
	Province           string  `gorm:"column:PROVINCE;size:45;default:null"`
	City               string  `gorm:"column:CITY;size:45;default:null"`
	Country            string  `gorm:"column:COUNTRY;size:45;default:null"`
	BankAccNumber      string  `gorm:"column:BANK_ACC_NUMBER;size:45;default:null"`
	WorkerID           int     `gorm:"column:WORKER_ID;size:11;default:null"`
	Worker             *Worker `gorm:"foreignKey:WORKER_ID;references:ID"`
	RegistrationNumber int64   `gorm:"column:REGISTRATION_NUMBER;type:BIGINT;size:12"`
	BusinessType       string  `gorm:"column:BUSINESS_TYPE;size:45;default:null"`
	KRS                string  `gorm:"column:KRS;size:45;default:null"`
	LineOfBusiness     string  `gorm:"column:LINE_OF_BUSINESS;size:45;default:null"`
}

func NewCustomerAccount() *CustomerAccount {
	return &CustomerAccount{}
}

func (c *CustomerAccount) TableName() string {
	return CUSTOMER_ACCOUNT
}

func (c *CustomerAccount) SetProviderID(i int) {
	c.ProviderID = i
}

func (c *CustomerAccount) SetCustomerTypeName(s string) {
	c.CustomerTypeName = s
}

func (c *CustomerAccount) SetFirstName(s string) {
	c.FirstName = s
}

func (c *CustomerAccount) SetLastName(s string) {
	c.LastName = s
}

func (c *CustomerAccount) SetStatus(b bool) {
	c.Status = b
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

func (c *CustomerAccount) SetBankAccNumber(s string) {
	c.BankAccNumber = s
}

func (c *CustomerAccount) SetWorkerID(i int) {
	c.WorkerID = i
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
