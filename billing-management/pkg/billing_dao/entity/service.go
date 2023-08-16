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

type ServiceType string

const (
	SERVICE = "SERVICE"

	SALE       ServiceType = "SALE"
	REPURCHASE ServiceType = "REPURCHASE"
)

// TODO: error handling
func (s ServiceType) Value() string {
	return string(s)
}

type ServiceEntity interface {
	SetProviderID(int)
	SetName(string)
	SetType(ServiceType)
}

type Service struct {
	ID         int         `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	ProviderID int         `gorm:"column:PROVIDER_ID;size:11;not null"`
	Name       string      `gorm:"column:NAME;size:45;default:null"`
	Type       ServiceType `gorm:"column:TYPE" sql:"type:ENUM('SALE', 'REPURCHASE')"`
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) TableName() string {
	return SERVICE
}

func (s *Service) SetProviderID(i int) {
	s.ProviderID = i
}

func (s *Service) SetName(name string) {
	s.Name = name
}

func (s *Service) SetType(t ServiceType) {
	s.Type = t
}
