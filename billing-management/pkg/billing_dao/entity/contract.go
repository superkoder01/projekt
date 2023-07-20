package entity

type ContractStatus string

const (
	CONTRACT = "CONTRACT"

	CS_DRAFT     ContractStatus = "DRAFT"
	CS_FINAL     ContractStatus = "FINAL"
	CS_SENT      ContractStatus = "SENT"
	CS_DELIVERED ContractStatus = "DELIVERED"
	CS_ACCEPTED  ContractStatus = "ACCEPTED"
)

// TODO: error handling
func (c ContractStatus) Value() string {
	return string(c)
}

type ContractEntity interface {
	SetDistributionNetworkOperatorID(int)
}

type Contract struct {
	ID     int            `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	Number string         `gorm:"column:NUMBER;size:45;default:null"`
	Status ContractStatus `gorm:"column:STATUS" sql:"type:ENUM('DRAFT', 'FINAL', 'SENT', 'DELIVERED', 'ACCEPTED')"`
}

func NewContract() *Contract {
	return &Contract{}
}

func (c *Contract) TableName() string {
	return CONTRACT
}

func (c *Contract) SetStatus(cs ContractStatus) {
	c.Status = cs
}

func (c *Contract) SetNumber(s string) {
	c.Number = s
}
