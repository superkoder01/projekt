package entity

const (
	ENERGY_EXCESS = "ENERGY_EXCESS"
)

type EnergyExcessEntity interface {
}

type EnergyExcess struct {
	ID                       int `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	ServiceAccessPointID     int `gorm:"column:SERVICE_ACCESS_POINT_ID;size:11;not null"`
	Period                   int `gorm:"column:PERIOD;size:11;not null"`
	CurrentEnergyConsumption int `gorm:"column:CURRENT_ENERGY_CONSUMPTION;size:11"`
	CurrentEnergyExcess      int `gorm:"column:CURRENT_ENERGY_EXCESS;size:11"`
	PreviousEnergyExcess     int `gorm:"column:PREVIOUS_ENERGY_EXCESS;size:11"`
	EnergyExcessUsage        int `gorm:"column:ENERGY_EXCESS_USAGE;size:11"`
	EnergyExcessTransfer     int `gorm:"column:ENERGY_EXCESS_TRANSFER;size:11"`
}

func NewEnergyExcess() *EnergyExcess {
	return &EnergyExcess{}
}

func (ea *EnergyExcess) TableName() string {
	return ENERGY_EXCESS
}

func (ea *EnergyExcess) SetServiceAccessPointID(i int) {
	ea.ServiceAccessPointID = i
}

func (ea *EnergyExcess) SetPeriod(i int) {
	ea.Period = i
}

func (ea *EnergyExcess) SetCurrentEnergyConsumption(i int) {
	ea.CurrentEnergyConsumption = i
}

func (ea *EnergyExcess) SetCurrentEnergyExcess(i int) {
	ea.CurrentEnergyExcess = i
}

func (ea *EnergyExcess) SetPreviousEnergyExcess(i int) {
	ea.PreviousEnergyExcess = i
}

func (ea *EnergyExcess) SetEnergyExcessUsage(i int) {
	ea.EnergyExcessUsage = i
}

func (ea *EnergyExcess) SetEnergyExcessTransfer(i int) {
	ea.EnergyExcessTransfer = i
}
