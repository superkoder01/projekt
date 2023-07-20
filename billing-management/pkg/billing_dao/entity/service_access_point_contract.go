package entity

const (
	SERVICE_ACCESS_POINT_CONTRACT = "SERVICE_ACCESS_POINT_CONTRACT"
)

type ServiceAccessPointContractEntity interface {
	SetContractID(int)
	SetEnergyPointID(int)
}

type ServiceAccessPointContract struct {
	ID                     int   `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	ContractID             int   `gorm:"column:CONTRACT_ID;size:11;not null"`
	AnnualEnergyUsage      int   `gorm:"column:ANNUAL_ENERGY_USAGE;size:11"`
	AnnualEnergyUsageUnits Units `gorm:"column:ANNUAL_ENERGY_USAGE_UNITS"`
	ServiceAccessPointID   int   `gorm:"column:SERVICE_ACCESS_POINT_ID;size:11;not null"`
}

func NewServiceAccessPointContract() *ServiceAccessPointContract {
	return &ServiceAccessPointContract{}
}

func (sapc *ServiceAccessPointContract) TableName() string {
	return SERVICE_ACCESS_POINT_CONTRACT
}

func (sapc *ServiceAccessPointContract) SetContractID(i int) {
	sapc.ContractID = i
}

func (sapc *ServiceAccessPointContract) SetAnnualEnergyUsage(i int) {
	sapc.AnnualEnergyUsage = i
}

func (sapc *ServiceAccessPointContract) SetAnnualEnergyUsageUnits(i Units) {
	sapc.AnnualEnergyUsageUnits = i
}

func (sapc *ServiceAccessPointContract) SetServiceAccessPointID(i int) {
	sapc.ServiceAccessPointID = i
}
