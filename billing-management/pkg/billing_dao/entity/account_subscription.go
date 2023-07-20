package entity

import "time"

const (
	ACCOUNT_SUBSCRIPTION = "ACCOUNT_SUBSCRIPTION"
)

type AccountSubscriptionEntity interface {
	SetServiceOfferGroupID(int)
	SetAccountID(int)
	SetProviderID(int)
	SetStatus(bool)
	SetPriority(int)
	SetStartDate(time.Time)
	SetEndDate(time.Time)
}

type AccountSubscription struct {
	Name                string    `gorm:"column:NAME;size:45;primaryKey;not null"`
	ServiceOfferGroupID int       `gorm:"column:SERVICE_OFFER_GROUP_ID;size:11;not null"`
	AccountID           int       `gorm:"column:ACCOUNT_ID;size:11;not null"`
	ProviderID          int       `gorm:"column:PROVIDER_ID;size:11;not null"`
	Status              bool      `gorm:"column:STATUS;default:1"`
	Priority            int       `gorm:"column:PRIORITY;size:11;not null"`
	StartDate           time.Time `gorm:"column:START_DATE"`
	EndDate             time.Time `gorm:"column:END_DATE"`
}

func NewAccountSubscription() *AccountSubscription {
	return &AccountSubscription{}
}

func (as *AccountSubscription) TableName() string {
	return ACCOUNT_SUBSCRIPTION
}

func (as *AccountSubscription) SetServiceOfferGroupID(i int) {
	as.ServiceOfferGroupID = i
}

func (as *AccountSubscription) SetAccountID(i int) {
	as.AccountID = i
}

func (as *AccountSubscription) SetProviderID(i int) {
	as.ProviderID = i
}

func (as *AccountSubscription) SetStatus(b bool) {
	as.Status = b
}

func (as *AccountSubscription) SetPriority(i int) {
	as.Priority = i
}

func (as *AccountSubscription) SetStartDate(t time.Time) {
	as.StartDate = t
}

func (as *AccountSubscription) SetEndDate(t time.Time) {
	as.EndDate = t
}
