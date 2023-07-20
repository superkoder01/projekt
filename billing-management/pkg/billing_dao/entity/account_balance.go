package entity

const (
	ACCOUNT_BALANCE = "ACCOUNT_BALANCE"
)

type AccountBalanceEntity interface {
	SetAccountID(int)
	SetProviderID(int)
	SetBalanceTypeName(string)
	SetName(string)
	SetStatus(bool)
	SetUnits(int)
}

type AccountBalance struct {
	ID              int    `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	ProviderID      int    `gorm:"column:PROVIDER_ID;size:11;not null"`
	AccountID       int    `gorm:"column:ACCOUNT_ID;size:11;not null"`
	BalanceTypeName string `gorm:"column:BALANCE_TYPE_NAME;size:45;not null"`
	Name            string `gorm:"column:NAME;size:45;default:null"`
	Status          bool   `gorm:"column:STATUS;default:1"`
	Units           int    `gorm:"column:UNITS;size:11;not null"`
}

func NewAccountBalance() *AccountBalance {
	return &AccountBalance{}
}

func (ab *AccountBalance) TableName() string {
	return ACCOUNT_BALANCE
}

func (ab *AccountBalance) SetProviderID(i int) {
	ab.ProviderID = i
}

func (ab *AccountBalance) SetAccountID(i int) {
	ab.AccountID = i
}

func (ab *AccountBalance) SetBalanceTypeName(s string) {
	ab.BalanceTypeName = s
}

func (ab *AccountBalance) SetName(s string) {
	ab.Name = s
}

func (ab *AccountBalance) SetStatus(b bool) {
	ab.Status = b
}

func (ab *AccountBalance) SetUnits(i int) {
	ab.Units = i
}
