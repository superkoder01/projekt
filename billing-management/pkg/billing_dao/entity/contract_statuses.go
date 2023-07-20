package entity

const (
	CONTRACT_STATUSES = "CONTRACT_STATUSES"
)

type ContractStatuses struct {
	Name string `gorm:"column:NAME;primaryKey;size:45"`
}

func NewContractStatuses() *ContractStatuses {
	return &ContractStatuses{}
}

func (tt *ContractStatuses) TableName() string {
	return CONTRACT_STATUSES
}
