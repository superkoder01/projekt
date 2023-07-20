package mariadb_dao

import (
	"encoding/json"
	"fmt"
	"time"
)

type EnergyProduction struct {
	StartDate            time.Time `gorm:"column:START_DATE"`
	EndDate              time.Time `gorm:"column:END_DATE"`
	EnergyAmount         int       `gorm:"column:ENERGY_AMOUNT"`
	EnergyAmountUnits    string    `gorm:"column:ENERGY_AMOUNT_UNITS"`
	NetPrice             float32   `gorm:"column:NET_PRICE"`
	NetValue             float32   `gorm:"column:NET_VALUE"`
	ServiceAccessPointId uint      `gorm:"column:SERVICE_ACCESS_POINT_ID"`
}

func (EnergyProduction) TableName() string {
	return "ENERGY_PRODUCTION"
}

func (p EnergyProduction) String() string {
	//data, _ := json.Marshal(p)
	data, _ := json.MarshalIndent(p, "", "\t")
	return fmt.Sprintf("%s", data)
}
