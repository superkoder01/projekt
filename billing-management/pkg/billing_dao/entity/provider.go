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
package entity

import "time"

const (
	PROVIDER = "PROVIDER"
)

type ProviderEntity interface {
	SetName(string)
	SetType(string)
	SetStatus(int)
	SetNIP(string)
	SetREGON(string)
	SetKRS(string)
	SetEmail(string)
	SetPhoneNumber(string)
	SetStreet(string)
	SetBuildingNumber(string)
	SetApartmentNumber(string)
	SetPostalCode(string)
	SetProvince(string)
	SetCity(string)
	SetCountry(string)
	SetLicenseID(string)
	SetLicenseExpirationDate(time.Time)
	SetLicenseArea(string)
}

type Provider struct {
	ID                    int       `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	Name                  string    `gorm:"column:NAME;size:255;not null;unique"`
	Type                  string    `gorm:"column:TYPE;size:5;default:SOE"`
	Status                bool      `gorm:"column:STATUS;default:0"`
	NIP                   string    `gorm:"column:NIP;size:45;not null;unique"`
	REGON                 string    `gorm:"column:REGON;size:45;default:null"`
	KRS                   string    `gorm:"column:KRS;size:45;default:null"`
	Email                 string    `gorm:"column:EMAIL;size:45;default:null"`
	PhoneNumber           string    `gorm:"column:PHONE_NUMBER;size:45;default:null"`
	Street                string    `gorm:"column:STREET;size:45;default:null"`
	BuildingNumber        string    `gorm:"column:BUILDING_NUMBER;size:45;default:null"`
	ApartmentNumber       string    `gorm:"column:APARTMENT_NUMBER;size:45;default:null"`
	PostalCode            string    `gorm:"column:POSTAL_CODE;size:45;default:null"`
	Province              string    `gorm:"column:PROVINCE;size:45;default:null"`
	City                  string    `gorm:"column:CITY;size:45;default:null"`
	Country               string    `gorm:"column:COUNTRY;size:45;default:null"`
	BlockchainAccAddress  string    `gorm:"column:BLOCKCHAIN_ACC_ADDRESS;size:255;default:null"`
	LicenseID             string    `gorm:"column:LICENSE_ID;size:45;default:null"`
	LicenseExpirationDate time.Time `gorm:"column:LICENSE_EXPIRATION_DATE;default:null"`
	LicenseArea           string    `gorm:"column:LICENSE_AREA;size:45;default:null"`
	WWW                   string    `gorm:"column:WWW;size:45;default:null"`
}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) TableName() string {
	return PROVIDER
}

func (p *Provider) SetName(s string) {
	p.Name = s
}

func (p *Provider) SetType(s string) {
	p.Type = s
}

func (p *Provider) SetStatus(b bool) {
	p.Status = b
}

func (p *Provider) SetNIP(s string) {
	p.NIP = s
}

func (p *Provider) SetREGON(s string) {
	p.REGON = s
}

func (p *Provider) SetKRS(s string) {
	p.KRS = s
}

func (p *Provider) SetEmail(s string) {
	p.Email = s
}

func (p *Provider) SetPhoneNumber(s string) {
	p.PhoneNumber = s
}

func (p *Provider) SetStreet(s string) {
	p.Street = s
}

func (p *Provider) SetBuildingNumber(s string) {
	p.BuildingNumber = s
}

func (p *Provider) SetApartmentNumber(s string) {
	p.ApartmentNumber = s
}

func (p *Provider) SetPostalCode(s string) {
	p.PostalCode = s
}

func (p *Provider) SetProvince(s string) {
	p.Province = s
}

func (p *Provider) SetCity(s string) {
	p.City = s
}

func (p *Provider) SetCountry(s string) {
	p.Country = s
}

func (p *Provider) SetBlockchainAccAddress(s string) {
	p.BlockchainAccAddress = s
}

func (p *Provider) SetLicenseID(s string) {
	p.LicenseID = s
}

func (p *Provider) SetLicenseExpirationDate(t time.Time) {
	p.LicenseExpirationDate = t
}

func (p *Provider) SetLicenseArea(s string) {
	p.LicenseArea = s
}
