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

const (
	REGION = "REGION"
)

type RegionEntity interface {
	SetProviderID(int)
	SetName(string)
}

type Region struct {
	ID                            int    `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	DistributionNetworkOperatorID int    `gorm:"column:DISTRIBUTION_NETWORK_OPERATOR_ID;size:11;not null"`
	Name                          string `gorm:"column:NAME;size:45;default:null"`
}

func NewRegion() *Region {
	return &Region{}
}

func (r *Region) TableName() string {
	return REGION
}

func (r *Region) SetDistributionNetworkOperatorID(i int) {
	r.DistributionNetworkOperatorID = i
}

func (r *Region) SetName(s string) {
	r.Name = s
}
