package provider

import (
	"fmt"
	"time"
)

type Provider struct {
	ID                    int       `json:"id,omitempty"`
	Name                  string    `json:"name"`
	Type                  string    `json:"type,omitempty"`
	Status                bool      `json:"status,omitempty"`
	NIP                   string    `json:"nip"`
	REGON                 string    `json:"regon,omitempty"`
	KRS                   string    `json:"krs,omitempty"`
	Email                 string    `json:"email,omitempty"`
	PhoneNumber           string    `json:"phoneNumber,omitempty"`
	Street                string    `json:"street,omitempty"`
	BuildingNumber        string    `json:"buildingNumber,omitempty"`
	ApartmentNumber       string    `json:"apartmentNumber,omitempty"`
	PostalCode            string    `json:"postalCode,omitempty"`
	Province              string    `json:"province,omitempty"`
	City                  string    `json:"city,omitempty"`
	Country               string    `json:"country,omitempty"`
	LicenseID             string    `json:"licenseId,omitempty"`
	LicenseExpirationDate time.Time `json:"licenseExpirationDate"`
	LicenseArea           string    `json:"licenseArea,omitempty"`
	WWW                   string    `json:"www,omitempty"`
}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) String() string {
	return fmt.Sprintf("%s", *p)
}
