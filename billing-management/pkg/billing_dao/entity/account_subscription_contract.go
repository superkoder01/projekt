package entity

const (
	ACCOUNT_SUBSCRIPTION_CONTRACT = "ACCOUNT_SUBSCRIPTION_CONTRACT"
)

type AccountSubscriptionContractEntity interface {
	SetAccountSubscriptionServiceOfferGroupID(int)
	SetAccountSubscriptionAccountID(int)
	SetAccountSubscriptionProviderID(int)
}

type AccountSubscriptionContract struct {
	ContractID                             int `gorm:"column:CONTRACT_ID;size:11;not null"`
	AccountSubscriptionServiceOfferGroupID int `gorm:"column:ACCOUNT_SUBSCRIPTION_SERVICE_OFFER_GROUP_ID;size:11;not null"`
	AccountSubscriptionAccountID           int `gorm:"column:ACCOUNT_SUBSCRIPTION_ACCOUNT_ID;size:11;not null"`
	AccountSubscriptionProviderID          int `gorm:"column:ACCOUNT_SUBSCRIPTION_PROVIDER_ID;size:11;not null"`
}

func NewAccountSubscriptionContract() *AccountSubscriptionContract {
	return &AccountSubscriptionContract{}
}

func (as *AccountSubscriptionContract) TableName() string {
	return ACCOUNT_SUBSCRIPTION_CONTRACT
}

func (as *AccountSubscriptionContract) SetAccountSubscriptionServiceOfferGroupID(i int) {
	as.AccountSubscriptionServiceOfferGroupID = i
}

func (as *AccountSubscriptionContract) SetAccountSubscriptionAccountID(i int) {
	as.AccountSubscriptionAccountID = i
}

func (as *AccountSubscriptionContract) SetAccountSubscriptionProviderID(i int) {
	as.AccountSubscriptionProviderID = i
}
