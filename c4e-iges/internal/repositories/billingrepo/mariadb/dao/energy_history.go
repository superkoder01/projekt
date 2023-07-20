package mariadb_dao

import (
	"encoding/json"
	"fmt"
)

type EnergyHistory struct {
	Period                   int  `gorm:"column:PERIOD"`
	CurrentEnergyConsumption int  `gorm:"column:CURRENT_ENERGY_CONSUMPTION"`
	CurrentEnergyExcess      int  `gorm:"column:CURRENT_ENERGY_EXCESS"`
	PreviousEnergyExcess     int  `gorm:"column:PREVIOUS_ENERGY_EXCESS"`
	EnergyExcessUsage        int  `gorm:"column:ENERGY_EXCESS_USAGE"`
	EnergyExcessTransfer     int  `gorm:"column:ENERGY_EXCESS_TRANSFER"`
	ServiceAccessPointId     uint `gorm:"column:SERVICE_ACCESS_POINT_ID"`
}

func (EnergyHistory) TableName() string {
	return "ENERGY_EXCESS"
}

func (p EnergyHistory) String() string {
	//data, _ := json.Marshal(p)
	data, _ := json.MarshalIndent(p, "", "\t")
	return fmt.Sprintf("%s", data)
}
