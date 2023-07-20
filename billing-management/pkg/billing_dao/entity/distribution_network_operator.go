package entity

const (
	DISTRIBUTION_NETWORK_OPERATOR = "DISTRIBUTION_NETWORK_OPERATOR"
)

type DistributionNetworkOperatorEntity interface {
	SetName(string)
	SetNip(string)
	SetAddress(string)
	SetCity(string)
}

type DistributionNetworkOperator struct {
	ID      int    `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	Name    string `gorm:"column:NAME;size:45;default:null"`
	Nip     string `gorm:"column:NIP;size:45;default:null"`
	Address string `gorm:"column:ADDRESS;size:45;default:null"`
	City    string `gorm:"column:CITY;size:45;default:null"`
}

func NewDistributionNetworkOperator() *DistributionNetworkOperator {
	return &DistributionNetworkOperator{}
}

func (dno *DistributionNetworkOperator) TableName() string {
	return DISTRIBUTION_NETWORK_OPERATOR
}

func (dno *DistributionNetworkOperator) SetName(s string) {
	dno.Name = s
}

func (dno *DistributionNetworkOperator) SetNip(s string) {
	dno.Nip = s
}

func (dno *DistributionNetworkOperator) SetAddress(s string) {
	dno.Address = s
}

func (dno *DistributionNetworkOperator) SetCity(s string) {
	dno.City = s
}
