package entity

type BalanceTypeName string

const (
	BALANCE_TYPE = "BALANCE_TYPE"
)

type BalanceType struct {
	Name string `gorm:"column:NAME;primaryKey;size:45"`
}

func NewBalanceType() *BalanceType {
	return &BalanceType{}
}

func (tt *BalanceType) TableName() string {
	return BALANCE_TYPE
}
