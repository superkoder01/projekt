package entity

const (
	SERVICE_OFFER = "SERVICE_OFFER"
)

type ServiceOfferEntity interface {
	SetServiceID(int)
	SetProviderID(int)
	SetRatingPlanID(int)
	SetServiceOfferGroupID(int)
	SetName(string)
}

type ServiceOffer struct {
	ID                  int    `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	ServiceID           int    `gorm:"column:SERVICE_ID;size:11;not null"`
	ProviderID          int    `gorm:"column:PROVIDER_ID;size:11;not null"`
	RatingPlanID        int    `gorm:"column:RATING_PLAN_ID;size:11;not null"`
	ServiceOfferGroupID int    `gorm:"column:SERVICE_OFFER_GROUP_ID;size:11;not null"`
	Name                string `gorm:"column:NAME;size:45;default:null"`
}

func NewServiceOffer() *ServiceOffer {
	return &ServiceOffer{}
}

func (so *ServiceOffer) TableName() string {
	return SERVICE_OFFER
}

func (so *ServiceOffer) SetServiceID(i int) {
	so.ServiceID = i
}

func (so *ServiceOffer) SetProviderID(i int) {
	so.ProviderID = i
}

func (so *ServiceOffer) SetRatingPlanID(i int) {
	so.RatingPlanID = i
}

func (so *ServiceOffer) SetServiceOfferGroupID(i int) {
	so.ServiceOfferGroupID = i
}

func (so *ServiceOffer) SetName(s string) {
	so.Name = s
}
