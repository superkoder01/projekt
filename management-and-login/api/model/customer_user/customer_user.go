package customer_user

import (
	"fmt"
)

type CustomerUser struct {
	ProviderID         int     `json:"providerId"`
	CustomerAccountID  int     `json:"customerAccountId,omitempty"`
	CustomerTypeName   string  `json:"customerTypeName"`
	UserID             int     `json:"userId,omitempty"`
	Login              *string `json:"login"`
	Password           *string `json:"password"`
	Active             bool    `json:"active,omitempty"`
	MustChangePassword bool    `json:"mustChangePassword,omitempty"`
	FirstName          string  `json:"firstName,omitempty"`
	LastName           string  `json:"lastName,omitempty"`
	Status             bool    `json:"status,omitempty"`
	Street             string  `json:"street,omitempty"`
	BuildingNumber     string  `json:"buildingNumber,omitempty"`
	ApartmentNumber    string  `json:"apartmentNumber,omitempty"`
	City               string  `json:"city,omitempty"`
	PostalCode         string  `json:"postalCode,omitempty"`
	Province           string  `json:"province,omitempty"`
	Country            string  `json:"country,omitempty"`
	Phone              string  `json:"phone,omitempty"`
	Email              string  `json:"email,omitempty"`
	NIP                string  `json:"nip,omitempty"`
	REGON              string  `json:"regon,omitempty"`
	PESEL              string  `json:"pesel,omitempty"`
	BankAccNumber      string  `json:"bankAccNumber,omitempty"`
	WorkerID           int     `json:"workerId,omitempty"`
	RoleID             int     `json:"roleId,omitempty"`
	RegistrationNumber int64   `json:"registrationNumber,omitempty"`
	BusinessType       string  `json:"businessType,omitempty"`
	KRS                string  `json:"krs,omitempty"`
	LineOfBusiness     string  `json:"lineOfBusiness,omitempty"`
}

func NewCustomerUser() *CustomerUser {
	return &CustomerUser{}
}

func (c *CustomerUser) String() string {
	return fmt.Sprintf("%s", *c)
}
