package mariadb_dao

import (
	"encoding/json"
	"fmt"
)

type ServiceAccessPoint struct {
	ID          uint   `gorm:"column:ID"`
	SapCode     string `gorm:"column:SAP_CODE"`
	MeterNumber string `gorm:"column:METER_NUMBER"`
	City        string `gorm:"column:CITY"`
	Address     string `gorm:"column:ADDRESS"`
	AccountId   int    `gorm:"column:ACCOUNT_ID"`
}

func (ServiceAccessPoint) TableName() string {
	return "SERVICE_ACCESS_POINT"
}

func (p ServiceAccessPoint) String() string {
	//data, _ := json.Marshal(p)
	data, _ := json.MarshalIndent(p, "", "\t")
	return fmt.Sprintf("%s", data)
}
