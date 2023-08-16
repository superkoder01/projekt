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

type SettingType string

const (
	SETTINGS = "SETTINGS"

	INTEGER SettingType = "INTEGER"
	BOOLEAN SettingType = "BOOLEAN"
	FLOAT   SettingType = "FLOAT"
)

// TODO: error handling
func (c SettingType) Value() string {
	return string(c)
}

type SettingsEntity interface {
	SetName(string)
	SetValue(string)
	SetType(SettingType)
}

type Settings struct {
	ID    int         `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	Name  string      `gorm:"column:NAME;size:45;default:null"`
	Value string      `gorm:"column:VALUE;size:45;default:null"`
	Type  SettingType `gorm:"column:TYPE" sql:"type:ENUM('INTEGER','BOOLEAN','FLOAT')"`
}

func NewSettings() *Settings {
	return &Settings{}
}

func (z *Settings) TableName() string {
	return ZONE
}

func (z *Settings) SetName(s string) {
	z.Name = s
}

func (z *Settings) SetValue(s string) {
	z.Value = s
}

func (z *Settings) SetType(s SettingType) {
	z.Type = s
}
